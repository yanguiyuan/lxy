// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/yanguiyuan/lxy/biz/dal"
	"github.com/yanguiyuan/yuan/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 创建一个新的配置对象
	c := config.NewConfig()

	// 使用 GORM 打开 MySQL 数据库连接
	db, err := gorm.Open(mysql.Open(c.GetString("mysql.dsn")))
	if err != nil {
		// 如果连接出错，则打印错误信息并返回
		hlog.Error(err)
		return
	}

	// 将默认的数据访问层对象设置为 GORM 数据库连接
	dal.SetDefault(db)

	// 创建一个默认的服务器对象
	h := server.Default()

	// 注册处理程序到服务器对象
	register(h)

	// 启动服务器对象的事件循环
	h.Spin()
}