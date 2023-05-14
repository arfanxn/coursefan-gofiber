# Coursefan (Gofiber/Backend)

Coursefan is online course application built with Gofiber, Mysql, Ngrok, and Midtrans

## Installation

Install all package dependencies and tidy
```sh
go get
go mod tidy
```

Create environtment files from example.env file
```sh
cp example.env local.env  // Local
cp example.env test.env  // Test 
```

Migrate up Coursefan required tables and seed data to tables
```sh
go run . --mifgrate-up
go run . --seed
```

Migrate down Coursefan tables
```sh
go run . --mifgrate-down
```

Migrate fresh Coursefan tables
```sh
go run . --mifgrate-fresh
```

## Docker

You can also deploy to docker container (Dockerizing/Containerizing)

By default, the Docker will expose port 8000, so change this within the
Dockerfile if necessary. When ready, simply use the Dockerfile to
build the image.

```sh
docker-compose build 
docker-compose docker-compose.yaml up -d
```

Verify the deployment by navigating to your server address in
your preferred browser.

```sh
127.0.0.1:8000
```

## License
MIT && COURSEFAN

**Open Source**
