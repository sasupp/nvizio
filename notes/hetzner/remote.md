# Scylla Benchmark
## Tool
https://github.com/scylladb/scylla-bench

scylla-bench -workload sequential -mode write -nodes 10.0.0.3


# Setup dev cluser
## Scylla
sudo docker run --name dev-scylla --hostname dev-scylla -d scylladb/scylla --smp 1

## Install 7z
sudo apt install p7zip-full -y

## Install lzma
sudo apt install lzma

## Rsync sync files
rsync -az feeds sadmin@192.168.1.21:~/sync --delete --exclude '__pycache__' --exclude '.idea' --exclude 'cache' --exclude 'venv'
rsync -az lib sadmin@192.168.1.21:~/sync --delete --exclude '__pycache__' --exclude '.idea' --exclude 'cache' --exclude 'venv'
rsync -az xbrl sadmin@192.168.1.21:~/sync --delete --exclude '__pycache__' --exclude '.idea' --exclude 'cache' --exclude 'venv'
rsync -az download_feed_files.py sadmin@192.168.1.21:~/sync --delete --exclude '__pycache__' --exclude '.idea' --exclude 'cache' --exclude 'venv'
rsync -az feeder.py sadmin@192.168.1.21:~/sync --delete --exclude '__pycache__' --exclude '.idea' --exclude 'cache' --exclude 'venv'
rsync -az process_filings.py sadmin@192.168.1.21:~/sync --delete --exclude '__pycache__' --exclude '.idea' --exclude 'cache' --exclude 'venv'
rsync -az main.py sadmin@192.168.1.21:~/sync --delete --exclude '__pycache__' --exclude '.idea' --exclude 'cache' --exclude 'venv'
rsync -az compress.py sadmin@192.168.1.21:~/sync --delete --exclude '__pycache__' --exclude '.idea' --exclude 'cache' --exclude 'venv'
rsync -az ./main sadmin@192.168.1.21:~/sync
rsync -az download_filings.py sadmin@192.168.1.21:~/sync --delete --exclude '__pycache__' --exclude '.idea' --exclude 'cache' --exclude 'venv'
rsync -az house-keep.py sadmin@192.168.1.21:~/sync --delete --exclude '__pycache__' --exclude '.idea' --exclude 'cache' --exclude 'venv'


## Install pip modules
sudo apt install python3-pip -y
pip3 install xmltodict
pip3 install cassandra-driver
pip3 install python-dateutil

## Install sqlite3
apt install sqlite3 -y

## cqlsh
snap install cqlsh -y

## unzip all
find . -name "*.zip" | while read filename; do unzip -o -d "`dirname "$filename"`" "$filename"; done;
# remove all zip files
find . -name "*.zip" | xargs rm

# Install node.js
curl -fsSL https://deb.nodesource.com/setup_20.x | bash - &&\
apt-get install -y nodejs

# Install pnpm
sudo apt install pnpm -y