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


| Environment Variables | Introduction                                       |
|-----------------------|----------------------------------------------------|
| VERSION               | Version number such as 0.0.1                       |
| DOCKER_USERNAME       | docker repository username                         |
| DOCKER_PASSWORD       | docker repository password                         |
| REPO                  | docker repository address， official hub: docker.io |

> If you want to publish docker images, you must set all environment variables above.