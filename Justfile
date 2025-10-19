
# Shows this help
default:
    just -l --unsorted

# Shows the command's help
help:
    go run . -h

# Runs an advent of code solution
run year day part input_file:
    go run . -y {{year}} -d {{day}} -p {{part}} -i {{input_file}}

# Runs all unit tests
test:
    go test ./...

# Runs all unit tests, without using cache
test-no-cache:
    go test ./... -count=1
