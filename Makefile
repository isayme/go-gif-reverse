COMMIT_HASH := $(shell git rev-parse HEAD)

.PHONY: gifr install clean
gifr:
	go build -ldflags "-X main.commitHash=${COMMIT_HASH}" -o gifr gifr.go

install:
	mv ./gifr ${GOBIN}/

clean:
	rm -rf gifr