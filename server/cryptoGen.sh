#!/bin/bash

if [ -d ./ssl ]; then
    echo Already exists
else
    mkdir ssl
    openssl genrsa -out ./ssl/server.key 2048
    openssl req -new -x509 -key ./ssl/server.key -out ./ssl/server.pem -days 365
fi