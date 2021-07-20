env GOOS=windows GOARCH=amd64 GIN_MODE=release go build -ldflags "-s -w" -o build/windows/amd64/
env GOOS=linux GOARCH=amd64 GIN_MODE=release go build -ldflags "-s -w" -o build/linux/amd64/