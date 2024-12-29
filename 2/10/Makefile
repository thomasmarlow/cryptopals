.PHONY: build
build:
	docker build --tag cryptopals-2-`basename $(shell pwd)` .

.PHONY: run
run:
	docker run cryptopals-2-`basename $(shell pwd)`:latest