package routes

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
	"url-shortner-fiber-redis/database"
	"url-shortner-fiber-redis/helpers"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"customShort"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"customShort"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"x_rate_remaining"`
	XRateLimitReset time.Duration `json:"x_rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot Parse Json",
		})
	}

	//rate limiting

	r2 := database.CreateClient(1)
	defer r2.Close()
	val, err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":            "Rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
		}
	}

	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}

	if !helpers.CheckDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "",
		})
	}

	body.URL = helpers.EnforceHTTP(body.URL)

	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}
	r := database.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(database.Ctx, id).Result()
	if val != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Custom URL is already in use",
		})
	}
	if body.Expiry == 0 {
		body.Expiry = 24
	}

	err = r.Set(database.Ctx, id, body.URL, body.Expiry).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to connect to Server",
		})
	}

	resp := response{
		URL:             body.URL,
		CustomShort:     "",
		Expiry:          body.Expiry,
		XRateRemaining:  10,
		XRateLimitReset: 30,
	}
	r2.Decr(database.Ctx, c.IP())

	val, err = r2.Get(database.Ctx, c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)
	ttl, _ := r2.TTL(database.Ctx, c.IP()).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute
	resp.CustomShort = os.Getenv("APP_URL") + "/" + id
	return c.Status(fiber.StatusOK).JSON(resp)
}
