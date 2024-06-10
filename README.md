# SSO

The main task of the project is to create and maintain the functionality of a single entry point to various applications
and user management.

The main technologies that this project uses are `Go`, `gRPC`, `PSQL`, `Redis`, `Protobuf` and there are many more
dependencies for this project, which can be seen in the file `go.mod`

The remote procedures that are used in this project are described in the repository:
<a href="https://github.com/Pashgunt/Sso-Protobuf-Golang.git" taget="_blank">Repository</a> and are pulled up through dependencies

### Usage

To start the server and migrations, there is a Makefile at the root of the project. It contains the main commands for
starting the project.

1. `make start` - start command `go run cmd/sso/main.go --config=./config/local.yml`
2. `make migrations` - start command `go run cmd/migrations/migrations.go --config=./config/local.yml`

In these commands, the path to the config file is specified as a parameter, which can be changed if necessary.

To launch docker containers for additional services, go to the docker folder and run the `docker-compose.yml` file there:

```Bash
cd docker && docker compose up -d
```

or start command

```Bash
make up
```

### Endpoints

1. **Registration**
    ```   
    GRPC grpcs://localhost:44044/Auth/Register
    
    {
        "email": "test@mail.ru",
        "password": "test"
    }
    ```

2. **Login**
    ```   
    GRPC grpcs://localhost:44044/Auth/Login
    
    {
        "email": "test@mail.ru",
        "password": "test",
        "appUuid": "fb84ec3b-0040-4046-8319-f685763eb19a"
    }
    ```
