SHELL := /bin/bash

ROOT_DIR:= $(patsubst %/,%,$(dir $(realpath $(lastword $(MAKEFILE_LIST)))))
include build/check_de***REMOVED***ned.mk

# Package
GO_PKG := github.com/operator-framework/operator-metering
REPORTING_OPERATOR_PKG := $(GO_PKG)/cmd/reporting-operator

DOCKER_BUILD_ARGS ?=
DOCKER_CACHE_FROM_ENABLED =
ifdef BRANCH_TAG
ifeq ($(BRANCH_TAG_CACHE), true)
	DOCKER_CACHE_FROM_ENABLED=true
endif
ifdef DOCKER_CACHE_FROM_ENABLED
	DOCKER_BUILD_ARGS += --cache-from $(IMAGE_NAME):$(BRANCH_TAG)
endif
endif

GO_BUILD_ARGS := -ldflags '-extldflags "-static"'
GOOS = "linux"
CGO_ENABLED = 0

REPORTING_OP_BIN_OUT = images/reporting-operator/bin/reporting-operator

DOCKER_BASE_URL := quay.io/coreos

METERING_HELM_OPERATOR_IMAGE := $(DOCKER_BASE_URL)/chargeback-helm-operator
REPORTING_OPERATOR_IMAGE := $(DOCKER_BASE_URL)/chargeback
HELM_OPERATOR_IMAGE := $(DOCKER_BASE_URL)/helm-operator
HADOOP_IMAGE := $(DOCKER_BASE_URL)/chargeback-hadoop
HIVE_IMAGE := $(DOCKER_BASE_URL)/chargeback-hive
PRESTO_IMAGE := $(DOCKER_BASE_URL)/chargeback-presto
METERING_INTEGRATION_TEST_IMAGE := $(DOCKER_BASE_URL)/chargeback-integration-tests

GIT_SHA    := $(shell git rev-parse HEAD)
GIT_TAG    := $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
RELEASE_TAG := $(shell hack/print-version.sh)

PULL_TAG_IMAGE_SOURCE ?= false
USE_LATEST_TAG ?= false
USE_RELEASE_TAG = true
PUSH_RELEASE_TAG = false

DOCKER_BUILD_CONTEXT = $(dir $(DOCKERFILE))
IMAGE_TAG = $(GIT_SHA)
TAG_IMAGE_SOURCE = $(IMAGE_NAME):$(GIT_SHA)

# Hive Git repository for Thrift de***REMOVED***nitions
HIVE_REPO := "git://git.apache.org/hive.git"
HIVE_SHA := "1fe8db618a7bbc09e041844021a2711c89355995"

JQ_DEP_SCRIPT = '.Deps[] | select(. | contains("$(GO_PKG)"))'
REPORTING_OPERATOR_GO_FILES := $(shell go list -json $(REPORTING_OPERATOR_PKG) | jq $(JQ_DEP_SCRIPT) -r | xargs -I{} ***REMOVED***nd $(GOPATH)/src/$(REPORTING_OPERATOR_PKG) $(GOPATH)/src/{} -type f -name '*.go' | sort | uniq)

CODEGEN_SOURCE_GO_FILES := $(shell $(ROOT_DIR)/hack/codegen_source_***REMOVED***les.sh)

CODEGEN_OUTPUT_GO_FILES := $(shell $(ROOT_DIR)/hack/codegen_output_***REMOVED***les.sh)

# TODO: Add tests
all: fmt test docker-build-all

# Usage:
#	make docker-build DOCKERFILE= IMAGE_NAME=

docker-build:
ifdef DOCKER_CACHE_FROM_ENABLED
	docker pull $(IMAGE_NAME):$(BRANCH_TAG) || true
endif
	docker build $(DOCKER_BUILD_ARGS) -t $(IMAGE_NAME):$(GIT_SHA) -f $(DOCKERFILE) $(DOCKER_BUILD_CONTEXT)
ifdef BRANCH_TAG
	$(MAKE) docker-tag IMAGE_NAME=$(IMAGE_NAME) IMAGE_TAG=$(BRANCH_TAG)
endif
ifdef DEPLOY_TAG
	$(MAKE) docker-tag IMAGE_NAME=$(IMAGE_NAME) IMAGE_TAG=$(DEPLOY_TAG)
