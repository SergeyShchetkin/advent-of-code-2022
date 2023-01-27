# Usage: make help

## Build and start docker containers, init project
default: build

#############################################
#				Docker-compose				#
#############################################

## Build docker containers
build:
	docker-compose build

run-task:
	docker-compose run go /bin/sh -c 'cd $(DIR)/ && go run $(TASK).go'

clear-data:
	docker-compose rm -f go

#############################################
#				Internals					#
#############################################

# hash for dynamic arguments. tab must be before @
%:
	@:

# Colors
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=50
## Help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Available targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 1, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)