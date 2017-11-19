rm -rf hvd_web_server 2> /dev/null;echo "building linux executable";CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o hvd_web_server;
