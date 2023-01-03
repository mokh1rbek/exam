package helper

import (
	"database/sql"
	"strconv"
	"strings"
)

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" {
			if v == "" {
				// oldsize := len(namedQuery)
				namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))

				// if oldsize != len(namedQuery) {
				args = append(args, v)
				i++
				// }
			}
		}
	}

	return namedQuery, args
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}

	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
