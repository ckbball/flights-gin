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

const listFlightsStatement = `SELECT * FROM flights ORDER BY airline`

func GetRows(table string, args ...interface{}) (*sql.Rows, error) {
  if args != nil {
    return dbHandle.listBy.Query(args...)
  }

  return dbHandle.list.Query()
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
  if DB.insert, err = db.Prepare(insertFlightStatement); err != nil {
    return nil, fmt.Errorf("ERROR --> mysql: prepare insert flight: %v", err)
  }
  if DB.list, err = db.Prepare(listFlightsStatement); err != nil {
    return nil, fmt.Errorf("ERROR --> mysql: prepare list flights: %v", err)
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
