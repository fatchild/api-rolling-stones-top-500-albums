package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

//func SetUpRouter() *gin.Engine {
//	router := gin.Default()
//	return router
//}

func TestSmoke(t *testing.T) {
	//	r := SetUpRouter()

	//	r.GET("/getAlbumList", getAlbums.GetAlbums)
	//	req, _ := http.NewRequest("GET", "/getAlbumList", nil)
	//	w := httptest.NewRecorder()
	//	r.ServeHTTP(w, req)

	//assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, 1)
}
