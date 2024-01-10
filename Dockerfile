# Start from the latest Golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Dave Augustus <dave@themonkeys.live>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN mkdir -p /build/config
COPY ./config/config.yml /build/config/config.yml
# Build the Go app
RUN cd microservices/the_monkeys_gateway && go build -o /build/the_monkeys_gateway
# Expose port 8081 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["/build/the_monkeys_gateway"]


