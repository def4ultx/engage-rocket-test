default: run

.PHONY: run
run: 
	go mod vendor
	go run main.go

.PHONY: build
build: 
	CGO_ENABLED=0 go build -ldflags="-s -w" -o server

.PHONY: test
test:
	go test -mod vendor -race -count=1 ./...

.PHONY: integration
integration:
	docker build -t engagerocket/score-server:latest .
	docker-compose up -d
	go test -v -tags=integration -count=1 ./integration/...
	docker-compose down