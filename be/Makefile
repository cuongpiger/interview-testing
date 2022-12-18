build_service:
	docker build -t cylo-service .

run_service:
	docker run -d --env-file=./build/.env --name cylo-service cylo-service

rm_service:
	docker rm -f cylo-service