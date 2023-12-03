DAY=$$(date +%d)

run:
	go run day${DAY}/cmd/main.go

run-all:
	go run cmd/runall/main.go

test:
	@go test -json ./day* | gotestfmt

input:
	find ./day* -maxdepth 0 -type d -exec touch {}/input.txt \;
