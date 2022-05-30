// 自动生成模板WeMrdk1
package autocode

// WeMrdk1 结构体
// 如果含有time.Time 请自行import time包
type WeMrdk1 struct {
      LogId  uint `json:"logId" form:"logId" gorm:"column:log_id;comment:用户日志id;size:10;primarykey"`
      Xh  string `json:"xh" form:"xh" gorm:"column:xh;comment:学号;size:40;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:姓名;size:255;"`
      Szdq  string `json:"szdq" form:"szdq" gorm:"column:szdq;comment:目前居住地;size:255;"`
      Sfhxdgr  string `json:"sfhxdgr" form:"sfhxdgr" gorm:"column:sfhxdgr;comment:是否与湖北人员接触;size:255;"`
      Stsfjk  string `json:"stsfjk" form:"stsfjk" gorm:"column:stsfjk;comment:健康状况;size:255;"`
      Jrtw  string `json:"jrtw" form:"jrtw" gorm:"column:jrtw;comment:今日体温;size:255;"`
}


// TableName WeMrdk1 表名
func (WeMrdk1) TableName() string {
  return "we_mrdk_1"
}

