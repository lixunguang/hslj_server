package dao

import (
	"hslj/internal/model/mysql"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	mysql.OpenDB()

	exitCode := m.Run()

	os.Exit(exitCode)

}
