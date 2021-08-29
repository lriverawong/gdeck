.DEFAULT_GOAL := run

.PHONY: run
run:
	go run main.go deck.go card.go

.PHONY: build
build:
	go build

.PHONY: clean
clean:
	rm ./gdeck ./cards_db.csv

.PHONY: test
test:
	go test