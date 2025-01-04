.PHONY: build test make-pngs lint lint-ci tidy tidy-ci

build: ./snackmgr

./snackmgr: test
	go build -o ./snackmgr

test:
	go test -cover ./api/... ./internal/... ./cmd/... # for now: ./pkg/... 

make-pngs: doc/structure.png

doc/structure.png: doc/structure.drawio
	drawio -x -f png -o $@ $<

lint:
	$(GOBIN)/golangci-lint run ./...

lint-ci:
	$(GOBIN)/golangci-lint run ./... --out-format=colored-line-number --timeout=5m

generate:
	go generate ./...

tidy:
	go mod tidy

tidy-ci:
	tidied -verbose

build-image:
	docker build -f deploy/docker/Dockerfile .