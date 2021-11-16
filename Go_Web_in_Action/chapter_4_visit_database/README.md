# Visit Database by Golang

### Visit Mysql by Golang

1. `database/sql` package in Golang provides several universal interface to connect SQL database, but does not provide the specific database driver. Therefore at least one database driver need to be imported before usage.

2. Use `sql.Open()` to get a DB instance, which is safe for concurrent use by multiple goroutines and maintains its own pool of idle connections. Thus, the Open function should be called just once. It is rarely necessary to close a DB.

3. [1.mysqlOperation.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_4_visit_database/1.mysqlOperation.go) shows the most common database operation by Golang. It should be ensured that call `Scan()` after QueryRow(), otherwise the connection would not be released.

4. A prepared statement could be fulfilled by using `*DB.prepare()`, which could save compilation cost and avoid SQL injection. Database transaction could be implemented by `*DB.Begin()`, `*Tx.rollback()` and `*Tx.Commit()`.

### Common ORM Library in Golang

1. ORM makes a mapping between the relational database and the object, which acts as a bridge between the business logic layer and the database layer.

2. Sample of gorm, see: [2.gormSample.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_4_visit_database/2.gormSample.go); sample of beegoORM, see: [3.beegoORM.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_4_visit_database/3.beegoORM.go)