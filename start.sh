#!/bin/bash
cd ${SSR_HOME}
sudo  make run 
sudo chmod a+x ${SSR_HOME}/bin/simpleskiresort
${SSR_HOME}/bin/simpleskiresort > /dev/null 2> /dev/null < /dev/null &