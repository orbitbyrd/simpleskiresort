#!/bin/bash
cd /home/ec2-user/simpleskiresort/
sudo go build -o bin/main main.go 
sudo chmod a+x /home/ec2-user/simpleskiresort/bin/main
/home/ec2-user/simpleskiresort/bin/main > /dev/null 2> /dev/null < /dev/null &