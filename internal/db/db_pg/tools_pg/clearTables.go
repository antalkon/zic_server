package toolspg

import (
	"errors"

	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func ClearCpuTable() error {
	db := pggorm.DB
	if db == nil {
		return errors.New("Error connecting to database")
	}
	err := db.Exec("TRUNCATE TABLE cpu_loads").Error
	if err != nil {
		return err
	}
	return nil
}

func ClearNetworkTable() error {
	db := pggorm.DB
	if db == nil {
		return errors.New("Error connecting to database")
	}
	err := db.Exec("TRUNCATE TABLE network_loads").Error
	if err != nil {
		return err
	}
	return nil
}

func ClearRamTable() error {
	db := pggorm.DB
	if db == nil {
		return errors.New("Error connecting to database")
	}
	err := db.Exec("TRUNCATE TABLE ram_loads").Error
	if err != nil {
		return err
	}
	return nil
}

func ClearDTasksTable() error {
	db := pggorm.DB
	if db == nil {
		return errors.New("Error connecting to database")
	}
	err := db.Exec("TRUNCATE TABLE delivered_tasks").Error
	if err != nil {
		return err
	}
	return nil
}

func ClearPTasksTable() error {
	db := pggorm.DB
	if db == nil {
		return errors.New("Error connecting to database")
	}
	err := db.Exec("TRUNCATE TABLE public_tasks").Error
	if err != nil {
		return err
	}
	return nil
}
