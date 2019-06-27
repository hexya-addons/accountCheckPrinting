package account_check_printing

import (
	"github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME string = "account_check_printing"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})

}
