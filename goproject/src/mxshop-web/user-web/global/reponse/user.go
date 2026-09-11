package reponse

import (
	"fmt"
	"time"
)

//import "time"
type JsonTime time.Time
func (j JsonTime)MarshalJSON() ([]byte,error){
	var stmp = fmt.Sprintf("\"%s\"",time.Time(j).Format("2006-01-02"))
	return []byte(stmp),nil
}
 
type UserResPonse struct {
	Id       int32  `json:"id"`
	NickName string `json:"name"`
	//BirthDay string `json:"birthday"`
	//BirthDay time.Time `json:"birthday"`
	BirthDay JsonTime `json:"birthday"`
	Gender   string `json:"gender"`
	Mobile string `json:"mobile"`
	
}