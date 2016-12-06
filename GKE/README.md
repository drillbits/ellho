# Run on GKE

## Build on locally

```
$ cd $GOPATH/src/github.com/drillbits/ellho
$ export PROJECT_ID=<GCP_PROJECT_NAME>
$ export IMAGE_NAME=asia.gcr.io/$PROJECT_ID/ellho
$ docker build -t $IMAGE_NAME .
:
Successfully built 9a466f0d293b
$ docker images
REPOSITORY     TAG      IMAGE ID       CREATED         SIZE
<IMAGE_NAME>   latest   9a466f0d293b   1 minutes ago   265 MB
:
```

## Push to Container Registry

```
$ gcloud docker -- push $IMAGE_NAME
```

## Create Kubernetes cluster

```
$ gcloud config set compute/zone asia-northeast1-a
$ gcloud container clusters create ellho
Creating cluster ellho...done.
Created [https://container.googleapis.com/v1/projects/<PROJECT_ID>/zones/asia-northeast1-a/clusters/ellho].
kubeconfig entry generated for ellho.
NAME   ZONE               MASTER_VERSION  MASTER_IP      MACHINE_TYPE   NODE_VERSION  NUM_NODES  STATUS
ellho  asia-northeast1-a  1.4.6           104.198.92.99  n1-standard-1  1.4.6         3          RUNNING
$ gcloud auth application-default login
$ gcloud container clusters get-credentials ellho
Fetching cluster endpoint and auth data.
kubeconfig entry generated for ellho.
```

## Create Kubernetes pod

```
$ kubectl run ellho --image=$IMAGE_NAME --port=8600
deployment "ellho" created
$ kubectl get deployments
NAME    DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
ellho   1         1         1            1           1m
$ kubectl get pods
NAME                     READY     STATUS    RESTARTS   AGE
ellho-2461584618-18swg   1/1       Running   0          1m
```

## Allow external traffic

```
$ kubectl expose deployment ellho --type="LoadBalancer"
service "ellho" exposed
$ kubectl get services ellho
NAME    CLUSTER-IP       EXTERNAL-IP      PORT(S)    AGE
ellho   10.111.251.106   104.198.112.14   8600/TCP   1m
```

## Cleaning it Up

```
$ kubectl delete service,deployment ellho
$ gcloud container clusters delete ellho
```
