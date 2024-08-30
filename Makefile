.PHONY: build args install

build:
	@rm dcshort && go build -o dcshort

push:
	@sh scripts/push.sh "$M"

install:
	@chmod +x scripts/install.sh && sh scripts/install.sh