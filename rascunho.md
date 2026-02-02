
### COMANDOS ÚTEIS NO MINIKUBE:

```bash
minikube start --cpus=4 --memory=8192 --driver=docker
```

~~~~BASH
## Abrir Dashboard do k8s
minikube dashboard

## Ativar o tunnel
minikube tunnel

## Fornecer url e porta para acessar o serviço, OBS: Gera porta aleatória acessível
minikube service app-a --url
minikube service grafana --url
~~~~

## Run lgtm in Kubernetes

```sh
# Create k8s resources
cd /home/fernando/cursos/opentelemetry/lab-docker-otel-lgtm
kubectl apply -f k8s/lgtm.yaml

# Configure port forwarding
kubectl port-forward service/lgtm 3000:3000 4040:4040 4317:4317 4318:4318 9090:9090
# abra: <http://localhost:3000>

# Using mise
mise k8s-apply
mise k8s-port-forward

## App em Go
kubectl port-forward service/rolldice 8081:8081
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



## Dia 31/01/2026

- TSHOOT Porta 3000


>
> sudo grep -R --line-number '"HostPort":"3000"' /var/lib/docker/containers 2>/dev/null
/var/lib/docker/containers/dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1/config.v2.json:1:{"StreamConfig":{},"State":{"Running":true,"Paused":false,"Restarting":false,"OOMKilled":false,"RemovalInProgress":false,"Dead":false,"Pid":40708,"ExitCode":0,"Error":"","StartedAt":"2026-01-31T15:55:49.832936683Z","FinishedAt":"2026-01-31T15:55:49.936046015Z","Health":{"Status":"starting","FailingStreak":0,"Log":[{"Start":"2026-01-31T12:53:55.24765491-03:00","End":"2026-01-31T12:53:55.298302547-03:00","ExitCode":0,"Output":"true\n"},{"Start":"2026-01-31T12:54:22.577534444-03:00","End":"2026-01-31T12:54:22.627184374-03:00","ExitCode":0,"Output":"true\n"},{"Start":"2026-01-31T12:54:49.910598269-03:00","End":"2026-01-31T12:54:49.960754345-03:00","ExitCode":0,"Output":"true\n"},{"Start":"2026-01-31T12:55:17.238577061-03:00","End":"2026-01-31T12:55:17.288109259-03:00","ExitCode":0,"Output":"true\n"},{"Start":"2026-01-31T12:55:44.562054354-03:00","End":"2026-01-31T12:55:44.612553597-03:00","ExitCode":0,"Output":"true\n"}]}},"ID":"dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1","Created":"2025-08-31T18:38:59.11161328Z","Managed":false,"Path":"bash","Args":["start.sh"],"Config":{"Hostname":"dde7a8080851","Domainname":"","User":"0:0","AttachStdin":false,"AttachStdout":false,"AttachStderr":false,"ExposedPorts":{"8080/tcp":{}},"Tty":false,"OpenStdin":false,"StdinOnce":false,"Env":["OLLAMA_BASE_URL=http://host.docker.internal:11434","PATH=/usr/local/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin","LANG=C.UTF-8","GPG_KEY=A035C8C19219BA821ECEA86B64E628F8D684696D","PYTHON_VERSION=3.11.13","PYTHON_SHA256=8fb5f9fbc7609fa822cb31549884575db7fd9657cbffb89510b5d7975963a83a","ENV=prod","PORT=8080","USE_OLLAMA_DOCKER=false","USE_CUDA_DOCKER=false","USE_SLIM_DOCKER=false","USE_CUDA_DOCKER_VER=cu128","USE_EMBEDDING_MODEL_DOCKER=sentence-transformers/all-MiniLM-L6-v2","USE_RERANKING_MODEL_DOCKER=","OPENAI_API_BASE_URL=","OPENAI_API_KEY=","WEBUI_SECRET_KEY=","SCARF_NO_ANALYTICS=true","DO_NOT_TRACK=true","ANONYMIZED_TELEMETRY=false","WHISPER_MODEL=base","WHISPER_MODEL_DIR=/app/backend/data/cache/whisper/models","RAG_EMBEDDING_MODEL=sentence-transformers/all-MiniLM-L6-v2","RAG_RERANKING_MODEL=","SENTENCE_TRANSFORMERS_HOME=/app/backend/data/cache/embedding/models","TIKTOKEN_ENCODING_NAME=cl100k_base","TIKTOKEN_CACHE_DIR=/app/backend/data/cache/tiktoken","HF_HOME=/app/backend/data/cache/embedding/models","HOME=/root","WEBUI_BUILD_VERSION=2407d9b905978d68619bdce4021e424046ec8df9","DOCKER=true"],"Cmd":["bash","start.sh"],"Healthcheck":{"Test":["CMD-SHELL","curl --silent --fail http://localhost:${PORT:-8080}/health | jq -ne 'input.status == true' || exit 1"]},"Image":"ghcr.io/open-webui/open-webui:main","Volumes":null,"WorkingDir":"/app/backend","Entrypoint":null,"OnBuild":null,"Labels":{"org.opencontainers.image.created":"2025-08-28T10:40:43.547Z","org.opencontainers.image.description":"User-friendly AI Interface (Supports Ollama, OpenAI API, ...)","org.opencontainers.image.licenses":"NOASSERTION","org.opencontainers.image.revision":"2407d9b905978d68619bdce4021e424046ec8df9","org.opencontainers.image.source":"https://github.com/open-webui/open-webui","org.opencontainers.image.title":"open-webui","org.opencontainers.image.url":"https://github.com/open-webui/open-webui","org.opencontainers.image.version":"main"}},"Image":"sha256:1df9c23da52bf54e82590b721f9e091e3e9865271f14cdad364dee4a2d257661","ImageManifest":null,"NetworkSettings":{"Bridge":"","SandboxID":"d33feda521f8a56cebf1ca2aac56de087b93de0494ea80c20654acf86b107ff6","SandboxKey":"/var/run/docker/netns/d33feda521f8","HairpinMode":false,"LinkLocalIPv6Address":"","LinkLocalIPv6PrefixLen":0,"Networks":{"bridge":{"IPAMConfig":null,"Links":null,"Aliases":null,"MacAddress":"02:42:ac:12:00:02","DriverOpts":null,"NetworkID":"a7c0932a2acbbb9e19501a90f34108c290f4de0698eccd17d757d3606be0a8bc","EndpointID":"ff3d49d85b2fa549deb0228b7eb48395408e78ac01b723eb7e0c1791619f1dfa","Gateway":"172.18.0.1","IPAddress":"172.18.0.2","IPPrefixLen":16,"IPv6Gateway":"","GlobalIPv6Address":"","GlobalIPv6PrefixLen":0,"DNSNames":null,"IPAMOperational":false,"DesiredMacAddress":""}},"Service":null,"Ports":{"8080/tcp":[{"HostIp":"0.0.0.0","HostPort":"3000"}]},"SecondaryIPAddresses":null,"SecondaryIPv6Addresses":null,"HasSwarmEndpoint":false},"LogPath":"/var/lib/docker/containers/dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1/dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1-json.log","Name":"/open-webui","Driver":"overlay2","OS":"linux","RestartCount":0,"HasBeenStartedBefore":true,"HasBeenManuallyStopped":false,"MountPoints":{"/app/backend/data":{"Source":"/var/lib/docker/volumes/open-webui/_data","Destination":"/app/backend/data","RW":true,"Name":"open-webui","Driver":"local","Type":"volume","Relabel":"z","ID":"cede6c78f60adc202cb6630ad65d2051352f26ea754bf05057ed12a8a9e8fc7e","Spec":{"Type":"volume","Source":"open-webui","Target":"/app/backend/data"},"SkipMountpointCreation":false}},"SecretReferences":null,"ConfigReferences":null,"MountLabel":"","ProcessLabel":"","AppArmorProfile":"","SeccompProfile":"","NoNewPrivileges":false,"HostnamePath":"/var/lib/docker/containers/dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1/hostname","HostsPath":"/var/lib/docker/containers/dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1/hosts","ShmPath":"","ResolvConfPath":"/var/lib/docker/containers/dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1/resolv.conf","LocalLogCacheMeta":{"HaveNotifyEnabled":false}}
/var/lib/docker/containers/dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1/hostconfig.json:1:{"Binds":["open-webui:/app/backend/data"],"ContainerIDFile":"","LogConfig":{"Type":"json-file","Config":{"max-size":"100m"}},"NetworkMode":"bridge","PortBindings":{"8080/tcp":[{"HostIp":"","HostPort":"3000"}]},"RestartPolicy":{"Name":"always","MaximumRetryCount":0},"AutoRemove":false,"VolumeDriver":"","VolumesFrom":null,"ConsoleSize":[36,120],"CapAdd":null,"CapDrop":null,"CgroupnsMode":"host","Dns":[],"DnsOptions":[],"DnsSearch":[],"ExtraHosts":["host.docker.internal:host-gateway"],"GroupAdd":null,"IpcMode":"private","Cgroup":"","Links":null,"OomScoreAdj":0,"PidMode":"","Privileged":false,"PublishAllPorts":false,"ReadonlyRootfs":false,"SecurityOpt":null,"UTSMode":"","UsernsMode":"","ShmSize":67108864,"Runtime":"runc","Isolation":"","CpuShares":0,"Memory":0,"NanoCpus":0,"CgroupParent":"","BlkioWeight":0,"BlkioWeightDevice":[],"BlkioDeviceReadBps":[],"BlkioDeviceWriteBps":[],"BlkioDeviceReadIOps":[],"BlkioDeviceWriteIOps":[],"CpuPeriod":0,"CpuQuota":0,"CpuRealtimePeriod":0,"CpuRealtimeRuntime":0,"CpusetCpus":"","CpusetMems":"","Devices":[],"DeviceCgroupRules":null,"DeviceRequests":null,"MemoryReservation":0,"MemorySwap":0,"MemorySwappiness":null,"OomKillDisable":false,"PidsLimit":null,"Ulimits":[],"CpuCount":0,"CpuPercent":0,"IOMaximumIOps":0,"IOMaximumBandwidth":0,"MaskedPaths":["/proc/asound","/proc/acpi","/proc/kcore","/proc/keys","/proc/latency_stats","/proc/timer_list","/proc/timer_stats","/proc/sched_debug","/proc/scsi","/sys/firmware","/sys/devices/virtual/powercap"],"ReadonlyPaths":["/proc/bus","/proc/fs","/proc/irq","/proc/sys","/proc/sysrq-trigger"]}
> date
Sat Jan 31 13:19:09 -03 2026

 ~                                  





Aí está a “origem do fantasma”: **existe sim um container**, gravado no state do Docker, chamado **`/open-webui`**, imagem `ghcr.io/open-webui/open-webui:main`, com **PortBindings 8080/tcp → HostPort 3000** e **RestartPolicy `"always"`**. Por isso ele ressuscita e recria NAT. (Você viu isso em `/var/lib/docker/containers/.../hostconfig.json`.)

O estranho é ele **não aparecer no `docker ps`**, mas o state mostra **Running=true** e PID **40708** — então o Docker está em estado inconsistente ou você está consultando um daemon diferente (menos provável). De qualquer forma, dá pra eliminar **definitivo**.

## ✅ Opção 1 (preferida): matar pelo Docker (se ele aceitar)

Tente primeiro o caminho “limpo”:

```bash
docker rm -f open-webui 2>/dev/null || docker rm -f dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1
docker volume rm open-webui 2>/dev/null || true
```

E desabilite restart (caso ele não remova de primeira):

```bash
docker update --restart=no open-webui 2>/dev/null || docker update --restart=no dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1
```

Depois confira:

```bash
sudo ss -ltnp | grep ':3000' || echo "porta 3000 livre"
```

Se isso funcionar, acabou.

---

## ✅ Opção 2 (definitiva mesmo se o Docker estiver bugado): remover o container “na unha” com segurança

Como o restart policy é `always`, faça assim:

1. **Pare o Docker**

```bash
sudo systemctl stop docker
```

2. **Backup do container (opcional, mas recomendado)**

```bash
CID=dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1
sudo tar -czf /root/docker-container-$CID-backup.tgz -C /var/lib/docker/containers $CID
```

3. **Remova o diretório do container**

```bash
sudo rm -rf /var/lib/docker/containers/$CID
```

4. **(Opcional) Remova o volume de dados do Open WebUI**

> Isso apaga os dados do WebUI. Como você não usa mais, faz sentido.

```bash
sudo rm -rf /var/lib/docker/volumes/open-webui
```

5. **Suba o Docker**

```bash
sudo systemctl start docker
```

6. **Cheque a porta**

```bash
sudo ss -ltnp | grep ':3000' || echo "porta 3000 livre"
```

---

## ✅ Garantir que o Grafana use a 3000

Assim que a porta estiver livre, suba o Grafana com `3000:3000`.

---

## Por que isso está acontecendo?

Você provou que o container tem:

* `RestartPolicy: "always"` → ele sempre volta quando o Docker sobe
* `PortBindings: HostPort 3000` → o Docker cria `docker-proxy` e regra NAT pra 3000

Então “impedir de subir” = **remover ou alterar esse container** (ou ao menos trocar restart policy pra `no` e remover port binding).

---

## Se quiser um check final (rapidinho)

Depois de iniciar o Docker, rode:

```bash
docker ps -a | grep -i webui || echo "sem webui"
sudo ss -ltnp | grep ':3000' || echo "3000 livre"
```

Se você executar a **Opção 2**, ele não volta mais — e o Grafana fica com a 3000.





> sudo grep -R --line-number '"HostPort":"3000"' /var/lib/docker/containers 2>/dev/null
/var/lib/docker/containers/dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1/config.v2.json:1:{"StreamConfig":{},"State":{"Running":true,"Paused":false,"Restarting":false,"OOMKilled":false,"RemovalInProgress":false,"Dead":false,"Pid":40708,"ExitCode":0,"Error":"","StartedAt":"2026-01-31T15:55:49.832936683Z","FinishedAt":"2026-01-31T15:55:49.936046015Z","Health":{"Status":"starting","FailingStreak":0,"Log":[{"Start":"2026-01-31T12:53:55.24765491-03:00","End":"2026-01-31T12:53:55.298302547-03:00","ExitCode":0,"Output":"true\n"},{"Start":"2026-01-31T12:54:22.577534444-03:00","End":"2026-01-31T12:54:22.627184374-03:00","ExitCode":0,"Output":"true\n"},{"Start":"2026-01-31T12:54:49.910598269-03:00","End":"2026-01-31T12:54:49.960754345-03:00","ExitCode":0,"Output":"true\n"},{"Start":"2026-01-31T12:55:17.238577061-03:00","End":"2026-01-31T12:55:17.288109259-03:00","ExitCode":0,"Output":"true\n"},{"Start":"2026-01-31T12:55:44.562054354-03:00","End":"2026-01-31T12:55:44.612553597-03:00","ExitCode":0,"Output":"true\n"}]}},"ID":"dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1","Created":"2025-08-31T18:38:59.11161328Z","Managed":false,"Path":"bash","Args":["start.sh"],"Config":{"Hostname":"dde7a8080851","Domainname":"","User":"0:0","AttachStdin":false,"AttachStdout":false,"AttachStderr":false,"ExposedPorts":{"8080/tcp":{}},"Tty":false,"OpenStdin":false,"StdinOnce":false,"Env":["OLLAMA_BASE_URL=http://host.docker.internal:11434","PATH=/usr/local/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin","LANG=C.UTF-8","GPG_KEY=A035C8C19219BA821ECEA86B64E628F8D684696D","PYTHON_VERSION=3.11.13","PYTHON_SHA256=8fb5f9fbc7609fa822cb31549884575db7fd9657cbffb89510b5d7975963a83a","ENV=prod","PORT=8080","USE_OLLAMA_DOCKER=false","USE_CUDA_DOCKER=false","USE_SLIM_DOCKER=false","USE_CUDA_DOCKER_VER=cu128","USE_EMBEDDING_MODEL_DOCKER=sentence-transformers/all-MiniLM-L6-v2","USE_RERANKING_MODEL_DOCKER=","OPENAI_API_BASE_URL=","OPENAI_API_KEY=","WEBUI_SECRET_KEY=","SCARF_NO_ANALYTICS=true","DO_NOT_TRACK=true","ANONYMIZED_TELEMETRY=false","WHISPER_MODEL=base","WHISPER_MODEL_DIR=/app/backend/data/cache/whisper/models","RAG_EMBEDDING_MODEL=sentence-transformers/all-MiniLM-L6-v2","RAG_RERANKING_MODEL=","SENTENCE_TRANSFORMERS_HOME=/app/backend/data/cache/embedding/models","TIKTOKEN_ENCODING_NAME=cl100k_base","TIKTOKEN_CACHE_DIR=/app/backend/data/cache/tiktoken","HF_HOME=/app/backend/data/cache/embedding/models","HOME=/root","WEBUI_BUILD_VERSION=2407d9b905978d68619bdce4021e424046ec8df9","DOCKER=true"],"Cmd":["bash","start.sh"],"Healthcheck":{"Test":["CMD-SHELL","curl --silent --fail http://localhost:${PORT:-8080}/health | jq -ne 'input.status == true' || exit 1"]},"Image":"ghcr.io/open-webui/open-webui:main","Volumes":null,"WorkingDir":"/app/backend","Entrypoint":null,"OnBuild":null,"Labels":{"org.opencontainers.image.created":"2025-08-28T10:40:43.547Z","org.opencontainers.image.description":"User-friendly AI Interface (Supports Ollama, OpenAI API, ...)","org.opencontainers.image.licenses":"NOASSERTION","org.opencontainers.image.revision":"2407d9b905978d68619bdce4021e424046ec8df9","org.opencontainers.image.source":"https://github.com/open-webui/open-webui","org.opencontainers.image.title":"open-webui","org.opencontainers.image.url":"https://github.com/open-webui/open-webui","org.opencontainers.image.version":"main"}},"Image":"sha256:1df9c23da52bf54e82590b721f9e091e3e9865271f14cdad364dee4a2d257661","ImageManifest":null,"NetworkSettings":{"Bridge":"","SandboxID":"d33feda521f8a56cebf1ca2aac56de087b93de0494ea80c20654acf86b107ff6","SandboxKey":"/var/run/docker/netns/d33feda521f8","HairpinMode":false,"LinkLocalIPv6Address":"","LinkLocalIPv6PrefixLen":0,"Networks":{"bridge":{"IPAMConfig":null,"Links":null,"Aliases":null,"MacAddress":"02:42:ac:12:00:02","DriverOpts":null,"NetworkID":"a7c0932a2acbbb9e19501a90f34108c290f4de0698eccd17d757d3606be0a8bc","EndpointID":"ff3d49d85b2fa549deb0228b7eb48395408e78ac01b723eb7e0c1791619f1dfa","Gateway":"172.18.0.1","IPAddress":"172.18.0.2","IPPrefixLen":16,"IPv6Gateway":"","GlobalIPv6Address":"","GlobalIPv6PrefixLen":0,"DNSNames":null,"IPAMOperational":false,"DesiredMacAddress":""}},"Service":null,"Ports":{"8080/tcp":[{"HostIp":"0.0.0.0","HostPort":"3000"}]},"SecondaryIPAddresses":null,"SecondaryIPv6Addresses":null,"HasSwarmEndpoint":false},"LogPath":"/var/lib/docker/containers/dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1/dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1-json.log","Name":"/open-webui","Driver":"overlay2","OS":"linux","RestartCount":0,"HasBeenStartedBefore":true,"HasBeenManuallyStopped":false,"MountPoints":{"/app/backend/data":{"Source":"/var/lib/docker/volumes/open-webui/_data","Destination":"/app/backend/data","RW":true,"Name":"open-webui","Driver":"local","Type":"volume","Relabel":"z","ID":"cede6c78f60adc202cb6630ad65d2051352f26ea754bf05057ed12a8a9e8fc7e","Spec":{"Type":"volume","Source":"open-webui","Target":"/app/backend/data"},"SkipMountpointCreation":false}},"SecretReferences":null,"ConfigReferences":null,"MountLabel":"","ProcessLabel":"","AppArmorProfile":"","SeccompProfile":"","NoNewPrivileges":false,"HostnamePath":"/var/lib/docker/containers/dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1/hostname","HostsPath":"/var/lib/docker/containers/dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1/hosts","ShmPath":"","ResolvConfPath":"/var/lib/docker/containers/dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1/resolv.conf","LocalLogCacheMeta":{"HaveNotifyEnabled":false}}
/var/lib/docker/containers/dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1/hostconfig.json:1:{"Binds":["open-webui:/app/backend/data"],"ContainerIDFile":"","LogConfig":{"Type":"json-file","Config":{"max-size":"100m"}},"NetworkMode":"bridge","PortBindings":{"8080/tcp":[{"HostIp":"","HostPort":"3000"}]},"RestartPolicy":{"Name":"always","MaximumRetryCount":0},"AutoRemove":false,"VolumeDriver":"","VolumesFrom":null,"ConsoleSize":[36,120],"CapAdd":null,"CapDrop":null,"CgroupnsMode":"host","Dns":[],"DnsOptions":[],"DnsSearch":[],"ExtraHosts":["host.docker.internal:host-gateway"],"GroupAdd":null,"IpcMode":"private","Cgroup":"","Links":null,"OomScoreAdj":0,"PidMode":"","Privileged":false,"PublishAllPorts":false,"ReadonlyRootfs":false,"SecurityOpt":null,"UTSMode":"","UsernsMode":"","ShmSize":67108864,"Runtime":"runc","Isolation":"","CpuShares":0,"Memory":0,"NanoCpus":0,"CgroupParent":"","BlkioWeight":0,"BlkioWeightDevice":[],"BlkioDeviceReadBps":[],"BlkioDeviceWriteBps":[],"BlkioDeviceReadIOps":[],"BlkioDeviceWriteIOps":[],"CpuPeriod":0,"CpuQuota":0,"CpuRealtimePeriod":0,"CpuRealtimeRuntime":0,"CpusetCpus":"","CpusetMems":"","Devices":[],"DeviceCgroupRules":null,"DeviceRequests":null,"MemoryReservation":0,"MemorySwap":0,"MemorySwappiness":null,"OomKillDisable":false,"PidsLimit":null,"Ulimits":[],"CpuCount":0,"CpuPercent":0,"IOMaximumIOps":0,"IOMaximumBandwidth":0,"MaskedPaths":["/proc/asound","/proc/acpi","/proc/kcore","/proc/keys","/proc/latency_stats","/proc/timer_list","/proc/timer_stats","/proc/sched_debug","/proc/scsi","/sys/firmware","/sys/devices/virtual/powercap"],"ReadonlyPaths":["/proc/bus","/proc/fs","/proc/irq","/proc/sys","/proc/sysrq-trigger"]}
> date
Sat Jan 31 13:19:09 -03 2026
> docker rm -f open-webui 2>/dev/null || docker rm -f dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1
docker volume rm open-webui 2>/dev/null || true

> docker update --restart=no open-webui 2>/dev/null || docker update --restart=no dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1
Error response from daemon: No such container: dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1
> sudo ss -ltnp | grep ':3000' || echo "porta 3000 livre"
porta 3000 livre
> CID=dde7a808085163a0680de9bcf9495927334728320258c66faf0f5fb018a919f1
sudo tar -czf /root/docker-container-$CID-backup.tgz -C /var/lib/docker/containers $CID

> sudo rm -rf /var/lib/docker/containers/$CID
> sudo rm -rf /var/lib/docker/volumes/open-webui

 ~                   





## PENDENTE
- Revisar como subir app de exemplo em Go da pasta "lab-docker-otel-lgtm/examples/go"
- Efetuar um cenário variado diferente.




### - Revisar como subir app de exemplo em Go da pasta "lab-docker-otel-lgtm/examples/go"

- Revisar como subir app de exemplo em Go da pasta "lab-docker-otel-lgtm/examples/go"



Para subir esta aplicação Go no Kubernetes, você precisa seguir estes passos:

## 1. Criar um Dockerfile

````dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o rolldice .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/rolldice .
EXPOSE 8080
CMD ["./rolldice"]
````

## 2. Criar manifestos Kubernetes

````yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: rolldice-config
data:
  OTEL_METRIC_EXPORT_INTERVAL: "5000"
  OTEL_EXPORTER_OTLP_INSECURE: "true"
  OTEL_EXPORTER_OTLP_ENDPOINT: "http://otel-collector:4317"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: rolldice
spec:
  replicas: 1
  selector:
    matchLabels:
      app: rolldice
  template:
    metadata:
      labels:
        app: rolldice
    spec:
      containers:
      - name: rolldice
        image: rolldice:latest
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 8080
        env:
        - name: OTEL_RESOURCE_ATTRIBUTES
          value: "service.name=rolldice,service.instance.id=$(POD_NAME)"
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        envFrom:
        - configMapRef:
            name: rolldice-config
---
apiVersion: v1
kind: Service
metadata:
  name: rolldice
spec:
  selector:
    app: rolldice
  ports:
  - protocol: TCP
    port: 8080
    targetPort: 8080
  type: ClusterIP
````

## 3. Comandos para deploy

```bash
# Build da imagem
docker build -t rolldice:latest .

