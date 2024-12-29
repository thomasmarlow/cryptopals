.PHONY: build
build:
	docker build --tag cryptopals-1-`basename $(shell pwd)` .

.PHONY: run
run:
	docker run cryptopals-1-`basename $(shell pwd)`:latest