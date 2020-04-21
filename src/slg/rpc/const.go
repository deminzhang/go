package Rpc

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
	TileFind_C       = 209 //资源/怪点查找
	March_C          = 211 //行军
	Recall_C         = 213 //
	MarchCmd_C       = 215 //
	SpeedUp_C        = 217 //加速
	CityMoveRandom_C = 219 //随机迁城

	//item
	ItemUse_C         = 301 //道具使用
	ItemDel_C         = 303 //删除道具
	ItemRead_C        = 305 //新道具已读
	ShopBuy_C         = 311 //商店购买
	VipShopBuy_C      = 313 //VIP商店购买
	MysticShopBuy_C   = 315 //神秘商店购买
	MysticShopView_C  = 317 //神秘商店查看
	MysticShopFlush_C = 319 //神秘商店刷新
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

	AllianceCreate_C        = 701 //创
	AllianceApply_C         = 703 //申入
	AllianceOut_C           = 705 //退
	AllianceSelect_C        = 707 //搜索联盟
	AllianceSelect_S        = 708 //
	AllianceAutoApply_C     = 709 //一键加入
	AllianceApproval_C      = 711 /*** 审批申请*/
	AllianceMemberView_C    = 713 //查看成员列表
	AllianceAppoint_C       = 715 //任命职位
	AllianceSetPermit_C     = 717 //更改招募状态
	AllianceMail_C          = 719 //发送全盟邮件
	AllianceInvite_C        = 721 //发送联盟邀请
	AllianceNeedHelp_C      = 723 //联盟请求帮助
	AllianceHelp_C          = 725 //联盟帮助
	AllianceTechRecommend_C = 727 //科技推荐
	AllianceTechUpgrade_C   = 729 //科技升级
	AllianceDonate_C        = 731 //联盟捐献
	AllianceKick_C          = 733 //联盟踢人
	AllianceRenotice_C      = 735 //联盟修改公告
	AllianceWelfare_C       = 737 //联盟领取福利
	AllianceShop_C          = 739 //联盟商店
	AllianceShutUp_C        = 741 //联盟禁言
	AllianceTechView_C      = 743 //联盟科技查看刷新

)
