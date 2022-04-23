@echo off
echo Start Building...

go-winres make

go build -ldflags="-H windowsgui -s -w" -buildvcs=false

echo Success