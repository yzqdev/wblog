package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strings"
	"wblog-server/models"
	"wblog-server/system"
)

const (
	SESSION_KEY          = "UserID"       // session key
	CONTEXT_USER_KEY     = "User"         // context user key
	SESSION_GITHUB_STATE = "GITHUB_STATE" // github state session key
	SESSION_CAPTCHA      = "GIN_CAPTCHA"  // captcha session key
)

func Handle404(c *gin.Context) {
	HandleMessage(c, "Sorry,I lost myself!")
}

func HandleMessage(c *gin.Context, message string) {
	c.HTML(http.StatusNotFound, "errors/error.html", gin.H{
		"message": message,
	})
}

func SendEmail(to, subject, body string) error {
	c := system.GetConfiguration()
	return SendToMail(c.SmtpUsername, c.SmtpPassword, c.SmtpHost, to, subject, body, "html")
}

func NotifyEmail(subject, body string) error {
	notifyEmailsStr := system.GetConfiguration().NotifyEmails
	if notifyEmailsStr != "" {
		notifyEmails := strings.Split(notifyEmailsStr, ";")
		emails := make([]string, 0)
		for _, email := range notifyEmails {
			if email != "" {
				emails = append(emails, email)
			}
		}
		if len(emails) > 0 {
			return SendEmail(strings.Join(emails, ";"), subject, body)
		}
	}
	return nil
}

func CreateXMLSitemap() {
	configuration := system.GetConfiguration()
	folder := path.Join(configuration.Public, "helpers")
	os.MkdirAll(folder, os.ModePerm)
	domain := configuration.Domain
	now := GetCurrentTime()
	items := make([]Item, 0)

	items = append(items, Item{
		Loc:        domain,
		LastMod:    now,
		Changefreq: "daily",
		Priority:   1,
	})

	posts, err := models.ListPublishedPost("", 0, 0)
	if err == nil {
		for _, post := range posts {
			items = append(items, Item{
				Loc:        fmt.Sprintf("%s/post/%d", domain, post.ID),
				LastMod:    post.UpdatedAt,
				Changefreq: "weekly",
				Priority:   0.9,
			})
		}
	}

	pages, err := models.ListPublishedPage()
	if err == nil {
		for _, page := range pages {
			items = append(items, Item{
				Loc:        fmt.Sprintf("%s/page/%d", domain, page.ID),
				LastMod:    page.UpdatedAt,
				Changefreq: "monthly",
				Priority:   0.8,
			})
		}
	}

	if err := SiteMap(path.Join(folder, "helpers1.xml.gz"), items); err != nil {
		return
	}
	if err := SiteMapIndex(folder, "helpers_index.xml", domain+"/static/helpers/"); err != nil {
		return
	}
}

func WriteJson(ctx *gin.Context, h gin.H) {
	if _, ok := h["succeed"]; !ok {
		h["succeed"] = false
	}
	ctx.JSON(http.StatusOK, h)
}
