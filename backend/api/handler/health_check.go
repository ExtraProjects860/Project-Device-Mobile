package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary      API status check
// @Description  Returns a simple message to confirm API is running
// @Tags         health
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /health/api [get]
func ApiHandler(ctx *gin.Context) {
	sendStatus(ctx, http.StatusOK, "Api is ok")
}

// @Summary      Database status check
// @Description  Tests connection to the database and returns status
// @Tags         health
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /health/database [get]
func DatabaseHandler(ctx *gin.Context, db *gorm.DB) {
	if err := config.TestConnectionSQL(db); err != nil {
		logger.Errorf("database connection test failed: %v", err)
		sendErr(
			ctx,
			http.StatusInternalServerError,
			fmt.Errorf("database connection failed: %v", err.Error()),
		)
		return
	}

	sendStatus(ctx, http.StatusOK, "Database connected")
}

// @Summary      Email service check
// @Description  Calls external email service to verify availability
// @Tags         health
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /health/email [get]
func EmailServiceHandler(ctx *gin.Context, serverDomain string) {
	resp, err := http.Get(serverDomain)
	if err != nil {
		logger.Errorf("failed to call email service at %s: %v", serverDomain, err)
		sendErr(ctx, http.StatusInternalServerError, errors.New("unable to reach email service"))
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("failed to read response body from %s: %v", serverDomain, err.Error())
		sendErr(ctx, http.StatusInternalServerError, errors.New("failed to read response from email service"))
		return
	}

	sendStatus(ctx, http.StatusOK, "Email service called successfully")
}
