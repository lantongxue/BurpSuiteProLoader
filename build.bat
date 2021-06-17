@echo off
echo Start Building...
go build -ldflags="-s -w -H windowsgui"

echo Success

pause