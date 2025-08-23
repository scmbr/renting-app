package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/pkg/error"
)

const (
	DefaultLimit = 20
	MaxLimit     = 100
)

type AdvertListResponse struct {
	Total   int64                    `json:"total"`
	Adverts []*dto.GetAdvertResponse `json:"adverts"`
}

// @Summary      Получить все объявления пользователя
// @Description  Возвращает список объявлений, созданных авторизованным пользователем, с поддержкой фильтров, сортировки и пагинации
// @Tags         adverts
// @Security     ApiKeyAuth
// @Produce      json
// @Param        city query string false "Город"
// @Param        district query string false "Район"
// @Param        bathroom_type query string false "Тип санузла"
// @Param        remont query string false "Ремонт"
// @Param        rental_type query string false "Тип аренды"
// @Param        rooms query int false "Количество комнат"
// @Param        price_from query int false "Минимальная цена"
// @Param        price_to query int false "Максимальная цена"
// @Param        floor_min query int false "Минимальный этаж"
// @Param        floor_max query int false "Максимальный этаж"
// @Param        year_min query int false "Минимальный год постройки"
// @Param        year_max query int false "Максимальный год постройки"
// @Param        apartment_rating_min query number false "Минимальный рейтинг квартиры"
// @Param        landlord_rating_min query number false "Минимальный рейтинг арендодателя"
// @Param        elevator query boolean false "Наличие лифта"
// @Param        concierge query boolean false "Наличие консьержа"
// @Param        pets query boolean false "Можно с животными"
// @Param        babies query boolean false "Можно с детьми"
// @Param        smoking query boolean false "Можно курить"
// @Param        internet query boolean false "Наличие интернета"
// @Param        washing_machine query boolean false "Наличие стиральной машины"
// @Param        tv query boolean false "Наличие телевизора"
// @Param        conditioner query boolean false "Наличие кондиционера"
// @Param        dishwasher query boolean false "Наличие посудомойки"
// @Param        lat query number false "Широта для геолокации"
// @Param        lng query number false "Долгота для геолокации"
// @Param        limit query int false "Лимит записей, по умолчанию 20"
// @Param        offset query int false "Смещение, по умолчанию 0"
// @Param        sort_by query string false "Поле сортировки, по умолчанию created_at"
// @Param        order query string false "Направление сортировки: asc или desc, по умолчанию desc"
// @Success      200 {object} AdvertListResponse
// @Failure      401 {object} error.Response "Пользователь не авторизован"
// @Failure      500 {object} error.Response "Внутренняя ошибка сервера"
// @Router       /my/advert [get]
func (h *Handler) getAllUserAdverts(c *gin.Context) {
	filter := bindAdvertFilter(c)

	userId := c.MustGet("userId").(int)

	adverts, total, err := h.services.Advert.GetAllUserAdverts(c.Request.Context(), userId, &filter)
	if err != nil {
		error.Internal(c, err)
		return
	}
	c.JSON(http.StatusOK, AdvertListResponse{
		Total:   total,
		Adverts: adverts,
	})
}

// @Summary      Получить объявление пользователя по ID
// @Description  Возвращает объявление пользователя по ID
// @Tags         adverts
// @Security     ApiKeyAuth
// @Produce      json
// @Param        id   path      int  true  "Advert ID"
// @Success      200  {object}  dto.GetAdvertResponse
// @Failure      400  {object}  error.Response "Некорректный ID объявления"
// @Failure      401  {object}  error.Response "Пользователь не авторизован"
// @Failure      500  {object}  error.Response "Внутренняя ошибка сервера"
// @Router       /my/advert/{id} [get]
func (h *Handler) getUserAdvertById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		error.Send(c, http.StatusBadRequest, error.ErrInvalidAdvertID)
		return
	}

	userId := c.MustGet("userId").(int)

	advert, err := h.services.Advert.GetUserAdvertById(c.Request.Context(), userId, id)
	if err != nil {
		error.Internal(c, err)
		return
	}
	c.JSON(http.StatusOK, advert)
}

