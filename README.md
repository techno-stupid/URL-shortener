# URL Shortener Project

A simple URL Shortener built using Go Fiber and Redis.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This URL Shortener project is a web application built using the Go Fiber framework and Redis database. It allows you to shorten long URLs into easily shareable and manageable short URLs. The project demonstrates how to create a RESTful API for URL shortening, using Redis for storing and retrieving the shortened URLs.

## Features

- Shorten long URLs into short and manageable links.
- Redirect users from short URLs to the original long URLs.
- Statistics tracking for the shortened URLs.
- RESTful API for creating and managing short URLs.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- Go (1.16 or later): https://golang.org/doc/install
- Redis: https://redis.io/download

### Installation

1. Clone this repository:

   ```sh
   git clone git@github.com:techno-stupid/URL-shortener.git```

2. Navigate to the project directory:

   ```sh
   cd url-shortener

3. Install dependencies:


   ```sh
   go mod tidy

4. Build the application:

   ```sh
   go build

## Usage

1. Start the Redis server.

2. Run the URL Shortener application:

  ```sh
      ./url-shortener
  ```

3. The application will start on http://localhost:3000.

4. Use your preferred API client (such as curl or Postman) to interact with the URL Shortener API. Refer to the API documentation for available endpoints and usage.
   
## Configuration
The application's configuration can be modified by editing the config.json file. You can specify the Redis server details, application port, and other settings in this file.

# Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request.
1. Fork the repository.
2. Create a new branch: git checkout -b feature/your-feature-name.
3. Make your changes and commit them: git commit -m 'Add some feature'.
4. Push to the branch: git push origin feature/your-feature-name.
5. Create a pull request.


