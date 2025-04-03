# Running ScyllaDB on docker for development
sudo docker run --name dev-Scylla -d -p 9042:9042 scylladb/scylla --smp 1 --broadcast-address 127.0.0.1 --listen-address 0.0.0.0 --broadcast-rpc-address 127.0.0.1
