build:
	bee pack -be GOOS=linux -be GOARCH=amd64 -a prototype
deploy:
	bee pack -be GOOS=linux -be GOARCH=amd64 -a prototype
	scp prototype.tar.gz root@192.168.5.5:/aichenk/service/prototype
	echo 'restart'
	ssh root@192.168.5.5 'bash -s' < restart.sh
clean:
	rm -f prototype prototype.tar.gz lastupdate.tmp