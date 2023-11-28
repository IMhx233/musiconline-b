package main
import (
"github.com/gin-gonic/gin"
"musiconline/controller"
)


func main(){
	controller.Login()
	r:=gin.Default()
	r.POST("password")
	
	r.Run(":5173")
	
}