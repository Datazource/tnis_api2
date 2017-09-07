package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountDepositController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountDepositController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountDepositController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountDepositController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountDepositController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountDepositController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountDepositController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountDepositController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountDepositController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountDepositController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingHistoryController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingHistoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingHistoryController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingHistoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingHistoryController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingHistoryController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingHistoryController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingHistoryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingHistoryController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:AccountSavingHistoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:CustomerAccountController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:CustomerAccountController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:CustomerAccountController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:CustomerAccountController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:CustomerAccountController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:CustomerAccountController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:CustomerAccountController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:CustomerAccountController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:CustomerAccountController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:CustomerAccountController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:CustomerController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:CustomerController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:CustomerController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:CustomerController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:CustomerController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:UsersController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:UsersController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:UsersController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:UsersController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["tnis_api2/controllers:UsersController"] = append(beego.GlobalControllerRouter["tnis_api2/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
