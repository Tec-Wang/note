package cmd

import (
	"context"
	"wangzheng/brain/config"
	"wangzheng/brain/internal/common"
)

var ctx = context.Background()
var Mysql = common.NewMysql()
var Mongo = common.NewMongoClient(ctx, config.MongoConfig().ConnectionString())
