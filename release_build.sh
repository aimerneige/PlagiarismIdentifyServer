export GIN_MODE=release
export GOARCH=amd64
export GOOS=linux
go build -o ./release/PlagiarismIdentidyServer
export GOOS=windows
go build -o ./release/PlagiarismIdentidyServer.exe
