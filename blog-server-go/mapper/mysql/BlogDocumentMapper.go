package mysql

import (
	"blog-server-go/models/config"
	"blog-server-go/models/entity"
)

func Insert(document entity.BlogDocument)  {
	config.DB.Save(document)
}

