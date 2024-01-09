#!/bin/bash
cd microservices/the_monkeys_gateway && CompileDaemon -build="go build -o /build/the_monkeys_gateway ." -command="/build/the_monkeys_gateway" &
cd microservices/the_monkeys_authz && CompileDaemon -build="go build -o /build/the_monkeys_authz ." -command="/build/the_monkeys_authz" &
wait
