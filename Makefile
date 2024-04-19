all: dir build publish

CONTAINER_TAG := $(shell grep -Eoh '\[\d.\d.\d+\]' CHANGELOG.md | head -1 | tr -d '[]\n')
REPOSITORY := log-generator

dir:
	mkdir -p bin/

build:
	docker pull $(REPOSITORY):$(CONTAINER_TAG) ||\
	docker build -t $(REPOSITORY):$(CONTAINER_TAG) .

test:
	GO111MODULE=off go test -v

publish:
	docker pull $(REPOSITORY):$(CONTAINER_TAG) ||\
	docker push $(REPOSITORY):$(CONTAINER_TAG)
