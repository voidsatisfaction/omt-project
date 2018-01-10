if [ ${APP_ENV="PROD"} = "PROD" ]; then
  echo "hello, PROD"
  go run server.go
else
  echo "hello, DEV"
  gin -p 9000 -a 9001 server.go
fi
