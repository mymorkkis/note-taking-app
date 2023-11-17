# Note Taking App

## To Develop Locally

Copy `.env.example` file to `.env` and update any necessary values.

`docker compose up` will run the applicaton and all db migrations.

The app has live reloading using [Air](https://github.com/cosmtrek/air)

### DB Migrations

[golang-migrate](https://github.com/golang-migrate/migrate) is used.

You can add a new migration with the following make command:

`NAME=THE_MIGRATION_NAME make create_migration`

There is also a make command to update the migrations, down/up etc, E.G:

`CMD=up make run_migrations`

### DBAL

The code in `/internal/dbal/` is generated by by [sqlc](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html) and not to be updated manually.

To generate the code:

- create a `.sql` file in `/sql/queries` and add any applicable queries
- run `sqlc generate`