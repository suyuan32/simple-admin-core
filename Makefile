version := 0.0.8
docker:
	sudo docker build -f Dockerfile-api -t coreapi:$(version) .
	sudo docker build -f Dockerfile-rpc -t corerpc:$(version) .

run-docker:
	sudo docker run -d --name corerpc-$(version) --network docker-compose_simple-admin --network-alias corerpc -p 9101:9101 corerpc:$(version)
	sudo docker run -d --name coreapi-$(version) --network docker-compose_simple-admin --network-alias coreapi -p 9100:9100 coreapi:$(version)

run-docker-rpc:
	sudo docker run -d --name corerpc-$(version) --network docker-compose_simple-admin --network-alias corerpc -p 9101:9101 corerpc:$(version)

run-docker-api:
	sudo docker run -d --name coreapi-$(version) --network docker-compose_simple-admin --network-alias coreapi -p 9100:9100 coreapi:$(version)
