RED  =  "\e[31;1m"
GREEN = "\e[32;1m"
YELLOW = "\e[33;1m"
BLUE  = "\e[34;1m"
PURPLE = "\e[35;1m"
CYAN  = "\e[36;1m"

docker:
	docker build -f Dockerfile-api -t ${DOCKER_USERNAME}/core-api:${VERSION} .
	docker build -f Dockerfile-rpc -t ${DOCKER_USERNAME}/core-rpc:${VERSION} .
	# docker build -f Dockerfile-job -t ${DOCKER_USERNAME}/core-job:${VERSION} .
	@printf $(GREEN)"[SUCCESS] build docker successfully"

publish-docker:
	echo "${DOCKER_PASSWORD}" | docker login --username ${DOCKER_USERNAME} --password-stdin https://${REPO}
	docker push ${DOCKER_USERNAME}/core-rpc:${VERSION}
	docker push ${DOCKER_USERNAME}/core-api:${VERSION}
	# docker push ${DOCKER_USERNAME}/core-job:${VERSION}
	@printf $(GREEN)"[SUCCESS] publish docker successfully"

gen-api:
	goctls api go --api ./api/desc/all.api --dir ./api --trans_err=true
	swagger generate spec --output=./core.yml --scan-models
	@printf $(GREEN)"[SUCCESS] generate API successfully"

gen-rpc:
	goctls rpc protoc ./rpc/core.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc
	@printf $(GREEN)"[SUCCESS] generate rpc successfully"

gen-ent:
	go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./pkg/ent/template/*.tmpl" ./pkg/ent/schema
	@printf $(GREEN)"[SUCCESS] generate ent successfully"

gen-rpc-ent-logic:
	goctls rpc ent --schema=./ent/schema  --style=go_zero --service_name=example --o=./ --model=$(model) --group=$(group)
	@printf $(GREEN)"[SUCCESS] generate ent logic codes successfully"

gen-swagger:
	swagger generate spec --output=./core.yml --scan-models
	@printf $(GREEN)"[SUCCESS] generate swagger successfully"

serve-swagger:
	lsof -i:36666 | awk 'NR!=1 {print $2}' | xargs killall -9 || true
	@printf $(GREEN)"[SUCCESS] serve swagger-ui successfully"
	swagger serve -F=swagger --port 36666 core.yml

doc:
	docsify serve ./docs