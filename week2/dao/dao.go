package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

//GetCount
//
func GetCount() (int64, error) {
	var count int64
	//TODO: 执行了查询
	return count, errors.Wrap(sql.ErrNoRows, "jikego/week2/GetCount")
}
