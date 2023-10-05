package controllers

import (
	"cd-docker/common"
	"cd-docker/configs"
	"cd-docker/docker"
	"github.com/gin-gonic/gin"
	"net/http"
)

type reqFaceUpdateServiceLatest struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Token string `json:"token" form:"token" validate:"required"`
}

type resFaceUpdateServiceLatest struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func UpdateServiceLatest() gin.HandlerFunc {
	return func(c *gin.Context) {
		configs.IniSetup()
		var reqBody reqFaceUpdateServiceLatest
		if !common.ValidBindForm(c, &reqBody) {
			return
		}

		section, err := configs.IniData.GetSection(reqBody.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, resFaceUpdateServiceLatest{
				Name:    reqBody.Name,
				Status:  "failed",
				Message: "service does not exist!",
			})
			return
		}

		if section.Key("token").String() != reqBody.Token {
			c.JSON(http.StatusUnauthorized, resFaceUpdateServiceLatest{
				Name:    reqBody.Name,
				Status:  "failed",
				Message: "token is invalid!",
			})
			return
		}

		err = docker.UpdateDockerServiceLatest(section.Key("serviceName").String())
		if err == nil {
			c.JSON(http.StatusOK, resFaceUpdateServiceLatest{Name: reqBody.Name, Status: "ok"})
			return
		}

		c.JSON(http.StatusInternalServerError, resFaceUpdateServiceLatest{
			Name:    reqBody.Name,
			Status:  "failed",
			Message: "see logs in cd-docker logs",
		})
		return
	}
}
