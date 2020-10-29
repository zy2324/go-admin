package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/engine"
	//"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	//"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/language"
	//"net"
	//"net/http"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	// global config
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:         "0.0.0.0",
				Port:         "6006",
				User:         "root",
				Pwd:          "zhaoyi007",
				Name:         "goadmin",
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:       "mysql",
			},
        	},
		UrlPrefix: "admin",
		// STORE is important. And the directory should has permission to write.
		Store: config.Store{
		    Path:   "./uploads", 
		    Prefix: "uploads",
		},
		Language: language.EN,
		// debug mode
		Debug: true,
		// log file absolute path
		InfoLogPath: "/var/logs/info.log",
		AccessLogPath: "/var/logs/access.log",
		ErrorLogPath: "/var/logs/error.log",
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}

	// add component chartjs
	template.AddComp(chartjs.NewChart())

	_ = eng.AddConfig(cfg).
		AddGenerators(datamodel.Generators).
	        // add generator, first parameter is the url prefix of table when visit.
    	        // example:
    	        //
    	        // "user" => http://localhost:9033/admin/info/user
    	        //		
		AddGenerator("user", datamodel.GetUserTable).
		Use(r)
	
	// customize your pages
	eng.HTML("GET", "/admin", datamodel.GetContent)

	
	_ = r.Run(":8080")

	//启动tcp4
	//server := &http.Server{Handler: r}
	//l, err := net.Listen("tcp4", "0.0.0.0:8080")
	//if err != nil {
	//    return
	//}
	//err = server.Serve(l)
	//if err != nil {
	//	return
	//}
}
