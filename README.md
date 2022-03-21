This is the sqlx prototype

sqlx is an extension toolkit library built on database/sql, where you have to write SQL queries to interact with your database

Basic Use:

The handle types in sqlx implement the same basic verbs for querying your database:

- **`Exec(...) (sql.Result, error)`** - unchanged from database/sql
- **`Query(...) (*sql.Rows, error)`** - unchanged from database/sql
- **`QueryRow(...) *sql.Row`** - unchanged from database/sql

These extensions to the built-in verbs:

- **`MustExec() sql.Result`** -- Exec, but panic on error
- **`Queryx(...) (*sqlx.Rows, error)`** - Query, but return an sqlx.Rows
- **`QueryRowx(...) *sqlx.Row`** -- QueryRow, but return an sqlx.Row

And these new semantics:

- **`Get(dest interface{}, ...) error`**
- **`Select(dest interface{}, ...) error`**


Models in this prototype
 - Person
 
