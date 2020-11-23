package PrimaryController

import (
	"Gocassa/create"
	"github.com/gocql/gocql"
)

func Controller(session *gocql.Session) {
	create.Controller(session)
}
