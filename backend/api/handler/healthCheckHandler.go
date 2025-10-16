package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/ExtraProjects860/Project-Device-Mobile/handler/response"
	"github.com/gin-gonic/gin"
)

// @Summary      API status check
// @Description  Returns a simple message to confirm API is running
// @Tags         health
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /health/api [get]
func ApiHandler(ctx *gin.Context) {
	response.SendStatus(ctx, http.StatusOK, "Api is ok")
}

// @Summary      Database status check
// @Description  Tests connection to the database and returns status
// @Tags         health
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /health/database [get]
func DatabaseHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := config.TestConnectionSQL(appCtx.DB); err != nil {
			logger.Errorf("database connection test failed: %v", err)
			response.SendErr(
				ctx,
				http.StatusInternalServerError,
				fmt.Errorf("database connection failed: %v", err.Error()),
			)
			return
		}

		response.SendStatus(ctx, http.StatusOK, "Database connected")
	}

}

// @Summary      Email service check
// @Description  Calls external email service to verify availability
// @Tags         health
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /health/email [get]
func EmailServiceHandler(appCtx *appcontext.AppContext, logger *config.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		serverDomain := appCtx.Env.API.EmailService

		resp, err := http.Get(serverDomain)
		if err != nil {
			logger.Errorf("failed to call email service at %s: %v", serverDomain, err)
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("unable to reach email service"))
			return
		}
		defer resp.Body.Close()

		_, err = io.ReadAll(resp.Body)
		if err != nil {
			logger.Errorf("failed to read response body from %s: %v", serverDomain, err.Error())
			response.SendErr(ctx, http.StatusInternalServerError, errors.New("failed to read response from email service"))
			return
		}

		response.SendStatus(ctx, http.StatusOK, "Email service called successfully")
	}

}
