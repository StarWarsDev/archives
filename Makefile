include hellogopher.mk

regenerate:
	$(q) $(GO) run -v github.com/99designs/gqlgen $1

run:
	$(q) $(GO) run main.go