# Se estiver usando Minikube
minikube image load rolldice:latest

# Aplicar os manifestos
kubectl apply -f k8s-deployment.yaml

# Verificar o status
kubectl get pods
kubectl logs -f deployment/rolldice
```

**Nota**: Ajuste o `OTEL_EXPORTER_OTLP_ENDPOINT` para apontar para o seu OpenTelemetry Collector no cluster.










preciso usar port-forward?


## PENDENTE
- Subir infra
- Subir app
- Fazer port-forward
- Validar app



- Subir infra
- Subir app

> kubectl get pods
NAME                       READY   STATUS    RESTARTS       AGE
lgtm-5f966c7986-kk6gv      1/1     Running   4 (3h6m ago)   14d
rolldice-cdc95cf9b-mt5b7   1/1     Running   1 (3h6m ago)   6h56m
>
>
> kubectl get svc
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                                        AGE
kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP                                        20d
lgtm         ClusterIP   10.97.4.8       <none>        3000/TCP,4040/TCP,4317/TCP,4318/TCP,9090/TCP   14d
rolldice     ClusterIP   10.105.69.199   <none>        8080/TCP                                       6h57m




- Fazer port-forward


kubectl port-forward service/lgtm 3000:3000 4040:4040 4317:4317 4318:4318 9090:9090

kubectl port-forward service/rolldice 8080:8080

http://localhost:8080/rolldice

> kubectl port-forward service/rolldice 8080:8080
Forwarding from 127.0.0.1:8080 -> 8080
Handling connection for 8080
E0131 21:04:57.067070    5067 portforward.go:424] "Unhandled Error" err=<
        an error occurred forwarding 8080 -> 8080: error forwarding port 8080 to pod baeed56b499a28020e9db082110cf9d29b0273c2264e517fe473602896dfb1b7, uid : exit status 1: 2026/02/01 00:04:57 socat[5491] E connect(5, AF=2 127.0.0.1:8080, 16): Connection refused
 >
error: lost connection to pod






- Analisando:

> kubectl get svc rolldice -o yaml
apiVersion: v1
kind: Service
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"v1","kind":"Service","metadata":{"annotations":{},"name":"rolldice","namespace":"default"},"spec":{"ports":[{"port":8080,"protocol":"TCP","targetPort":8080}],"selector":{"app":"rolldice"},"type":"ClusterIP"}}
  creationTimestamp: "2026-01-31T17:05:44Z"
  name: rolldice
  namespace: default
  resourceVersion: "16692"
  uid: 21fb77c9-b087-41fc-be9a-ec1622766bc7
spec:
  clusterIP: 10.105.69.199
  clusterIPs:
  - 10.105.69.199
  internalTrafficPolicy: Cluster
  ipFamilies:
  - IPv4
  ipFamilyPolicy: SingleStack
  ports:
  - port: 8080
    protocol: TCP
    targetPort: 8080
  selector:
    app: rolldice
  sessionAffinity: None
  type: ClusterIP
status:
  loadBalancer: {}
> kubectl get endpoints rolldice -o wide
# ou mais novo:
kubectl get endpointSlice -l kubernetes.io/service-name=rolldice -o wide

Warning: v1 Endpoints is deprecated in v1.33+; use discovery.k8s.io/v1 EndpointSlice
NAME       ENDPOINTS          AGE
rolldice   10.244.0.48:8080   7h1m
zsh: command not found: #
NAME             ADDRESSTYPE   PORTS   ENDPOINTS     AGE
rolldice-7n4xk   IPv4          8080    10.244.0.48   7h1m
> kubectl describe svc rolldice | sed -n '/Selector/,$p' | sed -n '/Endpoints/,+3p'
kubectl get pods -o wide --show-labels | grep -i rolldice

Endpoints:                10.244.0.48:8080
Session Affinity:         None
Internal Traffic Policy:  Cluster
Events:                   <none>
rolldice-cdc95cf9b-mt5b7   1/1     Running   1 (3h12m ago)   7h2m   10.244.0.48   minikube   <none>           <none>            app=rolldice,pod-template-hash=cdc95cf9b
> POD=$(kubectl get pods -l app=rolldice -o jsonpath='{.items[0].metadata.name}')
kubectl exec -it $POD -- sh -lc 'netstat -ltnp 2>/dev/null || ss -ltnp'

sh: 1: ss: not found
command terminated with exit code 127
> kubectl logs $POD --tail=200
2026/02/01 00:01:53 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
2026/02/01 00:02:23 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:02:53 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:03:23 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:03:53 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:04:23 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:04:53 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:05:23 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:05:53 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:06:23 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:06:53 failed to upload metrics: Post "http://otel-collector:4317/v1/metrics": reader collect and export timeout
2026/02/01 00:07:23 failed to upload metrics: Post "http://otel-collector:4317/v1/metrics": reader collect and export timeout
2026/02/01 00:07:53 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:08:23 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
>
> kubectl get svc rolldice -o yaml | sed -n '1,120p'
apiVersion: v1
kind: Service
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"v1","kind":"Service","metadata":{"annotations":{},"name":"rolldice","namespace":"default"},"spec":{"ports":[{"port":8080,"protocol":"TCP","targetPort":8080}],"selector":{"app":"rolldice"},"type":"ClusterIP"}}
  creationTimestamp: "2026-01-31T17:05:44Z"
  name: rolldice
  namespace: default
  resourceVersion: "16692"
  uid: 21fb77c9-b087-41fc-be9a-ec1622766bc7
spec:
  clusterIP: 10.105.69.199
  clusterIPs:
  - 10.105.69.199
  internalTrafficPolicy: Cluster
  ipFamilies:
  - IPv4
  ipFamilyPolicy: SingleStack
  ports:
  - port: 8080
    protocol: TCP
    targetPort: 8080
  selector:
    app: rolldice
  sessionAffinity: None
  type: ClusterIP
status:
  loadBalancer: {}
>
>
>
> kubectl get endpoints rolldice -o wide
Warning: v1 Endpoints is deprecated in v1.33+; use discovery.k8s.io/v1 EndpointSlice
NAME       ENDPOINTS          AGE
rolldice   10.244.0.48:8080   7h3m
>
>
> kubectl get pods -o wide | grep -i rolldice
rolldice-cdc95cf9b-mt5b7   1/1     Running   1 (3h13m ago)   7h3m   10.244.0.48   minikube   <none>           <none>



- No código da aplicação tem trecho:

~~~~go

	// Start HTTP server.
	srv := &http.Server{
		Addr:         ":8081",
~~~~


- Ajustando manifesto k8s


cd /home/fernando/cursos/opentelemetry/lab-docker-otel-lgtm/examples/go
kubectl delete -f k8s-deployment.yaml
kubectl apply -f k8s-deployment.yaml


kubectl port-forward service/rolldice 8081:8081

http://127.0.0.1:8081/rolldice

Go 	
curl http://127.0.0.1:8081/rolldice


> curl http://127.0.0.1:8081/rolldice
3



2026/02/01 00:22:12 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:22:42 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:23:12 failed to upload metrics: Post "http://otel-collector:4317/v1/metrics": reader collect and export timeout
2026/02/01 00:23:42 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:24:02 processor export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/logs": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:24:02 traces export: processor export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/traces": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:24:12 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:24:42 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:25:12 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:25:42 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:26:12 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving
2026/02/01 00:26:42 failed to upload metrics: reader collect and export timeout: retry-able request failure: Post "http://otel-collector:4317/v1/metrics": dial tcp: lookup otel-collector on 10.96.0.10:53: server misbehaving





Show — o HTTP ficou certo. Agora sobre os logs: sua app está tentando enviar OTLP para **`http://otel-collector:4317/...`**, mas **não existe nenhum Service chamado `otel-collector`** no cluster. Você só tem o Service **`lgtm`** que expõe **4317/4318**. Então o ajuste é: **apontar o exporter OTLP da app para `lgtm` (ou criar um Service `otel-collector` apontando para o `lgtm`)**.

