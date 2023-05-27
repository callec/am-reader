# Info
The database is wrapped using the `sqlc` program.
More information can be found here: https://sqlc.dev.
But basically, the initial database can be found in `init.sql`, and all
possible queries in `queries.sql`.

If you want to change or add new queries you must use the specific syntax
defined by `sqlc` and then re-run `sqlc generate` in this directory.

# Initialisation
`sqlc generate`.
