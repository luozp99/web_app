package context

//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"web_app/response"
//)
//
//type AppContext struct {
//	*gin.Context
//}
//
//func (c *AppContext) Success(){
//	c.JSON(http.StatusOK,response.Success())
//}
//
//func (c *AppContext) SuccessMsg(msg string){
//	c.JSON(http.StatusOK,response.SuccessMsg(msg))
//}
//
//func (c *AppContext) Error(){
//	c.JSON(http.StatusInternalServerError,response.Error())
//}
//
//func (c *AppContext) ErrorMsg(msg string){
//	c.JSON(http.StatusInternalServerError,response.ErrorMsg(msg))
//}