Você tem 2 jeitos bons:

---

## Opção A (recomendada): ajustar as env vars da app para usar `lgtm:4317`

No deployment do `rolldice`, configure o endpoint OTLP para o service `lgtm`.

### 1) Confirme que o `lgtm` tem 4317

Já tem: `lgtm ... 4317/TCP, 4318/TCP`.

### 2) Patch no deployment (exemplo genérico)

A maioria dos exemplos Go + OTel respeita **uma** dessas variáveis (depende do sample):

* `OTEL_EXPORTER_OTLP_ENDPOINT`
* `OTEL_EXPORTER_OTLP_TRACES_ENDPOINT`
* `OTEL_EXPORTER_OTLP_METRICS_ENDPOINT`
* `OTEL_EXPORTER_OTLP_LOGS_ENDPOINT`
* `OTEL_EXPORTER_OTLP_PROTOCOL` (grpc/http)

Como seu log mostra **HTTP** (`Post "http://.../v1/metrics"`), ele está usando **OTLP/HTTP** (porta típica **4318**).
Mas ele está apontando pra **4317** (que normalmente é gRPC). Então corrija **endpoint + porta** para bater com o protocolo.

### ✅ Se você quer OTLP/HTTP:

Use **4318**:

```bash
kubectl set env deploy/rolldice \
  OTEL_EXPORTER_OTLP_ENDPOINT=http://lgtm:4318 \
  OTEL_EXPORTER_OTLP_PROTOCOL=http/protobuf
```



