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

	existingUrl, _ := model.GetUrl(form.Url)
	if len(existingUrl) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong, please try again.",
		})
		return
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	id := r.Uint64()
	encodedId := shortener.Encode(id)

	_, err := model.InsertUrl(model.UrlModel{
		UrlID: id,
		Url: encodedId,
		Target: form.Url,
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

	fmt.Println(c.Param("url"))
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	model, err := model.GetUrl(params.Url)
	if len(model) == 0 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid URL",
		})
		return
	}

	c.Redirect(http.StatusPermanentRedirect, model[0].Target)
}
