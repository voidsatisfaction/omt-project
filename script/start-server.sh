#!/bin/sh
if [ ${APP_ENV="PROD"} = "PROD" ]; then
  echo "hello, PROD"
  go run src/server.go
else
  echo "hello, DEV"
  cd src
  gin -p 9000 -a 9001 src/server.go
fi
