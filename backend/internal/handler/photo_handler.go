package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/renting-app/internal/dto"
)

func (h *Handler) getAllPhotos(c *gin.Context) {
	apartmentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid apartment id")
		return
	}

	photos, err := h.services.ApartmentPhoto.GetAllPhotos(c, apartmentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, photos)
}

func (h *Handler) getPhotoById(c *gin.Context) {
	apartmentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid apartment id")
		return
	}

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid photo id")
		return
	}

	photo, err := h.services.ApartmentPhoto.GetPhotoById(c, apartmentId, photoId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, photo)
}

func (h *Handler) addPhotos(c *gin.Context) {
	apartmentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid apartment id")
		return
	}

	userId := c.GetInt("userId")

	form, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid multipart form")
		return
	}

	files := form.File["photos"]
	if len(files) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "no photos provided")
		return
	}

	var inputs []dto.CreatePhotoInput

	for i, fileHeader := range files {

		url, err := h.services.ApartmentPhoto.UploadPhotoToS3(c.Request.Context(), fileHeader)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, "failed to upload photo to S3")
			return
		}

		input := dto.CreatePhotoInput{
			URL:      url,
			FileName: fileHeader.Filename,
			IsCover:  i == 0,
		}
		inputs = append(inputs, input)
	}

	err = h.services.ApartmentPhoto.AddPhotoBatch(c, userId, apartmentId, inputs)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"uploaded_photos": inputs,
	})
}

func (h *Handler) deletePhoto(c *gin.Context) {
	apartmentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid apartment id")
		return
	}

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid photo id")
		return
	}

	userId := c.GetInt("userId")

	err = h.services.ApartmentPhoto.DeletePhoto(c, userId, apartmentId, photoId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) setCover(c *gin.Context) {
	apartmentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid apartment id")
		return
	}

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid photo id")
		return
	}

	userId := c.GetInt("userId")

	err = h.services.ApartmentPhoto.SetCover(c, userId, apartmentId, photoId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
