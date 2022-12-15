IMAGE ?= uyinn28/server-opentelemetry

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server-linux-amd64 .
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o server-linux-arm64 .


docker:
	docker buildx build --push \
		--platform=linux/amd64,linux/arm64 \
		-t $(IMAGE) .
