build:
	docker build --build-arg GITHUB_USER=${TR_GIT_USER} --build-arg GITHUB_TOKEN=${TR_GIT_TOKEN} -t github.com/turistikrota/service.listing . 

run:
	docker service create --name listing-api-turistikrota-com --network turistikrota --secret jwt_private_key --secret jwt_public_key --env-file .env --publish 6023:6023 --publish 7023:7023 github.com/turistikrota/service.listing:latest

remove:
	docker service rm listing-api-turistikrota-com

stop:
	docker service scale listing-api-turistikrota-com=0

start:
	docker service scale listing-api-turistikrota-com=1

restart: remove build run
	