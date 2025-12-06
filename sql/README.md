Here is a comprehensive cheat sheet for Go’s standard `database/sql` library, focusing on advanced concepts, performance patterns, and robust error handling.

### 1. Advanced Connection Pooling
Don't just use `sql.Open`. Configure the pool to prevent connection leaks and timeouts under load.

```go
db, err := sql.Open("postgres", "postgres://user:pass@localhost/db")
if err != nil {
    log.Fatal(err)
}

// Maximum number of open connections to the database.
db.SetMaxOpenConns(25)

// Maximum number of connections in the idle connection pool.
db.SetMaxIdleConns(25)

// Maximum amount of time a connection may be reused.
db.SetConnMaxLifetime(5 * time.Minute)

// Maximum amount of time a connection may be idle.
db.SetConnMaxIdleTime(5 * time.Minute)
```

### 2. Context & Timeouts (Crucial for Production)
Always use `*Context` methods to prevent hanging queries from blocking goroutines forever.

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()

// ExecContext handles cancellation automatically
query := "INSERT INTO users (name) VALUES ($1)"
_, err := db.ExecContext(ctx, query, "Alice")

if ctx.Err() == context.DeadlineExceeded {
    log.Println("Query timed out")
}
```

### 3. Transactions with Isolation Levels
Standard `db.Begin()` uses the default isolation level. Use `BeginTx` to specify strict isolation.

```go
ctx := context.Background()
opts := &sql.TxOptions{
    Isolation: sql.LevelSerializable, // Strict isolation
    ReadOnly:  false,
}

tx, err := db.BeginTx(ctx, opts)
if err != nil {
    log.Fatal(err)
}
defer tx.Rollback() // Safe to call even if committed (no-op)

// ... execute queries using tx ...
// _, err = tx.ExecContext(...)

if err := tx.Commit(); err != nil {
    log.Fatal("Commit failed:", err)
}
```

### 4. Handling Nulls & Custom Types
#### Standard Null Types
Use `sql.Null*` types when columns are nullable to avoid Scan errors.
```go
var name sql.NullString
err := db.QueryRow("SELECT name FROM users WHERE id=1").Scan(&name)
if name.Valid {
    fmt.Println(name.String)
} else {
    fmt.Println("Name is NULL")
}
```

#### Custom JSON Type (Implementing `sql.Scanner` / `driver.Valuer`)
Store Go structs as JSON in the DB but work with them as objects in Go.

```go
type UserMeta struct {
    Theme string `json:"theme"`
    Role  string `json:"role"`
}

// Make UserMeta implement sql.Scanner (Reading from DB)
func (m *UserMeta) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }
    return json.Unmarshal(bytes, &m)
}

// Make UserMeta implement driver.Valuer (Writing to DB)
func (m UserMeta) Value() (driver.Value, error) {
    return json.Marshal(m)
}

// Usage
meta := UserMeta{Theme: "dark", Role: "admin"}
db.Exec("INSERT INTO users (meta) VALUES ($1)", meta)
```

### 5. High Performance Scanning
#### `sql.RawBytes`
Use `sql.RawBytes` to avoid memory allocation for copying data. The data is valid **only** until the next call to `Next()`.

```go
rows, _ := db.Query("SELECT large_text_col FROM logs")
defer rows.Close()

var raw sql.RawBytes
for rows.Next() {
    // raw points directly to the network buffer
    rows.Scan(&raw) 
    
    // Process immediately (e.g., write to stream)
    os.Stdout.Write(raw) 
}
```

#### Dynamic Columns (Unknown Scans)
Scan a row when you don't know the column names or count ahead of time.
```go
rows, _ := db.Query("SELECT * FROM unknown_table")
cols, _ := rows.Columns()
vals := make([]interface{}, len(cols))
dest := make([]interface{}, len(cols))

for i := range vals {
    dest[i] = &vals[i] // Pointers to interface{}
}

for rows.Next() {
    rows.Scan(dest...)
    // vals[] now holds the data
}
```

### 6. Prepared Statements
Prepare once, execute many times. Reduces parsing overhead and completely prevents SQL injection.

```go
stmt, err := db.PrepareContext(ctx, "SELECT email FROM users WHERE id = $1")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close() // Important!

// Execute multiple times
stmt.QueryRowContext(ctx, 1).Scan(&email1)
stmt.QueryRowContext(ctx, 2).Scan(&email2)
```

### 7. Bulk Inserts (Pattern)
The standard library doesn't have a specific `BulkInsert` method. You must build the query manually.

```go
// Pattern: INSERT INTO users (name, age) VALUES ($1, $2), ($3, $4), ...
users := []User{{"Alice", 30}, {"Bob", 25}}
valueStrings := make([]string, 0, len(users))
valueArgs := make([]interface{}, 0, len(users)*2)

for i, u := range users {
    // PostgreSQL uses $1, $2; MySQL uses ?
    // This example assumes PostgreSQL placeholder logic
    n := i * 2
    valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", n+1, n+2))
    valueArgs = append(valueArgs, u.Name, u.Age)
}

stmt := fmt.Sprintf("INSERT INTO users (name, age) VALUES %s", 
    strings.Join(valueStrings, ","))
    
_, err := db.Exec(stmt, valueArgs...)
```

### 8. Multiple Result Sets
Advanced databases (like SQL Server or stored procs) can return multiple tables in one query.

```go
rows, err := db.Query("SELECT * FROM users; SELECT * FROM orders;")
defer rows.Close()

// First Result Set (Users)
for rows.Next() {
    // Scan users...
}

// Advance to next result set
if rows.NextResultSet() {
    // Second Result Set (Orders)
    for rows.Next() {
        // Scan orders...
    }
}
```

### 9. Robust Error Handling Patterns
Always check for `sql.ErrNoRows` and `rows.Err()` after iteration.

```go
var name string
err := db.QueryRow("SELECT name FROM users WHERE id = 999").Scan(&name)

if err == sql.ErrNoRows {
    // Handle "Not Found" logic specifically
    return nil, fmt.Errorf("user not found")
} else if err != nil {
    // Handle other DB errors (connection, syntax, etc.)
    return nil, err
}

// Iteration Error Handling
rows, _ := db.Query("SELECT ...")
defer rows.Close()
for rows.Next() {
    // scan...
}
// MUST check rows.Err() after loop finishes
if err := rows.Err(); err != nil {
    log.Println("Error during iteration:", err)
}
```

### Note on Named Parameters
Go's standard `database/sql` **does not** support named parameters (e.g., `:name`, `@val`).
- **PostgreSQL:** Uses `$1`, `$2`
- **MySQL / SQLite:** Uses `?`
- **SQL Server:** Uses `@p1`

*If you strictly require named parameters, you must use a third-party wrapper like `sqlx`.*