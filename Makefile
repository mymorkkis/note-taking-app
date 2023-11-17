create_migration:
	CMD="create -seq -ext=.sql -dir=./sql/migrations $(NAME)" docker compose up migrate

run_migrations:
	CMD="$(CMD)" docker compose up migrate
