.PHONY: ${TARGETS}
.DEFAULT_GOAL := help

export GOPATH := $(shell pwd)/build

help:
	@echo -e 'Usage: make TARGET\n'
	@printf '\033[36m%-30s\033[0m %s\n' 'TARGET' 'DESCRIPTION'
	@printf '%80s\n' | tr ' ' -
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

clean: ## Removes garbage files
	@echo "Removing garbage files ..."
	@chmod u+w $(GOPATH) -R
	@rm -rf $(GOPATH)

build: ## Build executable
	mkdir -p $(GOPATH)
	go build -o $(GOPATH)
	chmod u+w $(GOPATH) -R
