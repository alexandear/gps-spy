#!/usr/bin/env bash

curl -X POST \
  http://localhost:8080/ourell/bbinput \
  -H 'Content-Type: application/json' \
  -d '{
    "number": "+380991926482",
    "ip": "35.25.21.123",
    "imei": "502507345219189",
    "timestamp": "2019/03/22-15:50:20",
    "coordinates": {
        "longitude": 22.1832284135991,
        "latitude": 60.4538416572538
    }
}' -v

