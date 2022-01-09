package persistence

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"skiresorts/domain"
	"time"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newPostgresDB() *sql.DB {
	yfile, err := ioutil.ReadFile("../../ssrdb.yml")
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]string)
	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		log.Fatal(err)
	}

	dbHost := data["dbHost"]
	dbPort := data["dbPort"]
	dbName := data["dbName"]
	dbUser := data["dbUser"]
	dbPswd := data["dbPswd"]

	fmt.Printf("dbHost= %s\n", dbHost)
	fmt.Printf("dbPort= %s\n", dbPort)
	fmt.Printf("dbName= %s\n", dbName)
	fmt.Printf("dbUser= %s\n", dbUser)
	fmt.Printf("dbPswd= %s\n", dbPswd)

	dbConnectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPswd,
		dbName,
	)

	conn, err := sql.Open("postgres", dbConnectionStr)
	if err != nil {
		fmt.Printf("DB error = %s\n", err.Error())
	}
	if err = conn.Ping(); err != nil {
		fmt.Printf("Connection dropped")
	}
	return conn
}

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo() *postgresRepo {
	return &postgresRepo{
		db: newPostgresDB(),
	}
}

func (pr *postgresRepo) GetHill(id int) (*domain.Hill, error) {
	const query = `
		select length, slope
		from hills
		where id=%d
	`
	results, err := pr.db.Query(fmt.Sprintf(query, id))

	if err != nil {
		return nil, err
	}

	if results.Next() {
		var hill struct {
			Length float64
			Slope  float64
		}
		err = results.Scan(&hill.Length, &hill.Slope)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		return &domain.Hill{
			Length: hill.Length,
			Slope:  hill.Slope,
		}, nil
	}
	return nil, fmt.Errorf("Hill with id %d is not found", id)
}

func (pr *postgresRepo) UpdateHill(id int, hill *domain.Hill) error {
	gormLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		Colorful:                  false,
		IgnoreRecordNotFoundError: true,
	})
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: pr.db,
	}), &gorm.Config{Logger: gormLogger})

	if err != nil {
		return err
	}

	const query = `
		update hills set length = %f, slope = %f
		where id=%d
	`
	result := db.Exec(fmt.Sprintf(query, hill.Length, hill.Slope, id))
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("no update happened for hill %d", id)
	}
	return nil
}
