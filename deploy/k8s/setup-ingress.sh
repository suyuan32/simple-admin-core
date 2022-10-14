alias kubectl="minikube kubectl --"

# 注册 core-api
kubectl patch configmap tcp-services -n ingress-nginx --patch '{"data":{"9100":"simple-admin/core-api-svc:9100"}}'

# 注册 backend-ui, 将 8080 端口映射到backend-ui的80端口
kubectl patch configmap tcp-services -n ingress-nginx --patch '{"data":{"8080":"simple-admin/backend-ui-svc:80"}}'

kubectl patch deployment ingress-nginx-controller --patch "$(cat ingress-patch.yaml)" -n ingress-nginx