endif
ifneq ($(GIT_TAG),)
	$(MAKE) docker-tag IMAGE_NAME=$(IMAGE_NAME) IMAGE_TAG=$(GIT_TAG)
endif
ifeq ($(USE_RELEASE_TAG), true)
	$(MAKE) docker-tag IMAGE_NAME=$(IMAGE_NAME) IMAGE_TAG=$(RELEASE_TAG)
endif
ifeq ($(USE_LATEST_TAG), true)
	$(MAKE) docker-tag IMAGE_NAME=$(IMAGE_NAME) IMAGE_TAG=latest
endif

# Usage:
#	make docker-tag SOURCE_IMAGE=$(IMAGE_NAME):$(GIT_SHA) IMAGE_NAME= IMAGE_TAG=
docker-tag:
ifeq ($(PULL_TAG_IMAGE_SOURCE), true)
	$(MAKE) docker-pull IMAGE=$(TAG_IMAGE_SOURCE)
endif
	docker tag $(TAG_IMAGE_SOURCE) $(IMAGE_NAME):$(IMAGE_TAG)

# Usage:
#	make docker-pull IMAGE=

docker-pull:
	docker pull $(IMAGE)

# Usage:
#	make docker-push IMAGE_NAME= IMAGE_TAG=

docker-push:
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
ifeq ($(PUSH_RELEASE_TAG), true)
	docker push $(IMAGE_NAME):$(RELEASE_TAG)
endif
ifeq ($(USE_LATEST_TAG), true)
	docker push $(IMAGE_NAME):latest
endif
ifneq ($(GIT_TAG),)
	docker push $(IMAGE_NAME):$(GIT_TAG)
endif
ifdef BRANCH_TAG
	docker push $(IMAGE_NAME):$(BRANCH_TAG)
endif
ifdef DEPLOY_TAG
	docker push $(IMAGE_NAME):$(DEPLOY_TAG)
endif

DOCKER_TARGETS := \
	reporting-operator \
	metering-integration-tests \
	hadoop \
	hive \
	presto \
	helm-operator \
	metering-helm-operator
# These generate new make targets like metering-helm-operator-docker-build
# which can be invoked.
DOCKER_BUILD_TARGETS := $(addsuf***REMOVED***x -docker-build, $(DOCKER_TARGETS))
DOCKER_PUSH_TARGETS := $(addsuf***REMOVED***x -docker-push, $(DOCKER_TARGETS))
DOCKER_TAG_TARGETS := $(addsuf***REMOVED***x -docker-tag, $(DOCKER_TARGETS))
DOCKER_PULL_TARGETS := $(addsuf***REMOVED***x -docker-pull, $(DOCKER_TARGETS))

# The steps below run for each value of $(DOCKER_TARGETS) effectively, generating multiple Make targets.
# To make it easier to follow, each step will include an example after the evaluation.
# The example will be using the metering-helm-operator targets as it's example.
#
# The pattern/string manipulation below does the following (starting from the inner most expression):
# 1) strips -docker-push, -docker-tag, or -docker-pull from the target name ($@) giving us the non suf***REMOVED***xed value from $(TARGETS)
# ex: metering-helm-operator-docker-build -> metering-helm-operator
# 2) Replaces - with _
# ex: metering-helm-operator -> metering_helm_operator
# 3) Uppercases letters
# ex: metering_helm_operator -> METERING_HELM_OPERATOR
# 4) Appends _IMAGE
# ex: METERING_HELM_OPERATOR -> METERING_HELM_OPERATOR_IMAGE
# That gives us the value for the docker-build, docker-tag, or docker-push IMAGE_NAME variable.

$(DOCKER_PUSH_TARGETS)::
	$(MAKE) docker-push IMAGE_NAME=$($(addsuf***REMOVED***x _IMAGE, $(shell echo $(subst -,_,$(subst -docker-push,,$@)) | tr a-z A-Z)))

$(DOCKER_TAG_TARGETS)::
	$(MAKE) docker-tag IMAGE_NAME=$($(addsuf***REMOVED***x _IMAGE, $(shell echo $(subst -,_,$(subst -docker-tag,,$@)) | tr a-z A-Z)))

$(DOCKER_PULL_TARGETS)::
	$(MAKE) docker-pull IMAGE_NAME=$($(addsuf***REMOVED***x _IMAGE, $(shell echo $(subst -,_,$(subst -docker-pull,,$@)) | tr a-z A-Z)))

