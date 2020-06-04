


cspider:
	cd spider && go build -o spider_svc
	cd spider && docker build . -t spider:test
cweb:
	cd webserver && go build -o web
	cd webserver && docker build . -t web:test

call:cspider cweb


clean:
	find . -name app | xargs rm -rf
	find . -name web | xargs rm -rf
	find . -name spider_svc | xargs rm -rf

start:

stop:
