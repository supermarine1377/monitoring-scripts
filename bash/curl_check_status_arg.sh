#!/bin/bash
# script to execute cURL request to a specified URL every minute and check HTTP status code

usage() {
  echo "Usage: $0 <URL>"
  exit 1
}

if [ -z "$1" ]; then
  usage
fi

URL="$1"

while true; do
  HTTP_STATUS_CODE=$(curl -s -o /dev/null -w "%{http_code}" "$URL")

  if (( HTTP_STATUS_CODE >= 500 )); then
    echo "Error: HTTP status code $HTTP_STATUS_CODE"
  elif (( HTTP_STATUS_CODE == 200 )); then
    echo "OK"
  else
    echo "Other HTTP status code: $HTTP_STATUS_CODE"
  fi

  sleep 60
done