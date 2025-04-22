APP_NAME=middleware-client
APP_PATH=./app/
GIT_COMMIT=$(shell git rev-parse --short HEAD)
BUILD_TIME=$(shell TZ=Asia/Shanghai date +'%Y-%m-%d.%H:%M:%S%Z')
BUILD_FLAGS=-ldflags "-X 'github.com/memoio/middleware-client/cmd.BuildFlag=$(GIT_COMMIT)+$(BUILD_TIME)'"

all: clean  build

clean:
	rm -f ${APP_NAME}

build:
	go build $(BUILD_FLAGS) -o ${APP_PATH}${APP_NAME}

	
.PHONY: all clean build