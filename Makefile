GO_IMAGE=golang:1.17-alpine3.16

docker-compose:
	@which docker-compose > /dev/null || echo "Please install docker-compose as described at https://docs.docker.com/compose/install/"

up: docker-compose
	docker-compose up -d

down: docker-compose
	docker-compose down

test:
	docker run --rm -v $(PWD):/app -w /app -e CGO_ENABLED=0 -e GOOS=linux -it ${GO_IMAGE} go test -v ./... -coverprofile cover.out