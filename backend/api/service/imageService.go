package service

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/ExtraProjects860/Project-Device-Mobile/appcontext"
	"github.com/ExtraProjects860/Project-Device-Mobile/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

var (
	FolderUser    string = "project-device-mobile/image/upload/users"
	FolderProduct        = "project-device-mobile/image/upload/product"
)

type ImageService struct {
	cld    *cloudinary.Cloudinary
	logger *config.Logger
}

func GetImageService(appCtx *appcontext.AppContext) ImageService {
	return ImageService{
		cld:    appCtx.Cloudinary,
		logger: config.NewLogger("SERVICE - IMAGE"),
	}
}

func (s *ImageService) UploadImage(ctx *gin.Context, folderToSave string) (*string, string, error) {
	file, err := s.managerPhotoUrl(ctx, "image")
	if err != nil {
		s.logger.Errorf("Error managing photo from context: %v", err)
		return nil, "", err
	}

	if file == nil {
		s.logger.Info("No photo provided, skipping upload.")
		return nil, "", nil
	}

	s.logger.Infof("Photo provided, uploading to folder: %s", folderToSave)
	secureURL, publicID, err := s.savePhoto(file, folderToSave)
	if err != nil {
		return nil, "", err
	}

	return &secureURL, publicID, nil
}

func (s *ImageService) RemoveImage(publicID string) error {
	if publicID == "" {
		return nil
	}

	s.logger.Warningf("Rolling back photo upload: %s", publicID)
	response, err := s.cld.Upload.Destroy(
		context.Background(),
		uploader.DestroyParams{
			PublicID: publicID,
		},
	)
	if err != nil {
		s.logger.Errorf("Failed to rollback photo (API error): %v", err)
		return err
	}

	if response.Error.Message != "" && response.Result != "ok" {
		err := fmt.Errorf("failed to delete image with public ID %s. Result: %s", publicID, response.Result)
		s.logger.Error(err.Error())
		return err
	}

	s.logger.Infof("Successfully rolled back photo: %s", publicID)
	return nil
}

func (s *ImageService) managerPhotoUrl(ctx *gin.Context, imageKey string) (*multipart.FileHeader, error) {
	fileAny, exists := ctx.Get(imageKey)
	if !exists {
		return nil, nil
	}

	file, ok := fileAny.(*multipart.FileHeader)
	if !ok {
		s.logger.Error("Error to convert fileAny to FileHeader type")
		return nil, fmt.Errorf("invalid photo type in context")
	}

	return file, nil
}

func (s *ImageService) savePhoto(file *multipart.FileHeader, folderToSave string) (string, string, error) {
	src, err := file.Open()
	if err != nil {
		s.logger.Errorf("Failed to open uploaded file: %v", err)
		return "", "", err
	}
	defer src.Close()

	uploadParams := uploader.UploadParams{
		Folder: folderToSave,
	}

	uploadResult, err := s.cld.Upload.Upload(
		context.Background(),
		src,
		uploadParams,
	)
	if err != nil {
		s.logger.Errorf("Failed to upload to Cloudinary: %v", err)
		return "", "", err
	}

	return uploadResult.SecureURL, uploadResult.PublicID, nil
}
