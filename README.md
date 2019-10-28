# golden-headed-quetzal

Parse and store a structured CSV in a SQL database

## Setup

Run this commands if you want to just use the application

```bash
docker-compose build
```

```bash
docker-compose up
```

## Install dependencies for development

```bash
 make deps
```

## Config Env

- If you are creating a docker you have to copy config-example.toml to config.toml on root directory and alter the fields as you like
- Otherwise, to run the binary directly on the server, you can copy config-example.toml to bin/config/config.toml
- After that, edit config.toml with your database and server settings. You can use docker-compose.yml as a reference.

## Run project in debug mode

First you need to run a postgres

```bash
docker run -it --name pq -e POSTGRES_PASSWORD=golden-headed-quetzal -p 5432:5432 -d postgres:12-alpine
```

### HTTP server started on **1337** port

```bash
 make run-debug
```

### Upload File

To upload a file you can type the command below

```bash
curl -F 'file=@base_teste.txt' http://localhost:1337/upload
```
