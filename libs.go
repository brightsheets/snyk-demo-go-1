package snykdemogo1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v2"
	"strings"

	privatelib "github.com/brightsheets/private-lib-01"
)

// Libs exercises every dependency, so each one stays in go.mod. Several are pinned to old
// versions with known vulnerabilities on purpose (Snyk demo), and private-lib-01 is a private
// module in the brightsheets org: a cross-org dependency the test-snyk-demo Snyk app can't read.
func Libs() []string {
	out := []string{privatelib.Greet(ID)}

	r := gin.New()
	out = append(out, "gin routes: "+string(rune('0'+len(r.Routes()))))

	log := logrus.New()
	log.SetLevel(logrus.WarnLevel)
	out = append(out, "logrus level: "+log.GetLevel().String())

	y, _ := yaml.Marshal(map[string]string{"module": ID})
	out = append(out, "yaml: "+strings.TrimSpace(string(y)))

	t := jwt.New(jwt.SigningMethodHS256)
	out = append(out, "jwt alg: "+t.Method.Alg())

	doc, err := html.Parse(strings.NewReader("<p>hi</p>"))
	if err == nil && doc != nil {
		out = append(out, "html parsed")
	}
	return out
}
