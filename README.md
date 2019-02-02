
# Get Build Lambda Zip

```
go -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
```

# Build on Windows


```
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
go build -o main main.go
~\Go\Bin\build-lambda-zip.exe -o main.zip main
```
