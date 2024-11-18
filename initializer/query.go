package initializer

import "example.com/m/query"

func InitQuery() *query.Queries {
	return query.New(InitDb())
}
