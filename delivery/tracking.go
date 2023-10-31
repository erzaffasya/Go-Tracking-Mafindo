package delivery

import (
	"net/http"
	"strconv"

	"github.com/erzaffasya/Go-Tracking-Mafindo/middlewares"
	"github.com/erzaffasya/Go-Tracking-Mafindo/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Tracking struct {
	cuc models.TrackingUsecase
}

func NewSocialMediaRoute(handlers *gin.Engine, cuc models.TrackingUsecase) {
	route := &Tracking{cuc}

	handler := handlers.Group("/socialmedias")
	{
		handler.Use(middlewares.Authentication())
		handler.GET("/", route.Fetch)
		handler.POST("/", route.Store)
		handler.PUT("/:id", middlewares.SocialMediaAuthorization(route.cuc), route.Update)
		handler.DELETE("/:id", middlewares.SocialMediaAuthorization(route.cuc), route.Delete)
	}
}

// ? perlu melewati proses autentikasi dan autorisasi terlebih dahulu
// ? hanya get socialMedia sendiri?
// Fetch godoc
// @Summary      Fetch socialMedias
// @Description  get socialMedias
// @Tags         socialMedias
// @Accept       json
// @Produce      json
// @Success      200	{object}	[]models.Tracking
// @Failure      400	{object}	ErrorResponse
// @Failure      401	{object}	ErrorResponse
// @Security     Bearer
// @Router       /socialMedias  [get]
func (route *Tracking) Fetch(c *gin.Context) {
	var (
		socialMedias []models.Tracking
		err          error
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err = route.cuc.Fetch(c.Request.Context(), &socialMedias, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, socialMedias)
}

// Store godoc
// @Summary      Create an socialMedia
// @Description  create and store an socialMedia
// @Tags         socialMedias
// @Accept       json
// @Produce      json
// @Param        json  body  models.Tracking true  "SocialMedia"
// @Success      201  {object}  models.Tracking
// @Failure      400	{object}	ErrorResponse
// @Failure      401	{object}	ErrorResponse
// @Security     Bearer
// @Router       /socialMedias  [post]
func (route *Tracking) Store(c *gin.Context) {
	var (
		socialMedia models.Tracking
		err         error
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err = c.ShouldBindJSON(&socialMedia)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	Tracking.UsersId = userID

	err = route.cuc.Store(c.Request.Context(), &socialMedia)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               Tracking.ID,
		"user_id":          Tracking.UserID,
		"name":             Tracking.Name,
		"social_media_url": Tracking.SocialMediaUrl,
		"created_at":       Tracking.CreatedAt,
	})
}

// Update godoc
// @Summary      Update an socialMedia
// @Description  update an socialMedia by ID
// @Tags         socialMedias
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "SocialMedia ID"
// @Success      200  {string}  string
// @Failure      400  {object}	ErrorResponse
// @Failure      401  {object}	ErrorResponse
// @Failure      404  {object}	ErrorResponse
// @Security     Bearer
// @Router       /socialMedias/{id} [put]
func (route *Tracking) Update(c *gin.Context) {
	var (
		socialMedia models.Tracking
		err         error
	)

	socialMediaIDStr := c.Param("id")
	socialMediaIDInt, err := strconv.Atoi(socialMediaIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Failed to cast social media id to int",
		})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err = c.ShouldBindJSON(&socialMedia)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	socialMediaID := uint(socialMediaIDInt)

	updatedSocialMedia := models.Tracking{
		UserID:         userID,
		Name:           Tracking.Name,
		SocialMediaUrl: Tracking.SocialMediaUrl,
	}

	socialMedia, err = route.cuc.Update(c.Request.Context(), updatedSocialMedia, socialMediaID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               Tracking.ID,
		"user_id":          Tracking.UserID,
		"name":             Tracking.Name,
		"social_media_url": Tracking.SocialMediaUrl,
		"updated_at":       Tracking.UpdatedAt,
	})
}

// Delete godoc
// @Summary      Delete an socialMedia
// @Description  delete an socialMedia by ID
// @Tags         socialMedias
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "SocialMedia ID"
// @Success      200  {string}  string
// @Failure      400  {object}	ErrorResponse
// @Failure      401  {object}	ErrorResponse
// @Failure      404  {object}	ErrorResponse
// @Security     Bearer
// @Router       /socialMedias/{id} [delete]
func (route *Tracking) Delete(c *gin.Context) {
	socialMediaIDStr := c.Param("id")
	socialMediaIDInt, err := strconv.Atoi(socialMediaIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Failed to cast social media id to int",
		})
		return
	}

	socialMediaID := uint(socialMediaIDInt)

	err = route.cuc.Delete(c.Request.Context(), socialMediaID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Your social media has been successfully deleted",
		},
	)
}
