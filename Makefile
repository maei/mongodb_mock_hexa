GOC=go
GO111MODULE=on

DIRECTORIES=$(sort $(dir $(wildcard src/domain)))
MOCKS=$(foreach x, $(DIRECTORIES), mocks/$(x))

.PHONY: all test clean-mocks mocks

all: demo

demo:
	$(GOC) build -o demo

test: | mocks
	go test ./...

clean-mocks:
	rm -rf mocks

mocks: $(MOCKS)

$(MOCKS): mocks/% : %
	mockery -output=$@ -dir=$^ -all