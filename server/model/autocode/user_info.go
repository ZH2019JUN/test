// 自动生成模板UserInfo
package autocode

// UserInfo 结构体
// 如果含有time.Time 请自行import time包
type UserInfo struct {
      Xh  string `json:"xh" form:"xh" gorm:"column:xh;comment:学号;size:20;primarykey"`
      Xm  string `json:"xm" form:"xm" gorm:"column:xm;comment:姓名;size:255;"`
      Xb  string `json:"xb" form:"xb" gorm:"column:xb;comment:性别;size:255;"`
      Mz  string `json:"mz" form:"mz" gorm:"column:mz;comment:民族;size:255;"`
      Syd  string `json:"syd" form:"syd" gorm:"column:syd;comment:生源地;size:255;"`
      Csrq  string `json:"csrq" form:"csrq" gorm:"column:csrq;comment:出生日期;size:255;"`
      Sfzh  string `json:"sfzh" form:"sfzh" gorm:"column:sfzh;comment:身份证号;size:255;"`
      Nj  string `json:"nj" form:"nj" gorm:"column:nj;comment:年级;size:255;"`
      Bj  string `json:"bj" form:"bj" gorm:"column:bj;comment:班级;size:255;"`
      Zy  string `json:"zy" form:"zy" gorm:"column:zy;comment:专业;size:255;"`
      Xy  string `json:"xy" form:"xy" gorm:"column:xy;comment:学院;size:255;"`
      Tysbm  string `json:"tysbm" form:"tysbm" gorm:"column:tysbm;comment:统一识别码;size:255;"`
      Zyh  string `json:"zyh" form:"zyh" gorm:"column:zyh;comment:专业号;size:255;"`
      Yxh  string `json:"yxh" form:"yxh" gorm:"column:yxh;comment:院系号;size:255;"`
      Xjzt  string `json:"xjzt" form:"xjzt" gorm:"column:xjzt;comment:学籍状态;size:1;"`
      Pjjdcj  string `json:"pjjdcj" form:"pjjdcj" gorm:"column:pjjdcj;comment:平均绩点成绩;size:255;"`
}


// TableName UserInfo 表名
func (UserInfo) TableName() string {
  return "user_info"
}

