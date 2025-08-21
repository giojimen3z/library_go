package mock

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SetupGormWithSQLMock creates a sqlmock-backed *gorm.DB for tests and returns the
// GORM DB along with the sqlmock handle. It registers cleanup to close the
// underlying sql.DB automatically when the test finishes.
func SetupGormWithSQLMock(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	t.Helper()

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("no se pudo crear sqlmock: %v", err)
	}
	// Ensure the underlying sql DB is closed after the test
	t.Cleanup(func() { _ = sqlDB.Close() })

	dialectic := postgres.New(postgres.Config{
		Conn:                 sqlDB,
		PreferSimpleProtocol: true,
	})
	gdb, err := gorm.Open(dialectic, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("no se pudo abrir gorm con postgres sqlmock: %v", err)
	}

	return gdb, mock
}
