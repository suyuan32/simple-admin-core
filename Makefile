docker:
	docker build -f Dockerfile-api -t ${DOCKER_USERNAME}/coreapi:${VERSION} .
	docker build -f Dockerfile-rpc -t ${DOCKER_USERNAME}/corerpc:${VERSION} .

publish-docker:
	echo "${DOCKER_PASSWORD}" | docker login --username ${DOCKER_USERNAME} --password-stdin http://${REPO}
	docker push ${REPO}/${DOCKER_USERNAME}/corerpc:${VERSION}
	docker push ${REPO}/${DOCKER_USERNAME}/coreapi:${VERSION}