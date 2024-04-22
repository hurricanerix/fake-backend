.PHONY: default run build release mkdirs clean 

OS=`uname | tr '[:upper:]' '[:lower:]'`
ARCH=`arch`
CMD=fb
BUILD_DIR=bin
HOME_DIR=`echo $(HOME)`

default: run

run:
	go run cmd/fb/main.go

build: clean mkdirs
	GOOS=linux GOARCH=amd64 go build -o "$(BUILD_DIR)/$(CMD)-linux-amd64" cmd/fb/main.go
	GOOS=windows GOARCH=amd64 go build -o "$(BUILD_DIR)/$(CMD)-windows-amd64.exe" -buildmode=exe cmd/fb/main.go
	GOOS=darwin GOARCH=amd64 go build -o "$(BUILD_DIR)/$(CMD)-darwin-amd64" cmd/fb/main.go
	GOOS=darwin GOARCH=arm64 go build -o "$(BUILD_DIR)/$(CMD)-darwin-arm64" cmd/fb/main.go

release: build
	@tar czf "$(BUILD_DIR)/$(CMD)-linux-amd64.tar.gz" --directory="$(BUILD_DIR)" "$(CMD)-linux-amd64"
	@cd $(BUILD_DIR); shasum -a 256  "$(CMD)-linux-amd64.tar.gz"
	@zip -q -r -j "$(BUILD_DIR)/$(CMD)-windows-amd64.zip" "$(BUILD_DIR)/$(CMD)-windows-amd64.exe"
	@cd $(BUILD_DIR); shasum -a 256 "$(CMD)-windows-amd64.zip"
	@tar czf "$(BUILD_DIR)/$(CMD)-darwin-amd64.tar.gz" --directory="$(BUILD_DIR)" "$(CMD)-darwin-amd64"
	@cd $(BUILD_DIR); shasum -a 256 "$(CMD)-darwin-amd64.tar.gz"
	@tar czf "$(BUILD_DIR)/$(CMD)-darwin-arm64.tar.gz" --directory="$(BUILD_DIR)" "$(CMD)-darwin-arm64"
	@cd $(BUILD_DIR); shasum -a 256 "$(CMD)-darwin-arm64.tar.gz"

mkdirs:
	mkdir -p $(BUILD_DIR)/

clean: 
	rm -f $(BUILD_DIR)/fb-linux-amd64
	rm -f $(BUILD_DIR)/fb-windows-amd64.exe
	rm -f $(BUILD_DIR)/fb-darwin-amd64
	rm -f $(BUILD_DIR)/fb-darwin-arm64
	rm -f $(BUILD_DIR)/fb-linux-amd64.tar.gz
	rm -f $(BUILD_DIR)/fb-windows-amd64.zip
	rm -f $(BUILD_DIR)/fb-darwin-amd64.tar.gz
	rm -f $(BUILD_DIR)/fb-darwin-arm64.tar.gz