- Ajustei no manifesto.

  OTEL_EXPORTER_OTLP_ENDPOINT: "http://lgtm:4318"


cd /home/fernando/cursos/opentelemetry/lab-docker-otel-lgtm/examples/go
kubectl delete -f k8s-deployment.yaml
kubectl apply -f k8s-deployment.yaml


kubectl port-forward service/rolldice 8081:8081

http://127.0.0.1:8081/rolldice

Go 	
curl http://127.0.0.1:8081/rolldice



>
> curl http://127.0.0.1:8081/rolldice
2
> curl http://127.0.0.1:8081/rolldice
6
> curl http://127.0.0.1:8081/rolldice
6
> date
Sat Jan 31 21:31:44 -03 2026




- Pod sem logs de erro:
The selected container has not logged any messages yet.



- No painel do Grafana, verificando o Loki, tem logs do rolldice:

~~~~bash
2026-01-31 21:32:05.155infoRolled a dice: 3
2026-01-31 21:31:43.699infoRolled a dice: 6
2026-01-31 21:31:42.196infoRolled a dice: 6
2026-01-31 21:31:40.173infoRolled a dice: 2
~~~~





----------------------------------------------------------
----------------------------------------------------------
----------------------------------------------------------
----------------------------------------------------------
## PENDENTE

