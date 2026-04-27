@echo off

echo ===== DEPLOY AWS =====

cd manochy-api

set GOOS=linux
set GOARCH=amd64

go build -o bootstrap main.go

cd ..

sam deploy --no-build --no-confirm-changeset