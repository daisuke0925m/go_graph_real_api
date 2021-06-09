#!/bin/bash

set -e

until mysqladmin ping -h ${DB_HOST} -P ${DB_PORT} --silent; do
  echo "waiting for mysql..."
  sleep 2
done
echo "success to connect mysql."

# if [ $GO_ENV = "development" ]; then
#   echo 'start setup test db.'
#   mysql -h ${DB_HOST} -P ${DB_PORT} -u ${DB_USER} -p${DB_PASSWORD} -e "CREATE DATABASE IF NOT EXISTS "$TEST_DB_NAME";"
#   goose -env test up
#   echo "created & migrated test db"
# fi

if [ $GO_ENV = "production" ]; then
  /app/build
elif [ $GO_ENV = "development" ]; then
  air
fi
