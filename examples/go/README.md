# Go + OpenTelemetry — `rolldice`

Exemplo de aplicação **Go** instrumentada com **OpenTelemetry** (traces, métricas e logs) para enviar telemetria via **OTLP/HTTP** para um collector/stack (ex.: **Grafana LGTM**).

---

## O que esta aplicação faz

- Sobe um **servidor HTTP** na porta **`8081`**
- Expõe o endpoint:
  - `GET /rolldice` → retorna um número `1..6`
- Simula uma operação “demorada” no `roll()` com **busy-wait ~1s** (útil para ver spans mais claros em traces/flamegraph)

### Instrumentação OTel no código

- **Traces**
  - Cria um span manual `roll` no handler (`tracer.Start(...)`)
  - Instrumenta HTTP automaticamente com `otelhttp.NewHandler(...)`
  - Define `http.route` usando `otelhttp.WithRouteTag("/rolldice", ...)`
- **Métricas**
  - Registra um histograma `dice.roll` com o valor do dado (1..6)
- **Logs**
  - Usa `slog` + bridge `otelslog` para logs com contexto (correlação com trace)
  - Exporta logs via **OTLP/HTTP**
- **Runtime metrics**
  - Habilita `go.opentelemetry.io/contrib/instrumentation/runtime`

---

## Requisitos

- Go instalado (veja o `Dockerfile` e/ou seu ambiente local)
- Para ver telemetria: um endpoint OTLP/HTTP acessível (ex.: collector, Grafana LGTM)
  - Por padrão, os exporters OTLP/HTTP do Go usam `localhost:4318` se você não configurar o endpoint.
  - Neste exemplo, Docker/K8s já configuram `OTEL_EXPORTER_OTLP_ENDPOINT` para `http://lgtm:4318`.

---

## Rodando local (recomendado: `run.sh`)

O script `run.sh` já configura variáveis úteis:

- reduz o intervalo de exportação de métricas para **5s**
- seta `OTEL_RESOURCE_ATTRIBUTES` (service name e instance id)
- força **OTLP insecure** (HTTP), necessário para o exporter OTLP/HTTP

````bash
cd /home/fernando/cursos/opentelemetry/lab-docker-otel-lgtm/examples/go
chmod +x run.sh
./run.sh
````

Em outro terminal:

````bash
curl -s http://localhost:8081/rolldice
````

Gerando carga:

````bash
for i in $(seq 1 20); do curl -s http://localhost:8081/rolldice; done
````

### Ajustando o destino OTLP localmente

Se seu collector não está em `localhost:4318`, configure:

````bash
export OTEL_EXPORTER_OTLP_ENDPOINT="http://SEU_COLLECTOR:4318"
export OTEL_EXPORTER_OTLP_INSECURE="true"
./run.sh
````

---

## Rodando com Docker

### Build da imagem

````bash
cd /home/fernando/cursos/opentelemetry/lab-docker-otel-lgtm/examples/go
docker build -t rolldice:latest .
````

### Run (apontando para um collector OTLP/HTTP)

````bash
docker run --rm -p 8081:8081 \
  -e OTEL_EXPORTER_OTLP_ENDPOINT="http://host.docker.internal:4318" \
  -e OTEL_EXPORTER_OTLP_INSECURE="true" \
  -e OTEL_METRIC_EXPORT_INTERVAL="5000" \
  -e OTEL_RESOURCE_ATTRIBUTES="service.name=rolldice,service.instance.id=docker:8081" \
  rolldice:latest
````

> Em Linux, `host.docker.internal` pode não existir dependendo da versão/configuração do Docker. Alternativas comuns:
> - usar o IP do host na interface bridge
> - rodar o collector no mesmo `docker network` e usar o nome do container/serviço

---

## Rodando com Docker Compose (arquivo `docker-compose.oats.yml`)

Este compose foi feito para rodar o container **go** apontando para `http://lgtm:4318` e expor a aplicação em **`localhost:8080`** (mapeando para `8081` no container).

````bash
cd /home/fernando/cursos/opentelemetry/lab-docker-otel-lgtm/examples/go
docker compose -f docker-compose.oats.yml up --build
````

Teste:

````bash
curl -s http://localhost:8080/rolldice
````

Parar:

````bash
docker compose -f docker-compose.oats.yml down
````

> Observação: esse compose **não** define o serviço `lgtm` aqui. Você precisa subir a stack LGTM/collector no mesmo network (ou ajustar `OTEL_EXPORTER_OTLP_ENDPOINT`).

---

## Kubernetes: subindo a infra LGTM + port-forward (Minikube)

