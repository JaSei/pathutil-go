HELP?=$$(go run main.go --help 2>&1)
VERSION?=$$(cat VERSION)
LINTER?=$$(which golangci-lint)
LINTER_VERSION=1.50.0

ifeq ($(OS),Windows_NT)
	LINTER_FILE=golangci-lint-$(LINTER_VERSION)-windows-amd64.zip
	LINTER_UNPACK= >| app.zip; unzip -j app.zip -d $$GOPATH/bin; rm app.zip
else ifeq ($(OS), Darwin)
	LINTER_FILE=golangci-lint-$(LINTER_VERSION)-darwin-amd64.tar.gz
	LINTER_UNPACK= | tar xzf - -C $$GOPATH/bin --wildcards --strip 1 "**/golangci-lint"
else
	LINTER_FILE=golangci-lint-$(LINTER_VERSION)-linux-amd64.tar.gz
	LINTER_UNPACK= | tar xzf - -C $$GOPATH/bin --wildcards --strip 1 "**/golangci-lint"
endif

setup: ## Install all the build and lint dependencies
	go install golang.org/x/tools/cmd/cover@latest
	go install github.com/robertkrimen/godocdown/godocdown@latest

	go mod download


generate: ## Generate README.md
	godocdown >| README.md

test: generate test_only ## Run all the tests

test_only: ## Run tests only (skip generate readme)
	echo 'mode: atomic' > coverage.txt && go list ./... | grep -v vendor | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp

cover: test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

lint: ## Run all the linters
	@if [ "$(LINTER)" = "" ]; then\
		curl -L https://github.com/golangci/golangci-lint/releases/download/v$(LINTER_VERSION)/$(LINTER_FILE) $(LINTER_UNPACK) ;\
		chmod +x $$GOPATH/bin/golangci-lint;\
	fi

	golangci-lint run

ci: test_only ## Run all the tests but no linters (use https://golangci.com integration instead) and no generate

build: ## Build the app
	go build

release: ## Release new version
	git tag | grep -q $(VERSION) && echo This version was released! Increase VERSION! || git tag $(VERSION) && git push origin $(VERSION)

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
