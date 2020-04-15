package Const

//RPC协议号
const (
	Test          = 0
	Ping          = 1
	Pong          = 2
	Response_S    = 12
	Error_S       = 14
	Login_C       = 91 //登陆
	Login_S       = 92 //登陆
	GetRoleInfo_C = 93
	GetRoleInfo_S = 94
	Cmd_C         = 99
	//user
	Rename_C   = 103 //改名
	ReIcon_C   = 105 //改头像
	ReIconB_C  = 107 //改头像框
	UserView_C = 109 //查看角色
	//world
	View_C           = 203 //换视野
	CityMove_C       = 207 //迁城
	March_C          = 211 //行军
	Recall_C         = 213 //
	MarchCmd_C       = 215 //
	SpeedUp_C        = 217 //加速
	CityMoveRandom_C = 219 //随机迁城
	//item
	ItemUse_C = 301 //道具使用
	ItemDel_C = 303 //删除道具
	ShopBuy_C = 311 //商店购买
	//city
	Build_C     = 401 //建筑初建
	BuildUp_C   = 403 //建筑升级
	BuildDel_C  = 405 //建筑拆除
	BuildMove_C = 409 //建筑移动
	//hero
	HeroStepUp_C  = 421 //英雄升阶
	HeroAddExp_C  = 423 //英雄加经验
	HeroView_C    = 425 //英雄查看
	HeroCompose_C = 427 //英雄合成
	HeroRecruit_C = 429 //英雄招募
	HeroRecruit_S = 430 //
	//unit
	UnitTrain_C   = 451 //兵训练
	UnitDisMiss_C = 459 //兵遣散
	UnitUp_C      = 461 //兵升级
	UnitHeal_C    = 471 //兵治疗
	//tech
	Research_C = 481 //研究科技
	//job
	JobDone_C   = 491 //通用队列验收
	JobCancel_C = 493 //通用队列取消
	JobFast_C   = 495 //通用队列加速

	//mail
	MailGet_C    = 501 //邮件请求
	MailDel_C    = 503 //邮件删除
	MailRead_C   = 505 //邮件读
	MailTake_C   = 507 //邮件收附件
	MailFavor_C  = 509 //邮件收藏
	ReadReport_C = 511 //战报读取

	//task
	UserTaskReward_C  = 601 //领取主线支线任务奖励
	DailyTaskToDay_C  = 603 //获取今天每日任务
	DailyTaskReward_C = 605 //领取每日任务奖励
	DailyBoxReward_C  = 607 //领取活跃度宝箱

	AllianceCreate_C    = 701 //创
	AllianceApply_C     = 703 //申入
	AllianceOut_C       = 705 //退
	AllianceSelect_C    = 707 //搜索联盟
	AllianceSelect_S    = 708 //
	AllianceAutoApply_C = 709 //一键加入
	AllianceApproval_C  = 777 //审批

)
