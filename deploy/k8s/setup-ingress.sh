alias kubectl="minikube kubectl --"

# register core-api into ingress tcp config map
kubectl patch configmap tcp-services -n ingress-nginx --patch '{"data":{"9100":"simple-admin/core-api-svc:9100"}}'

# register backend-ui, port 8080 is mapped to port 80 of backend-ui
kubectl patch configmap tcp-services -n ingress-nginx --patch '{"data":{"8080":"simple-admin/backend-ui-svc:80"}}'

# register service into ingress controller
kubectl patch deployment ingress-nginx-controller --patch "$(cat ingress-patch.yaml)" -n ingress-nginx