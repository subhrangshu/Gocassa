package create

import (
	"github.com/gocql/gocql"
	"./table"
)

func Controller(session *gocql.Session) {
	table.CreateTable(session);
}
