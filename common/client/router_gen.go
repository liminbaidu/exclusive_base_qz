package client

import (
	"exclusive_base_qz/handler"

	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func register(s *ghttp.Server) {
	s.Group("/common", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.GET("/login", CommonLogin)
		group.GET("/signout", CommonSignout)
		group.GET("/updateUserInfo", CommonUpdateUserInfo)
	})
	s.Group("/income", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.GET("/createIncome", CreateIncome)
		group.GET("/updateIncome", UpdateIncome)
		group.GET("/deleteIncome", DeleteIncome)
		group.GET("/queryIncome", QueryIncome)
	})
	s.Group("/diary", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.GET("/createDiary", CreateDiary)
		group.GET("/updateDiary", UpdateDiary)
		group.GET("/deleteDiary", DeleteDiary)
		group.GET("/queryDiary", QueryDiary)
	})
}

func CommonLogin(r *ghttp.Request) {
	resp := handler.CommonLogin(r)
	r.Response.Writef(resp)
}

func CommonSignout(r *ghttp.Request) {
	resp := handler.CommonSignout(r)
	r.Response.Writef(resp)
}

func CommonUpdateUserInfo(r *ghttp.Request) {
	resp := handler.CommonUpdateUserInfo(r)
	r.Response.Writef(resp)
}

func CreateIncome(r *ghttp.Request) {
	resp := handler.CreateIncome(r)
	r.Response.Writef(resp)
}

func UpdateIncome(r *ghttp.Request) {
	resp := handler.UpdateIncome(r)
	r.Response.Writef(resp)
}

func DeleteIncome(r *ghttp.Request) {
	resp := handler.DeleteIncome(r)
	r.Response.Writef(resp)
}

func QueryIncome(r *ghttp.Request) {
	resp := handler.QueryIncome(r)
	r.Response.Writef(resp)
}

func CreateDiary(r *ghttp.Request) {
	resp := handler.CreateDiary(r)
	r.Response.Writef(resp)
}

func UpdateDiary(r *ghttp.Request) {
	resp := handler.UpdateDiary(r)
	r.Response.Writef(resp)
}

func DeleteDiary(r *ghttp.Request) {
	resp := handler.DeleteDiary(r)
	r.Response.Writef(resp)
}
func QueryDiary(r *ghttp.Request) {
	resp := handler.QueryDiary(r)
	r.Response.Writef(resp)
}