- Material bom: https://medium.com/@ancilartech/the-ultimate-go-observability-cheat-sheet-opentelemetry-edition-9e020eecc747 

- Explorar Lab de aplicação com outras funcionalidades. Adicionar um script que pode ajudar a simular latencias na aplicação em Go "Rolldice" ou criar app que fala com ela.

- Explorar cenário diferente

- Lab do guia https://www.elastic.co/observability-labs/blog/manual-instrumentation-apps-opentelemetry , usando Elastic e instrumentação manual. 
App “Elastiflix” do guia da Elastic, repo: https://github.com/elastic/observability-examples/blob/main/Elastiflix/README.md
Aplicação tem vários microserviços em várias linguagens.
Estrutura pronta é disponível em Docker-compose.
Adaptar ela para Kubernetes.


- Lab para migração de Grafana para ELK e vice-versa.





## Dia 01/02/2026

- Explorar cenário diferente
Lab do guia https://www.elastic.co/observability-labs/blog/manual-instrumentation-apps-opentelemetry , usando Elastic e instrumentação manual. 
App “Elastiflix” do guia da Elastic, repo: https://github.com/elastic/observability-examples/blob/main/Elastiflix/README.md 

lab é extenso, deixar para segundo momento



- Explorar Lab de aplicação com outras funcionalidades. Adicionar um script que pode ajudar a simular latencias na aplicação em Go "Rolldice" ou criar app que fala com ela.






