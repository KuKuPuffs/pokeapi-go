GOCMD = go
GOTEST = $(GOCMD) test

test:
	echo "testing poke api..."
	$(GOTEST) -v ./... -count=1