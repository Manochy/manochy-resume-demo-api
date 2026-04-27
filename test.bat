@echo off

echo ===== LOCAL TEST =====

cd manochy-api

set GOOS=linux
set GOARCH=amd64

go build -o bootstrap main.go

cd ..

sam local start-api -p 3000