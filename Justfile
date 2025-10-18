
# Shows this help
default:
    just -l --unsorted

# Shows the command's help
help:
    go run . -h

# Runs an advent of code solution
run year day part:
    go run . -y {{year}} -d {{day}} -p {{part}}
