#!/bin/sh
migrate -path=/migrations -database "postgres://${DB_USER}:${DB_PASS}@${DB_HOST}/${DB_NAME}?sslmode=${DB_SSL}" $@
