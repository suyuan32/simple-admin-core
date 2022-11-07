# Quick Commands

We offer several commands in makefile ， you can just run in bash:

```shell
# build docker image, require environment variables DOCKER_USERNAME VERSION 
make docker  

# publish docker, require environment variables DOCKER_USERNAME VERSION DOCKER_PASSWORD
make publish-docker

# generate api code with files in api/desc, and generate swagger file
make gen-api

# generate code by rpc/core.proto
make gen-rpc

# generate ent code
make gen-ent

# generate swagger
make gen-swagger

# run swagger service
make serve-swagger

# visit doc locally
make doc

```