package shorten

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opn-ooo/gin-boilerplate/app/model"
	"github.com/opn-ooo/gin-boilerplate/config"
	"github.com/opn-ooo/gin-boilerplate/helpers/local"
	"github.com/opn-ooo/gin-boilerplate/helpers/shortener"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type PostShortenForm struct {
	Url string `form:"url" binding:"required"`
}

type GetShortenRedirectParam struct {
	Url string `uri:"url"`
}

func PostShorten(c *gin.Context) {
	form := PostShortenForm{}

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if _, err := url.ParseRequestURI(form.Url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid URL",
		})
		return
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	id := r.Uint64()
	encodedId := shortener.Encode(id)

	_, err := model.InsertUrl(model.UrlModel{
		UrlID: id,
		Url: form.Url,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong, please try again.",
		})
		return
	}

	goDotENV := config.GetGoDotENV()
	result := url.URL{
		Scheme: "http",
		Host: local.GetLocalIP() + goDotENV.Port,
		Path: fmt.Sprintf("/s/%s", encodedId),
	}

	c.JSON(http.StatusOK, gin.H{
		"url": result.String(),
	})
}

func GetShortenRedirect(c *gin.Context) {
	params := GetShortenRedirectParam{}

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	decodedId, err := shortener.Decode(params.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong, please try again.",
		})
		return
	}

	model, err := model.GetUrl(decodedId)
	if len(model) == 0 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid URL",
		})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, model[0].Url)
}
