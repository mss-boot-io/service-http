package system

import (
	"runtime"

	"github.com/mss-boot-io/mss-boot/pkg/migration"
	migrationModel "github.com/mss-boot-io/mss-boot/pkg/migration/models"
	"gorm.io/gorm"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1691804837583Tables)
}

func _1691804837583Tables(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		err := tx.Migrator().AutoMigrate()
		if err != nil {
			return err
		}

		return tx.Create(&migrationModel.Migration{
			Version: version,
		}).Error
	})
}
