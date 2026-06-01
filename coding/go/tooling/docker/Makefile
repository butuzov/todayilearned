VERSION ?= 1.16


# ~~~ Build ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

version:              ## Prints Go version we going to work with.
	echo $(VERSION)

bases-pull:           ## Downlaods Base Images
	@docker pull golang:$(VERSION)-alpine
	@docker pull golang:$(VERSION)-buster


bases-list:            ## List Base Images
	@docker images | grep golang

buster:  # Build Debian Buster Image
	docker build -t butuzov/goapp-debian -f debian.Dockerfile .

distroless: # Build Distroless Image
	docker build -t butuzov/goapp-distroless -f distroless.Dockerfile .

alpine:        ## Build Alpine Image
	docker build -t butuzov/goapp-alpine -f alpine.Dockerfile .


all: alpine buster distroless ## Build all images

cleanfailed:          ## Clean untagged images
	@docker images -q -f "dangling=true" | xargs -L1 docker rmi
	@docker ps -q -f "status=exited" | xargs -L1 docker rm


# ~~~ Docker Compose ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

up: # Startup Environment
	docker-compose up --no-recreate -d

down: # Shutdown Environment
	docker-compose down -v

logs: # See Logs
	docker-compose logs

ls:  # Show resulted images
	@docker images | grep "butuzov/goapp"

clean: down cleanfailed    ## Teardown All Bases/Failed/Resulted
	@ echo "Cleanup Docker Images"
	@ docker images | grep "butuzov/goapp" | awk '{print $$3}' | xargs -L1 docker rmi
	@ @docker images | grep "golang" | awk '{print $$3}' | xargs -L1 docker rmi
