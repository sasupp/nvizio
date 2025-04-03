# Infrastrucre

# Initialization
## Update
apt update
apt upgrade

## Firewall
apt install ufw
ufw allow OpenSSH

## SSH hardening
https://community.hetzner.com/tutorials/securing-ssh#step-11---deactivate-the-root-login

Use a new root user instead of root

adduser sasupp
usermod -aG sudo sasupp

PermitRootLogin no
ClientAliveInterval 300
ClientAliveCountMax 1
AllowUsers sasupp

### Remove password for sudo
sudo visudo
add this line at end: sasupp ALL=(ALL) NOPASSWD: ALL

## scylla ports
