.PHONY: build args install alias

build:
	@rm dcshort && go build -o dcshort

push:
	@sh scripts/push.sh "$M"

install:
	@chmod +x scripts/install.sh && sh scripts/install.sh

alias:
	@chmod +x scripts/alias.sh && sh scripts/alias.sh