package routers

import (
	"blog-server-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/blog/document/?:documentId", &controllers.BlogDocumentController{})
    beego.Router("/blog/tag", &controllers.BlogTagController{})
    beego.Router("/blog/audio", &controllers.BlogAudioController{})
}
