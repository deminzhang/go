
wget https://dl.google.com/go/go1.12.6.linux-amd64.tar.gz

tar -C /usr/local -xzf go1.12.6.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

export GOPATH=/data/app/slg_dev/Server
export PATH=$PATH:/data/app/slg_dev/Server/bin
