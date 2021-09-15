# ./Dockerfile

FROM golang:1.16-alpine AS builder-env

RUN apk add --no-cache mysql-client

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image 
# and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64
RUN go build -ldflags="-s -w" -o go-fiber-api .

# for building minimal docker image
FROM scratch

# Copy binary and config files from /build 
# to root folder of scratch container.
COPY --from=builder-env ["/build/go-fiber-api", "/build/.env", "/"]

# Export necessary port.
EXPOSE 8080

# Command to run when starting the container.
ENTRYPOINT ["/go-fiber-api"]