# Build stage
FROM golang:alpine AS build-env
# Install git
RUN apk add --no-cache git

# Copy source code
ADD . /golden-headed-quetzal
WORKDIR /golden-headed-quetzal

# Install dependencies
RUN go get -v ./...

# Build static binary with debug disabled 
RUN go build -ldflags "-w -s" -o golden-headed-quetzal *.go

# Final stage
FROM alpine

# Install bash and ca-certificates to run ssl if we need
RUN apk add --no-cache bash ca-certificates
WORKDIR /app

# Copying the app and the configuration file
COPY --from=build-env /golden-headed-quetzal/golden-headed-quetzal /app/
COPY --from=build-env /golden-headed-quetzal/config.toml /app/

# Wait-for-it is a shell script to wait for the start of the postgres server
COPY --from=build-env /golden-headed-quetzal/wait-for-it.sh /app/

RUN mkdir config
RUN mv config.toml config/config.toml

CMD ["./wait-for-it.sh", "db:5432", "--", "./golden-headed-quetzal"]