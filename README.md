### Min ukube important commands
```
minikube start
minikube status
minikube stop

minikube service hello-world --url
minikube service list

eval $(minikube docker-env)
```
### K8s cluster related commands
```
kubectl cluster-info
kubectl config get-contexts
```
### K8s deployment commands
```
kubectl apply -f hello-world-deployment.yaml
kubectl apply -f hello-world-service.yaml
```

### Docker commands
```
docker build -t hello-world-image .
docker tag my-local-image:latest my-local-repo/my-image:latest
```

