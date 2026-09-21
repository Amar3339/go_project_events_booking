package main
import("github.com/gin-gonic/gin",
"fmt",
"net/http"
)
func main(){
	server:=gin.default()
	server.GET("events",getEvents)
	server.Run(":8080") //localhost 8080
}
getEvents(context *gin.Contenxt){
	context.JSON(http.StatusOK,gin.H{"message:My first api response""jai shree mahakal "})
}