package cloudpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func GetGolderIDByClassID(classID string) (string, error) {
	db := pggorm.DB
	if db == nil {
		return "", errors.New("database connection is not established")
	}
	var c models.Class
	db.Where("id =?", classID).First(&c)
	return c.FolderID, nil
}
