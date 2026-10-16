package main

import (
	"fmt"
	"net/http"

	"example.com/db"
	"example.com/models"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/v2/event"
)
func main(){
	server:=gin.Default()
	db.InitDB()
	server.GET("/events",getEvents) //if /events api endpoints is hit the getEvents function call and run the function.this is total meaning of this line 
	server.GET("/home",gethome)

	server.POST("/events",createEvents) //events is the api endpoint 
	server.Run(":8080") //localhost 8080
}
 func getEvents(context *gin.Context){
	events,err:=models.GetAllEvents()
	if err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"could not fetch data it is server error  "})
		return

	}
	
	context.JSON(http.StatusOK, events)


}
func gethome(context *gin.Context){
	context.JSON(http.StatusOK,gin.H{"message":"this is my home page"})
}
func createEvents(context *gin.Context){
	var event models.Event
	err:=context.ShouldBindJSON(&event)
	if err !=nil{
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message":"could not fetch data it is server error  "})
		return
	}
	event.ID=1
	event.UserID=1
	event.Save()
	fmt.Print(err)
	context.JSON(http.StatusCreated,gin.H{"message":"event created","event":event})

}
