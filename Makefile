NAME=square
VERSION=0.1
QUALITY_THRESHOLD=80

clean:
	rm -f build/square.out
	rm -f coverage.out
	go clean -cache -testcache
	docker rmi -f ${NAME}:${VERSION} || true

compile:
	CGO_ENABLED=0 go build -o build/square.out cmd/main.go

localbuild: compile
	docker build -t ${NAME}:${VERSION} build

run:
	docker stop ${NAME} || true
	docker run --rm -p 8530:3000 --name ${NAME} ${NAME}:${VERSION}

stagedbuild:
	docker build -t ${NAME}:${VERSION} .

test:
	go test ./...

checkcoverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func coverage.out | awk -F'\t' -v threshold=${QUALITY_THRESHOLD} '/^total:/{print $$0; overall_percent=$$NF; if (overall_percent >= threshold) {exit 0} else {exit 1}}'
