GOCMD=go
GOTEST=$(GOCMD) test
LOCAL_BUILD=./scripts/build-go.sh
LATEST_BUILD=$(LOCAL_BUILD) latest
LATEST_DARWIN_BUILD=$(LATEST_BUILD) darwin
LATEST_LINUX_BUILD=$(LATEST_BUILD) linux
NEW_VERSION_BUILD=$(LOCAL_BUILD) new-version
NEW_VERSION_BUILD_DARWIN=$(NEW_VERSION_BUILD) darwin
NEW_VERSION_BUILD_LINUX=$(NEW_VERSION_BUILD) linux

.PHONY: install-test-dependencies
install-test-dependencies:
	env GO111MODULE=on go get -u github.com/client9/misspell/cmd/misspell
	env GO111MODULE=on go get -u golang.org/x/lint/golint

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: lint
lint:
	golint -set_exit_status ./...

.PHONY: check-misspell
check-misspell:
	misspell ./**/* -error
