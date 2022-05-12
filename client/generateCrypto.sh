#!/bin/bash

if [ -d ./nginx/ssl ]; then 
    echo "Already exists"
else
    mkdir ./nginx/ssl

    openssl dhparam -out ./nginx/ssl/dhparam.pem 2048

    openssl req -x509 -days 10 -nodes -newkey rsa:2048 -keyout ./nginx/ssl/self.key -out ./nginx/ssl/self.crt
fi