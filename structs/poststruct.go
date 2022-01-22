package structs


type PostStruct struct{  
	User string `json:"-"`
	Image string `json:"image"`
	Comment string `json:"comment"`
	//TODO create like count in postresq default 0
	//TODO create created at
}
