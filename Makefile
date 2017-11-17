export GOPATH := $(CURDIR)
export CRAZY_PRO_ROOT := $(CURDIR)
BINARY=producer

nsq_producer:
	go build -o bin/producer stress_suit/nsq/producer

nsq_consumer:
	go build -o bin/consumer stress_suit/nsq/consumer

deps:
	@echo "Install Installing dependencies"
	@go get -u github.com/golang/dep/cmd/dep
	cd src/stress_suit; ${GOPATH}/bin/dep init; ${GOPATH}/bin/dep ensure -v

clean:
	-rm -f bin/nsq_producer_*