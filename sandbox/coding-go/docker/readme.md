<!-- menu: Docker Images -->
# Building Docker Images

I this example we are building a multistage small image for go binary. This is usually a goto example on devops classes/tutorials etc.

## Keypoints

1. Multistage builds are used.
2. UPX used to reduce binary size
3. Running Go Infrastructure within docker image.


## Read More

  * https://ardanlabs.com/blog/2020/02/docker-images-part1-reducing-image-size.html
  * https://ardanlabs.com/blog/2020/02/docker-images-part2-details-specific-to-different-languages.html

## Examples

{{% list "Makefile,main.go,docker-compose.yaml,alpine.Dockerfile,debian.Dockerfile,distroless.Dockerfile" %}}
