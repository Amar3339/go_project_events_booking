package main

import (
	"net/http"

	"example.com/models"
	"github.com/gin-gonic/gin"
)
func main(){
	server:=gin.Default()
	server.GET("/events",getEvents)
	server.GET("/home",gethome)
	server.POST("/create",create)
	server.POST("/events",createEvents)
	server.Run(":8080") //localhost 8080
}
 func getEvents(context *gin.Context){
	context.JSON(http.StatusOK,gin.H{"message":"My first api response jai shree mahakal"})
}
func gethome(context *gin.Context){
	context.JSON(http.StatusOK,gin.H{"message":"this is my home page"})
}
func createEvents(context *gin.Context){
	var event models.Event
	err:=context.ShouldBindJSON(&event)
	if err !=nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"bad response"})
		return
	}
	event.ID=1
	event.UserID=1
	event.Save()
	context.JSON(http.StatusCreated,gin.H{"message":"event created","event":event})

}
func create(context *gin.Context){
	var event models.Event
	err:=context.ShouldBindJSON(&event)
	if err !=nil{
context.JSON(http.StatusCreated,gin.H{"message":"api not found ankit"})
return
	}
	event.ID=2
	event.UserID=2
	event.Save()
	context.JSON(http.StatusCreated,gin.H{"mesaage":"anit event is creeate ed","event":event})
}
