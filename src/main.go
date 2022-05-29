package main

// import (
// 	"fmt"

// 	"practice/controllers"
// 	"practice/controllers/admins"
// 	"practice/models"
// 	// 同じパッケージ名がある場合、一番最後にimportしたものが使用される
// 	"practice/models/admins"
// )

import (
	"fmt"

	"practice/controllers"
	admin_controller "practice/controllers/admins"
	"practice/models"
	admin_model "practice/models/admins"
)

func main() {
	fmt.Printf("%v\n", controllers.GetController());
	fmt.Printf("%v\n", admin_controller.GetController());
	fmt.Printf("%v\n", models.GetModel());
	fmt.Printf("%v\n", admin_model.GetModel());
}


