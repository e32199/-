cd /tmp/
wget https://go.dev/dl/go1.20.3.linux-amd64.tar.gz

rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz

vim /etc/profile
export PATH=$PATH:/usr/local/go/bin

go version

mkdir -pv HOME/Pictures/codesty/go
mkdir -pv HOME/Pictures/devops/go

vim ~/.bashrc
export GOPATH=$HOME/Pictures/codesty/go

source ~/.bashrc

echo $GOPATH
