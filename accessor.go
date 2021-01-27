package database

import (
	"log"
	"strings"

	convert "github.com/leguminosa/golang-convert"
	"github.com/leguminosa/kururu-engine/driver"
)

type (
	databaseAccessor struct {
		connections map[string]*database
	}
)

func (accessor *databaseAccessor) GetDB(name string) DatabaseItf {
	log.Printf("Getting database %q...\n", name)
	instances, ok := accessor.connections[name]
	if !ok {
		log.Fatalln("Invalid database name given.")
	}

	return instances
}

func (accessor *databaseAccessor) GetDBInstance(name, replication string) driver.OrmItf {
	instances := accessor.GetDB(name)

	log.Printf("Getting replication %q...\n", replication)
	if replication == "master" {
		return instances.GetMasterDB()
	}

	slaveStr := strings.Split(replication, "_")
	if slaveStr[0] != "slave" {
		log.Fatalln("Invalid replication name given.")
	}

	var index int
	if len(slaveStr) > 1 {
		index = convert.ToInt(slaveStr[1])
	}

	return instances.GetSlaveDB(index)
}