// @Summary      Создать объявление
// @Description  Создает объявление авторизированному пользователю
// @Tags         adverts
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        input  body      dto.CreateAdvertInput  true  "Advert input"
// @Success      201  {object}  dto.GetAdvertResponse
// @Failure      400  {object}  error.Response "Некорректные входные данные объявления"
// @Failure      401  {object}  error.Response "Пользователь не авторизован"
// @Failure      500  {object}  error.Response "Внутренняя ошибка сервера"
// @Router       /advert [post]
func (h *Handler) createAdvert(c *gin.Context) {
	userId := c.MustGet("userId").(int)

	var input dto.CreateAdvertInput
	if err := c.ShouldBindJSON(&input); err != nil {
		error.Send(c, http.StatusBadRequest, error.ErrInvalidAdvertInput)
		return
	}

	advert, err := h.services.Advert.CreateAdvert(c.Request.Context(), userId, input)
	if err != nil {
		error.Internal(c, err)
		return
	}
	c.JSON(http.StatusCreated, advert)
}

// @Summary      Удалить объявление
// @Tags         adverts
// @Security     ApiKeyAuth
// @Produce      json
// @Param        id   path      int  true  "Advert ID"
// @Success      200
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /advert/{id} [delete]
func (h *Handler) deleteAdvert(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid advert id")
		return
	}

	userId := c.MustGet("userId").(int)

	err = h.services.Advert.DeleteAdvert(c.Request.Context(), userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

// @Summary      Обновить объявление
// @Tags         adverts
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        id     path      int                   true  "Advert ID"
// @Param        input  body      dto.UpdateAdvertInput true  "Update input"
// @Success      200
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /advert/{id} [patch]
func (h *Handler) updateAdvert(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid advert id")
		return
	}

	userId := c.MustGet("userId").(int)

	var input dto.UpdateAdvertInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Advert.UpdateAdvert(c.Request.Context(), userId, id, &input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// @Summary      Получение списка объявлений
// @Description  Возвращает список всех объявлений с фильтрацией. Если пользователь авторизован — добавляет флаг избранного.
// @Tags         adverts
// @Accept       json
// @Produce      json
// @Param        city query string false "Город"
// @Param        district query string false "Район"
// @Param        bathroom_type query string false "Тип санузла"
// @Param        remont query string false "Ремонт"
// @Param        rental_type query string false "Тип аренды"
// @Param        rooms query int false "Количество комнат"
// @Param        price_from query int false "Минимальная цена"
// @Param        price_to query int false "Максимальная цена"
// @Param        floor_min query int false "Минимальный этаж"
// @Param        floor_max query int false "Максимальный этаж"
// @Param        year_min query int false "Минимальный год постройки"
// @Param        year_max query int false "Максимальный год постройки"
// @Param        apartment_rating_min query number false "Минимальный рейтинг квартиры"
// @Param        landlord_rating_min query number false "Минимальный рейтинг арендодателя"
// @Param        elevator query boolean false "Наличие лифта"
// @Param        concierge query boolean false "Наличие консьержа"
// @Param        pets query boolean false "Можно с животными"
// @Param        babies query boolean false "Можно с детьми"
// @Param        smoking query boolean false "Можно курить"
// @Param        internet query boolean false "Наличие интернета"
// @Param        washing_machine query boolean false "Наличие стиральной машины"
// @Param        tv query boolean false "Наличие телевизора"
// @Param        conditioner query boolean false "Наличие кондиционера"
// @Param        dishwasher query boolean false "Наличие посудомойки"
// @Param        lat query number false "Широта для геолокации"
// @Param        lng query number false "Долгота для геолокации"
// @Param        limit query int false "Лимит записей, по умолчанию 20"
// @Param        offset query int false "Смещение, по умолчанию 0"
// @Param        sort_by query string false "Поле сортировки, по умолчанию created_at"
// @Param        order query string false "Направление сортировки: asc или desc, по умолчанию desc"
// @Success      200  {object}  AdvertListResponse
// @Failure      500  {object}  error.Response "Внутренняя ошибка сервера"
// @Router       /adverts [get]
func (h *Handler) getAllAdverts(c *gin.Context) {
	filter := bindAdvertFilter(c)
	userIdVal, exists := c.Get("userId")
	var userId *int
	if exists {
		if uid, ok := userIdVal.(int); ok {
			userId = &uid
		}
	}
	adverts, total, err := h.services.Advert.GetAllAdverts(c.Request.Context(), userId, &filter)
	if err != nil {
		error.Internal(c, err)
		return
	}

	c.JSON(http.StatusOK, AdvertListResponse{
		Total:   total,
		Adverts: adverts,
	})
}

// @Summary      Получение объявления по ID
// @Description  Возвращает подробности одного объявления по его идентификатору
// @Tags         adverts
// @Accept       json
// @Produce      json
// @Param        id path int true "ID объявления"
// @Success      200  {object}  dto.GetAdvertResponse
// @Failure      400  {object}  error.Response "Неверный ID объявления"
// @Failure      500  {object}  error.Response "Внутренняя ошибка сервера"
// @Router       /adverts/{id} [get]
func (h *Handler) getAdvertById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		error.Send(c, http.StatusBadRequest, error.ErrInvalidAdvertID)
		return
	}

	advert, err := h.services.Advert.GetAdvertById(c.Request.Context(), id)
	if err != nil {
		error.Internal(c, err)
		return
	}
	c.JSON(http.StatusOK, advert)
}

