docker:
	docker build -f Dockerfile-api -t ${DOCKER_USERNAME}/core-api:${VERSION} .
	docker build -f Dockerfile-rpc -t ${DOCKER_USERNAME}/core-rpc:${VERSION} .
	# docker build -f Dockerfile-job -t ${DOCKER_USERNAME}/core-job:${VERSION} .

publish-docker:
	echo "${DOCKER_PASSWORD}" | docker login --username ${DOCKER_USERNAME} --password-stdin http://${REPO}
	docker push ${REPO}/${DOCKER_USERNAME}/core-rpc:${VERSION}
	docker push ${REPO}/${DOCKER_USERNAME}/core-api:${VERSION}
	# docker push ${REPO}/${DOCKER_USERNAME}/core-job:${VERSION}