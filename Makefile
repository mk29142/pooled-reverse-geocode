
ginkgo := go run github.com/onsi/ginkgo/ginkgo -r --randomizeAllSpecs --failOnPending

build:
	go build -o .

init:
	go mod download

test-unit:
	$(ginkgo) --skipPackage benchmark

test-benchmark:
	$(ginkgo) ./workpool/benchmark
