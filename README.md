# BurpSuiteProLoader
BurpSuiteProLoader Source

# Dependents
### Windows
```shell
go install github.com/tc-hib/go-winres@latest
```


# Build
### Windows
```batch
go-winres make
go build -ldflags="-H windowsgui -s -w" -buildvcs=false
```

`-buildvcs=false` is an optional

### Linux & MacOS
```shell
go build
```
