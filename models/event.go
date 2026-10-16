package models

import (
	// "os/exec"
	"time"

	"example.com/db"
	// "github.com/go-playground/validator/translations/id"
)

type Event struct {
	ID          int64
	Name        string
	Address     string
	Description string
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query:=`insert into events(Name,Address,Description,DateTime,UserID)
	values ('ceremony','lucknow tiwariganj','this is the celebration of shop ceremony','10-05-2026',01)`
	  stmt, err:=db.DB.Prepare(query)
	  if err!=nil{
		panic("the query doesn't executed ")
	  }
	  defer stmt.Close()
	 result, err:= stmt.Exec(e.Name,e.Address,e.DateTime,e.DateTime,e.Description)
	 if err!=nil{
		return err
	 }
	 
	id, err:=result.LastInsertId()
	e.ID= id;
	return err
}

	

func GetAllEvents()([]Event,error){ 		//This []Event returns the entire slice containing all the events. and event is custom type which i created 
	query:="select * from events"	
	rows, err:=db.DB.Query(query)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next(){
		var event Event //here Event is type of event 
		err:=rows.Scan(&event.ID,&event.Name,&event.Description,&event.Address,&event.DateTime,&event.UserID)
		if err!=nil{
			return nil,err
		}
		events=append(events, event)
	}
	return events,nil

} 