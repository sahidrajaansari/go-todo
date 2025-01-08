package persistence

func AppendIdInQuery(query string, id string) string {
	queryStr := "_id=" + id
	if len(query) != 0 {
		queryStr = queryStr + "&" + query
	}
	return queryStr
}
