.PHONY: bench

bench:
	@go test ./... -bench=. -benchtime=1s -benchmem -count 2
