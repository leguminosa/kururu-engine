package database

import "github.com/leguminosa/kururu-engine/driver"

type (
	DatabaseAccessorItf interface {
		GetDB(name string) DatabaseItf
		GetDBInstance(name, replication string) driver.OrmItf
	}
	DatabaseItf interface {
		GetMasterDB() driver.OrmItf
		GetSlaveDB(index int) driver.OrmItf
	}
	DatabaseConnectionConfig struct {
		Dialect string
		Master  string
		Slave   []string
	}
)
