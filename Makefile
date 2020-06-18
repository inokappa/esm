STAGE = dev
.PHONY: build clean deploy gomodgen

default: ## Display Help
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: gomodgen ## Run go build
	@export GO111MODULE=on
	@./build.sh

clean: ## Run remove some files
	rm -vrf ./bin ./vendor Gopkg.lock

gomodgen: ## Run go mod
	@chmod u+x gomod.sh
	@./gomod.sh
