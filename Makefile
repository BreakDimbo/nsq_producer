export GOPATH := $(CURDIR)
export CRAZY_PRO_ROOT := $(CURDIR)
BINARY=producer

nsq_producer:
	go build -o bin/producer nsq_stress/nsq_producer

nsq_consumer:
	go build -o bin/producer nsq_stress/nsq_consumer

deps:
	@echo "Install Installing dependencies"
	@go get -u github.com/golang/dep/cmd/dep
	cd src/nsq_stress; ${GOPATH}/bin/dep init; ${GOPATH}/bin/dep ensure -v

clean:
	-rm -f bin/nsq_producer_*