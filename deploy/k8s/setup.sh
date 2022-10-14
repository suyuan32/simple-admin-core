alias kubectl="minikube kubectl --"

kubectl apply -f auth.yaml

kubectl apply -f core-rpc.yaml

kubectl apply -f core-api.yaml

kubectl apply -f backend-ui.yaml
