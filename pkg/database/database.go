package database

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/SamuelTissot/sqltime"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ErrorLogger func(err error, msg string)

func ConnectInit(dsn, user, pass string, retries int) (*gorm.DB, error) {
	if len(dsn) == 0 {
		return nil, errors.New("missing dsn configuration")
	}

	if retries == 0 {
		return nil, errors.New("missing retry configuration")
	}

	dsnpg, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}

	dsnpg.User = url.UserPassword(user, pass)
	db, err := connect(dsnpg.String(), retries, func(err error, msg string) {
		fmt.Printf("%s, %s\n", msg, err.Error())
	})

	return db, err
}

func connect(dsn string, retries int, logger ErrorLogger) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	gormConfig := &gorm.Config{
		NowFunc: func() time.Time {
			return sqltime.Now().UTC()
		},
	}

	db, err = gorm.Open(postgres.Open(dsn), gormConfig)
	for err != nil {
		if logger != nil {
			logger(err, fmt.Sprintf("Database connect failed (%d)", retries))
		}

		if retries > 1 {
			retries--
			time.Sleep(5 * time.Second)
			db, err = gorm.Open(postgres.Open(dsn), gormConfig)
			continue
		}

		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
