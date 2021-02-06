package docs

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/swaggo/swag"
)

type s struct{}

func (s *s) ReadDoc() string {
	res, err := http.Get("https://squaaat-lambda.s3.ap-northeast-2.amazonaws.com/serverless/jeonong-api/alpha/swagger.yml")
	b, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(b))
	res.Body.Close()
	if err != nil {
		panic(err.Error())
	}
	return string(b)
}

func init() {
	swag.Register(swag.Name, &s{})
}
