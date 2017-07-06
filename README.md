# inety
**inety** is a raw tcp client.

It simply connects to a server on a port you specify and starts sending data from stdin to the server and put the data from server to the stdout.

It is useful for debugging tcp connections. It is a simple client made using Go.

There are two flavours of inety:
- Release 
- Debug

## Compilation

### Linux
To compile inety, install go in your system and execute:
```
make release
```
If you want debug version, simply execute
```
make
```

### Windows
On windows, use of cygwin or msys2 is recommended for easier commandline, though you can compile and use inety directly from PowerShell or Cmd. If you use cygwin or msys2, then compilation is same as for Linux. If you want to use from pure Windows, then:
1. Install go from https://golang.org
2. Download source code and from source directory execute
```
go build
```
3. If you want debug version, execute
```
go build -tags debug
```

## Usage
Simply execute inety from commandline, it will ask for address and port on execution.
Or you can specify the address and port number from commandline:
```
inety ftp.gnu.org ftp
```
```
inety -p 21 -a ftp.gnu.org
```

For port you can use port numbers directly or words like `ftp`, `http`, etc. 
On connection, you can type in to send to server and result will be received and displayed.
To end the connection, send the EOF: 
- ^D in unix
- ^Z in Windows

## Bugs
If you have found any bugs, please open an issue.
