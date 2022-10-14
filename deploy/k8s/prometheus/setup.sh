alias kubectl="minikube kubectl --"

kubectl apply -f rbac.yaml

kubectl apply -f prometheus-service-monitor.yaml

