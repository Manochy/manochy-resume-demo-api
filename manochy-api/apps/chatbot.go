package apps

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type History struct {
	Company string `json:"company_en"`
	Detail  string `json:"company_detail_en"`
}

var data []History

func loadData() {
	if len(data) > 0 {
		return
	}

	file, _ := os.ReadFile("manochy-history.json")
	json.Unmarshal(file, &data)
}

func Chatbot(c *gin.Context) {
	loadData()

	var req struct {
		Question string `json:"question"`
	}

	c.BindJSON(&req)

	var result []History

	for _, d := range data {
		if strings.Contains(strings.ToLower(d.Company), strings.ToLower(req.Question)) ||
			strings.Contains(strings.ToLower(d.Detail), strings.ToLower(req.Question)) {
			result = append(result, d)
		}
	}

	c.JSON(http.StatusOK, result)
}
