package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

func getAcceptLanguage(acceptLanguate string) string {
	var serverLangs = []language.Tag{
		language.SimplifiedChinese, // zh-Hans fallback
		language.AmericanEnglish,   // en-US
		language.Korean,            // de
	}

	// 也可以不定义 serverLangs 用下面一行选择支持所有语种。
	// var matcher = language.NewMatcher(message.DefaultCatalog.Languages())
	var matcher = language.NewMatcher(serverLangs)
	t, _, _ := language.ParseAcceptLanguage(acceptLanguate)
	tag, index, confidence := matcher.Match(t...)

	fmt.Printf("best match: %s (%s) index=%d confidence=%v\n",
		display.English.Tags().Name(tag),
		display.Self.Name(tag),
		index, confidence)

	str := fmt.Sprintf("tag is %s", tag)
	fmt.Println(str)
	fmt.Printf("best match: %s\n", display.Self.Name(tag))
	return str
}

func I18nMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.Query("locale")
		if locale != "" {
			c.Request.Header.Set("Accept-Language", locale)
		}
		lang := getAcceptLanguage(c.GetHeader("Accept-Language"))

		// NOTE: On June 2012, the deprecation of recommendation to use the "X-" prefix has become official as RFC 6648.
		// https://stackoverflow.com/questions/3561381/custom-http-headers-naming-conventions
		c.Request.Header.Set("I18n-Language", lang)
		c.Next()
	}
}
