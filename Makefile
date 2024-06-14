run:
	docker-compose up -d 
run-local:
	go mod download
	go run main.go