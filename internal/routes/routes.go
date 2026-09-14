package routes

import (
	controllers "book.api/train4/internal/controller"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/members", controllers.CreateMember)
	r.GET("/members", controllers.ShowMembers)
	r.GET("/members/:id", controllers.ShowMember)
	r.PATCH("/members/:id", controllers.ChangeInfo)
	r.DELETE("/members/:id", controllers.DeleteMember)

	r.POST("/trainers", controllers.CreateTrainer)
	r.GET("/trainers", controllers.ShowTrainers)
	r.GET("/trainers/:id", controllers.ShowTrainer)
	r.PATCH("/trainers/:id", controllers.ChangeInfoTrainer)
	r.DELETE("/trainers/:id", controllers.DeleteTrainer)

	r.POST("/classes", controllers.CreateClass)
	r.GET("/classes", controllers.ShowClasses)
	r.GET("/classes/:id", controllers.ShowClass)

	r.POST("/register", controllers.Register)
	r.GET("/members/:id/classes", controllers.ClassesforMember)
	r.GET("/classes/:id/members", controllers.MembersOfclass)
	r.GET("/members/search", controllers.SearchMembers)
}
