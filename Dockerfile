FROM golang:1.19-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Copy the .env file
COPY .env ./.env

# Install dotenv package
RUN go get github.com/joho/godotenv/cmd/godotenv

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /openai ./cmd/web/main.go

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/openai"]
