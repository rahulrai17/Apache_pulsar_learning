@echo off

start cmd /k "cd hello-service && go run main.go"
start cmd /k "cd world-service && go run main.go"
start cmd /k "cd whatsup-service && go run main.go"
start cmd /k "cd web-service && go run main.go"
