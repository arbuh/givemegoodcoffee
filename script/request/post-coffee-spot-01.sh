#!/bin/bash

# Default host
HOST="${1:-http://localhost:8080}"

JSON_DATA='{
  "name": "Blue Bottle Coffee",
  "type": "cafe",
  "location": {
    "countryCode": "NL",
    "address": "Nijnte pleintje 7, Utrecht"
  }
}'

curl -X POST "$HOST/api/v1/coffeespot" \
  -H "Content-Type: application/json" \
  -d "$JSON_DATA" \
  -v
