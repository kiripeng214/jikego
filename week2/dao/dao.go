package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

//GetCount
//
// 按照业务性来做就是需要把code.ErrNoRows来作为ddd业务
//if err == sql.ErrNoRows {
//
//return errors.Wrapf(code.NotFound, fmt.Sprintf("sql: %s error: %v", sql, err))
//
//} else {
//
//return errors.Wrapf(code.Internal, fmt.Sprintf("sql: %s error: %v", sql, err))
//
//}
func GetCount() (int64, error) {
	var count int64
	//TODO: 执行了查询
	return count, errors.Wrap(sql.ErrNoRows, "jikego/week2/GetCount")
}
