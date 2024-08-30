.PHONY: build args

build:
	@rm docker-short && go build -o docker-short

push:
	@sh scripts/push.sh "$M"