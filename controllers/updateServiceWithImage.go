package controllers

import (
	"cd-docker/common"
	"cd-docker/configs"
	"cd-docker/docker"
	"github.com/gin-gonic/gin"
	"net/http"
)

type reqFaceUpdateServiceWithImage struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Token string `json:"token" form:"token" validate:"required"`
	Image string `json:"image" form:"image" validate:"required"`
}

type resFaceUpdateServiceWithImage struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func UpdateServiceWithImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		configs.IniSetup()
		var reqBody reqFaceUpdateServiceWithImage
		if !common.ValidBindForm(c, &reqBody) {
			return
		}

		section, err := configs.IniData.GetSection(reqBody.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, resFaceUpdateServiceWithImage{
				Name:    reqBody.Name,
				Status:  "failed",
				Message: "service does not exist!",
			})
			return
		}

		if section.Key("token").String() != reqBody.Token {
			c.JSON(http.StatusUnauthorized, resFaceUpdateServiceWithImage{
				Name:    reqBody.Name,
				Status:  "failed",
				Message: "token is invalid!",
			})
			return
		}

		err = docker.UpdateDockerService(section.Key("serviceName").String(), reqBody.Image)

		if err == nil {
			c.JSON(http.StatusOK, resFaceUpdateServiceWithImage{Name: reqBody.Name, Status: "ok"})
			return
		}

		c.JSON(http.StatusInternalServerError, resFaceUpdateServiceWithImage{
			Name:    reqBody.Name,
			Status:  "failed",
			Message: "see logs in cd-docker logs",
		})
		return
	}
}