Quando você sobe o LGTM no Kubernetes, normalmente ele fica como **Service `ClusterIP`**. Isso significa que, **do seu notebook**, você **não acessa diretamente** sem um:
- `kubectl port-forward` (mais simples para laboratório), ou
- `minikube service ... --url`, ou
- Ingress/LoadBalancer (mais “real”, mas fora do foco do exemplo)

### Comandos úteis no Minikube

````bash
minikube start --cpus=4 --memory=8192 --driver=docker
````

Dashboard:

````bash
minikube dashboard
````

Tunnel (útil se você estiver testando Services do tipo LoadBalancer):

````bash
minikube tunnel
````

Gerar uma URL acessível para um Service (porta aleatória, depende do serviço/tipo):

````bash
minikube service grafana --url
minikube service lgtm --url
````

### Subir LGTM no Kubernetes

A infra LGTM do repositório fica no nível acima deste exemplo.

````bash
cd /home/fernando/cursos/opentelemetry/lab-docker-otel-lgtm
kubectl apply -f k8s/lgtm.yaml
````

Aguarde o Pod ficar **Running** antes do port-forward (se tentar antes, pode aparecer: `pod is not running. Current status=Pending`):

````bash
kubectl get pods
kubectl get pods -w
````

### Port-forward do LGTM (Grafana + OTLP + etc.)

Mantenha este comando rodando em um terminal:

````bash
kubectl port-forward service/lgtm 3000:3000 4040:4040 4317:4317 4318:4318 9090:9090
````

- Grafana: <http://localhost:3000> (default: `admin` / `admin`)
- OTLP/gRPC: `localhost:4317`
- OTLP/HTTP: `localhost:4318`

> Dica: se a porta `3000` estiver ocupada, mude apenas o lado esquerdo. Ex.: `3001:3000`.

---

## Rodando no Kubernetes (manifest `k8s-deployment.yaml`)

O manifesto cria:

- `ConfigMap` com:
  - `OTEL_EXPORTER_OTLP_ENDPOINT=http://lgtm:4318`
  - `OTEL_EXPORTER_OTLP_INSECURE=true`
  - `OTEL_METRIC_EXPORT_INTERVAL=5000`
- `Deployment` `rolldice` (1 réplica)
- `Service` `rolldice` (ClusterIP na porta 8081)

### 1) Construir a imagem para o cluster

Você precisa que o cluster consiga “ver” a imagem `rolldice:latest`.

- **Minikube**:
  ````bash
  eval "$(minikube docker-env)"
  docker build -t rolldice:latest .
  ````

- **Kind** (alternativa):
  ````bash
  docker build -t rolldice:latest .
  kind load docker-image rolldice:latest
  ````

### 2) Aplicar no cluster

````bash
cd /home/fernando/cursos/opentelemetry/lab-docker-otel-lgtm/examples/go
kubectl apply -f k8s-deployment.yaml
````

### 3) Acessar a aplicação (port-forward)

Como o Service é `ClusterIP`, use port-forward para testar localmente:

````bash
kubectl port-forward svc/rolldice 8081:8081
````

Teste:

````bash
curl -s http://localhost:8081/rolldice
````

### 4) Logs e troubleshooting rápido

````bash
kubectl get pods -l app=rolldice
kubectl logs -l app=rolldice -f
````

Se o `port-forward` retornar `Connection refused`, verifique se o Service/targetPort bate com a porta real do app (a app escuta em `:8081`).

---

## Variáveis de ambiente OTel usadas aqui

- `OTEL_EXPORTER_OTLP_ENDPOINT`  
  Endpoint OTLP/HTTP (ex.: `http://lgtm:4318`)
- `OTEL_EXPORTER_OTLP_INSECURE=true`  
  Força HTTP (sem TLS) para exporters OTLP/HTTP
- `OTEL_METRIC_EXPORT_INTERVAL=5000`  
  Intervalo (ms) para exportação periódica de métricas (default costuma ser 60000ms)
- `OTEL_RESOURCE_ATTRIBUTES`  
  Define atributos do recurso, usado aqui para:
  - `service.name=rolldice`
  - `service.instance.id=...` (no K8s usa o nome do Pod)

---

## O que procurar no Grafana/Tempo/Mimir/Loki (LGTM)

Depois de gerar tráfego em `/rolldice`, procure por:

- **Traces**: service `rolldice`, span `roll`
- **Métricas**: instrumento `dice.roll` (histograma)
- **Logs**: logs do serviço `rolldice` correlacionados ao trace (via contexto)

---

## Notas de implementação (para estudo)

- O servidor usa `signal.NotifyContext` para shutdown gracioso com `CTRL+C`
- `otelhttp.NewHandler(mux, "/")` instrumenta todas as rotas do mux
- `otelhttp.WithRouteTag("/rolldice", ...)` ajuda a preencher `http.route` corretamente
- `runtime.Start(...)` habilita métricas do runtime Go

---