A forma mais útil de “evoluir” esse `rolldice` é transformar o exemplo em um **mini-sistema** (2–4 serviços) com **dependências reais** (HTTP, fila, banco/cache) e com cenários de **latência, erro, retry e fan-out**. Isso gera traces mais interessantes (árvore de spans), métricas acionáveis (RED/USE) e logs correlacionados.

## 1) Evoluir a aplicação atual (mantendo 1 serviço)

### A. Mais rotas com comportamento diferente
Adicionar endpoints que produzam padrões distintos:
- `GET /healthz` e `GET /readyz` (sem instrumentação “pesada”)
- `GET /roll?faces=6&delayMs=100&failRate=0.1`
- `POST /orders` (gera um “workflow” com vários passos/spans)
- `GET /panic` (simula crash/controlado para ver erros e logs)

**Por quê:** você testa `http.route`, status codes, erros, latência e cardinalidade de atributos.

### B. Spans internos (subspans) e eventos
Hoje existe 1 span `roll`. Criar subspans e eventos:
- `validate_request`
- `compute_roll` (com atributos `faces`, `rng`)
- `write_response`
- `span.AddEvent("cache_hit")`, `span.RecordError(err)`

**Por quê:** isso cria “história” no trace, útil em troubleshooting.

### C. Métricas mais “operacionais”
Além do histograma `dice.roll`, adicionar métricas estilo **RED**:
- `http.server.requests` (já pode vir do `otelhttp`, dependendo da config) + **labels/attributes controlados**
- contador `rolldice.requests_total` por resultado (`success/failure`)
- histograma `rolldice.roll.duration_ms`
- gauge/observable: tamanho de fila interna, cache hits, etc.

