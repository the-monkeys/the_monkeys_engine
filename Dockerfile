FROM golang:latest
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0
WORKDIR /app
RUN mkdir -p /build/config/
COPY . .
RUN chmod +x startup.sh

# Create the directory


# Copy the file
COPY ./config/config.yml /build/config/config.yml

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
ENTRYPOINT ["./startup.sh"]
