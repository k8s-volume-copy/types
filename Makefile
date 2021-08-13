QUAY_USERNAME ?= k8s-volume-copy
LATEST_TAG ?= ci
IMAGE_TAG ?= $(shell git rev-parse --short HEAD)

.PHONY: crd-gen
crd-gen:
	rm -rf config/crd
	controller-gen crd:crdVersions=v1 paths=./apis/demo.io/v1
	controller-gen object paths=./apis/demo.io/v1

.PHONY: vendor
vendor:
	@GO111MODULE=on go mod tidy
	@GO111MODULE=on go mod vendor

