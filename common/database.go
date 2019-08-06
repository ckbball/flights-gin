package common

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

type Database struct {
  db *sql.DB

  insert *sql.Stmt
  list   *sql.Stmt
  listBy *sql.Stmt
  get    *sql.Stmt
  filter *sql.Stmt
}

var dbHandle *Database

const insertFlightStatement = `
  INSERT INTO Flights (departurecity, arrivalcity, airline, airlineid, departat, arriveat) VALUES(?, ?, ?, ?, ?, ?)`

func ExecAffectingOneRow(stmt string, args ...interface{}) (sql.Result, error) {
  var r1 sql.Result
  if stmt == "insert" {
    r2, err := dbHandle.insert.Exec(args...)
    if err != nil {
      return r2, fmt.Errorf("mysql: could not execute statement: %v", err)
    }
    r1 = r2
  }

  return r1, nil
}

const filterFlightStatement = `
  SELECT * FROM flights WHERE departurecity=? and arrivalcity=? and departat=?`

func FilterRows(table string, dc string, ac string, da string) (*sql.Rows, error) {
  return dbHandle.filter.Query(dc, ac, da)
}

// GetRows - takes the table to get rows from and the args for the query

const listFlightsStatement = `SELECT * FROM flights ORDER BY airline`

func GetRows(table string, args ...interface{}) (*sql.Rows, error) {
  if args != nil {
    return dbHandle.listBy.Query(args...)
  }

  return dbHandle.list.Query()
}

// GetRow - takes an id and gets the row that corresponds to the id

const getStatement = `SELECT * FROM flights WHERE id = ?`

func GetRow(table string, id int) (*sql.Row, error) {
  return dbHandle.get.QueryRow(id), nil
}

// -------------- BEGINNING OF DB INITIALIZATION AND SHUT DOWN -------------------

// Open a database and save reference to `Database` struct
func Init() (*Database, error) {
  db, err := sql.Open("mysql", "dev:dev-user5@/flighttest")
  if err != nil {
    fmt.Println("db err: ", err)
  }
  fmt.Println("DB connected")
  /*err = db.Ping()
    if err != nil {
      fmt.Println("DB is not connecting: in database.go - ", err)
    }*/

  db.SetMaxIdleConns(10)
  DB := &Database{
    db: db,
  }

  // Prepared statements

  // Register the insert statement for flights
  if DB.insert, err = db.Prepare(insertFlightStatement); err != nil {
    return nil, fmt.Errorf("ERROR --> mysql: prepare insert flight: %v", err)
  }
  // Register the list statement for flights
  if DB.list, err = db.Prepare(listFlightsStatement); err != nil {
    return nil, fmt.Errorf("ERROR --> mysql: prepare list flights: %v", err)
  }
  // Register the get by id statement for flights
  if DB.get, err = db.Prepare(getStatement); err != nil {
    return nil, fmt.Errorf("ERROR --> mysql: prepare get flight: %v", err)
  }
  // Register the filter statement for flights
  if DB.filter, err = db.Prepare(filterFlightStatement); err != nil {
    return nil, fmt.Errorf("ERROR --> mysql: prepare filter flights: %v", err)
  }
  dbHandle = DB
  return dbHandle, nil
}

// This function will create a temporary database for test cases
func TestDBInit() (*Database, error) {
  test_db, err := sql.Open("mysql", "web:poggerspeepo@/flighttest")
  if err != nil {
    fmt.Println("db err: ", err)
  }

  test_db.SetMaxIdleConns(3)
  //test_db.LogMode(true)
  DB := &Database{
    db: test_db,
  }

  // Prepared statements
  if DB.insert, err = test_db.Prepare(insertFlightStatement); err != nil {
    return nil, fmt.Errorf("mysql: prepare insert flight: %v", err)
  }

  dbHandle = DB
  return dbHandle, nil
}

func Close() {
  dbHandle.db.Close()
}

func TestDBFree(test_db *Database) {
  test_db.db.Close()
}

func GetDB() *Database {
  return dbHandle
}

// ----------------- END OF DB SET UP AND SHUT DOWN ------------------------------
