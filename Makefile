.PHONY: test test-race lint lintci-deps goimports clean

test:
	go test --timeout 1m -shuffle=on -cover -coverprofile=cover.out ./...

test-race:
	go test --timeout 1m -race -shuffle=on ./...

lint:
	@./build/bin/golangci-lint run --config ./.golangci.yml

lintci-deps:
	rm -f ./build/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./build/bin v1.46.2

goimports:
	goimports -local "$(PACKAGE)" -w .

clean:
	env GO111MODULE=on go clean -cache
	rm -fr build/_workspace/pkg/ $(GOBIN)/*
