.PHONY: build run stop start-docker stop-docker

build:
	cd wbl && go build -o wbl.exe ./cmd/main.go

send-message:
	cd cmd/sender && go run sender.go

clear-cache:
	docker-compose exec redis redis-cli FLUSHALL
	
run:
	docker-compose up -d
	cd cmd && go run .

start-docker:
	docker-compose up -d

stop-docker:
	docker-compose down -v
