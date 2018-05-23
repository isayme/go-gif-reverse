COMMIT_HASH := $(shell git rev-parse HEAD)

.PHONY: gifr
gifr:
	go build -ldflags "-X main.commitHash=${COMMIT_HASH}" -o gifr gifr.go

.PHONY: clean
clean:
	rm -rf gifr