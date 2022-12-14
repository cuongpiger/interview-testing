build_db:
	docker build --no-cache -t cylo-mysql .

run_db:
	docker run -d -p 3306:3306 --name cylo_db cylo-mysql

exec_db:
	docker exec -it cylo_db mysql -u root -p