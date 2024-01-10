#!/bin/bash

# Start the first service in the background
/build/the_monkeys_gateway &

# Start the second service in the foreground
# This will cause the script to block until this service exits
/build/the_monkeys_authz &

wait