
tag=test

cspider:
	cd spider && go build -o spider_svc
	cd spider && docker build . -t unistd1999/spider:${tag}
cweb:
	cd webserver && go build -o web
	cd webserver && docker build . -t unistd1999/web:${tag}

call:cspider cweb

push:
	docker push unistd1999/spider
	docker push unistd1999/web

rm:
	docker-compose stop
	docker-compose rm
up:
	docker-compose up -d

rebuild: rm call up

clean:
	find . -name app | xargs rm -rf
	find . -name web | xargs rm -rf
	find . -name spider_svc | xargs rm -rf

start:

stop:
