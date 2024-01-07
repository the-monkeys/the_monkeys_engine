#!/bin/bash
cd microservices/api_gateway/cmd && CompileDaemon -build="go build -o /build/service1 ." -command="/build/service1" &
cd microservices/auth_service/cmd && CompileDaemon -build="go build -o /build/service2 ." -command="/build/service2" &
wait
