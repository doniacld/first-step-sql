package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/doniacld/first-step-sql/internal/db"
	_ "github.com/lib/pq" // to establish the connection with the driver even it is not instantiate
)

/*
type studentDB struct {
	pool *pgxpool.Pool
}

func New() (db.StudentDB, error) {
	// parse the configuration to retrieve the database URL

//	config, err := pgxpool.ParseConfig(os.Getenv("postgres://root:root@locahost:5432/test_db"))
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	// establish a pool connection with the database
	postgresPool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	d := studentDB{pool: postgresPool}
	return &d, nil
}
*/
const (
	// should be the name of the database service in the docker compose
	HOST = "db"
	PORT = 5432
)

type studentDB struct {
	conn *sql.DB
}

func New(username, password, database string) (db.StudentDB, error) {

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", HOST, PORT, username, password, database)

	log.Printf(connStr)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not open a new connection: %s", err)
	}

	err = pingDB(conn, 3)
	if err != nil {
	// TODO DONIA
			return nil, fmt.Errorf("ping database failed: %s", err)
	}

	log.Println("Database connection established")

	d := studentDB{conn: conn}
	return &d, nil
}

func pingDB(conn *sql.DB, n int) (err error) {
	for i := 0; i < n; i++ {
		err = conn.Ping()
		if err == nil {
			return nil
		}
		log.Printf("try to connect to the database: %d ", i)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("something wrong happened when pinging database: %s", err)
	}
	return nil
}

// Close the pool connection
func (s *studentDB) Close() error {
	return s.conn.Close()
}
