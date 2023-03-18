init_dir := $(CURDIR)/db/init_db

start_db:
	(docker stop postgres_web101 && docker rm postgres_web101) || true 
	docker run \
	--name postgres_web101 \
	-e POSTGRESQL_USERNAME=web101 \
	-e POSTGRESQL_DATABASE=sport \
	-e POSTGRESQL_PASSWORD=123 \
	-p 5433:5432 \
	-v $(init_dir):/docker-entrypoint-initdb.d \
	-d bitnami/postgresql:latest