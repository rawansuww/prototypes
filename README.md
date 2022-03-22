    sqlc is a SQL-to-Go code generator

Steps:
- Installed sqlc with 'brew install sqlc' (can also install with go)
- Added a sqlc.yaml, schema.sql and query.sql to define a basic schema and some desired queries
- Ran the command 'sqlc generate' which produced three files under the 'tutorial' package
    
    
    The 'tutorial' package is purely generated code. The main.go file is where you can call and use the auto-generated functions. 
    
    
Problems:
- the codegen actually did not create a postgres table, i had to do it manually each time, not sure if there was a generated 'create table' function at all
- it was actually a pain to use
