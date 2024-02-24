
gcc ./title.c -o ./build/c/title

GOOS=linux GOARCH=arm64 go build -v -x -o ./build/go/title_phone ./title.go
GOOS=linux GOARCH=arm GOARM=7 go build -v -x -o ./build/go/title_tablet ./title.go
GOOS=linux GOARCH=arm GOARM=5 go build -v -x -o ./build/go/title_server ./title.go
GOOS=windows GOARCH=amd64 go build -v -x -o ./build/go/title_windows.exe ./title.go
