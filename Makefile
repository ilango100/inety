.PHONY: all clean

all: inety.exe

inety.exe: inety.go
	go build $^

clean: 
	rm -f inety.exe


