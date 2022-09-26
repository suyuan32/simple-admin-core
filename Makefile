version := 0.0.1
docker:
	sudo docker build -f Dockerfile-api -t simpleadminapi:$(version) .
	sudo docker build -f Dockerfile-rpc -t simpleadminrpc:$(version) .