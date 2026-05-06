# Auto-Scaling URL Shortener 🚀

A cloud-native URL shortener built using Go, Docker, Kubernetes, Prometheus, Grafana, GitHub Actions, and Terraform.

This project demonstrates:

* Microservice deployment with Kubernetes
* Containerization with Docker
* Monitoring using Prometheus + Grafana
* CI/CD concepts with GitHub Actions
* Infrastructure as Code using Terraform
* Service discovery and observability in Kubernetes

---

# 📌 Features

* Shorten long URLs
* Redirect shortened URLs
* Kubernetes-based deployment
* Dockerized backend + frontend
* Prometheus metrics endpoint
* Grafana dashboards
* Kubernetes Services and networking
* NodePort exposure
* Monitoring stack integration

---

# 🏗️ Architecture

```text
                        +-------------------+
                        |     Browser       |
                        +---------+---------+
                                  |
                                  v
                    +--------------------------+
                    | url-frontend-service     |
                    | (NodePort Service)       |
                    +------------+-------------+
                                 |
                                 v
                      +--------------------+
                      | Frontend UI        |
                      | HTML/CSS/JS        |
                      +---------+----------+
                                |
                                v
                  +-----------------------------+
                  | url-shortener-service       |
                  | Kubernetes NodePort Service |
                  +--------------+--------------+
                                 |
                                 v
                    +------------------------+
                    | Go URL Shortener API   |
                    | /shorten               |
                    | /r/{code}              |
                    | /metrics               |
                    +------------+-----------+
                                 |
                                 v
                     +----------------------+
                     | In-Memory Storage    |
                     +----------------------+

Monitoring Stack:
-----------------
Prometheus ---> Scrapes /metrics
Grafana -----> Visualizes metrics
```

---

# 🛠️ Tech Stack

| Technology     | Purpose                  |
| -------------- | ------------------------ |
| Go             | Backend API              |
| Docker         | Containerization         |
| Kubernetes     | Orchestration            |
| Prometheus     | Metrics collection       |
| Grafana        | Visualization            |
| GitHub Actions | CI/CD                    |
| Terraform      | Infrastructure as Code   |
| Minikube       | Local Kubernetes Cluster |

---

# 📂 Project Structure

```text
auto-scale-URL/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── handler/
│   │   └── handler.go
│   ├── service/
│   │   └── url_service.go
│   └── storage/
│       └── memory.go
│
├── pkg/
│   └── utils/
│       └── generator.go
│
├── frontend/
│   ├── index.html
│   ├── Dockerfile
│
├── k8s/
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── frontend-deployment.yaml
│   ├── frontend-service.yaml
│   └── servicemonitor.yaml
│
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

---

# 🚀 Running the Project

# 1️⃣ Clone Repository

```bash
git clone https://github.com/YOUR_USERNAME/auto-scale-URL.git

cd auto-scale-URL
```

---

# 2️⃣ Start Minikube

```bash
minikube start
```

Verify:

```bash
kubectl get nodes
```

---

# 3️⃣ Build Backend Docker Image

```bash
docker build -t YOUR_USERNAME/url-shortener .
```

Push image:

```bash
docker push YOUR_USERNAME/url-shortener
```

---

# 4️⃣ Build Frontend Docker Image

```bash
cd frontend

docker build -t YOUR_USERNAME/url-frontend .

docker push YOUR_USERNAME/url-frontend

cd ..
```

---

# 5️⃣ Deploy Backend

```bash
kubectl apply -f k8s/deployment.yaml

kubectl apply -f k8s/service.yaml
```

---

# 6️⃣ Deploy Frontend

```bash
kubectl apply -f k8s/frontend-deployment.yaml

kubectl apply -f k8s/frontend-service.yaml
```

---

# 7️⃣ Verify Pods

```bash
kubectl get pods
```

Expected:

```text
url-shortener-xxxxx   Running
url-frontend-xxxxx    Running
```

---

# 8️⃣ Access Frontend

Get Minikube IP:

```bash
minikube ip
```

Open:

```text
http://MINIKUBE_IP:30008
```

---

# 📊 Monitoring Setup

# 1️⃣ Install Helm

```bash
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```

Verify:

```bash
helm version
```

---

# 2️⃣ Add Prometheus Helm Repo

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

helm repo update
```

---

# 3️⃣ Install Prometheus + Grafana

```bash
helm install monitoring prometheus-community/kube-prometheus-stack
```

---

# 4️⃣ Verify Monitoring Pods

```bash
kubectl get pods
```

Expected:

```text
monitoring-grafana
monitoring-kube-prometheus-prometheus
```

---

# 5️⃣ Access Grafana

```bash
kubectl port-forward svc/monitoring-grafana 3000:80
```

Open:

```text
http://localhost:3000
```

---

# 6️⃣ Get Grafana Password

```bash
kubectl get secret monitoring-grafana -o jsonpath="{.data.admin-password}" | base64 -d
```

Login:

```text
Username: admin
Password: <output>
```

---

# 7️⃣ Access Prometheus

```bash
kubectl port-forward svc/monitoring-kube-prometheus-prometheus 9090:9090
```

Open:

```text
http://localhost:9090
```

---

# 📈 Prometheus ServiceMonitor

`k8s/servicemonitor.yaml`

```yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: url-shortener-monitor
  labels:
    release: monitoring
spec:
  selector:
    matchLabels:
      app: url-shortener
  endpoints:
    - port: http
      path: /metrics
      interval: 10s
```

Apply:

```bash
kubectl apply -f k8s/servicemonitor.yaml
```

---

# 🔍 Verify Metrics

Prometheus:

```text
Status -> Targets
```

Expected:

```text
url-shortener-monitor -> UP
```

---

# 📊 Example Grafana Queries

Requests per second:

```promql
rate(http_requests_total[1m])
```

CPU usage:

```promql
sum(rate(container_cpu_usage_seconds_total[1m])) by (pod)
```

Memory usage:

```promql
container_memory_usage_bytes
```

---

# 🧪 Testing Backend

Test shortening:

```bash
curl "http://localhost:8081/shorten?url=https://google.com"
```

Test redirect:

```bash
curl -v "http://localhost:8081/r/abc123"
```

---

# 🐳 Docker Commands

Build backend:

```bash
docker build -t YOUR_USERNAME/url-shortener .
```

Run locally:

```bash
docker run -p 8081:8081 YOUR_USERNAME/url-shortener
```

---

# ☸️ Useful Kubernetes Commands

Get pods:

```bash
kubectl get pods
```

Get services:

```bash
kubectl get svc
```

View logs:

```bash
kubectl logs deployment/url-shortener
```

Restart deployment:

```bash
kubectl rollout restart deployment url-shortener
```

Delete all:

```bash
kubectl delete all --all
```

---

# 🔥 Future Improvements

* Redis integration
* Persistent database
* JWT authentication
* Ingress controller
* HTTPS/TLS
* Custom Grafana dashboards
* Rate limiting
* Helm charts
* Cloud deployment (AWS EKS/GKE)

---

# 👨‍💻 Author

Built by Raunak Rose 🚀

---

# 📜 License

MIT License
