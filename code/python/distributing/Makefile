all: build-docker cleanup

build-docker:
	docker build -t docker.io/butuzov/python  .

	@ docker images --format "table {{.ID}}\t{{.Repository}}:{{.Tag}}\t{{.Size}}" \
		| grep "butuzov/python"
	@ echo "Run Some Script Inside"
	@ docker run -it docker.io/butuzov/python
	# @ echo "Run Command Inside"
	# @  docker run -it docker.io/butuzov/python -m pip list

cleanup:
	@ echo "@ ----------------------------------------------------- "
	@ echo "List Dangling Images"
	@ docker images -f "dangling=true" -q
	@ echo "@ ----------------------------------------------------- "
	@ echo "Cleanup Dangling Images"
	@ docker images -f "dangling=true" -q | xargs -I {} sh -c "docker rmi {} -f"

