package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"book.api/train4/internal/model"
	"book.api/train4/internal/services"
)

func CreateMember(c *gin.Context) {

	var member model.Member

	err := c.ShouldBindJSON(&member)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	newMember, err := services.CreateMember(member)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newMember)
}

func ShowMembers(c *gin.Context) {
	members, err := services.GetMembers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, members)
}

func ShowMember(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	member, err := services.GetMember(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "member not found"})
		return
	}

	c.JSON(http.StatusOK, member)
}
func ChangeInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var member model.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	changedMember, err := services.ChangeMember(id, member)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, changedMember)
}

func DeleteMember(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := services.DeleteMember(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "member not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

func CreateTrainer(c *gin.Context) {

	var trainer model.Trainer

	err := c.ShouldBindJSON(&trainer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	newTrainer, err := services.CreateTrainer(trainer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newTrainer)
}

func ShowTrainers(c *gin.Context) {
	trainer, err := services.GetTrainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, trainer)
}

func ShowTrainer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	trainer, err := services.GetTrainer(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "trainer not found"})
		return
	}

	c.JSON(http.StatusOK, trainer)
}
func ChangeInfoTrainer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var trainer model.Trainer
	if err := c.ShouldBindJSON(&trainer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	changedTrainer, err := services.ChangeTrainer(id, trainer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, changedTrainer)
}

func DeleteTrainer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := services.DeleteTrainer(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "trainer not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

func CreateClass(c *gin.Context) {

	var class model.Class

	err := c.ShouldBindJSON(&class)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid input",
		})
		return
	}

	newClass, err := services.CreateClass(c.Request.Context(), class)

	if err != nil {

		if err.Error() == "trainer not found" {
			c.JSON(404, gin.H{
				"error": "trainer not found",
			})
			return
		}

		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, newClass)
}

func ShowClasses(c *gin.Context) {
	classes, err := services.ShowClasses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, classes)
}
func ShowClass(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	class, err := services.ShowClass(id)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"errer": "class not found"})
	}
	c.JSON(http.StatusOK, class)

}
func Register(c *gin.Context) {
	var register model.Register
	err := c.ShouldBindJSON(&register)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid input",
		})
		return
	}

	newRegister, err := services.Register(register.MemberID, register.ClassID)
	if err != nil {
		if err.Error() == "member not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "member not found"})
			return
		}
		if err.Error() == "class not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "class not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newRegister)
}
func ClassesforMember(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	classes, err := services.ClassesOfthisMember(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, classes)
}
func MembersOfclass(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	members, err := services.MemberOfthisClass(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, members)

}
func SearchMembers(c *gin.Context){}