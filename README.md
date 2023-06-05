# Usage

1. Build from root folder of the repo
```
docker build -t hello-world-image .
```

2. Docker run
```
docker run -p 8080:8080 hello-world-image:latest
```

3. Test

```
curl http://localhost:8080/hello
```

### Miniukube important commands
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
### Exec into containers
```
docker exec -it b3d284774829 bash
kubectl exec -it hello-world-deployment-59c8dd6d58-xhwdh -- /bin/sh
```



