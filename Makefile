# Paths
BIN = $(CURDIR)/bin
SERVER = $(CURDIR)/server

.PHONY: tools
tools:
	@mkdir -p $(BIN)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.43.0 go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/cmd/goimports@latest
	cd $(SERVER) && go generate -tags tools ./tools/...