func parseIntOrZero(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func parseFloat32OrZero(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
}

func parseIntOrDefault(s string, def int) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return n
}

func parseBoolPointer(s string) *bool {
	if s == "" {
		return nil
	}
	b, _ := strconv.ParseBool(s)
	return &b
}
func bindAdvertFilter(c *gin.Context) dto.AdvertFilter {
	filter := dto.AdvertFilter{
		City:               c.Query("city"),
		District:           c.Query("district"),
		BathroomType:       c.Query("bathroom_type"),
		Remont:             c.Query("remont"),
		RentalType:         c.Query("rental_type"),
		Rooms:              parseIntOrZero(c.Query("rooms")),
		PriceMin:           parseIntOrZero(c.Query("price_from")),
		PriceMax:           parseIntOrZero(c.Query("price_to")),
		FloorMin:           parseIntOrZero(c.Query("floor_min")),
		FloorMax:           parseIntOrZero(c.Query("floor_max")),
		YearMin:            parseIntOrZero(c.Query("year_min")),
		YearMax:            parseIntOrZero(c.Query("year_max")),
		ApartmentRatingMin: parseFloat32OrZero(c.Query("apartment_rating_min")),
		LandlordRatingMin:  parseFloat32OrZero(c.Query("landlord_rating_min")),
		Limit:              parseIntOrDefault(c.Query("limit"), DefaultLimit),
		Offset:             parseIntOrDefault(c.Query("offset"), 0),
		SortBy:             c.DefaultQuery("sort_by", "created_at"),
		Order:              c.DefaultQuery("order", "desc"),
		Lat:                parseFloat32OrZero(c.Query("lat")),
		Lng:                parseFloat32OrZero(c.Query("lng")),
	}

	boolFields := map[string]**bool{
		"elevator":        &filter.Elevator,
		"concierge":       &filter.Concierge,
		"pets":            &filter.Pets,
		"babies":          &filter.Babies,
		"smoking":         &filter.Smoking,
		"internet":        &filter.Internet,
		"washing_machine": &filter.WashingMachine,
		"tv":              &filter.TV,
		"conditioner":     &filter.Conditioner,
		"dishwasher":      &filter.Dishwasher,
	}

	for key, ptr := range boolFields {
		*ptr = parseBoolPointer(c.Query(key))
	}

	if filter.Limit <= 0 || filter.Limit > MaxLimit {
		filter.Limit = DefaultLimit
	}
	if filter.Offset < 0 {
		filter.Offset = 0
	}
	validOrders := map[string]bool{"asc": true, "desc": true}
	if !validOrders[filter.Order] {
		filter.Order = "desc"
	}

	return filter
}
