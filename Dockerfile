FROM golang:1.24-alpine

WORKDIR /app

# Install air for hot reloading
RUN go install github.com/air-verse/air@v1.61.0

# Install make for the build command
RUN apk add --no-cache make

COPY go.mod ./
# RUN go mod download # Standard practice, but if go.mod is empty/new, might fail. 
# Let's check go.mod first.

COPY . .

CMD ["air", "-c", ".air.toml"]