**Cuidado:** evite atributos de alta cardinalidade (ex.: `user_id`, `request_id`).

### D. Logs estruturados com correlação e níveis
- Logar `trace_id`/`span_id` (o bridge geralmente já correlaciona, mas vale garantir campos úteis)
- Padronizar chaves: `route`, `latency_ms`, `error.type`, `error.msg`

---

## 2) Criar um sistema com múltiplos serviços (recomendado)

### Arquitetura mínima (2 serviços)
1. **api-gateway** (HTTP público)
2. **rolldice** (serviço atual)

Fluxo:
`gateway -> rolldice` via HTTP, propagando contexto.

**O que testar:**
- Propagação W3C (`traceparent`) entre serviços
- Spans client/server (HTTP client instrumentado)
- Erros e retries no client (ex.: timeout)

### Arquitetura “boa de laboratório” (3–4 serviços)
1. **gateway**: recebe requests e inicia o trace
2. **rolldice**: calcula, registra métrica e log
3. **scoreboard**: persiste resultados (Postgres) e expõe `GET /stats`
4. **worker**: consome eventos (Redis Streams / NATS / Kafka) e faz pós-processamento

**O que isso habilita:**
- Traces distribuídos com fan-out (HTTP + mensagem)
- Métricas por dependência (DB, cache, fila)
- Logs do worker correlacionados a um trace originado no gateway (via propagação em mensagens)

---

## 3) Dependências “reais” para enriquecer observabilidade

### A. Banco de dados (Postgres)
- Gravar cada roll (timestamp, resultado, request_id)
- Endpoint `GET /stats` (p95, contagem por face, últimos N rolls)

Instrumentação:
- `database/sql` + instrumentação OTel (ou spans manuais se preferir aprender na raça)

Métricas úteis:
- `db.query.duration`
- `db.pool.in_use`, `db.pool.open_connections`

### B. Cache (Redis)
- Cachear `GET /stats` por 5–10s
- Span/eventos: `cache_hit`, `cache_miss`

### C. Resiliência (timeout/retry/circuit breaker)
- Timeout curto no client HTTP do gateway
- Retry com backoff (e registrar atributos/events)
- Contadores: `retries_total`, `timeouts_total`

---

## 4) Experimentos especificamente bons para OTel

- **Sampling**: configurar head/tail sampling no Collector e comparar o que chega no Tempo
- **Exemplars**: correlacionar métricas com trace (quando suportado)
- **Baggage**: passar `tenant`/`experiment` via baggage e ver nos spans (com cuidado de cardinalidade)
- **Custom propagator** (opcional): entender propagação além do básico
- **SLO/SLI**: medir p95 de `/rolldice` e taxa de erro; criar um dashboard simples

---

## 5) Próximo passo “prático” (sugestão de roadmap)
1. Criar **gateway** chamando `rolldice` (HTTP client instrumentado).
2. Adicionar **Postgres** + endpoint `/stats`.
3. Introduzir **fila** (worker) para processamento assíncrono.
4. Criar **cenários de falha** (delay, erro, timeout) com flags/env vars.
5. Montar um dashboard no Grafana (RED + logs/traces).

Se você quiser, eu preparo um esqueleto de projeto (ex.: `examples/go/gateway`, `examples/go/scoreboard`, manifests K8s e `docker compose`) já com propagação de contexto e alguns spans/métricas padrão.







1. Criar **gateway** chamando `rolldice` (HTTP client instrumentado).

