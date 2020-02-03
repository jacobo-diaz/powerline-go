.PHONY: ${TARGETS}
.DEFAULT_GOAL := help

export GOPATH := $(shell pwd)/build

RELEASE_TAG = $(strip $(word 2,$(MAKECMDGOALS)))

help:
	@echo -e 'Usage: make TARGET [tag]\n'
	@printf '\033[36m%-30s\033[0m %s\n' 'TARGET' 'DESCRIPTION'
	@printf '%80s\n' | tr ' ' -
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

clean: ## Removes garbage files
	@echo "Removing garbage files ..."
	@chmod -R u+w $(GOPATH)
	@rm -rf $(GOPATH)

build: ## Build executable
	mkdir -p $(GOPATH)
	go build -o $(GOPATH)
	chmod -R u+w $(GOPATH)

$(RELEASE_TAG):
	@:

list: ## Shows a list of existing releases on GitHub
	@hub release

release: build ## Create a release on GitHub of the current commit in the current branch with the given tag
	@[ "$(RELEASE_TAG)" ] || ( echo "Release tag is not set"; exit 1 )
	@echo "Creating release $(RELEASE_TAG) on GitHub ..."
	@git tag -a $(RELEASE_TAG) -m "$(RELEASE_TAG)"
	@git push origin --tags
	@hub release create -a $(GOPATH)/powerline-go -m "$(RELEASE_TAG)" $(RELEASE_TAG)

unrelease: ## Removes a release on GitHub with the given tag
	@[ "$(RELEASE_TAG)" ] || ( echo "Release tag is not set"; exit 1 )
	@echo "Removing release $(RELEASE_TAG) on GitHub ..."
	@hub release delete $(RELEASE_TAG)
	@git tag -d $(RELEASE_TAG)
	@git push origin --delete $(RELEASE_TAG)