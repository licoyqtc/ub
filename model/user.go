package model


type User struct {


	Username	string	`json:"username"`
	Password	string	`json:"password"`
}




func (uo User) QueryByid()(u User, e error) {




}