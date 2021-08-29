.DEFAULT_GOAL := run

.PHONY: run
run:
	go run main.go deck.go card.go
	jq '.[]' cards_db.json > cards_db_pretty.json

.PHONY: build
build:
	go build

.PHONY: clean
clean:
	rm ./gdeck ./cards_db.json ./cards_db_pretty.json

.PHONY: test
test:
	go test