run:
	docker-compose up -d

build:
	docker-compose build

go_install:
	docker-compose exec cao go mod vendor

yarn_install:
	docker-compose exec frontend yarn

gen_ssl:
	openssl dhparam -out nginx/dhparam.pem 2048
