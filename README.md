## Testing

Create [KinD](https://book.kubebuilder.io/reference/kind.html) cluster:

```shell
kind delete cluster --name kind
kind create cluster --config hack/kind-config.yaml --image=kindest/node:v1.20.0
kind export kubeconfig

make install
kubectl get crd
NAME                                CREATED AT
applications.apps.martinheinz.dev   2020-12-31T12:07:25Z

# Make sure to export GOROOT and GOPATH
# export GOROOT=$HOME/go
# export GOPATH="$HOME/go"
# export GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"

export USERNAME=martinheinz
export IMAGE=docker.io/$USERNAME/application-operator:v0.0.1

docker build -t $IMAGE .
docker push $IMAGE # kind load docker-image $IMAGE
make deploy

kubectl create -f config/samples/apps_v1alpha1_application.yaml
kubectl get pods
NAME                                  READY   STATUS    RESTARTS   AGE
application-sample-55bf9d85b7-kvh2j   1/1     Running   0          94s
application-sample-55bf9d85b7-z98sw   1/1     Running   0          94s
```

## Testing Webhooks

```shell
# ... Create Kind Cluster (See above)

# Install cert-manager
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.1.0/cert-manager.yaml
kubectl get pods --namespace cert-manager
NAME                                      READY   STATUS    RESTARTS   AGE
cert-manager-5597cff495-d8mmx             1/1     Running   0          34s
cert-manager-cainjector-bd5f9c764-mssm2   1/1     Running   0          34s
cert-manager-webhook-5f57f59fbc-m8j2j     1/1     Running   0          34s

make docker-build

# Doesn't work (because of imagePullPolicy?) 
kind load docker-image martinheinz/application-operator:latest
docker push martinheinz/application-operator:latest
make deploy IMG=martinheinz/application-operator:latest

kubectl get pods -n application-operator-system
NAME                                                       READY   STATUS    RESTARTS   AGE
application-operator-controller-manager-6d4878c964-8hqlh   2/2     Running   0          2m48s

kubectl logs application-operator-controller-manager-6d4878c964-8hqlh -n application-operator-system -c manager
...
```


## Issues

- `make test` doesn't work:
    - `kubebuilder` has to be downloaded and installed using [docs](https://book.kubebuilder.io/quick-start.html#installation) (no need to `export PATH...`)
    