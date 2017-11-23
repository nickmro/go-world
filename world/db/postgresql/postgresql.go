package postgresql

import "github.com/drinkin/di/env"

func init() {
	env.MustSet("LOGXI", "*,-dat*")
	env.MustSet("LOGXI_FORMAT", "happy")
}
