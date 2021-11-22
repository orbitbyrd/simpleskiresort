#!/bin/bash
cd /home/ec2-user/simpleskiresort/
sudo go build -o bin/simpleskiresort main.go 
sudo chmod a+x /home/ec2-user/simpleskiresort/bin/simpleskiresort
/home/ec2-user/simpleskiresort/bin/simpleskiresort > /dev/null 2> /dev/null < /dev/null &