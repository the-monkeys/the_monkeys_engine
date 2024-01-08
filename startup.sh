#!/bin/bash
cd microservices/the_monkeys_gateway && CompileDaemon -build="go build -o /build/the_monkeys_gateway ." -command="/build/the_monkeys_gateway" &
cd microservices/auth_service/cmd && CompileDaemon -build="go build -o /build/service2 ." -command="/build/service2" &
wait
