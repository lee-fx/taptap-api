#! /bin/bash

cp -R ./templates ./bin/

makdir ./bin/videos

cd bin

nohup ./api &
nohup ./schduler &
nohup ./streamserver &

echo "deploy finished"