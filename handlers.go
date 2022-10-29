package main

import (
	"atomicgo.dev/assert"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Fields struct {
	Uri          string
	StatusCode   int
	ResponseBody map[string]interface{}
}

func ApiTester(c *gin.Context) {
	var input Fields

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req, err := Client(input.Uri)
	if err != nil {
		return
	}

	responseStatus := req.StatusCode
	response := req.Body

	compareStatus := assert.Equal(input.StatusCode, responseStatus)
	compareBody := assert.Equal(input.ResponseBody, response)

	if compareStatus == true && compareBody == true {
		c.JSON(http.StatusOK, gin.H{"message": "SUCCESS, TEST PASSED", "Status Expected": input.StatusCode, "Status Received": responseStatus,
			"Body Expected": input.ResponseBody, "Body Received": response})
	}

	if compareStatus == false || compareBody == false {
		c.JSON(http.StatusOK, gin.H{"message": "Test Failed", "Status Expected": input.StatusCode, "Status Received": responseStatus,
			"Body Expected": input.ResponseBody, "Body Received": response})
	}

}
