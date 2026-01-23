### **Challenge: "Dynamic SQL Query Builder CLI"**

**Objective:**  
Build a Go CLI application that acts as a dynamic SQL query builder. The application should allow users to construct complex SQL queries (SELECT, INSERT, UPDATE, DELETE) via command-line arguments and flags. The queries should be executed against a PostgreSQL database, and the results should be displayed in a formatted manner.

**Requirements:**

1. **CLI Interface:**
 - Use the `cobra` library to create a structured CLI.
 - Support the following commands:
 - `select`: Build and execute a SELECT query.
 - `insert`: Build and execute an INSERT query.
 - `update`: Build and execute an UPDATE query.
 - `delete`: Build and execute a DELETE query.

2. **Dynamic Query Building:**
 - Allow users to specify table names, columns, conditions, and values via flags.
 - For `SELECT`, support aggregation functions (e.g., `COUNT`, `SUM`, `AVG`) and `GROUP BY` clauses.
 - For `INSERT`, `UPDATE`, and `DELETE`, handle multiple columns and values dynamically.

3. **Database Interaction:**
 - Use the `database/sql` package with the `pgx` driver for PostgreSQL.
 - Ensure proper error handling for database operations.

4. **Output Formatting:**
 - Display query results in a tabular format using a library like `rodaine/table`.
 - For non-SELECT queries, display the number of rows affected.

5. **Advanced Features:**
 - Support transactions for multiple queries.
 - Allow users to save and reuse query templates.
 - Include a `--dry-run` flag to display the generated SQL query without executing it.

**Example Usage:**

```bash
# SELECT query with aggregation
$ querybuilder select --table users --columns "COUNT(id) as total" --group-by role

# INSERT query
$ querybuilder insert --table users --columns "name, email" --values "John Doe, john@example.com"

# UPDATE query with transaction
$ querybuilder update --table users --set "role='admin'" --where "id=1" --transaction

# DELETE query with dry-run
$ querybuilder delete --table users --where "id>10" --dry-run
```

**Evaluation Criteria:**
- Code cleanliness and modularity.
- Proper handling of edge cases (e.g., invalid SQL, empty inputs).
- Robust error handling and user-friendly messages.
- Efficient use of Go concurrency (if applicable).
- Comprehensive test coverage.

---

This challenge tests advanced Go programming skills, SQL knowledge, and CLI design principles, making it suitable for experienced developers.