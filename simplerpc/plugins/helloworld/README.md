# helloworld

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

## Build docker image

```shell
# add a builder first
docker buildx create --use --name=mybuilder --driver docker-container --driver-opt image=dockerpracticesig/buildkit:master

# build and load
docker buildx build --platform linux/amd64 --progress=plain -t helloworld:latest . --load
```

## Documents

https://docs.pzero.io