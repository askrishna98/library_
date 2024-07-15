package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/askrishna98/library_/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandlers(t *testing.T) {
	assert := assert.New(t)

	router := gin.Default()
	Handlers(router)

	t.Run("Test_Create_MemberAPI", func(t *testing.T) {

		testMember := &models.Member{
			Name:  "test",
			Phone: "123456789",
			Email: "test@example",
		}

		memberJson, _ := json.Marshal(testMember)

		req, _ := http.NewRequest("POST", "/api/members", strings.NewReader(string(memberJson)))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusOK, w.Code)
		var respose models.Member
		err := json.Unmarshal(w.Body.Bytes(), &respose)
		assert.NoError(err)
		assert.Equal("A011", respose.Member_id)
	})

	t.Run("Test_GetMemberinfobyIDAPI", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "/api/members/A001", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusOK, w.Code)
		var respose models.Member

		err := json.Unmarshal(w.Body.Bytes(), &respose)
		assert.NoError(err)
		assert.Equal("A001", respose.Member_id)

		// testing - accessing data which doesnt exist

		req, _ = http.NewRequest("GET", "/api/members/A012", nil)
		w = httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusInternalServerError, w.Code)
	})

	t.Run("Test_DeleteMemberAPI", func(t *testing.T) {

		memberID := "A011"
		phone := "123456789"

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/members?id=%s&phone=%s", memberID, phone), nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusOK, w.Code)

		// deleting data which doesnt exist
		req, _ = http.NewRequest("DELETE", "/api/members?id=A011&phone=123456789", nil)
		w = httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(http.StatusInternalServerError, w.Code)
	})
}
