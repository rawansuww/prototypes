sqlc is a SQL-to-Go code generator

Steps:
    1. Installed sqlc with 'brew install sqlc' (can also install with go)
    2. Added a sqlc.yaml, schema.sql and query.sql to define a basic schema and some desired queries
    3. Ran the command 'sqlc generate' which produced three files under the 'tutorial' package