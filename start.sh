#!/bin/bash
export SSR_HOME="/home/ec2-user/simpleskiresort"
cd /home/ec2-user/simpleskiresort/
sudo  make run 
sudo chmod a+x /home/ec2-user/simpleskiresort/bin/simpleskiresort
/home/ec2-user/simpleskiresort/bin/simpleskiresort > /dev/null 2> /dev/null < /dev/null &