package controller

import (
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"net/http"
)

func GetHome(request *http.Request, response http.ResponseWriter) (int, string) {
	var db = new(util.Database)
	defer db.Close()
	if post, success := service.GetPostByUrl(db, "/"); success {
		layout, template, params := GetViewParams(post)
		params["posts"] = service.GetPostListingHome(db)
		return util.RespondTemplate(http.StatusOK, layout, template, params)
	}

	return PageNotFound(request, response)
}