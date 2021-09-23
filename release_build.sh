export GIN_MODE=release
export GOARCH=amd64
export GOOS=linux
go build -o ./release/restfultemplate
export GOOS=windows
go build -o ./release/restfultemplate.exe
