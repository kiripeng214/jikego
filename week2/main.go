package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"kiripeng214/jikego/week2/dao"
)

//需要，方便定位位置错误，能够跟踪到根源错误头
func main() {
	count, err := dao.GetCount()
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("sql error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace: %+v\n", err)
		fmt.Println(errors.Unwrap(err))
	} else {
		fmt.Println(count)
	}
}
