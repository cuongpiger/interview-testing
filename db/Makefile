build_db:
	docker build --no-cache -t cylo-mysql .

run_db:
	docker run -d -p 3306:3306 --name cylo_db cylo-mysql

exec_db:
	docker exec -it cylo_db mysql -u root -p

build_db_postgre:
	docker build --no-cache -t cylo-postgre .

run_db_postgre:
	docker run -d -p 5432:5432 --name cylo_db_postgre cylo-postgre

rm_db_postgre:
	docker rm -f cylo_db_postgre

history_db_postgre:
	docker history cylo-postgre

logs_db_postgre:
	docker logs cylo_db_postgre