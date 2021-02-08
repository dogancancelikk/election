package application

import "github.com/dogancancelikk/election/controllers/parties"

func mapUrls() {

	router.POST("/", parties.Create)
	router.GET("/", parties.GetAll)

}
