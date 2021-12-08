SHA=$(shell git rev-parse --short HEAD)
VERSION=$(shell cat VERSION)
DIRTY=false
GO_PACKAGE=$(shell go list)
LDFLAGS=-ldflags "-w -s -X $(GO_PACKAGE)/pkg/util.GitSha=${SHA} -X $(GO_PACKAGE)/pkg/util.Version=${VERSION} -X $(GO_PACKAGE)/pkg/util.Dirty=${DIRTY}"
export GO111MODULE=on

clean: ## clean the repo
	rm happy 2>/dev/null || true
	go clean
	go clean -testcache
	rm -rf dist 2>/dev/null || true
	rm coverage.out 2>/dev/null || true
	if [ -e /tmp/happy.lock ]; then \
        rm /tmp/happy.lock; \
        fi

setup: # setup development dependencies
	export GO111MODULE=on
	curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh
.PHONY: setup

install:
	go install
.PHONY: install

test:
	CGO_ENABLED=1 go test -race ./...
.PHONY: test

lint:
	golangci-lint run -E whitespace --exclude-use-default
.PHONY: lint

generate-mocks:
	@go get github.com/golang/mock/mockgen@v1.5.0
	@go get -u github.com/aws/aws-sdk-go/...
	@rm -rf mocks/mock_*
	@cd mocks; go generate
	@go mod tidy
.PHONY: generate-mocks
