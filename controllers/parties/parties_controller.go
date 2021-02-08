package parties

import (
	"github.com/dogancancelikk/election/domain/parties"
	"github.com/dogancancelikk/election/services"
	"github.com/dogancancelikk/election/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var party parties.Party

	if err := c.ShouldBindJSON(&party); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateParty(party)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetAll(c *gin.Context) {
	parties, err := services.GetParties()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, parties)
}
