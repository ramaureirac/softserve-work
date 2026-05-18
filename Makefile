GO_BINARY_FOLDER = bin
GO_BINARY_NAME   = softserve
GO_SOURCES       = src/cmd
GO_VULN			 = ~/go/bin/govulncheck

DOCKER_TAG		 = ramaureirac/softserve
DOCKER_FILE		 = docker/Dockerfile
DOCKER_COMPOSE	 = docker/Compose.yml
DOCKER_VOLUMES	 = docker/Volumes


## pkg build
build:
	@rm -rf $(GO_BINARY_FOLDER)
	@go build -o $(GO_BINARY_FOLDER)/$(GO_BINARY_NAME) ./$(GO_SOURCES)


## docker 
docker-build:
	docker build -t ${DOCKER_TAG}:latest -f ./${DOCKER_FILE} .

