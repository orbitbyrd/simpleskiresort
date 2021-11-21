#!/bin/bash
sudo go build -o /home/ec2-user/simpleskiresort/bin/main /home/ec2-user/simpleskiresort/main.go 
chmod a+x /home/ec2-user/simpleskiresort/bin/main
/home/ec2-user/simpleskiresort/bin/main