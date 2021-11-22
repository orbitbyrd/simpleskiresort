#!/bin/bash
cd /home/ec2-user/simpleskiresort/
sudo go build -o bin/main main.go 
chmod a+x /home/ec2-user/simpleskiresort/bin/main
/home/ec2-user/simpleskiresort/bin/main