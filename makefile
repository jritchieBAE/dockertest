make up: build run

build:
	@docker build -t halloserver .

run:
	@docker-compose up