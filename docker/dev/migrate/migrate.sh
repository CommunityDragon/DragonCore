#!/bin/sh
migrate -path=/migrations -database postgres://postgres:postgres@db/app?sslmode=disable $@
