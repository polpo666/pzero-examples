# simpleapi-with-model-redis

## Install Pzero Framework

```shell
go install github.com/polpo666/pzero/cmd/pzero@latest

pzero check
```

## Generate code

### Generate server code

```shell
pzero gen
```

### Generate swagger code

```shell
pzero gen swagger
```

you can see generated swagger json in `desc/swagger`

## Build docker image

```shell
# add a builder first
docker buildx create --use --name=mybuilder --driver docker-container --driver-opt image=dockerpracticesig/buildkit:master

# build and load
docker buildx build --platform linux/amd64 --progress=plain -t simpleapi-with-model-redis:latest . --load
```

## Documents

https://docs.pzero.io