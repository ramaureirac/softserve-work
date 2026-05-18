SoftServe Work

Simple URL malware dectector. Uses a hashmap for storing hostnames and insecure paths. Uses `gin-gonic` to expose this db as an API server.

requirements:

- go
- make
- docker (optional)

build & run:

    git clone https://github.com/ramaureirac/softserve-work && cd softserve-work
    make build # or make docker-build
    ./bin/softserve # or docker run -p 8080:8080 -it ramaureirac/softserve:latest

test:

    make test


usage:

    # check url
    curl http://localhost:8080/urlinfo/hecker.info/dolphin.exe

    # register url
    curl -X POST http://localhost:8080/urlinfo/hecker.info/dolphin2.exe