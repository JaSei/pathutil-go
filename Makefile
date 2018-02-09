HELP?=$$(go run main.go --help 2>&1)
VERSION?=$$(cat VERSION)
GO18?=$(shell go version | grep -E "go1\.[89]")
DEP?=$$(which dep)

ifeq ($(OS),Windows_NT)
	DEP_VERS=dep-windows-amd64
else ifeq ($(OS), Darwin)
	DEP_VERS=dep-darwin-amd64
else
	DEP_VERS=dep-linux-amd64
endif

setup: ## Install all the build and lint dependencies
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/robertkrimen/godocdown/godocdown
	@if [ "$(DEP)" = "" ]; then\
		curl -L https://github.com/golang/dep/releases/download/v0.4.1/$(DEP_VERS) >| $$GOPATH/bin/dep;\
		chmod +x $$GOPATH/bin/dep;\
	fi
	dep ensure
ifeq ($(GO18),) 
	@echo no install metalinter, because metalinter need go1.8+
else
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
endif

generate: ## Generate README.md
	godocdown >| README.md

test: generate ## Run all the tests
	echo 'mode: atomic' > coverage.txt && go list ./... | grep -v vendor | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp

cover: test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

lint: ## Run all the linters
ifeq ($(GO18),) 
	@echo no run metalinter, because metalinter need go1.8+
else
	#https://github.com/golang/go/issues/19490
	#--enable=vetshadow \

	gometalinter --vendor --disable-all \
		--enable=deadcode \
		--enable=ineffassign \
		--enable=gosimple \
		--enable=staticcheck \
		--enable=gofmt \
		--enable=goimports \
		--enable=dupl \
		--enable=misspell \
		--enable=errcheck \
		--enable=vet \
		--deadline=10m \
		--enable=vetshadow \
		./...

endif

ci: test lint  ## Run all the tests and code checks

build: ## Build the app
	go build

release: ## Release new version
	git tag | grep -q $(VERSION) && echo This version was released! Increase VERSION! || git tag $(VERSION) && git push origin $(VERSION)

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
