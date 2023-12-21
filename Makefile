BIN := fulltexas

GOPATH := $(shell go env GOPATH)

ifeq ($(OS),Windows_NT)
    GO_FILES := $(shell dir /S /B *.go)
    GO_DEPS := $(shell dir /S /B go.mod go.sum)
    CLEAN := del
else
    GO_FILES := $(shell find . -name '*.go')
    GO_DEPS := $(shell find . -name go.mod -o -name go.sum)
    CLEAN := rm -f
endif

$(BIN): $(GO_FILES) $(GO_DEPS) images/turtle1.png images/turtle2.png
	gofumpt -w $(GO_FILES)
	golangci-lint run
	go build -o $(BIN) cmd/main.go

.PHONY: test
test: $(BIN)
	./$(BIN) --log-level=debug

images/turtle1.png:
	mkdir -p images/
	curl -Lo images/turtle1.png 'https://github.com/eliben/code-for-blog/blob/master/2023/go-google-ai-gemini/images/turtle1.png?raw=true'

images/turtle2.png:
	mkdir -p images/
	curl -Lo images/turtle2.png 'https://github.com/eliben/code-for-blog/blob/master/2023/go-google-ai-gemini/images/turtle2.png?raw=true'

.PHONY: install
install: $(BIN)
	mv $(BIN) $(GOPATH)/bin/$(BIN)

.PHONY: clean
clean:
	$(CLEAN) $(BIN)
	$(CLEAN) -r images
