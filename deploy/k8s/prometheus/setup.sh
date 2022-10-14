kubectl create namespace simpe-admin

kubectl apply -f rbac.yaml

kubectl apply -f prometheus-service-monitor.yaml

kubectl apply -f auth.yaml

kubectl apply

for ns in default kube-system monitoring test; do
  kubectl patch ns $ns --patch '{"metadata":{"labels":{"serviceMonitor": "prometheus" } } }'
done
