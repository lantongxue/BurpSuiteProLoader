# BurpSuiteProLoader
BurpSuiteProLoader Source

# Dependents
### Windows
```shell
go install github.com/tc-hib/go-winres@latest
```


# Build
### Windows
```shell
go-winres make && go build -ldflags="-H windowsgui"
```

### Linux & MacOS
```shell
go build
```
