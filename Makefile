docker:
	docker build -f Dockerfile-api -t coreapi:${VERSION} .
	docker build -f Dockerfile-rpc -t corerpc:${VERSION} .

run-docker:
	docker run -d --name corerpc-${VERSION} --network docker-compose_simple-admin --network-alias corerpc -p 9101:9101 ${DOCKER_USERNAME}/corerpc:${VERSION}
	docker run -d --name coreapi-${VERSION} --network docker-compose_simple-admin --network-alias coreapi -p 9100:9100 ${DOCKER_USERNAME}/coreapi:${VERSION}

run-docker-rpc:
	docker run -d --name corerpc-${VERSION} --network docker-compose_simple-admin --network-alias corerpc -p 9101:9101 ${DOCKER_USERNAME}/corerpc:${VERSION}

run-docker-api:
	docker run -d --name coreapi-${VERSION} --network docker-compose_simple-admin --network-alias coreapi -p 9100:9100 ${DOCKER_USERNAME}/coreapi:${VERSION}

