/*
*
@author: yudan.chen
@date: 2023/10/15
*
*/
package global

import (
	"blueegg/blog-service/pkg/logger"
	"blueegg/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
