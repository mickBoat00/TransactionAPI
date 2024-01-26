build:
	docker build -t golang-sqlc-goose -f Dockerfile_base .
	docker compose up --build

up:
	docker compose down
	docker compose up --build
	
push-base:
	docker build tag golang-sqlc-goose:latest mickeys0105/golang-transaction:golang-sqlc-goose
	docker push mickeys0105/golang-transaction:golang-sqlc-goose

push-main:
	docker tag transactionapi-app:latest mickeys0105/golang-transaction:transactionapi
	docker push mickeys0105/golang-transaction:transactionapi
