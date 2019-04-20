#!/usr/bin/env bash

curl -X POST \
  http://localhost:8080/ourell/bbs \
  -H 'Content-Type: application/json' \
  -d '{
    "number1": "+380991926482",
    "number2": "+380991926452",
    "from": "2019/03/22-15:50:20",
    "to": "2019/03/26-15:50:20"
}' -v
