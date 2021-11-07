package FirstWork

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

//遇到sql.ErrNoRows,Wrap error抛给上层
func GetData() error {
	return errors.Wrap(sql.ErrNoRows, "sql: no rows in result!")
}

func QurrySQL() error {

	return errors.WithMessage(GetData(), "Qurry:call failed")
}

//上层处理
func main() {

	err := QurrySQL()
	if err != nil {
		fmt.Printf("original error, %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
		os.Exit(1)
	}

}
