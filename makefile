.PHONY: install
build:
	go install
	mkdir -p $(HOME)/go/bin/scripts
	cp scripts/* $(HOME)/go/bin/scripts