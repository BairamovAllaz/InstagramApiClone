package structs

import "mime/multipart"

type PostStruct struct {
	User    int `form:"-"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
	Comment string `form:"comment"`
	//TODO create like count in postresq default 0
	//TODO create created at
}
