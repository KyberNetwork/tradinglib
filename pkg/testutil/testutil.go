package testutil

import (
	"encoding/json"
	"fmt"
	"maps"
	"math/big"
	"math/rand"

	"github.com/KyberNetwork/tradinglib/pkg/dbutil"
	"github.com/jmoiron/sqlx"
)

const (
	defaultHost     = "127.0.0.1"
	defaultPort     = 5432
	defaultUser     = "test"
	defaultPassword = "test"
)

func DefaultDSN() map[string]any {
	return map[string]any{
		"host":     defaultHost,
		"port":     defaultPort,
		"user":     defaultUser,
		"password": defaultPassword,
		"sslmode":  "disable",
		"TimeZone": "UTC",
	}
}

// MustNewDevelopmentDB creates a new development DB.
// It also returns a function to teardown it after the test.
func MustNewDevelopmentDB(migrationPath string, dsn map[string]any, dbName string) (*sqlx.DB, func() error) {
	copyDSN := maps.Clone(dsn)
	delete(copyDSN, dbName)

	// CREATE DB
	dsnStr := dbutil.FormatDSN(copyDSN)
	ddlDB, err := dbutil.NewDB(dsnStr)
	if err != nil {
		panic(err)
	}
	ddlDB.MustExec(fmt.Sprintf(`CREATE DATABASE "%s"`, dbName))
	if err := ddlDB.Close(); err != nil {
		panic(err)
	}

	// MIGRATE
	copyDSN["dbname"] = dbName
	dsnWithDB := dbutil.FormatDSN(copyDSN)
	db, err := dbutil.NewDB(dsnWithDB)
	if err != nil {
		panic(err)
	}
	m, err := dbutil.RunMigrationUp(db.DB, migrationPath, dbName)
	if err != nil {
		panic(err)
	}

	return db, func() error {
		if _, err := m.Close(); err != nil {
			return err
		}
		ddlDB, err := dbutil.NewDB(dsnStr)
		if err != nil {
			return err
		}
		if _, err = ddlDB.Exec(fmt.Sprintf(`DROP DATABASE "%s"`, dbName)); err != nil {
			return err
		}
		return ddlDB.Close()
	}
}

// RandomString generates a random string with given length.
// Notice: this function uses a pseudo random algorithm, only for use in test.
func RandomString(n int) string {
	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))] //nolint
	}

	return string(b)
}

func MustJsonify(data interface{}) string {
	d, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}
	return string(d)
}

func NewBig10(s string) *big.Int {
	b, _ := new(big.Int).SetString(s, 10)

	return b
}
