OUT := jenkobs
GO := go

SOURCES = \
	*.go

all: update-deps mainbuild strip

mainbuild: $(SOURCES)
	$(GO) build -x -o $(OUT) $(SOURCES)

strip:
	strip $(OUT)

update-deps:
	$(GO) get -u

deps:
	$(GO) get -v ./... 2>/dev/null ||:
strip:
	strip $(OUT)

clean:
	rm $(OUT)
	$(GO) clean -x -i

flush-cache:
	$(GO) clean -cache -modcache -i -r

.PHONY: all install clean
