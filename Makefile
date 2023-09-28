.PHONY: hurl-tests create-tables

down:
	docker compose down

up:
	docker compose up -d

create-tables:
	sleep 5
	docker compose exec postgres psql -U joesantos418 -d api_db -a -f /db-scripts/create_tables.sql

hurl-tests:
	sleep 10
	docker compose exec hurl hurl --test /api-tests/empty_email.hurl
	docker compose exec hurl hurl --test /api-tests/empty_name.hurl
	docker compose exec hurl hurl --test /api-tests/empty_request.hurl
	docker compose exec hurl hurl --test /api-tests/invalid_email.hurl
	docker compose exec hurl hurl --test /api-tests/method_not_allowed.hurl
	docker compose exec hurl hurl --test /api-tests/valid_request.hurl

integration-tests: down up create-tables hurl-tests
	docker compose exec integration-tests go install github.com/jstemmer/go-junit-report/v2@latest
	@mkdir -p test-results/api
	docker compose exec integration-tests ./integration-test.sh

ci-integration-tests: up create-tables hurl-tests
	docker compose exec integration-tests go install github.com/jstemmer/go-junit-report/v2@latest
	@mkdir -p test-results/api
	docker compose exec integration-tests ./integration-test.sh
