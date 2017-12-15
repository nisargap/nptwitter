BUILD_PATH=$(GOPATH)/src/github.com/nisargap/nptwitter
all_osx:
	go build -o bin/main $(BUILD_PATH)/main.go
all_linux:
	GOOS=linux GOARCH=amd64 go build -o bin/main_linux $(BUILD_PATH)/main.go

