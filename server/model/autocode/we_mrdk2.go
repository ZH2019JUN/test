// 自动生成模板WeMrdk2
package autocode

// WeMrdk2 结构体
// 如果含有time.Time 请自行import time包
type WeMrdk2 struct {
      LogId  uint `json:"logId" form:"logId" gorm:"column:log_id;comment:用户日志id;size:10;primarykey"`
      Xh  string `json:"xh" form:"xh" gorm:"column:xh;comment:学号;size:40;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:姓名;size:255;"`
      Xb  string `json:"xb" form:"xb" gorm:"column:xb;comment:性别;size:255;"`
      Lxdh  string `json:"lxdh" form:"lxdh" gorm:"column:lxdh;comment:联系电话;size:255;"`
      Szdq  string `json:"szdq" form:"szdq" gorm:"column:szdq;comment:目前居住地;size:255;"`
      Sfhxdgr  string `json:"sfhxdgr" form:"sfhxdgr" gorm:"column:sfhxdgr;comment:是否与湖北人员接触;size:255;"`
      Stsfjk  string `json:"stsfjk" form:"stsfjk" gorm:"column:stsfjk;comment:健康状况;size:255;"`
      Jrtw  string `json:"jrtw" form:"jrtw" gorm:"column:jrtw;comment:今日体温;size:255;"`
      Sffr  string `json:"sffr" form:"sffr" gorm:"column:sffr;comment:是否发热;size:255;"`
      Sfks  string `json:"sfks" form:"sfks" gorm:"column:sfks;comment:是否咳嗽;size:255;"`
      Sfxm  string `json:"sfxm" form:"sfxm" gorm:"column:sfxm;comment:是否胸闷;size:255;"`
      Qtycqk  string `json:"qtycqk" form:"qtycqk" gorm:"column:qtycqk;comment:其他异常情况;size:255;"`
      Ywjchblj  string `json:"ywjchblj" form:"ywjchblj" gorm:"column:ywjchblj;comment:14天内是否有中高风险地区旅居史;size:255;"`
      Ywjcqzbl  string `json:"ywjcqzbl" form:"ywjcqzbl" gorm:"column:ywjcqzbl;comment:目前居住地新冠肺炎疫情风险等级;size:255;"`
      Xjzdywqzbl  string `json:"xjzdywqzbl" form:"xjzdywqzbl" gorm:"column:xjzdywqzbl;comment:14天内是否接触中高风险地区旅居史人员;size:255;"`
      Twsfzc  string `json:"twsfzc" form:"twsfzc" gorm:"column:twsfzc;comment:今日体温是否正常;size:255;"`
      Ywytdzz  string `json:"ywytdzz" form:"ywytdzz" gorm:"column:ywytdzz;comment:今日是否有与新冠病毒感染有关的症状;size:255;"`
      Jbsks  string `json:"jbsks" form:"jbsks" gorm:"column:jbsks;comment:是否咳嗽;size:255;"`
      Jbsfl  string `json:"jbsfl" form:"jbsfl" gorm:"column:jbsfl;comment:是否乏力;size:255;"`
      Jbsbs  string `json:"jbsbs" form:"jbsbs" gorm:"column:jbsbs;comment:是否鼻塞;size:255;"`
      Jbslt  string `json:"jbslt" form:"jbslt" gorm:"column:jbslt;comment:是否流涕;size:255;"`
      Jbsyt  string `json:"jbsyt" form:"jbsyt" gorm:"column:jbsyt;comment:是否咽痛;size:255;"`
      Jbsfx  string `json:"jbsfx" form:"jbsfx" gorm:"column:jbsfx;comment:是否腹泻;size:255;"`
      Hjsfly  string `json:"hjsfly" form:"hjsfly" gorm:"column:hjsfly;comment:寒假是否离渝;size:255;"`
      Sfyfy  string `json:"sfyfy" form:"sfyfy" gorm:"column:sfyfy;comment:是否已返渝;size:255;"`
      Fylx  string `json:"fylx" form:"fylx" gorm:"column:fylx;comment:反渝路线（途径哪些地方）;size:255;"`
      Fyjtgj  string `json:"fyjtgj" form:"fyjtgj" gorm:"column:fyjtgj;comment:返回乘坐的交通工具（班次车牌）;size:255;"`
      Fyddsj  string `json:"fyddsj" form:"fyddsj" gorm:"column:fyddsj;comment:返渝到达时间;size:255;"`
      Sfbgsq  string `json:"sfbgsq" form:"sfbgsq" gorm:"column:sfbgsq;comment:是否报告社区;size:255;"`
      Sfjjgl  string `json:"sfjjgl" form:"sfjjgl" gorm:"column:sfjjgl;comment:是否居家隔离;size:255;"`
      Jjglqssj  string `json:"jjglqssj" form:"jjglqssj" gorm:"column:jjglqssj;comment:居家隔离起始时间;size:255;"`
      Wjjglmqqx  string `json:"wjjglmqqx" form:"wjjglmqqx" gorm:"column:wjjglmqqx;comment:未居家隔离的目前去向;size:255;"`
      Xxdz  string `json:"xxdz" form:"xxdz" gorm:"column:xxdz;comment:详细地址;size:255;"`
      Brsfqz  string `json:"brsfqz" form:"brsfqz" gorm:"column:brsfqz;comment:本人及家庭成员是否为确诊病例;size:255;"`
      Brsfys  string `json:"brsfys" form:"brsfys" gorm:"column:brsfys;comment:本人及家庭成员是否为疑似病例;size:255;"`
      Jbs  string `json:"jbs" form:"jbs" gorm:"column:jbs;comment:有无疾病史;size:255;"`
      Jkmresult  string `json:"jkmresult" form:"jkmresult" gorm:"column:jkmresult;comment:健康码识别信息;size:255;"`
}


// TableName WeMrdk2 表名
func (WeMrdk2) TableName() string {
  return "we_mrdk_2"
}

