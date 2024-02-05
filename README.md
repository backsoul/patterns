[![Go](https://github.com/backsoul/patterns/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/backsoul/patterns/actions/workflows/go.yml)

[![Coverage](https://github.com/backsoul/patterns/blob/master/badge.svg)](https://github.com/backsoul/patterns/blob/master/coverage.out)

# Go Database Factory Pattern Example

This is a simple example project demonstrating the use of the Factory Pattern in Go to connect to MySQL and PostgreSQL databases. The project provides a flexible way to switch between different types of databases by configuring the `.env` file.

## Prerequisites

Before running the project, make sure you have the following installed:

- Go (version 1.14 or later)
- MySQL (for MySQL example)
- PostgreSQL (for PostgreSQL example)

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/backsoul/patterns.git
   ```

2. Edit .env file:

   ```bash
   uncomment line for your favorite engine db.
   ```

3. up containers:

   ```bash
   docker-compose up -d
   ```
