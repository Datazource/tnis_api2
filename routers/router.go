// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"tnis_api2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/account_deposit",
			beego.NSInclude(
				&controllers.AccountDepositController{},
			),
		),

		beego.NSNamespace("/account_saving",
			beego.NSInclude(
				&controllers.AccountSavingController{},
			),
		),

		beego.NSNamespace("/account_saving_history",
			beego.NSInclude(
				&controllers.AccountSavingHistoryController{},
			),
		),

		beego.NSNamespace("/customer",
			beego.NSInclude(
				&controllers.CustomerController{},
			),
		),

		beego.NSNamespace("/customer_account",
			beego.NSInclude(
				&controllers.CustomerAccountController{},
			),
		),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