1. Criar **gateway** chamando `rolldice` (HTTP client instrumentado).
1. Criar **gateway** chamando `rolldice` (HTTP client instrumentado).
1. Criar **gateway** chamando `rolldice` (HTTP client instrumentado).
1. Criar **gateway** chamando `rolldice` (HTTP client instrumentado).




## Como rodar (rápido)

### Docker (tudo junto)
```bash
cd /home/fernando/cursos/opentelemetry/lab-docker-otel-lgtm/examples/go
docker compose -f docker-compose.demo.yml up --build
```

- Gateway: `http://localhost:8080/roll`
- Rolldice direto: `http://localhost:8081/rolldice`
- Grafana: `http://localhost:3000` (admin/admin)

### Kubernetes (assumindo que `lgtm` já está no cluster)
1) Build/load das imagens no Minikube:
```bash
eval "$(minikube docker-env)"
cd /home/fernando/cursos/opentelemetry/lab-docker-otel-lgtm/examples/go
docker build -t rolldice:latest .
docker build -t gateway:latest ./gateway
```

2) Apply:
```bash
kubectl apply -f k8s-deployment.yaml
kubectl apply -f k8s-gateway.yaml
```

3) Port-forward:
```bash
kubectl port-forward service/lgtm 3000:3000 4318:4318
kubectl port-forward service/gateway 8080:8080
```

4) Tráfego:
```bash
chmod +x generate-traffic.sh
./generate-traffic.sh
```

---







> docker build -t gateway:latest ./gateway
[+] Building 9.1s (10/10) FINISHED                                                                                                                                                                                                                                       docker:default
 => [internal] load build definition from Dockerfile                                                                                                                                                                                                                               0.0s
 => => transferring dockerfile: 207B                                                                                                                                                                                                                                               0.0s
 => [internal] load metadata for docker.io/library/golang:1.25                                                                                                                                                                                                                     0.7s
 => [internal] load .dockerignore                                                                                                                                                                                                                                                  0.0s
 => => transferring context: 2B                                                                                                                                                                                                                                                    0.0s
 => [1/6] FROM docker.io/library/golang:1.25@sha256:ce63a16e0f7063787ebb4eb28e72d477b00b4726f79874b3205a965ffd797ab2                                                                                                                                                               0.0s
 => => resolve docker.io/library/golang:1.25@sha256:ce63a16e0f7063787ebb4eb28e72d477b00b4726f79874b3205a965ffd797ab2                                                                                                                                                               0.0s
 => [internal] load build context                                                                                                                                                                                                                                                  0.0s
 => => transferring context: 2.55kB                                                                                                                                                                                                                                                0.0s
 => CACHED [2/6] WORKDIR /app                                                                                                                                                                                                                                                      0.0s
 => CACHED [3/6] COPY go.mod go.sum ./                                                                                                                                                                                                                                             0.0s
 => CACHED [4/6] RUN go mod download                                                                                                                                                                                                                                               0.0s
 => [5/6] COPY *.go ./                                                                                                                                                                                                                                                             0.0s
 => ERROR [6/6] RUN CGO_ENABLED=0 GOOS=linux go build -o /gateway                                                                                                                                                                                                                  8.3s
------
 > [6/6] RUN CGO_ENABLED=0 GOOS=linux go build -o /gateway:
8.036 # github.com/fernando/lab-docker-otel-lgtm/examples/go/gateway
8.036 ./otel.go:43:2: declared and not used: handleErr
------
Dockerfile:10
--------------------
   8 |     COPY *.go ./
   9 |
  10 | >>> RUN CGO_ENABLED=0 GOOS=linux go build -o /gateway
  11 |
  12 |     EXPOSE 8080
--------------------
ERROR: failed to solve: process "/bin/sh -c CGO_ENABLED=0 GOOS=linux go build -o /gateway" did not complete successfully: exit code: 1










----------------------------------------------------------
----------------------------------------------------------
----------------------------------------------------------
----------------------------------------------------------
## PENDENTE

- TSHOOT de erro no build do Gateway.
~~~~bash
8.036 # github.com/fernando/lab-docker-otel-lgtm/examples/go/gateway
8.036 ./otel.go:43:2: declared and not used: handleErr
ERROR: failed to solve: process "/bin/sh -c CGO_ENABLED=0 GOOS=linux go build -o /gateway" did not complete successfully: exit code: 1
~~~~

- Aplicar o Roadmap e Experimentos:

## 4) Experimentos especificamente bons para OTel

- **Sampling**: configurar head/tail sampling no Collector e comparar o que chega no Tempo
- **Exemplars**: correlacionar métricas com trace (quando suportado)
- **Baggage**: passar `tenant`/`experiment` via baggage e ver nos spans (com cuidado de cardinalidade)
- **Custom propagator** (opcional): entender propagação além do básico
- **SLO/SLI**: medir p95 de `/rolldice` e taxa de erro; criar um dashboard simples

---

## 5) Próximo passo “prático” (sugestão de roadmap)
1. Criar **gateway** chamando `rolldice` (HTTP client instrumentado).
2. Adicionar **Postgres** + endpoint `/stats`.
3. Introduzir **fila** (worker) para processamento assíncrono.
4. Criar **cenários de falha** (delay, erro, timeout) com flags/env vars.
5. Montar um dashboard no Grafana (RED + logs/traces).


- Explorar Lab de aplicação com outras funcionalidades. Adicionar um script que pode ajudar a simular latencias na aplicação em Go "Rolldice" ou criar app que fala com ela.

- Material bom: https://medium.com/@ancilartech/the-ultimate-go-observability-cheat-sheet-opentelemetry-edition-9e020eecc747 

- Explorar cenário diferente

- Lab do guia https://www.elastic.co/observability-labs/blog/manual-instrumentation-apps-opentelemetry , usando Elastic e instrumentação manual. 
App “Elastiflix” do guia da Elastic, repo: https://github.com/elastic/observability-examples/blob/main/Elastiflix/README.md
Aplicação tem vários microserviços em várias linguagens.
Estrutura pronta é disponível em Docker-compose.
Adaptar ela para Kubernetes.


- Lab para migração de Grafana para ELK e vice-versa.
