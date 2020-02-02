/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 19:33
 */

package main

import (
	"flag"
	"safe-community/core/controller"
	"safe-community/core/dao/mysql"
)

func main() {

	flag.Parse()
	r := controller.InitRouter()
	_ = mysql.SingleStore()

	r.Run(":8080")
}
