package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	en "github.com/go-playground/locales/en"
	zh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"strings"

	entranslations "github.com/go-playground/validator/v10/translations/en"
	cntranslations "github.com/go-playground/validator/v10/translations/zh"
)

type LoginForm struct {
	User     string `json:"user" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required"`
}

type SignUpForm struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

var trans ut.Translator

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := make(map[string]string,0)

	for field , value := range fields {
		fmt.Println(field, value)
		rsp[field[strings.Index(field, ".")+1:]] = value
	}

	return rsp

}

func InitTrans(locale string) (err error) {
	//to custom validator property
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterTagNameFunc(func (fld reflect.StructField)string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}

			return name
		})
		zh := zh.New()
		en := en.New()
		//default en, support cn, and en
		uni := ut.New(en, zh, en)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s", locale)

		}
		switch locale {
		case "en":
			entranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			cntranslations.RegisterDefaultTranslations(v, trans)
		default:
			entranslations.RegisterDefaultTranslations(v, trans)
		}
	}

	return nil

}

func main() {
	
	if err := InitTrans("zh"); err != nil {
		fmt.Println("init trans error")
		return
	}
	router := gin.Default()
	router.POST("/loginJSON", func(c *gin.Context) {
		var loginForm LoginForm
		if err := c.ShouldBind(&loginForm); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				fmt.Println(err.Error())
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"error": removeTopStruct(errs.Translate(trans)),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": "login ok",
		})

	})

	router.POST("/signup", func(c *gin.Context) {

		var signupForm SignUpForm
		if err := c.ShouldBind(&signupForm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": "registration ok",
		})

	})

	_ = router.Run(":8083")

}
