alias kubectl="minikube kubectl --"

# create service account
kubectl apply -f auth.yaml

# create persistent volume for log
kubectl apply -f pv.yaml

# create core rpc
kubectl apply -f core-rpc.yaml

# create core api
kubectl apply -f core-api.yaml

# create backend ui
kubectl apply -f backend-ui.yaml
