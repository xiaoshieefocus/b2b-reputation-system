.PHONY: all dev clean build env-up env-down run erase

all: clean build env-up run

dev: build run

##### BUILD
build:
	@echo "Build ..."
	# @dep ensure
	@go build
	@echo "Build done"

##### ENV
env-up:
	@echo "Start environment ..."
	@cd fixtures && docker-compose -f docker-compose-cli.yaml up --force-recreate -d
	@echo "Environment up"

env-down:
	@echo "Stop environment ..."
	@cd fixtures && docker-compose -f docker-compose-cli.yaml down --volumes --remove-orphans
	@echo "Environment down"

##### RUN
run:
	@echo "Start app ..."
	@./b2b-reputation-system

##### CLEAN
clean: env-down
	@echo "Clean up ..."
	@rm -rf /tmp/b2b-reputation-system-* b2b-reputation-system
	@docker rm -f -v `docker ps -a --no-trunc | grep "bom2buy" | cut -d ' ' -f 1` 2>/dev/null || true
	@docker rmi `docker images --no-trunc | grep "bom2buy" | cut -d ' ' -f 1` 2>/dev/null || true
	@echo "Clean up done"

