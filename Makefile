export GOPATH := $(CURDIR)
export CRAZY_PRO_ROOT := $(CURDIR)
BINARY=producer

nsq_producer:
	go build -o bin/producer stress_suit/nsq/pro

nsq_consumer:
	go build -o bin/consumer stress_suit/nsq/con

redis_pub:
	go build -o bin/redis_pub stress_suit/redis/pub

redis_sub:
	go build -o bin/redis_sub stress_suit/redis/sub

deps:
	@echo "Install Installing dependencies"
	@go get -u github.com/golang/dep/cmd/dep
	cd src/stress_suit; ${GOPATH}/bin/dep init; ${GOPATH}/bin/dep ensure -v

clean:
	-rm -f bin/nsq_producer_*