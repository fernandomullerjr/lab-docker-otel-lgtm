
## Run lgtm in Kubernetes

```sh
# Create k8s resources
cd /home/fernando/cursos/opentelemetry/lab-docker-otel-lgtm
kubectl apply -f k8s/lgtm.yaml

# Configure port forwarding
kubectl port-forward service/lgtm 3000:3000 4040:4040 4317:4317 4318:4318 9090:9090

# Using mise
mise k8s-apply
mise k8s-port-forward
```

### Run

Run the example REST service:

#### Unix/Linux

```sh
./run-example.sh

## Aplicação em Python
cd /home/fernando/cursos/opentelemetry/lab-docker-otel-lgtm/examples/python
./run.sh
curl http://127.0.0.1:8082/rolldice
```


### Generate traffic

#### Unix/Linux

```sh
./generate-traffic.sh
```





>
> kubectl apply -f k8s/lgtm.yaml
service/lgtm created
                    deployment.apps/lgtm created
                                                %                                                                                                                                                                                                                                       > kubectl port-forward service/lgtm 3000:3000 4317:4317 4318:4318
error: unable to forward port because pod is not running. Current status=Pending



>
> kubectl get pods -A
NAMESPACE              NAME                                         READY   STATUS              RESTARTS       AGE
default                lgtm-5f966c7986-kk6gv                        0/1     ContainerCreating   0              45s
kube-system            coredns-66bc5c9577-rmfg7                     1/1     Running             1 (112m ago)   6d4h
kube-system            etcd-minikube                                1/1     Running             1 (112m ago)   6d4h
kube-system            kube-apiserver-minikube                      1/1     Running             1 (112m ago)   6d4h
kube-system            kube-controller-manager-minikube             1/1     Running             1 (112m ago)   6d4h
kube-system            kube-proxy-ljs8m                             1/1     Running             1 (112m ago)   6d4h
kube-system            kube-scheduler-minikube                      1/1     Running             1 (112m ago)   6d4h
kube-system            storage-provisioner                          1/1     Running             3 (111m ago)   6d4h
kubernetes-dashboard   dashboard-metrics-scraper-77bf4d6c4c-s646h   1/1     Running             1 (112m ago)   6d4h
kubernetes-dashboard   kubernetes-dashboard-855c9754f9-5wv9d        1/1     Running             2 (111m ago)   6d4h
>
>
>
> kubectl get pods -A
NAMESPACE              NAME                                         READY   STATUS              RESTARTS       AGE
default                lgtm-5f966c7986-kk6gv                        0/1     ContainerCreating   0              47s
kube-system            coredns-66bc5c9577-rmfg7                     1/1     Running             1 (112m ago)   6d4h
kube-system            etcd-minikube                                1/1     Running             1 (112m ago)   6d4h
kube-system            kube-apiserver-minikube                      1/1     Running             1 (112m ago)   6d4h
kube-system            kube-controller-manager-minikube             1/1     Running             1 (112m ago)   6d4h
kube-system            kube-proxy-ljs8m                             1/1     Running             1 (112m ago)   6d4h
kube-system            kube-scheduler-minikube                      1/1     Running             1 (112m ago)   6d4h
kube-system            storage-provisioner                          1/1     Running             3 (111m ago)   6d4h
kubernetes-dashboard   dashboard-metrics-scraper-77bf4d6c4c-s646h   1/1     Running             1 (112m ago)   6d4h
kubernetes-dashboard   kubernetes-dashboard-855c9754f9-5wv9d        1/1     Running             2 (111m ago)   6d4h
>
>
> kubectl get pods -A
NAMESPACE              NAME                                         READY   STATUS    RESTARTS       AGE
default                lgtm-5f966c7986-kk6gv                        1/1     Running   0              3m17s
kube-system            coredns-66bc5c9577-rmfg7                     1/1     Running   1 (114m ago)   6d4h
kube-system            etcd-minikube                                1/1     Running   1 (114m ago)   6d4h
kube-system            kube-apiserver-minikube                      1/1     Running   1 (114m ago)   6d4h
kube-system            kube-controller-manager-minikube             1/1     Running   1 (114m ago)   6d4h
kube-system            kube-proxy-ljs8m                             1/1     Running   1 (114m ago)   6d4h
kube-system            kube-scheduler-minikube                      1/1     Running   1 (114m ago)   6d4h
kube-system            storage-provisioner                          1/1     Running   3 (113m ago)   6d4h
kubernetes-dashboard   dashboard-metrics-scraper-77bf4d6c4c-s646h   1/1     Running   1 (114m ago)   6d4h
kubernetes-dashboard   kubernetes-dashboard-855c9754f9-5wv9d        1/1     Running   2 (113m ago)   6d4h
> ./run-example.sh
+ [[ ! -f ./target/rolldice.jar ]]
+ ./mvnw clean package






- APP de exemplo só tem 1 Java
- E está com problemas

