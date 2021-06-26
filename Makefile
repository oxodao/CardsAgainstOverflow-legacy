run:
	docker-compose up -d

build:
	docker-compose build

go_install:
	docker-compose exec cao go mod vendor

yarn_install:
	docker-compose exec frontend yarn

gen_ssl:
	#mkcert stuff
	openssl dhparam -out nginx/dhparam.pem 2048
	echo "SSL dhparam certificate generated. I've not done the SSL certificate yet. To do so, just use mkcert and generate one, put files at nginx/{key,key-pub}.pem
