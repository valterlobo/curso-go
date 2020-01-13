package main

import (
	"fmt"
	"regexp"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

const tagName = "valid"

type User struct {
	Email string `json:"email" validate:"email"`
	Name  string `json:"name"  validate:"required"`
}

type Post struct {
	Title    string `valid:"stringlength(3|40)"`
	Message  string `valid:"duck,ascii"`
	Message2 string `valid:"animal(dog)"`
	AuthorIP string `valid:"ipv4"`
	Date     string `valid:"-"`
}

func main() {
	u := User{
		Name:  "John Doe Teste Silva 34343434",
		Email: "johnexample",
	}
	//err := validate.Validate(a)

	//
	translator := pt_BR.New()
	uni := ut.New(translator, translator)

	a := User{
		Email: "a"
	}
	err := validator.Struct(a)

	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e.Translate(trans))
	}

	v := validator.New()
	a := User{
		Email: "a",
	}
	err := v.Struct(a)

	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e)
	}
	//

	result, err := govalidator.ValidateStruct(u)

	if err != nil {
		println("error: " + err.Error())

	}

	println(result)

	post := &Post{
		Title:    "My Example Post",
		Message:  "duck2",
		Message2: "dog",
		AuthorIP: "123.234.54.3",
	}

	// Add your own struct validation tags
	govalidator.TagMap["duck"] = govalidator.Validator(func(str string) bool {
		return str == "duck"
	})

	// Add your own struct validation tags with parameter
	govalidator.ParamTagMap["animal"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		species := params[0]
		return str == species
	})
	govalidator.ParamTagRegexMap["animal"] = regexp.MustCompile("^animal\\((\\w+)\\)$")

	result2, err2 := govalidator.ValidateStruct(post)

	if err2 != nil {
		println("error: " + err2.Error())
	}

	println(result2)

}
