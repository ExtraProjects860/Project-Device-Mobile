package handler

import (
	"io"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/gin-gonic/gin"
)

func Api(ctx *gin.Context) {
	sendStatus(ctx, "Api is ok")
}

func Database(ctx *gin.Context) {
	if err := config.TestConnectionSQL(); err != nil {
		logger.Errorf("database connection test failed: %v", err)
		sendErr(
			ctx, 
			http.StatusInternalServerError, 
			gin.H{"status": "Database connection failed", "error": err.Error()},
		)
		return
	}

	sendStatus(ctx, "Database connected")
}

func EmailService(ctx *gin.Context) {
	url := config.GetEnv().API.EmailService

	resp, err := http.Get(url)
	if err != nil {
		logger.Errorf("failed to call email service at %s: %v", url, err)
		sendErr(ctx, http.StatusInternalServerError, gin.H{"error": "unable to reach email service"})
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("failed to read response body from %s: %v", url, err.Error())
		sendErr(ctx, http.StatusInternalServerError, gin.H{"error": "failed to read response from email service"})
		return
	}

	sendStatus(ctx, "email service called successfully")
}

