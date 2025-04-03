# Infrastructure

## Load Balancer
HA Proxy deployed on lowest ARM cloud server
Alternatives
Traefik: It is more suitable for service mesh with lots of micro services

## Key Value Store
etcd
This is needed to store infra wide key values as configuration
Alternatives
consul: It is more suitable for service mesh with lots of micro services

## Monitoring, Alerting and Dashboard
Prometheus, Grafana and Node Exporter

## Deployment automation and Inventory Management
Ansible: Very simple and SSH based deployment. Basically, a better alternative to scp
