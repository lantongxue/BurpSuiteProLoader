@echo off
echo Start Building...
go-winres make

go build -ldflags="-H windowsgui"

echo Success

pause