/*
 * Copyright (c) 2018
 * time:   6/24/18 3:22 PM
 * author: linhuanchao
 * e-mail: 873085747@qq.com
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"SensitiveWords/tool"
	"net/http"
)

func Check(context *gin.Context)  {
	content := context.Query("content")
	sensitiveMap := tool.GetMap()
	target, result := sensitiveMap.CheckSensitive(content)
	context.JSON(http.StatusOK, gin.H{
		"target" : target,
		"result" : result,
	})
}

func All(context *gin.Context)  {
	content := context.Query("content")
	sensitiveMap := tool.GetMap()
	target := sensitiveMap.FindAllSensitive(content)

	type Target struct {
		Word string `json:"word"`
		I []int `json:"i"`
		L int `json:"l"`
	}

	targetArray := []Target{}

	for key, value := range target{
		t := Target{
			Word: key,
			I:value.Indexes,
			L:value.Len,
		}
		targetArray = append(targetArray, t)
	}

	context.JSON(http.StatusOK,gin.H{
		"target": targetArray,
	})
}