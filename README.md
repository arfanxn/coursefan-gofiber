# Auth Module (GoFiber/Backend)


## Installation

Install the dependecies

```sh
go get
go mod tidy
```

Configure environment from example.env
```sh
cp example.env {your_env_file}
```

Migrate up Auth Module required tables 
```sh
go run . --env={your_env_file} --migrate-up
```

Run the app
```sh
go run . --env={your_env_file}
```

Remigrate Auth Module tables (if needed)
```sh
go run . --env={your_env_file} --migrate-down --migrate-up
```

Migrate drop Auth Module tables (if needed)
```sh
go run . --env={your_env_file} --migrate-down
```

### Guide
- __ENV__ : You can also set your default *Environtment File* on *./app/env.go*, so you doesn't have to set it manualy via console whenever you run the program


## License
MIT && Auth Module

**Open Source**
