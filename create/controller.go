package create

import (
	table "Gocassa/create/table"
	"github.com/gocql/gocql"
)

func Controller(session *gocql.Session) {
	table.CreateTable(session)
}
