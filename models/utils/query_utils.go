package utils

import "gensh.me/goforum/models/database"

type Count struct {
	Count int
}

const CountSQL = "select count(*) as count from "

func Counts(tableName string, filter string, args ...interface{}) int {
	count := Count{}
	err := database.O.Raw(CountSQL + tableName + " where " + filter, args).QueryRow(&count)
	if err == nil {
		return count.Count
	}
	return 0
}
