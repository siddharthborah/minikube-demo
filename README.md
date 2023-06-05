minikube start
minikube status
minikube stop


kubectl cluster-info
kubectl config get-contexts


kubectl apply -f hello-world-deployment.yaml
kubectl apply -f hello-world-service.yaml


minikube service hello-world --url
minikube service list


docker build -t hello-world-image .

minikube image repository my-local-repo
docker tag my-local-image:latest my-local-repo/my-image:latest


eval $(minikube docker-env)