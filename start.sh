#!/bin/bash
cd /home/ec2-user/simpleskiresort/
sudo  make run 
sudo chmod a+x /home/ec2-user/simpleskiresort/bin/simpleskiresort
sudo /home/ec2-user/simpleskiresort/bin/simpleskiresort > /dev/null 2> /dev/null < /dev/null &