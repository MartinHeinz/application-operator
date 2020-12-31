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

## Issues

- `make test` doesn't work:
    - `kubebuilder` has to be downloaded and installed using [docs](https://book.kubebuilder.io/quick-start.html#installation) (no need to `export PATH...`)
    