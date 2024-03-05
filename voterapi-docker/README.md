# How to run (makefile)

```
make docker-build
make run-docker-compose
make run-tests
```

CLEANUP (cleans go cache for tests, if thats a problem for you)

```
make cleanup-cache
```

CLEANUP (deletes background docker images if using )

```
make cleanup
```


# How to run

This program will use a redis cache to store information.

To run tests against the working docker image run:

```
go test ./tests/... -v
```

## Build docker image

RUN:

```
make docker-build
docker-compose up
```

TO RUN FROM DOCKER HUB:

```
docker-compose -f docker-compose-hub.yaml up
```

After run tests or commands 

## Docker hub link

https://hub.docker.com/repository/docker/maddchickenz/battaglialab_voterapi/general

## Build multi architecture 



```
# First create builder (only need to run once)
docker buildx create --name=multi-builder --driver=docker-container
# Build and push to my repository
docker buildx build --platform linux/arm/v7,linux/arm64/8,linux/amd64 --builder multi-builder --push  -t maddchickenz/battaglialab_voterapi:latest .
```


Link to listed OS's: https://hub.docker.com/repository/docker/maddchickenz/battaglialab_voterapi/tags?page=1&ordering=last_updated