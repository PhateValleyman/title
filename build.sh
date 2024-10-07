
gcc ./title.c -o ./bin/title

GOOS=android GOARCH=arm64 go build -trimpath -v -x -o ./bin/title_aarch64 ./title.go
GOOS=linux GOARCH=arm GOARM=7 go build -trimpath -v -x -o ./bin/title_armv7 ./title.go
GOOS=linux GOARCH=arm GOARM=5 go build -trimpath -v -x -o ./bin/title_armv5 ./title.go
