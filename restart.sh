#! /bin/bash
#默认进入的是登录用户的目录
cd /aichenk/service/prototype
tar -xzvf prototype.tar.gz
#remove conf of dev
rm -rf conf/app.conf
cp conf/app.conf.bak conf/app.conf
supervisorctl restart prototype