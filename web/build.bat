del /Q server.exe

go build -a --installsuffix cgo --ldflags="-s" -o server.exe