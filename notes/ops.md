# Update go modules
go get -u
go mod tidy

# Python Deps Install
pip install requests
pip install numpy
pip install xmltodict
pip install cassandra-driver
pip install python-dateutil
pip install lz4
pip install geomet
pip install numpy

# rsync contents to local server
rsync -avz --exclude '.git' -e ssh /home/sasupp/code/stock-api/ sadmin@192.168.1.21:/home/sadmin/nvizio/

export PYTHONPATH=$PYTHONPATH:$PWD
python3 app/nse/feeder.py --feed --scylla 172.17.0.4

# cronjob for nse feeder
*/15 * * * * PYTHONPATH=/home/sadmin/nvizio/filings /usr/bin/python3 /home/sadmin/nvizio/filings/app/nse/feeder.py --feed --scylla 172.17.0.4 >> /home/sadmin/logs/nse_feed.log 2>&1

# persistent scylla
sudo docker run -p 9042:9042 -d --name dev-scylla -v scylla-data:/var/lib/scylla scylladb/scylla