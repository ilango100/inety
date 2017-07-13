.PHONY: all clean debug release

SOURCES = inety.go

all: debug

debug: $(SOURCES)
	go build -tags debug

release: $(SOURCES)
	go build

clean: 
	rm -f inety inety.exe


