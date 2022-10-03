#!/bin/bash
# for i in {1..5}
# do
#    echo $(curl localhost:9000)
#    sleep 2
# done

while true
do
   echo $(curl localhost:9000)
   sleep 2
done