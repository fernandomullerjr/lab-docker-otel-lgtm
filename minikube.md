
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

### 🔹 Ativar o tunnel

```bash
minikube tunnel
```

🎉 Agora o LoadBalancer funciona.

📌 **Importante**

* O `minikube tunnel` precisa ficar rodando
* Ele cria rotas de rede temporárias