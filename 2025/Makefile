# Default variables
DAY := $(shell date '+%d')
YEAR := $(shell date +%Y)

# Create the directory path with a zero padded day
WORKING_DIR = day$(DAY)

.PHONY: fetch-input fetch-example fetch

# Fetch input for the specified day and year
fetch-input:
	@mkdir -p $(WORKING_DIR)
	@aocd $(DAY) $(YEAR) > "$(WORKING_DIR)/input.txt"
	@echo "Input fetched for Year $(YEAR), Day $(DAY)"

# Fetch example for the specified day and year
fetch-example:
	@mkdir -p $(WORKING_DIR)
	@aocd $(DAY) $(YEAR) --example > "$(WORKING_DIR)/example.txt"
	@echo "Example fetched for Year $(YEAR), Day $(DAY)"

fetch: fetch-input fetch-example

generate:
	go run _template/generate.go -day $(DAY)

run:
	go run day${DAY}/cmd/main.go

run-all:
	go run cmd/runall/main.go

test:
	@go test -json ./day* | gotestfmt

input:
	find ./day* -maxdepth 0 -type d -exec touch {}/input.txt \;

# BubbleTea TUI app
# Find all Go source files in the project
BUBBLETEA_SOURCES := $(shell find . -name '*.go' -not -path './_template/*' 2>/dev/null)

# Build the binary only if sources or dependencies have changed
.bin/bubbletea: $(BUBBLETEA_SOURCES) go.mod go.sum
	@mkdir -p .bin
	@echo "Building bubbletea..."
	@go build -o .bin/bubbletea cmd/bubbleapp/main.go

# Run bubbletea (will rebuild if needed)
.PHONY: bubbletea
bubbletea: .bin/bubbletea
	@./.bin/bubbletea
