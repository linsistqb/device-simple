.PHONY: build test clean docker

GO=CGO_ENABLED=0 GO111MODULE=on go

MICROSERVICES=cmd/device-dcs/device-dcs
.PHONY: $(MICROSERVICES)

VERSION=$(shell cat ./VERSION)

GOFLAGS=-ldflags "-X github.com/edgexfoundry/device-dcs.Version=$(VERSION)"

GIT_SHA=$(shell git rev-parse HEAD)

build: $(MICROSERVICES)
	$(GO) build ./...

cmd/device-dcs/device-dcs:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/device-dcs

docker:
	docker build \
		-f cmd/device-dcs/Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t edgexfoundry/docker-device-dcs:$(GIT_SHA) \
		-t edgexfoundry/docker-device-dcs:$(VERSION)-dev \
		.

test:
	$(GO) test ./... -cover

clean:
	rm -f $(MICROSERVICES)
