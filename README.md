# goswa

A simple go web application example for working out Go web stuffs

* As simple as possible
* No dependencies
* No surprise
* Extensible


Largely inspired by Mat Ryer videos on how he writes web API.

# Build

```shell
make build
```
Will create an executable named `server` in bin/ folder.

With Docker:

```shell
./scripts/build_image.sh
```
An image named `goswa` will be build.

# Run

```shell
make run
```

To test it:
```shell
curl localhost:8080?name=Dude
```

# Deploy on K8S

Tested with minikube.

Make you docker client point to run in minikube environment
with:

```shell
eval $(minikube docker-env)
```

Build your image:

```shell
./scripts/build_image.sh
```

Deploy it:

```shell
minikube kubectl -- apply -f build/package/deployment.yaml
```