docker-build-all: $(DOCKER_BUILD_TARGETS)

docker-push-all: $(DOCKER_PUSH_TARGETS)

docker-tag-all: $(DOCKER_TAG_TARGETS)

docker-pull-all: $(DOCKER_PULL_TARGETS)

reporting-operator-docker-build: images/reporting-operator/Docker***REMOVED***le $(REPORTING_OP_BIN_OUT)
	$(MAKE) docker-build DOCKERFILE=$< IMAGE_NAME=$(REPORTING_OPERATOR_IMAGE)

metering-integration-tests-docker-build: images/integration-tests/Docker***REMOVED***le
	$(MAKE) docker-build DOCKERFILE=$< IMAGE_NAME=$(METERING_INTEGRATION_TEST_IMAGE) DOCKER_BUILD_CONTEXT=$(ROOT_DIR)

metering-helm-operator-docker-build: \
		images/metering-helm-operator/Docker***REMOVED***le \
		helm-operator-docker-build \
		images/metering-helm-operator/tectonic-metering-0.1.0.tgz \
		images/metering-helm-operator/openshift-metering-0.1.0.tgz \
		images/metering-helm-operator/operator-metering-0.1.0.tgz \
		images/metering-helm-operator/metering-override-values.yaml
	$(MAKE) docker-build DOCKERFILE=$< IMAGE_NAME=$(METERING_HELM_OPERATOR_IMAGE)

helm-operator-docker-build: images/helm-operator/Docker***REMOVED***le
	$(MAKE) docker-build DOCKERFILE=$< IMAGE_NAME=$(HELM_OPERATOR_IMAGE) USE_LATEST_TAG=true

presto-docker-build: images/presto/Docker***REMOVED***le
	$(MAKE) docker-build DOCKERFILE=$< IMAGE_NAME=$(PRESTO_IMAGE)

hadoop-docker-build: images/hadoop/Docker***REMOVED***le
	$(MAKE) docker-build DOCKERFILE=$< IMAGE_NAME=$(HADOOP_IMAGE) USE_LATEST_TAG=true

hive-docker-build: images/hive/Docker***REMOVED***le hadoop-docker-build
	$(MAKE) docker-build DOCKERFILE=$< IMAGE_NAME=$(HIVE_IMAGE)

# Update dependencies
vendor: Gopkg.toml
	dep ensure -v

test:
	go test ./pkg/...

# Runs gofmt on all ***REMOVED***les in project except vendored source and Hive Thrift de***REMOVED***nitions
fmt:
	***REMOVED***nd . -name '*.go' -not -path "./vendor/*" -not -path "./pkg/hive/hive_thrift/*" | xargs gofmt -w

# validates no unstaged changes exist
ci-validate: verify-codegen all-charts metering-manifests fmt
	@echo Checking for unstaged changes
	git diff --stat HEAD --ignore-submodules --exit-code

reporting-operator-bin: $(REPORTING_OP_BIN_OUT)

reporting-operator-local: $(REPORTING_OPERATOR_GO_FILES)
	$(MAKE) build-reporting-operator REPORTING_OPERATOR_BIN_LOCATION=$@ GOOS=$(shell go env GOOS)

.PHONY: run-reporting-operator-local
run-reporting-operator-local:
	$(MAKE) reporting-operator-local
	./hack/run-local-with-port-forward.sh $(REPORTING_OPERATOR_ARGS)

$(REPORTING_OP_BIN_OUT): $(REPORTING_OPERATOR_GO_FILES)
	$(MAKE) build-reporting-operator REPORTING_OPERATOR_BIN_LOCATION=$@

build-reporting-operator:
	@:$(call check_de***REMOVED***ned, REPORTING_OPERATOR_BIN_LOCATION, Path to output binary location)
	$(MAKE) update-codegen
	mkdir -p $(dir $@)
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) go build $(GO_BUILD_ARGS) -o $(REPORTING_OPERATOR_BIN_LOCATION) $(REPORTING_OPERATOR_PKG)

images/metering-helm-operator/metering-override-values.yaml: ./hack/render-metering-chart-override-values.sh
	./hack/render-metering-chart-override-values.sh $(RELEASE_TAG) > $@

CHART_DEPS := images/metering-helm-operator/tectonic-metering-0.1.0.tgz \
	images/metering-helm-operator/openshift-metering-0.1.0.tgz \
	images/metering-helm-operator/operator-metering-0.1.0.tgz

