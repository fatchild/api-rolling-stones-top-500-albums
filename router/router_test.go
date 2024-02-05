package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/getAlbumByID"
	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/getAlbums"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetAlbumsStatusCode200(t *testing.T) {
	r := SetUpRouter()

	r.GET("/getAlbumList", getAlbums.GetAlbums)
	req, _ := http.NewRequest("GET", "/getAlbumList", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAlbumByIDStatusCode200(t *testing.T) {
	r := SetUpRouter()

	r.GET("/getAlbumAtPosition/:position", getAlbumByID.GetAlbumByID)
	req, _ := http.NewRequest("GET", "/getAlbumAtPosition/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