~~~~bash
> ./run-example.sh
+ [[ ! -f ./target/rolldice.jar ]]
+ ./mvnw clean package
The JAVA_HOME environment variable is not defined correctly,
this environment variable is needed to run this program.
~~~~


## PENDENTE
- Revisar como subir app de exemplo em Python da pasta "lab-docker-otel-lgtm/examples/python"
- Revisar como subir app de exemplo em Go da pasta "lab-docker-otel-lgtm/examples/go"





> ls
Dockerfile  app.py  docker-compose.oats.yml  oats.yaml  requirements.txt  run.sh
> ./run.sh
Requirement already satisfied: pip in ./venv/lib/python3.12/site-packages (24.2)
Collecting pip
  Downloading pip-25.3-py3-none-any.whl.metadata (4.7 kB)
Downloading pip-25.3-py3-none-any.whl (1.8 MB)
   ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 1.8/1.8 MB 5.3 MB/s eta 0:00:00
Installing collected packages: pip
  Attempting uninstall: pip
    Found existing installation: pip 24.2
    Uninstalling pip-24.2:
      Successfully uninstalled pip-24.2
Successfully installed pip-25.3
Collecting blinker==1.9.0 (from -r requirements.txt (line 1))

  Downloading opentelemetry_instrumentation_urllib3-0.60b1-py3-none-any.whl.metadata (4.2 kB)
Requirement already satisfied: opentelemetry-api~=1.12 in ./venv/lib/python3.12/site-packages (from opentelemetry-instrumentation-urllib3==0.60b1) (1.39.1)
Requirement already satisfied: opentelemetry-instrumentation==0.60b1 in ./venv/lib/python3.12/site-packages (from opentelemetry-instrumentation-urllib3==0.60b1) (0.60b1)
Requirement already satisfied: opentelemetry-semantic-conventions==0.60b1 in ./venv/lib/python3.12/site-packages (from opentelemetry-instrumentation-urllib3==0.60b1) (0.60b1)
Requirement already satisfied: opentelemetry-util-http==0.60b1 in ./venv/lib/python3.12/site-packages (from opentelemetry-instrumentation-urllib3==0.60b1) (0.60b1)
Requirement already satisfied: wrapt<2.0.0,>=1.0.0 in ./venv/lib/python3.12/site-packages (from opentelemetry-instrumentation-urllib3==0.60b1) (1.17.3)
Requirement already satisfied: packaging>=18.0 in ./venv/lib/python3.12/site-packages (from opentelemetry-instrumentation==0.60b1->opentelemetry-instrumentation-urllib3==0.60b1) (25.0)
Requirement already satisfied: typing-extensions>=4.5.0 in ./venv/lib/python3.12/site-packages (from opentelemetry-semantic-conventions==0.60b1->opentelemetry-instrumentation-urllib3==0.60b1) (4.15.0)
Requirement already satisfied: importlib-metadata<8.8.0,>=6.0 in ./venv/lib/python3.12/site-packages (from opentelemetry-api~=1.12->opentelemetry-instrumentation-urllib3==0.60b1) (8.7.1)
Requirement already satisfied: zipp>=3.20 in ./venv/lib/python3.12/site-packages (from importlib-metadata<8.8.0,>=6.0->opentelemetry-api~=1.12->opentelemetry-instrumentation-urllib3==0.60b1) (3.23.0)
Downloading opentelemetry_instrumentation_urllib3-0.60b1-py3-none-any.whl (13 kB)
Installing collected packages: opentelemetry-instrumentation-urllib3
Successfully installed opentelemetry-instrumentation-urllib3-0.60b1
 * Debug mode: off
INFO:werkzeug:WARNING: This is a development server. Do not use it in a production deployment. Use a production WSGI server instead.
 * Running on http://127.0.0.1:8082
INFO:werkzeug:Press CTRL+C to quit
WARNING:app:Anonymous player is rolling the dice: 6
INFO:werkzeug:127.0.0.1 - - [17/Jan/2026 18:47:37] "GET /rolldice HTTP/1.1" 200 -
WARNING:app:Anonymous player is rolling the dice: 5
INFO:werkzeug:127.0.0.1 - - [17/Jan/2026 18:47:39] "GET /rolldice HTTP/1.1" 200 -
WARNING:app:Anonymous player is rolling the dice: 5
INFO:werkzeug:127.0.0.1 - - [17/Jan/2026 18:47:41] "GET /rolldice HTTP/1.1" 200 -
WARNING:app:Anonymous player is rolling the dice: 4
INFO:werkzeug:127.0.0.1 - - [17/Jan/2026 18:47:44] "GET /rolldice HTTP/1.1" 200 -
WARNING:app:Anonymous player is rolling the dice: 6
INFO:werkzeug:127.0.0.1 - - [17/Jan/2026 18:47:46] "GET /rolldice HTTP/1.1" 200 -
WARNING:app:Anonymous player is rolling the dice: 4




## PENDENTE
- Revisar como subir app de exemplo em Go da pasta "lab-docker-otel-lgtm/examples/go"
- Efetuar um cenário variado diferente.