all-charts: $(CHART_DEPS)

tectonic-metering-chart: images/metering-helm-operator/tectonic-metering-0.1.0.tgz

openshift-metering-chart: images/metering-helm-operator/openshift-metering-0.1.0.tgz

operator-metering-chart: images/metering-helm-operator/operator-metering-0.1.0.tgz

images/metering-helm-operator/tectonic-metering-0.1.0.tgz: images/metering-helm-operator/metering-override-values.yaml $(shell ***REMOVED***nd charts -type f)
	helm dep update --skip-refresh charts/tectonic-metering
	helm package --save=false -d images/metering-helm-operator charts/tectonic-metering

images/metering-helm-operator/openshift-metering-0.1.0.tgz: images/metering-helm-operator/metering-override-values.yaml $(shell ***REMOVED***nd charts -type f)
	helm dep update --skip-refresh charts/openshift-metering
	helm package --save=false -d images/metering-helm-operator charts/openshift-metering

images/metering-helm-operator/operator-metering-0.1.0.tgz: images/metering-helm-operator/metering-override-values.yaml $(shell ***REMOVED***nd charts -type f)
	helm dep update --skip-refresh charts/operator-metering
	helm package --save=false -d images/metering-helm-operator charts/operator-metering

metering-manifests:
	./hack/create-metering-manifests.sh $(RELEASE_TAG)

.PHONY: \
	test vendor fmt regenerate-hive-thrift thrift-gen \
	update-codegen verify-codegen \
	$(DOCKER_BUILD_TARGETS) $(DOCKER_PUSH_TARGETS) \
	$(DOCKER_TAG_TARGETS) $(DOCKER_PULL_TARGETS) \
	docker-build docker-tag docker-push \
	docker-build-all docker-tag-all docker-push-all \
	metering-integration-tests-docker-build \
	build-reporting-operator reporting-operator-bin reporting-operator-local \
	operator-metering-chart tectonic-metering-chart openshift-metering chart \
	images/metering-helm-operator/metering-override-values.yaml \
	metering-manifests bill-of-materials.json \
	install-kube-prometheus-helm

update-codegen: $(CODEGEN_OUTPUT_GO_FILES)
	./hack/update-codegen.sh

$(CODEGEN_OUTPUT_GO_FILES): $(CODEGEN_SOURCE_GO_FILES)

verify-codegen:
	./hack/verify-codegen.sh

# The results of these targets get vendored, but the targets exist for
# regenerating if needed.
regenerate-hive-thrift: pkg/hive/hive_thrift

# Download Hive git repo.
out/thrift.git:
	mkdir -p out
	git clone --single-branch --bare ${HIVE_REPO} $@

# Retrieve Hive thrift de***REMOVED***nition from git repo.
thrift/TCLIService.thrift: out/thrift.git
	mkdir -p $(dir $@)
	git -C $< show ${HIVE_SHA}:service-rpc/if/$(notdir $@) > $@

# Generate source from Hive thrift de***REMOVED***ntions and remove executable packages.
pkg/hive/hive_thrift: thrift/TCLIService.thrift thrift-gen

thrift-gen:
	thrift -gen go:package_pre***REMOVED***x=${GO_PKG}/pkg/hive,package=hive_thrift -out pkg/hive thrift/TCLIService.thrift
	for i in `go list -f '{{if eq .Name "main"}}{{ .Dir }}{{end}}' ./pkg/hive/hive_thrift/...`; do rm -rf $$i; done

bill-of-materials.json: bill-of-materials.override.json
	license-bill-of-materials --override-***REMOVED***le $(ROOT_DIR)/bill-of-materials.override.json ./... > $(ROOT_DIR)/bill-of-materials.json

kube-prometheus-helm-install:
	@echo "KUBECONFIG: $(KUBECONFIG)"
	helm ls
	helm version
	helm repo add coreos https://s3-eu-west-1.amazonaws.com/coreos-charts/stable/
	helm upgrade --install --namespace monitoring prometheus-operator coreos/prometheus-operator --wait
	# set https to false on kubelets for GKE and set the fullnameOverride for the
	# Prometheus resource so our service has a consistent name.
	helm upgrade --install --namespace monitoring kube-prometheus coreos/kube-prometheus --set 'prometheus.fullnameOverride=prometheus-k8s,exporter-kubelets.https=false' --wait
