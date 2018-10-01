#!/bin/bash

ssh -t kriptokuna@anonutopia 'export PATH=$PATH:/usr/local/go/bin && 
export GOPATH=$GOPATH:/home/kriptokuna/.go && 
cd build/anonutopia && git pull && 
go get && 
go build && 
cd && 
mkdir anonutopianew && 
cp -R anonutopia/config.json anonutopianew/ &&
cp -R build/anonutopia/templates/ anonutopianew/ &&
cp -R build/anonutopia/public/ anonutopianew/ &&
cp -R build/anonutopia/conf/ anonutopianew/ &&
cp -R build/anonutopia/anonutopia anonutopianew/ && 
rm -rf anonutopia && 
mv anonutopianew anonutopia && 
sudo supervisorctl reload'
