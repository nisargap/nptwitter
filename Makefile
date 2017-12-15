BUILD_PATH=$(GOPATH)/src/github.com/nisargap/nptwitter
all:
	go build -o bin/main $(BUILD_PATH)/main.go

