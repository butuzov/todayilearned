# DCA Exam

Article by Bhargav Bachina [link](https://medium.com/bb-tutorials-and-thoughts/250-practice-questions-for-the-dca-exam-84f3b9e8f5ce)


Docker is an essential tool in every organization nowadays. Every company is implementing DevOps and containerize their applications for easier and faster production deployments. Docker Certified Associate Exam is designed to validate the Docker skills necessary for any individual to succeed in today's DevOps world.

These practice questions are intended for those who want to take the DCA exam and entirely based on this study guide from the Docker. All the material for these questions is taken from the [official docker documentation](https://docs.docker.com/). These questions are divided based on the sections from the study guide.

* Orchestration (25%)
* Image Creation, Management, and Registry (20%)
* Installation and Configuration (15%)
* Networking (15%)
* Security (15%)
* Storage and Volumes (10%)
* Orchestration (25%)

##  Orchestration (25%)
1. Which command do you use to create a new swarm?

```
docker swarm init --advertise-addr <MANAGER-IP>
```

2. What is this flag --advertise-addr for?

> This flag configures the IP address for the manager node and The other nodes in the swarm must be able to access the manager at the IP address.

3. How do you know the current status of the swarm?
```
docker info // you can find the info under the swarm section
```

4. Which command do you use to find the information about the nodes in the swarm?
```
docker node ls
```

5. How to add another manager to the swarm?
```
// it generate the instructions for the manager to be added
docker swarm join-token manager
```

6. How to add another worker node to the swarm?
```
// it generate the instructions for the worker to be added
docker swarm join-token worker
```
7. How to run the container?
```
docker run <image>
```

8. What is the autolock feature in the Docker swarm?

> When Docker restarts, both the TLS key used to encrypt communication among swarm nodes, and the key used to encrypt and decrypt Raft logs `n` disk, are loaded into each manager node’s memory.
> Docker 1.13 introduces the ability to protect the mutual TLS encryption key and the key used to encrypt and decrypt Raft logs at rest, by allowing you to take ownership of these keys and to require manual unlocking of your managers. This feature is called autolock.


9. How to lock the swarm?
```
// This command produces unlock key. You need to place that in safe place
docker swarm init --autolock
```
10. How to unlock the swarm?
```
docker swarm unlock
```
11. Are we able to enable autolock feature only when we create a swarm for the first time?
> No. You can lock the existing swarm as well

12. How to enable or disable autolock on the existing swarm?
```
//enable autolock
docker swarm update --autolock=true
//disable autolock
docker swarm update --autolock=false
```

13. How to view the current unlock key for the running swarm?
```
docker swarm unlock-key
```

14. How to rotate the unlock key?
```
docker swarm unlock-key --rotate
```

15. If the key was rotated after one of the manager nodes became unavailable and if you don’t have access to the previous key you may need to force the manager to leave the swarm and join it back as a new manager. Is this statement correct?
 > Yes
16. How to deploy a service in the docker swarm?
```
// for the nginx image
docker create service --replicas 3 --name nginx-web nginx
```
17. How to list the services in the Docker swarm?
```
docker service ls
```

18. How to list the tasks of the service in the Docker swarm?

```
docker service ps <service name>
```

19. How to inspect the service on the swarm?

```
docker service inspect <service name>
```
20. How to inspect the service on the swarm so that it will print limited information in an  easily readable format?
```
docker service inspect <service> --pretty
```
21. How to find out which nodes are running the service?
```
docker service ps <service>
```
22. How to find out more details of the container running these tasks of the service?
```
// you need to run this command on the particular node
docker ps
```
23. If you are running co-related services in the docker swarm, what do you call this?
> stack
24. What is Docker stack?
> A stack is a group of interrelated services that share dependencies, and can be orchestrated and scaled together.
25. Explain the several commands associated with Docker stack?
```
// deploy the new stack or update
docker stack deploy -c <compose file>
// list services in the stack
docker stack services
// list the tasks in the stack
docker stack ps
// remove the stack
docker stack rm
//List stack
docker stack ls
```

26. How to filter the services in the stack?
```
// with the help of --filter flag
docker stack service nginx-web --filter name=web
```

27. How to format the output of the docker stack services command?
```
docker stack services --format "{{.ID}}: {{.Mode}} {{.Replicas}}"
```

28. How to increase the number of replicas?
```
docker service scale SERVICE=REPLICAS
// example
docker service scale frontend=50
// you can scale multiple services as well
docker service scale frontend=50 backend=30
// you can also scale with the update command
docker service update --replicas=50 frontend
```

29. How to revert the changes for the service configuration?
```
docker service rollback my-service
```

30. What are the networks available for the docker services?
  * overlay networks: manage communications among the Docker daemons participating in the swarm.You can attach a service to one or more existing overlay networks as well, to enable service-to-service communication.
  * ingress network: is a special overlay network that facilitates load balancing among a service’s nodes. When any swarm node receives a request on a published port, it hands that request off to a module called IPVS. IPVS keeps track of all the IP addresses participating in that service, selects one of them, and routes the request to it, over the ingress network.
  * docker_gwbridge: is a bridge network that connects the overlay networks (including the ingress network) to an individual Docker daemon’s physical network.

31. Is the ingress network created automatically when you initialize or join a swarm?
> Yes
32. Is docker_gwbridge network created automatically when you initialize or join a swarm?
> Yes

33. How to create an overlay network?
```
docker network create --driver overlay my-network
// you can customize it
 docker network create \
  --driver overlay \
  --subnet 10.0.9.0/24 \
  --gateway 10.0.9.99 \
  my-network
```
34. How to inspect the network?
```
docker network inspect my-network
```
35. How to attach a service to an overlay network?
```
docker service create \
  --replicas 3 \
  --name my-web \
  --network my-network \
  nginx
```
36. Can service containers connected to the overlay network communicate with each other?
> Yes
37. How to find which networks the service is connected to?
```
docker network inspect my-network
               or
docker service ls // for the name
docker service ps <SERVICE> // to list the networks
```
38. Customize the ingress network involves removing and creating a new one and you need to do that before you create any services in the swarm. Is this statement correct?
> Yes

39. How to remove and create an ingress network?
```
docker network rm ingress
docker network create \
  --driver overlay \
  --ingress \
  --subnet=10.11.0.0/16 \
  --gateway=10.11.0.2 \
  --opt com.docker.network.mtu=1200 \
  my-ingress
```
40. What is the difference between `-v` and `--mount` flags in terms of creating volumes?
```
Originally, the -v or --volume flag was used for standalone containers and the --mount flag was used for swarm services. However, starting with Docker 17.06, you can also use --mount with standalone containers. In general, --mount is more explicit and verbose.
```
41. How to create a service with volume?
```
docker service create -d \
  --replicas=4 \
  --name devtest-service \
  --mount source=myvol2,target=/app \
  nginx:latest
```
42. Does docker service create command supports `-v` or `—volume` flag?
> No

43. What are the volume drivers?

> When building fault-tolerant applications, you might need to configure multiple replicas of the same service to have access to the same files.
> Volume drivers allow you to abstract the underlying storage system from the application logic. For example, if your services use a volume with an NFS driver, you can update the services to use a different driver, as an example to store data in the cloud, without changing the application logic.

44. How to create a volume with the volume driver?
```
docker volume create --driver vieux/sshfs \
  -o sshcmd=test@node2:/home/test \
  -o password=testpassword \
  sshvolume
```
45. How to create a service with volume driver?
```
docker service create -d \
  --name nfs-service \
  --mount 'type=volume,source=nfsvolume,target=/app,volume-driver=local,volume-opt=type=nfs,volume-opt=device=:/var/docker-nfs,volume-opt=o=addr=10.0.0.10' \
  nginx:latest
```
46. I created a deployment that runs exactly one task on every node. which type of service deployment is this?
```
global
```
47. I created a deployment that runs several identical tasks on nodes. which type of service deployment is this?
```
replicated
```
48. If you want to troubleshoot the UCP clusters what is the best method?
> it's always best practice to use client bundle to troubleshoot UCP clusters

49. What is the general flow when troubleshooting services or clusters?
```
docker service ls
docker service ps <service>
docker service inspect <service>
docker inspect <task>
docker inspect <container>
docker logs <container>
```
50. How to update metadata about a node?
> you can use labels to add metadata about the node

51. How to add a label to the node?
```
docker node update --label-add foo worker1
// add multiple labels
docker node update --label-add foo --label-add bar worker1
```
52. How to remove the label from the node?
```
docker node update --label-rm foo worker1
```
53. How to set up the service to divide tasks evenly over different categories of nodes?
```
--placement-pref
// example: if we have three datacenters 3 replicas will be placed on each datacenter
docker service create \
  --replicas 9 \
  --name redis_2 \
  --placement-pref 'spread=node.labels.datacenter' \
  redis:3.0.6
```
53. How to limit your service on particular nodes?
```
--constraint
// example: the following limits tasks for the redis service to nodes where the node type label equals queue
docker service create \
  --name redis_2 \
  --constraint 'node.labels.type == queue' \
  redis:3.0.6
```
54. Which algorithm does the docker engine use when it is in swarm mode to manage the global cluster state?
> Raft Consensus Algorithm

55. What is a quorum and why it is important?
> Quorun ensure that the cluster state stays consistent in the presence of failures by requiring a majority of nodes to agree on values.
> Raft tolerates up to (N-1)/2 failures and requires a majority or quorum of (N/2)+1 members to agree on values proposed to the cluster. Without quorun swarm wont be able to serve the requests

56. What are the supported flags for creating services with templates?
```
--env
--mount
--hostname
// example
service create --name hosttempl \
    --hostname="{{.Node.Hostname}}-{{.Node.ID}}-{{.Service.Name}}"\
      busybox top
```

## Image Creation, Management, and Registry (20%)

57. Which instruction sets the base image for the subsequent builds in the Dokcerfile?
```
FROM
```

58. No instruction can precede FROM in the Dockerfile. Is this statement correct?
> No. ARG is the only instruction can precede FROM

59. What are the two forms for the RUN instruction?
```
shell form: RUN <command>
exec form: RUN ["executable", "param1", "param2"]
```
60. What does the RUN instruction do in the Dockerfile?

> The RUN instruction will execute any commands in a new layer on top of the current image and commit the results.

61. The RUN command normally utilizes cache from the previous build. Which flag should you specify for the build not to use cache?
```bash
# --no-cache
docker build --no-cache .
```
62. Is there any other instruction that can invalidate the cache?
> Yes. ADD

63. How many forms that CMD instruction has?

```
CMD ["executable","param1","param2"] (exec form, this is the preferred form)
CMD ["param1","param2"] (as default parameters to ENTRYPOINT)
CMD command param1 param2 (shell form)
```
64. If CMD instruction provides default arguments for the ENTRYPOINT instruction, both should be specified in JSON format. Is this statement correct?
> Yes

65. What is the purpose of the CMD instruction in the Dockerfile?

> The main purpose of a CMD is to provide defaults for an executing container. These defaults can include an executable, or they can omit the executable, in which case you must specify an ENTRYPOINT instruction as well.

66. How to make your container execute the same executable every time?
```bash
use ENTRYPOINT in combination with CMD
```
67. What is the purpose of the LABEL instruction in the Dockerfile?
> It adds metadata to the Image

68. How to check the labels for the current image?
```
docker inspect // Under Labels section
```
69. The EXPOSE instruction actually publish the port. Is this statement correct?
> No. It serves as a type of documentation between the image publisher and image consumer

70. What should you do to actually publish the ports?
```
use -p flag when running a container
```
71. What is the purpose of the ENV instruction in the Dockerfile?
```dockerfile
ENV <key> <value>
# an ENV instruction sets the enviroment value to the key
# and it is available for the subsequent build steps and in the running container as well.
```
72. How to change the environment variables when running containers?
```
docker run --env <key>=<value>
```
73. What is the difference between ADD and COPY instructions?
```Dockerfile
ADD [--chown=<user>:<group>] <src>... <dest>
# The ADD instruction copies new files, directories or remote file URLs
# from <src> and adds them to the filesystem of the image at the path <dest>.
COPY [--chown=<user>:<group>] <src>... <dest>
# The COPY instruction copies new files or directories from <src> and adds
# them to the filesystem of the container at the path <dest>.
```
74. What is ENTRYPOINT instruction in the Dockerfile?
```
# An ENTRYPOINT allows you to configure a container that will run as an executable.
# Command line arguments to docker run <image> will be appended after all elements
# in an exec form ENTRYPOINT, and will override all elements specified using CMD.
```
75. How can you override the ENTRYPOINT instruction?
```
docker run --entrypoint
```
76. What is the VOLUME instruction in the Dockerfile?
> The VOLUME instruction creates a mount point with the specified name and marks it as holding externally mounted volumes from native host or other containers.

77. What initializes the newly created Volume?
```
docker run -v
```
78. What is the USER instruction in the Dockerfile?
> The USER instruction sets the user name (or UID) and optionally the user group (or GID) to use when running the image and for any RUN, CMD and ENTRYPOINT instructions that follow it in the Dockerfile.

79. What is the WORKDIR instruction in the Dockerfile?

> The WORKDIR instruction sets the working directory for any RUN, CMD, ENTRYPOINT, COPY and ADD instructions that follow it in the Dockerfile.

80. You have specified multiple WORKDIR instructions in the Dockerfile what is the result WORKDIR?

```
WORKDIR /a
WORKDIR b
WORKDIR c
RUN pwd
result: /a/b/c
```

81. You have specified multiple WORKDIR instructions in the Dockerfile what is the result WORKDIR?
```
WORKDIR /a
WORKDIR /b
WORKDIR c
RUN pwd
result: /b/c
```
82. What is the ARG instruction in the Dockerfile?
```Dockerfile
ARG <name>[=<default value>]
# The ARG instruction defines a variable that users can pass at
# build-time to the builder with the docker build command using the
# --build-arg <varname>=<value> flag.
```
83. What is the ONBUILD instruction in the Dockerfile?
> The ONBUILD instruction adds to the image a trigger instruction to
be executed at a later time, when the image is used as the base for another build.

84. Which instruction sets the system call signal that will be sent to the container to exit?
```
STOPSIGNAL signal
```
85. Which instruction let Docker daemon know the health of the container?
> HEALTHCHECK

86. What are all the options that can be provided for the HEALTHCHECK instruction?
```
--interval=DURATION (default: 30s)
--timeout=DURATION (default: 30s)
--start-period=DURATION (default: 0s)
--retries=N (default: 3)
```
87. What is the SHELL instruction in the Dockerfile?
```
The SHELL instruction allows the default shell used for the shell form of commands to be overridden. The default shell on Linux is ["/bin/sh", "-c"], and on Windows is ["cmd", "/S", "/C"]. The SHELL instruction must be written in JSON form in a Dockerfile.
```

88. Create ephemeral containers is considered best practice?
> Yes

89. What should you do if you want to exclude some files while executing the docker build image and don’t want to send all the files to Docker daemon?
```
use .dockerignore file
```
90. What is the best way to drastically reduce the size of an image?
> Multi Stage Builds

91. How do you minimize the number of layers while building the image?
```
Only the instructions RUN, COPY, ADD create layers.
Where possible, use multi-stage builds, and only copy the artifacts you need into the final image.
sort multi line arguments
RUN apt-get update && apt-get install -y \
  bzr \
  cvs \
  git \
  mercurial \
  subversion
```
92. How to leverage the build cache?
> Put instructions that likely to change often at the bottom of the dockerfile.

93. How to remove unused images?
```
docker image prune
```

94. How to see the history of the image?
```
docker image history
```

95. How to format the output of the docker inspect command?

```
//by using --format flag
//examples
docker inspect --format='{{range .NetworkSettings.Networks}}{{.MacAddress}}{{end}}' $INSTANCE_ID
docker inspect --format='{{.LogPath}}' $INSTANCE_ID
```
96. How to tag an image?
```
docker tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]
docker tag 0e5574283393 fedora/httpd:version1.0 // by id
docker tag httpd fedora/httpd:version1.0 // by name
docker tag httpd:test fedora/httpd:version1.0.test // by name and tag
docker tag 0e5574283393 myregistryhost:5000/fedora/httpd:version1.0
```
97. How to run a local registry?
```
docker run -d -p 5000:5000 --restart=always --name registry registry:2
```

98. How to copy an image from the docker hub to a local repository?
```
// pull an image from the Docker Hub
docker pull ubuntu
// tag an image
docker tag ubuntu:16.04 localhost:5000/my-ubuntu
// push the image
docker push localhost:5000/my-ubuntu
```
99. How to stop and remove a local registry?
```
docker container stop registry && docker container rm -v registry
```
100. How to display the layers of the Docker image?
```
docker image inspect //under Layers section
```
101. How to create a Docker image from archive or stdin?
```
docker image load
// example
docker image load -i example.tar
```
102. How to modify an image to a single layer?
```
// take any multiple layer image
// run the container
docker export <container> > single-layer.tar
docker import /path/to/single-layer.tar
// check the history
docker image history
```
103. Each layer is only a set of differences from the layer before it. The layers are stacked on top of each other. Is this statement about the image correct?
> Yes

104. When you create a container It adds one writable layer on top of all the layers of the image. Is this statement about the image correct?
> yes

105. What is the copy-on-write (CoW) strategy?
> Copy-on-write is a strategy of sharing and copying files for maximum efficiency. If a file or directory exists in a lower layer within the image, and another layer (including the writable layer) needs read access to it, it just uses the existing file. The first time another layer needs to modify the file (when building the image or running the container), the file is copied into that layer and modified. This minimizes I/O and the size of each of the subsequent layers.

106. How to customize the registry while deploying?
```
// customize published port
docker run -d \
  -p 5001:5000 \
  --name registry-test \
  registry:2
// If you want to change the port the registry listens on within the container
docker run -d \
  -e REGISTRY_HTTP_ADDR=0.0.0.0:5001 \
  -p 5001:5001 \
  --name registry-test \
  registry:2
// storage customization
docker run -d \
  -p 5000:5000 \
  --restart=always \
  --name registry \
  -v /mnt/registry:/var/lib/registry \
  registry:2
```
107. How to configure a registry?
```
The Registry configuration is based on a YAML file. you can specify a configuration variable from the environment by passing -e arguments to your docker run stanza or from within a Dockerfile using the ENV instruction.
// for example you have a configuration like this for root directory
storage:
  filesystem:
    rootdirectory: /var/lib/registry
// you can create environment variable like this
REGISTRY_STORAGE_FILESYSTEM_ROOTDIRECTORY=/somewhere
it will change from /var/lib/registry to /somewhere
```
108. What is the location of the registry configuration file?
```
/etc/docker/registry/config.yml
```
109. How to customize an entire config file of registry?
```
docker run -d -p 5000:5000 --restart=always --name registry \
             -v `pwd`/config.yml:/etc/docker/registry/config.yml \
             registry:2
```
110. How to login to a self-hosted registry?
```
docker login localhost:5000
```

111. Where do you configure any credential helpers or credentials for the registry to prevent passing every time you log in?
```
/etc/docker/daemon.json
```

112. How to limit the number of records when docker search?
```
docker search nginx --limit=2
```

113. How to format the docker search?
```
docker search --format "{{.Name}}: {{.StarCount}}" nginx
```

114. How to disable Image signing while pushing an image to the repository?
```bash
docker push [OPTIONS] NAME[:TAG]
--disable-content-trust=true
```
115. How to enable docker content trust in the Docker CLI?

```bash
export DOCKER_CONTENT_TRUST=1
docker push <dtr-domain>/<repository>/<image>:<tag>
```

116. How to pull an image from the repository?
```
# docker pull [OPTIONS] NAME[:TAG|@DIGEST]
// pulling from docker hub by default
docker pull debian
// pulling from other repositories
docker pull myregistry.local:5000/testing/test-image
```

117. How to pull an image with multiple images?
```
# -a or --all-tags
docker pull --all-tags fedora
```

118. How to remove all images which are not used by existing containers?
```
docker image prune -a
```

119. How to limit the scope when pruning images?
```bash
# by uisng --filter flag
docker image prune -a --filter "until=24h"
```
120. How to remove an image?
```
docker rmi <IMAGE ID>
```
121. How to remove image without deleting the untagged parent images?
```
docker rmi --no-prune <IMAGE ID>
```
122. How to delete an image from the repository?
```
login into DTR web UI
go to the TAGS section delete the specific TAG
you can also delete all images by deleting the entire repository
```
##  Installation and Configuration (15%)

123. What is the recommended way of installing Docker
```
set up docker repositories
install from them for the ease of installation and upgrade tasks.
```
124. How to upgrade docker-engine?
```
sudo apt-get update
install docker from the instructions from here
```
125. How to uninstall docker?
```
sudo apt-get purge docker-ce
sudo rm -rf /var/lib/docker
```
126. Are Images, containers, volumes, or customized configuration files on your host are not automatically removed when you uninstall docker?

> No. You need to explicitly delete those

127. How to add the user to the Docker group and use docker as a non-root user?
```
sudo usermod -aG docker your-user
```
128. What are the ways to install docker?
  * using repositories
  * using DEB package
  * using convience scripts

129. How to install Docker CE on Centos?
```
// uninstall older versions
sudo yum remove docker \
                docker-client \
                docker-client-latest \
                docker-common \
                docker-latest \
                docker-latest-logrotate \
                docker-logrotate \
                docker-engine
// install required libs
sudo yum install -y yum-utils \
  device-mapper-persistent-data \
  lvm2
// set up the stable repo
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
// install
sudo yum install docker-ce docker-ce-cli containerd.io
// if you want to install specific versions
sudo yum install docker-ce-<VERSION_STRING> docker-ce-cli-<VERSION_STRING> containerd.io
// start docker
sudo systemctl start docker
```

130. How to install Docker CE on Debian?
```
// uninstall older versions
sudo apt-get remove docker docker-engine docker.io containerd runc
// update
sudo apt-get update
// install required
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg2 \
    software-properties-common
// add dockers official gpg key
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
// set up stable repo
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/debian \
   $(lsb_release -cs) \
   stable"
// update and install
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
// if you want to install specific versions
sudo apt-get install docker-ce=<VERSION_STRING> docker-ce-cli=<VERSION_STRING> containerd.io
```
131. How to install Docker CE on Fedora?
```
// uninstall old versions
sudo dnf remove docker \
                docker-client \
                docker-client-latest \
                docker-common \
                docker-latest \
                docker-latest-logrotate \
                docker-logrotate \
                docker-selinux \
                docker-engine-selinux \
                docker-engine
// install required packages
sudo dnf -y install dnf-plugins-core
// add the stable repo
sudo dnf config-manager \
    --add-repo \
    https://download.docker.com/linux/fedora/docker-ce.repo
// install community version
sudo dnf install docker-ce docker-ce-cli containerd.io
// if you want specific versions
sudo dnf -y install docker-ce-<VERSION_STRING> docker-ce-cli-<VERSION_STRING> containerd.io
// start docker
sudo systemctl start docker
```
132. How to install Docker CE on Ubuntu?
```
// uninstall old versions
sudo apt-get remove docker docker-engine docker.io containerd runc
// update and install required packages
sudo apt-get update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
// add official gpg key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
// stable repo
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
// update and install
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
// if you want specific versions
sudo apt-get install docker-ce=<VERSION_STRING> docker-ce-cli=<VERSION_STRING> containerd.io
```
133. What are the recommended storage drivers on different distributions?
```
Centos: overlay2
Ubuntu supports overlay2, aufs and btrfs storage drivers. Overlay2 is the default one
```
134. What are all the release channels that Docker CE supports?
```
Stable gives you latest releases for general availability.
Test gives pre-releases that are ready for testing before general availability.
Nightly gives you latest builds of work in progress for the next major release.
```
135. Where are the Docker-CE binaries available?
> Docker Engine - Community binaries for a release are available on download.docker.com as packages for the supported operating systems.
136. Where are the Docker-EE binaries available?
> Docker Hub
137. What are logging drivers?
> Docker has multiple mechanisms to get the logging information from running docker containers and services. These mechanisms are called logging drivers

138. How to configure a logging driver for the Docker daemon so that all the containers use it?
```
configure log-driver in /etc/docker/daemon.json
{
  "log-driver": "syslog"
}
```
139. Whats is the default logging driver?
> `json-file`

140. If you have configurable options for your logging driver how do you specify?
```
use log-opts in the daemon.json file
{
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "10m",
    "max-file": "3",
    "labels": "production_status",
    "env": "os,customer"
  }
}
```
141. How to find the logging driver for the Docker daemon?
```
docker info --format '{{.LoggingDriver}}'
```
142. How to configure a logging driver for a container?
```
docker run -it --log-driver json-file --log-opt max-size=10m alpine ash
```
143. What are the available logging drivers for the Docker CE edition?
```
json-file
local
journald
```
144. If the swarm loses the quorum of managers it loses the ability to perform management tasks. Is this statement correct?
> Yes

145. If the swarm loses quorum all the existing tasks and services are all deleted. Is this statement correct?
> No. All the existing tasks will continue to run. But, new nodes cannot be added and new tasks can't be created.

146. We should use a fixed IP address for the advertise address to prevent the swarm from becoming unstable on machine reboot. Is this statement correct?
> Yes. If the whole swarm restarts and every manager node subsequently gets a new IP address, there is no way for any node to contact an existing manager. Therefore the swarm is hung while nodes try to contact one another at their old IP addresses.

147. You should maintain an odd number of managers in the swarm to support manager node failures. Is this statement correct?
> Yes.

148. I have manager nodes 3, 5, 7, 9. How do you distribute these manager nodes on availability zones so that If you suffer a failure in any of those zones, the swarm should maintain the quorum of manager nodes
```
Manager Nodes           Availability Zones
    3                     1-1-1
    5                     2-2-1
    7                     3-2-2
    9                     3-3-3
```
149. How to drain the node
```
docker node update --availability drain <NODE>
```
150. How to cleanly rejoin a manager node in the cluster?

    1. To demote the node to a worker, run docker node demote <NODE>
    2. To remove the node from the swarm, run docker node rm <NODE>
    3. Re-join the node to the swarm with a fresh state using docker swarm jo

151. How to forcibly remove a node?
```
docker node rm --force <NODE>
```
152. If you want to remove a manager node you need to demote it to a worker role first. Is this statement correct?
> Yes. You must ensure that there is a quorum

153. What is the location where swarm managers save the swarm state?

```
/var/lib/docker/swarm
```

154. How to backup the swarm?

  * If autolock is enabled. You must unlock the swarm
  * stop the docker on the manager node so that you don't have unpredictable results
  * save the entire contents of /var/lib/docker/swarm
  * start the manager

155. How to restore swarm from the backup?

  1. shut down the docker on the targeted machine
  2. Remove the contents of /var/lib/docker/swarm
  3. Restore the /var/lib/docker/swarm directory from the backup
  4. Start the docker on the node so that it doesn't connect to old ones
  docker swarm init --force-new-cluster
  5. Verify the state of the swarm docker service ls
  6. rotate the autolock key
  7. Add manager and worker nodes for the required capacity
  8. backup this swarm

156. What is a team in the DTR?

> A team defines the permissions a set of users have for a set of repositories.

157. What are all the permission levels that teams could have?

```
Read Only: View repository and pull images.
Read Write: View repository, pull and push images.
Admin: Manage repository and change its settings, pull and push images.
```

158. Where is the Docker daemon directory?
```
/var/lib/docker on Linux
C:\ProgramData\docker on Windows
```
159. How to enable the debugging on Docker daemon
```
  1. add this flag in /etc/docker/daemon.json
  {
    "debug": true
  }
  2. Send a HUP signal to the daemon to cause it to reload its configuration.
  sudo kill -SIGHUP $(pidof dockerd)
```
160. How to check whether Docker is running?
```
// all these can be used depending on the operating system
docker info
sudo systemctl is-active docker
sudo status docker
sudo service docker status
```
161. What are the hardware and software requirements for UCP?
```
Minimum
1. 8GB of RAM for manager nodes or nodes running DTR
2. 4GB of RAM for worker nodes
3. 3GB of free disk space
Recommended
1. 16GB of RAM for manager nodes or nodes running DTR
2. 4 vCPUs for manager nodes or nodes running DTR
3. 25-100GB of free disk space
```
162. What products that Docker EE contains?
```
UCP
DTR
Docker Engine with enterprise-grade support,
```
163. Where is the location of the custom certificates?
```
/etc/docker/certs.d
```
164. What are the ports that DTR uses?
```
80/tcp     -     Web app and API client access to DTR.
443/tcp    -     Web app and API client access to DTR
```
165. DTR needs to be installed on a worker node that is being managed by UCP. You can't install DTR on a standalone Docker engine. Is this statement correct?
> Yes

166. How to backup the UCP
```bash
# To create a UCP backup, run the docker/ucp:2.2.22 backup command on a single UCP
# manager
docker container run \
  --log-driver none --rm \
  --interactive \
  --name ucp \
  -v /var/run/docker.sock:/var/run/docker.sock \
  docker/ucp:2.2.22 backup \
  --id <ucp-instance-id> \
  --passphrase "secret" > /tmp/backup.tar
```
167. How to restore the UCP
```bash
docker/ucp:2.2.22 restore --passphrase "secret"
docker container run --rm -i --name ucp \
  -v /var/run/docker.sock:/var/run/docker.sock  \
  docker/ucp:2.2.22 restore --passphrase "secret" < /tmp/backup.tar
```
168. You need to backup the UCP on the single manager node since Docker maintains the same  UCP state in all the manager nodes. Is this statement correct?
> Yes

169. How to backup the DTR?
> To perform a backup of a DTR node, run the docker/dtr backup command.

170. DTR requires that a majority (n/2 + 1) of its replicas are healthy at all times for it to work. So if a majority of replicas are unhealthy or lost, the only way to restore DTR to a working state is by recovering from a backup. Is this statement correct?
> Yes

171. How to configure Docker to start on boot?
```
sudo systemctl enable docker
```

## Networking (15%)

172. What is the default network that the docker creates automatically?
> Bridge

173. How to list the networks on the Docker machine?
```
docker netwrok ls
```
174. How to connect to the default bridge network when you create a container?
```
// since no network is specified, it will be connected to default bridge network
docker run -dit --name alpine1 alpine ash
```
175. How to inspect the default network bridge?
```
docker network inspect bridge
```
176. The default bridge network is not recommended for production. Is this statement correct?
> Yes

177. How to create a user-defined network?
```
docker network create --driver bridge my-network
```
178. How to inspect the user-defined network?
```
docker network inspect my-network
```
179. How to connect to the user-defined network while creating a container?
```
docker run -dit --name alpine1 --network my-network alpine ash
```
180. How to connect the existing container to the user-defined network?
```
docker netwrok connect my-network alpine2
```
181. How to troubleshoot a user-defined network?
```
// using  nicolaka/netshoot
docker run -it --rm --network container:<container_name> nicolaka/netshoot
```
182. How to publish a port so that it can be accessed externally?
```
docker run -p 127.0.0.1:$HOSTPORT:$CONTAINERPORT --name CONTAINER -t <image>
```
183. How to list port mappings or a specific mapping for the container?
```
// List the containers
docker ps
// use this command with container name
docker port <CONTAINER NAME>
// USE the specific port
docker port <CONTAINER NAME> <specific port>
```
184. What are all the different built-in network drivers?
```
Bridge Network Driver
Overlay Network Driver
MACVLAN Driver
Host
None
```
185. What are the Bridge network and its use case?
```
The bridge driver creates a private network internal to the host so containers on this network can communicate.
The bridge driver does the service discovery for us automatically if two containers are on the same network
The bridge driver is a local scope driver, which means it only provides service discovery, IPAM, and connectivity on a single host.
```
186. What is the scope of the bridge network?
```
local
```
187. What are the Overlay network and their use case?
> The built-in Docker overlay network driver radically simplifies many of the complexities in multi-host networking.
> It is a swarm scope driver, which means that it operates across an entire Swarm or UCP cluster rather than individual hosts.

188. What is the scope of the overlay network?
> swarm

189. What are the MACVLAN network and their use case?
> The macvlan driver is the newest built-in network driver and offers several unique characteristics.
> It’s a very lightweight driver, because rather than using any Linux bridging or port mapping, it connects container interfaces directly to host interfaces.

190. What is the scope of the macvlan network?
```local```
191. What are the Host network and its use case?
> With the host driver, a container uses the networking stack of the host. There is no namespace separation, and all interfaces on the host can be used directly by the container.

192. What is the scope of the host network?
```
local
```
193. What are the None network and its use case?
> The none driver gives a container its own networking stack and network namespace but does not configure interfaces inside the container. Without additional configuration, the container is completely isolated from the host networking stack.

194. What is the scope of the None network?
```
local
```
195. The Docker networking architecture is built on a set of interfaces called the Container Networking Model (CNM). Is this statement correct?
> Yes

196. What is a sandbox in the CNM model?
> A Sandbox contains the configuration of a container's network stack. This includes the management of the container's interfaces, routing table, and DNS settings. An implementation of a Sandbox could be a Windows HNS or Linux Network Namespace, a FreeBSD Jail, or other similar concept. A Sandbox may contain many endpoints from multiple networks.

197. What is an endpoint in the CNM model?
> An Endpoint joins a Sandbox to a Network. The Endpoint construct exists so the actual connection to the network can be abstracted away from the application. This helps maintain portability so that a service can use different types of network drivers without being concerned with how it's connected to that network.

198. What is a network in the CNM model?
> The CNM does not specify a Network in terms of the OSI model. An implementation of a Network could be a Linux bridge, a VLAN, etc. A Network is a collection of endpoints that have connectivity between them. Endpoints that are not connected to a network do not have connectivity on a network.

199. What part of the Docker that provides the actual implementation that makes networks work?

## Network Drivers

200. What is IPAM drivers?

> Docker has a native IP Address Management Driver that provides default
subnets or IP addresses for the networks and endpoints if they are not specified.

201. How to configure docker to use external DNS?
```
edit the /etc/docker/daemon.json
{
   "dns": ["10.0.0.2", "8.8.8.8"]
}
restart the docker
sudo systemctl docker restart
```

202. Which network should handles control and data traffic related to swarm services?
> ingress

203. Which network which connects the individual Docker daemon to the other daemons participating in the swarm?

```
docker_gwbridge
```

204. What is the default network created when you create a swarm cluster?
> ingress


205. How to create a user-defined overlay network for communication among services?
```
docker network create -d overlay my-overlay
```

206. How to create an overlay network which can be used by swarm services or standalone containers to communicate with other standalone containers running on other Docker daemons?
```
create with --attachable flag
docker network create -d overlay --attachable my-attachable-overlay
```

207. All the swarm management data is encrypted by default. Is this statement correct?
> Yes

208. is application data on the swarm encrypted by default?
> No

209. How to encrypt application data as well on the swarm?
```
// use --opt=encrypted
docker network create --opt encrypted --driver overlay --attachable my-attachable-multi-host-network
```
210. What is the host port publishing mode?

```
To publish a service’s port directly on the node where it is running, use the mode=host option to the --publish flag.
```

## Security (15%)

211. What is DCT?
```
Through DCT, image publishers can sign their images and image consumers can ensure that the images they use are signed.
```
212. What is DCT stand for?
```
Docker Content Trust
```
213. What is the command to generate delegation keys?
```
docker trust generate key
```

214. How to load if you have any existing keys?
```
docker trust key load
```
215. How to sign a particular tag and push it up to the registry?
```
docker trust sign dtr.example.com/admin/demo:1
```
216. How to enable docker content trust so that you can sign images automatically when you use docker push?
```
export DOKCER_CONTENT_TRUST=1
```

217. How to inspect remote trusted data for a tag?
```
docker trust inspect
```
218. How to remove remote trusted data for a tag?

```
docker trust revoke
```

219. What is a grant?
> A grant defines who has how much access to set of resources

220. What is the subject?
> A subject can be user, team, organization and is granted a role for set of resources

221. What is the role?
> A role is a set of permitted API operations that you can assign to a specific subject and collection by using a grant

222. What is a Client Bundle?
> A client bundle is a group of certificates downloadable directly from the Docker Universal Control Plane (UCP) user interface within the admin section for “My Profile”. This allows you to authorize a remote Docker engine to a specific user account managed in Docker EE, absorbing all associated RBAC controls in the process. You can now execute docker swarm commands from your remote machine that take effect on the remote cluster.

223. What is the easiest way to access or control the UCP?
> Client Bundle

224. What is the kernel feature that isolates the processes running in different containers?
> Namespaces

225. Which kernel feature limits the resources used by docker containers?
> Control Groups

226. What is the kernel feature that needed extra configuration?
> user

227. Docker swarm should be secure by default?
> yes

## Storage and Volumes (10%)

228. What is the pluggable architecture that Docker supports for the container writable layer storage?
> Storage Drivers

229. What is the preferred storage driver for all Linux distributions which need no extra configuration?
> Overlay2

230. Which device-mapper driver is used for production environments?
> direct-lvm

231. Which device-mapper driver is used for testing environments?
> loopback-lvm

232. How do you check the current storage driver information?
```
docker info
```
233. How do you configure device-mapper?
```
// stop docker
sudo systemctl stop docker
// set the device-mapper in /etc/docker/daemon.json file
{
  "storage-driver": "devicemapper"
}
//start docker
sudo systemctl start docker
```

234. what is the option that sets the direct-lvm for production device-mapper?
```
dm.directlvm_device
```

235. What do you set the device-mapper and all configurable options in the daemon.json?
```
{
  "storage-driver": "devicemapper",
  "storage-opts": [
    "dm.directlvm_device=/dev/xdf",
    "dm.thinp_percent=95",
    "dm.thinp_metapercent=1",
    "dm.thinp_autoextend_threshold=80",
    "dm.thinp_autoextend_percent=20",
    "dm.directlvm_device_force=false"
  ]
}
```
236. What are the different available storage options available for containers?
```
Block Storage
FiLE System Storage
Object Storage
```
237. Do containers create a writable layer on top of Image read-only layers?
> Yes

238. Where is the Docker’s local storage area?
```
/var/lib/docker/<storage-driver>
```
239. What is the difference between bind mounts and volumes?
```
Volumes are completely managed by docker
Bind Mounts are dependent on the host directory structure
```

240. Volumes don’t increase the size of the containers. Is this statement correct and why?
> Yes. Because volumes live outside of containers

241. What should we use if we want to persist the data beyond the lifecycle of the containers?
> Volumes

242. How to create a Volume?
```
docker volume create my-volume
```
243. How to list docker volumes?
```
docker volume ls
```
244. How to inspect docker volume?
```
docker volume inspect my-vol
```
245. How to remove docker volumes?
```
docker volume rm my-vol
```
246. If you start a container with a volume that does not yet exist, Docker creates the volume for you. Is this statement correct?
> Yes

247. How to create a volume myvol2 with a docker run?
```
docker run -d \
  --name devtest \
  -v myvol2:/app \
  nginx:latest
```
248. How to verify the volume is created with the container?
```
// Look for the mounts section
docker inspect devtest
```
249. How to create a volume with the --mount flag?
```
docker run -d \
  --name devtest \
  --mount source=myvol2,target=/app \
  nginx:latest
```
250. How to remove all unused images not just dangling images?
```
docker system prune --all
```