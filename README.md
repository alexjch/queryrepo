# Build repo query

```
go build cmd/queryrepo/main.go
```

# Build image

```
sudo docker build --force-rm --file deployments/Dockerfile .
```

to tag the image use

```
sudo docker build --force-rm --tag <repo>/queryrepo:<tag> --file deployments/Dockerfile .
```

# Deploy service using helm

```
helm install queryrepo deployments/helm-chart/
```

# If need to upgrade after chaging ```values.yaml```

```
helm upgrade -f deployments/helm-chart/values.yaml queryrepo deployments/helm-chart
```

# Expose service to the outside

```
kubectl expose deployments queryrepo-queryrepo-helm --type=LoadBalancer --name=lb-queryrepo-queryrepo-helm --port=9001 --target-port=39001
```

test using curl

```
curl http://<lb ip>:<lb port>/api/v1/ --data '{"name": "rpm"}'

```

if successful the returning value should be:

```
{"success":false,"data":"rpm-0:4.11.3-43.el7.src\n"}
```
