.PHONY: gifr
gifr:
	go build -o gifr gifr.go

.PHONY: clean
clean:
	rm -rf gifr