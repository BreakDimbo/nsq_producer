export GOPATH := $(CURDIR)
export CRAZY_PRO_ROOT := $(CURDIR)
BINARY=producer


nsq_producer:
	go build -o bin/producer nsq_producer/producer

deps:
	@echo "Install Installing dependencies"
	@go get -u github.com/golang/dep/cmd/dep
	cd src/nsq_producer; ${GOPATH}/bin/dep ensure -v

clean:
	-rm -f bin/nsq_producer_*