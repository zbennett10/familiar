package factory

import (
	"dungeon-comrade/daos/cassandra"
	"dungeon-comrade/daos/interfaces"
	"fmt"
	"os"
)

func UserFactoryDao(dbType string) interfaces.UserDao {
	var i interfaces.UserDao
	switch dbType {
	case "cassandra":
		i = cassandra.UserImpl{}
	default:
		fmt.Printf("This database is not implemented %s", dbType)
		os.Exit(1)
		return nil
	}

	return i
}
