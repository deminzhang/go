package defs

import "time"

/* gate 处理小于[,OpcodeGateHandleMax)协议
*  agent [OpcodeGateHandleMax+1, OpcodeAgentHandleMax)
*        (OpcodeAgentHandleWithBattleBegin,OpcodeAgentHandleMax)  player in battle
*  cell [OpcodeAgentHandleMax+1, OpcodeS2SBegin)
*  S2S [OpcodeS2SBegin, )
 */

const (
	App      = "pff"
	AppShort = "pf"

	EnvProd    = "prod"
	EnvTest    = "test"
	EnvDev     = "dev"
	EnvPre     = "pre"
	EnvBetaDev = "betadev"

	DefaultLoc = "Local"
	UTCLoc     = "UTC"

	ImFromSys = "sys"

	ImMsgTypeFriendAdd          = 501
	ImMsgTypeFriendIntimacy     = 502
	ImMsgTypeLogin              = 503
	ImMsgTypeLogout             = 504
	ImMsgTypeEnterMap           = 505
	ImMsgTypeFriendDelete       = 506
	ImMsgTypeFriendApplyInfo    = 507
	ImMsgTypeAnnounceInfo       = 508
	ImMsgTypeFriendUpdateAvatar = 509

	UGCRoomShardId = 10000

	DaySeconds = 86400

	ExpireAuth                = 24 * time.Hour
	ExpireUser2Player         = 24 * time.Hour
	ExpireAgentShard          = 130 * time.Second
	ExpireImShard             = 130 * time.Second
	ExpireGenAgentShard       = (1800 + 10) * time.Second
	IntervalUpdateGatePlayer  = 60 * 1000   // millisecond
	IntervalImHeartbeat       = 60 * 1000   // millisecond
	IntervalGenAgentHeartbeat = 1800 * 1000 // millisecond

	GptGrpcTimeout = 5 * 60 * time.Second

	CtxKeyUid       = "uid"
	CtxKeyCryptoKey = "crypto_key"

	MailPrefabOtpActivate      = "otp_activate"
	MailPrefabOtpResetPassword = "otp_resetpass"
	MailPrefabPreRegister      = "open_game"
	FrameQueueSize             = 512
	MatchMinTime               = 20  //自定议匹配时间下限()
	MatchTime                  = 120 //缺省匹配时间seconds
	MatchMaxTime               = 240 //自定议匹配时间上限

	FrameTypeKickOut    = 1
	FrameTypeClientReq  = 2
	FrameTypeClientResp = 3
	FrameTypeAgentReq   = 4
	FrameTypeAgentResp  = 5
	FrameTypeCellReq    = 6
	FrameTypeCellResp   = 7

	OpcodePing = 1
	OpcodePong = 2

	//gate
	OpcodeUserLogin     = 100
	OpcodeUserLoginResp = 101

	OpcodeImUserLogin         = 201
	OpcodeImUserLoginResp     = 202
	OpcodeImPing              = 203
	OpcodeImPong              = 204
	OpcodeImMsgResp           = 210
	OpcodeImPubPerson         = 221
	OpcodeImPubPersonResp     = 222
	OpcodeImPubChannel        = 223
	OpcodeImPubChannelResp    = 224
	OpcodeImBroadcast         = 225
	OpcodeImBroadcastResp     = 226
	OpcodeImSub               = 227
	OpcodeImSubResp           = 228
	OpcodeImUnsub             = 229
	OpcodeImUnsubResp         = 230
	OpcodeImFetchPrivChat     = 231
	OpcodeImFetchPrivChatResp = 232

	OpcodeGateHandleMax = 1000

	//agent
	OpcodePlayerLogin          = 1001
	OpcodePlayerLoginResp      = 1002
	OpcodePlayerRebuilding     = 1007 //重连
	OpcodePlayerRebuildingResp = 1008
	OpcodePlayerJoinBattle     = 1010 //新入
	OpcodeSelectHome           = 1011 //选家园地图
	OpcodeSelectHomeResp       = 1012
	OpcodeKickOutResp          = 1013

	OpcodeCellEnterResp  = 1301
	OpcodeCellLeaveResp  = 1302 //自己离开
	OpcodeCellChangeResp = 1303 //切换地图

	//agent and player in battling
	OpcodeAgentHandleInBattlingBegin = 3000
	OpcodeTryLogout                  = 3010 //主动退出游戏
	OpcodeTryLogoutResp              = 3011 //回完即断
	OpcodeMatch                      = 3101
	OpcodeMatchResp                  = 3102
	OpcodeMatchCancel                = 3103
	OpcodeMatchCancelResp            = 3104
	OpcodeMatchTimeoutResp           = 3105
	OpcodeMatchUpdateResp            = 3106 //匹配等待房间更新

	OpcodeQueryPlayersInfo     = 3200 //查询其它玩家基础信息
	OpcodeQueryPlayersInfoResp = 3201
	OpcodeHomeLevelUpResp      = 3202 // 家园/家/主角升级

	OpcodeItemData                 = 4001
	OpcodeItemDataResp             = 4002
	OpcodeItemDel                  = 4003
	OpcodeItemDelResp              = 4004
	OpcodeItemMove                 = 4005
	OpcodeItemMoveResp             = 4006
	OpcodeItemBagChangesResp       = 4007
	OpcodeItemWarehouseChangesResp = 4008
	OpcodeItemMoveStack            = 4009
	OpcodeItemMoveStackResp        = 4010

	//Game1修改时区
	OpCodeTimeZoneModify     = 4018 // 修改时区
	OpCodeTimeZoneModifyResp = 4019 // 修改时区
	//商店
	OpcodeShopBuy           = 4020 //购买
	OpcodeShopBuyResp       = 4021 //购买
	OpcodeShopData          = 4022 //获取消息
	OpcodeShopDataResp      = 4023 //获取消息
	OpcodeShopBuyRecord     = 4024 //获取玩家购买
	OpcodeShopBuyRecordResp = 4025 //获取玩家购买
	// 货币
	OpcodeCoinData       = 4030
	OpcodeCoinDataResp   = 4031 //获取消息
	OpcodeCoinUpdateResp = 4032

	// 任务
	OpcodeQuestData           = 4040 // 请求任务数据
	OpcodeQuestDataResp       = 4041 // 应答任务数据
	OpcodeQuestAccept         = 4042
	OpcodeQuestAcceptResp     = 4043
	OpcodeQuestCommit         = 4044
	OpcodeQuestCommitResp     = 4045
	OpcodeQuestReset          = 4046
	OpcodeQuestResetResp      = 4047
	OpcodeQuestChangeResp     = 4048 // 任务数据变化应答
	OpcodeQuestEvent          = 4049
	OpcodeQuestEventResp      = 4050
	OpcodeQuestTrace          = 4051
	OpcodeQuestTraceResp      = 4052
	OpcodeQuestCleanTrace     = 4053
	OpcodeQuestCleanTraceResp = 4054
	OpcodeQuestCommitItem     = 4055
	OpcodeQuestCommitItemResp = 4056

	//Game1时装
	OpcodeFashionPreset               = 4060 // 时装预设
	OpcodeFashionPresetResp           = 4061 // 预设返回
	OpcodeFashionPresetSve            = 4062 // 保存预设
	OpcodeFashionPresetSveResp        = 4063 // 保存预设返回
	OpcodeFashionPresetChangeName     = 4064 // 预设改名
	OpcodeFashionPresetChangeNameResp = 4065 // 预设改名返回
	OpcodeFashionPresetWear           = 4066 // 穿戴预设
	OpcodeFashionPresetWearResp       = 4067 // 穿戴返回
	OpcodeFashion                     = 4068 // 时装请求
	OpcodeFashionResp                 = 4069 // 时装应答
	OpcodeFashionWear                 = 4070 // 穿时装
	OpcodeFashionWearResp             = 4071 // 穿时装应答
	OpcodeFashionTakeOff              = 4072 // 脱时装
	OpcodeFashionTakeOffResp          = 4073 // 脱时装应答
	OpcodeFashionTakeOffAll           = 4074 // 脱下当前所有时装
	OpcodeFashionTakeOffAllResp       = 4075 // 应答
	OpcodeFashionBroadcast            = 4076 // 广播玩家当前的时装变化

	//Game1捏脸
	OpcodeCharacterCreatorUnlockInfo     = 4090 // 捏脸请求
	OpcodeCharacterCreatorUnlockInfoResp = 4091 // 捏脸应答
	OpcodeCharacterCreatorSave           = 4092 // 保存捏脸数据
	OpcodeCharacterCreatorSaveResp       = 4093 // 保存捏脸数据
	OpcodeCharacterCreatorBroadcastResp  = 4094 // 捏脸广播给场景中的玩家
	OpcodeCharacterCreatorUnlockResp     = 4095 // 解锁

	OpcodeCraftData      = 4100 // 请求物品制造数据
	OpcodeCraftDataResp  = 4101 // 应答物品制造数据
	OpcodeCraftLearnResp = 4102 // 物品制造学习配方之后返回
	OpcodeCraft          = 4103 // 物品制造
	OpcodeCraftResp      = 4104 // 应答物品制造

	// 物品出售
	OpcodeItemSellData     = 4110 // 请求高价物品信息
	OpcodeItemSellDataResp = 4111 // 返回高价物品信息
	OpcodeItemSell         = 4112 // 出售
	OpcodeItemSellResp     = 4113 // 出售返回

	// 功能开启
	OpcodeModuleSwitchData       = 4120 // 请求功能控制数据
	OpcodeModuleSwitchDataResp   = 4121 // 应答功能控制数据
	OpcodeModuleSwitchChangeResp = 4122 // 新功能开启应答
	// 邮件
	OpcodeMailQuery         = 4130 // 请求邮件查询
	OpcodeMailQueryResp     = 4131 // 应答邮件查询
	OpcodeMailGetAttach     = 4132 // 请求邮件领取
	OpcodeMailGetAttachResp = 4133 // 应答邮件领取
	OpcodeMailRead          = 4134 // 请求邮件阅读
	OpcodeMailReadResp      = 4135 // 应答邮件阅读
	OpcodeMailDelete        = 4136 // 请求邮件删除
	OpcodeMailDeleteResp    = 4137 // 应答邮件删除
	OpcodeMailReceiveResp   = 4138 // 邮件接收通知

	//宠物
	OpcodePetData                = 4140 // 请求宠物数据
	OpcodePetDataResp            = 4141 // 应答宠物数据
	OpcodePetFeed                = 4142 // 道具喂食请求
	OpcodePetFeedResp            = 4143 // 道具喂食应答
	OpcodePetRename              = 4144 // 宠物改名请求
	OpcodePetRenameResp          = 4145 // 宠物改名应答
	OpcodePetSetFollow           = 4146 // 设置宠物跟随请求
	OpcodePetSetFollowResp       = 4147 // 设置宠物跟随应答
	OpcodePetChangeResp          = 4148 // 宠物数据变化
	OpcodePetFollowChangeAoiResp = 4149 // 视野内跟随跟随宠物数据变化
	OpcodePetSetUsedOrder        = 4150 // 设置已用指令请求
	OpcodePetSetUsedOrderResp    = 4151 // 设置已用指令应答
	OpcodePetQuestActiveResp     = 4152 // 激活宠物任务
	OpcodePetFollowRemoveAoiResp = 4153 // 视野内跟随跟随宠物删除
	OpcodePetExpChangeResp       = 4154 // 宠物经验变化通知
	OpcodePetRead                = 4155 // 设置宠物已读
	OpcodePetReadRes             = 4156 // 设置宠物已读应答

	//烹饪
	OpcodeCookingData     = 4160 // 请求烹饪数据
	OpcodeCookingDataResp = 4161 // 返回烹饪数据
	OpcodeCookingCook     = 4162 // 烹饪
	OpcodeCookingCookResp = 4163 // 返回烹饪结果

	//图纸diy
	OpCodeTextureDiyData               = 4170 // 请求图纸Diy数据
	OpCodeTextureDiyDataResp           = 4171 // 返回图纸Diy数据
	OpCodeTextureDiySave               = 4172 // 更新图纸Diy数据
	OpCodeTextureDiySaveResp           = 4173 // 更新返回
	OpCodeTextureDiyMove               = 4174 // 移动
	OpCodeTextureDiyMoveResp           = 4175 // 移动返回
	OpCodeTextureDiyChangeName         = 4176 // 改名
	OpCodeTextureDiyChangeNameResp     = 4177 // 改名返回
	OpCodeTextureDiyChangeResp         = 4178 // 数据变更
	OpCodeTextureDiyFashionWear        = 4179 // 穿戴diy时装
	OpCodeTextureDiyFashionWearResp    = 4180 // 穿戴diy时装返回
	OpCodeTextureDiyFashionTakeoff     = 4181 // 脱下diy时装
	OpCodeTextureDiyFashionTakeoffResp = 4182 // 脱下diy时装返回
	OpCodeTextureDiyFashionBroadcast   = 4183 // 穿戴变更时广播给玩家

	//捐赠
	OpCodeContributeReq      = 4190 // 请求捐赠
	OpCodeContributeResp     = 4191 // 捐赠返回
	OpCodeContributeDataReq  = 4192 // 请求捐赠数据
	OpCodeContributeDataResp = 4193 // 请求捐赠数据返回
	//快捷栏
	OpCodeShortcutReq      = 4195 // 请求快捷栏
	OpCodeShortcutResp     = 4196 // 快捷栏返回
	OpCodeShortcutSaveReq  = 4197 // 保存
	OpCodeShortcutSaveResp = 4198 // 保存返回

	OpcodeBlueprintList             = 4300 //蓝图列表
	OpcodeBlueprintListResp         = 4301 //
	OpcodeBlueprintSave             = 4302 //蓝图保存
	OpcodeBlueprintSaveResp         = 4303 //
	OpcodeBlueprintDel              = 4304 //蓝图删除
	OpcodeBlueprintDelResp          = 4305 //
	OpcodeBlueprintPub              = 4306 //蓝图发布
	OpcodeBlueprintPubResp          = 4307 //
	OpcodeBlueprintUnPub            = 4308 //蓝图取消发布
	OpcodeBlueprintUnPubResp        = 4309 //
	OpcodeBlueprintSupportList      = 4310 //点赞过的蓝图id
	OpcodeBlueprintSupportListResp  = 4311 //
	OpcodeBlueprintSupport          = 4312 //蓝图点赞/取消
	OpcodeBlueprintSupportResp      = 4313 //
	OpcodeBlueprintFavorite         = 4314 //蓝图收藏/取消
	OpcodeBlueprintFavoriteResp     = 4315 //
	OpcodeBlueprintSearch           = 4316 //蓝图搜索
	OpcodeBlueprintSearchResp       = 4317 //
	OpcodeBlueprintSetPasswd        = 4318 //密码设置
	OpcodeBlueprintSetPasswdResp    = 4319 //
	OpcodeBlueprintPubAuthority     = 4320 //发布成官方蓝图(权限人员专用)
	OpcodeBlueprintPubAuthorityResp = 4321
	OpcodeBlueprintPubGet           = 4322 //按id单查
	OpcodeBlueprintPubGetResp       = 4323 //

	OpcodeIllustratedBookSumUpdate        = 4350 //图鉴总览下发
	OpcodeIllustratedBookList             = 4351 //图鉴分组请求
	OpcodeIllustratedBookListResp         = 4352 //图鉴更新
	OpcodeIllustratedBookRead             = 4353 //图鉴已读
	OpcodeIllustratedBookReadResp         = 4354
	OpcodeIllustratedBookSumGetReward     = 4355 //图鉴总览奖励
	OpcodeIllustratedBookSumGetRewardResp = 4356

	OpcodeAgentHandleMax = 5000 //==============================================

	//Cell 通用
	OpcodeCellQuit          = 5001 //主动离开当前场景
	OpcodeCellQuitResp      = 5002 //他人离开
	OpcodeCustomStatusInput = 5003

	OpcodeNtp                = 5005
	OpcodeNtpResp            = 5006
	OpcodeFrameInput         = 5007
	OpcodeFrameUpdateResp    = 5008
	OpcodePlayerEnterMapResp = 5009 //主角进场
	OpcodeObjectSpawnResp    = 5010 //在场单位增量

	//Cell.HomeSteadId

	OpcodeNpcTalkStart           = 5011
	OpcodeNpcTalkStartResp       = 5012
	OpcodeNpcTalkStop            = 5013
	OpcodeNpcTalkStopResp        = 5014
	OpcodeStateChangeResp        = 5015
	OpcodePlayerStrengthResp     = 5017
	OpcodePlayerJoinFinishedResp = 5018
	OpcodeAoiLeaveViewResp       = 5019 //谁退出我的视野
	OpcodeHouseRegionListResp    = 5020
	OpcodePlayerEmote            = 5021 //表情动作上报
	OpcodePlayerEmoteResp        = 5022 //表情动作广播
	OpcodeSetHouseLight          = 5023 //修改房间的灯光
	OpcodeSetHouseLightResp      = 5024 //房间的灯光变动广播
	OpcodeSceneOpenUpdate        = 5025 //解锁信息更新
	OpcodeEmotionUpdateResp      = 5026 //解锁的表情动作

	OpcodeFrameResetResp         = 5029 //重置玩家序列帧
	OpcodeGoToMap                = 5030 //请求跳图
	OpcodeGoToMapResp            = 5031
	OpcodeMapInvitePlayer        = 5032 //邀请到我当前所在地图实例
	OpcodeMapInvitePlayerResp    = 5033
	OpcodeJoinMapInviteResp      = 5034 //好友邀请进所在图
	OpcodeJoinMap                = 5035 //接受邀请跳图
	OpcodeJoinMapResp            = 5036
	OpcodeCreateHouse            = 5037 //请求预创建子房间地图
	OpcodeCreateHouseResp        = 5038
	OpcodePlayerJumpPosResp      = 5039 //玩家同图位置跳转
	OpcodeGM                     = 5040
	OpcodeGMResp                 = 5041
	OpcodeCellPlayerBriefMap     = 5042
	OpcodeCellPlayerBriefMapResp = 5043
	OpcodeRemoveObjectsResp      = 5044 //通用单位清场
	OpcodePlayerJumpPos          = 5045 //玩家同图位置跳转

	//Game1天氣相关
	OpcodeCellWeatherResp = 5049

	OpcodeUnitMoveResp       = 5052 //非帧控(NPC)单位移动
	OpcodeUnitStopResp       = 5054 //非帧控单位停止
	OpcodeUnitPlayActionResp = 5055 //非帧控(NPC)播放动作

	// 除Use,转agent,缓存换图不清
	OpcodeBlueprintUse     = 5070 //蓝图使用
	OpcodeBlueprintUseResp = 5071 //

	OpcodeVoxelDigPile         = 5101 //岛建操作
	OpcodeModifyVoxelSpanResp  = 5102
	OpcodeVoxelSpanAllDataResp = 5103
	OpcodeDropItemListResp     = 5104 //掉落物列表
	OpcodeDropItemAddResp      = 5105 //掉落物新加
	OpcodeDropItemPickUp       = 5106 //拾取掉落物
	OpcodeDropItemRemoveResp   = 5107 //掉落物消失
	OpcodeEstuaryListResp      = 5108 //入海口

	//Game1道具使用
	OpcodeItemBoxRewardsResp = 5109 // 宝箱开启后返回
	OpcodeItemUse            = 5110
	OpcodeItemUseResp        = 5111
	OpcodeItemUseBatch       = 5112
	OpcodeItemUseBatchResp   = 5113

	OpcodeCommonClientBroadcast     = 5114
	OpcodeCommonClientBroadcastResp = 5115

	//Game1工具系统
	OpcodeToolSystem            = 5120 // 请求数据
	OpcodeToolSystemResp        = 5121
	OpcodeToolSystemEquip       = 5122
	OpcodeToolSystemEquipResp   = 5123
	OpcodeToolSystemUnEquip     = 5124
	OpcodeToolSystemUnEquipResp = 5125

	//Game1种植
	OpcodePlantListResp    = 5130 //植物列表
	OpcodePlant            = 5131 //种植请求
	OpcodePlantResp        = 5132 //种植应答
	OpcodePlantRemoveResp  = 5133 //植物消失
	OpcodePlantDig         = 5134 //挖掘植物
	OpcodePlantDigResp     = 5135 //挖掘应答
	OpCodePlantGather      = 5136 //摘取请求
	OpCodePlantGatherResp  = 5137 //摘取应答
	OpCodePlantTrample     = 5139 //踩踏请求
	OpCodePlantTrampleResp = 5140 //踩踏应答
	OpCodePlantWater       = 5141 //浇水请求
	OpCodePlantWaterResp   = 5142 //浇水应答
	OpcodePlantAddResp     = 5144 //新增植物列表
	OpcodePlantUpdateResp  = 5145 //植物更新
	OpcodePlantCut         = 5156 //砍树
	OpcodePlantCutResp     = 5157 //砍树应答
	OpCodePlantShake       = 5161 //摇晃请求
	OpCodePlantShakeResp   = 5162 //摇晃应答

	OpcodeAvatarUpdateResp      = 5163
	OpcodeDisplayNameUpdateResp = 5164

	OpcodePlayersExtra     = 5165 // 获取多个玩家扩展信息请求
	OpcodePlayersExtraResp = 5166 // 获取多个玩家扩展信息应答

	// 事件玩法
	OpcodeGamePlayListResp                      = 5170 // 事件玩法列表
	OpcodeGamePlayActiveResp                    = 5171 // 事件玩法激活通知
	OpcodeGamePlayCloseResp                     = 5172 // 事件玩法结束通知
	OpcodeGamePlaySignup                        = 5173 // 玩法报名请求
	OpcodeGamePlaySignupResp                    = 5174 // 玩法报名应答
	OpcodeGamePlayAddContestantResp             = 5175 // 新增参数选手
	OpcodeGamePlayRemoveContestantResp          = 5176 // 移除参数选手
	OpcodeGamePlayContestantScoreResp           = 5177 // 参赛选手积分变化
	OpcodeGamePlayFishingContestShowOffResp     = 5178 // 钓鱼大赛捕获鱼炫耀飘字
	OpcodeGamePlayFishingContestSettlementResp  = 5179 // 钓鱼大赛结算界面
	OpcodeGamePlayFishingContestChangeStageResp = 5180 // 钓鱼大赛阶段变化通知

	//Game1体力推送
	OpcodeStrengthUpdateResp = 5190 //体力数据推送

	OpcodeBrushMaterial                = 5191 //刷地表
	OpcodeBrushMaterialResp            = 5192
	OpcodeBrushTexture                 = 5193 //刷贴图
	OpcodeBrushTextureResp             = 5194
	OpcodeOfficeMaterialUnlockListResp = 5195 //官方解锁的地表列表
	OpcodeOfficeMaterialUnlockAddResp  = 5196 //更新解锁的列表增加
	OpcodeTextureIdUrlListResp         = 5197 //贴图id2url 列表
	OpcodeVoxelTextureListResp         = 5198
	OpcodeEraseMaterial                = 5199
	OpcodeEraseMaterialResp            = 5200
	OpcodeTextureIdUrlUpdate           = 5201
	OpcodeVoxelTextureUpdateResp       = 5203
	OpcodeTextureIdUrlAddResp          = 5204
	OpcodeTextureIdUrlDeleteResp       = 5205

	//Game1摆放物
	OpcodeFurnitureListResp     = 5220 // 数据返回
	OpcodeFurnitureHangUp       = 5221 // 放上装饰物
	OpcodeFurnitureHangUpResp   = 5222 // 放上装饰物返回
	OpcodeFurnitureTakeDown     = 5223 // 取下
	OpcodeFurnitureTakeDownResp = 5224 // 取下返回
	OpcodeFurnitureUpdate       = 5225 // 更新数据
	OpcodeFurnitureUpdateResp   = 5226 // 更新数据返回
	// 家具的交互
	OpcodeFurnitureInteract         = 5227 // 开始交互
	OpcodeFurnitureInteractResp     = 5228 // 开始交互返回
	OpcodeFurnitureInteractOver     = 5229 // 取消交互
	OpcodeFurnitureInteractOverResp = 5230 //
	// 家具打地鼠中手动修改槽位的引用
	OpcodeFurnitureSetSlot     = 5231 // 设置槽位状态
	OpcodeFurnitureSetSlotResp = 5232 //
	// 交互冷却
	OpcodeInteractCoolDownChangeResp = 5233 // 冷却数据变更

	//Game1钓鱼
	OpCodeFishingFallRod     = 5240 //抛杆
	OpCodeFishingFallRodResp = 5241 //抛杆返回
	OpCodeFishingResult      = 5242 //钓鱼结果上报
	OpCodeFishingResultResp  = 5243 //钓鱼结果上报返回
	OpCodeFishingParade      = 5244 //炫耀
	OpCodeFishingParadeResp  = 5245 // 炫耀

	//Game1建筑
	OpCodeBuildingListResp    = 5250 // 建筑数据返回
	OpCodeBuildingBuild       = 5251 //建造
	OpCodeBuildingBuildResp   = 5252 //
	OpCodeBuildingSaveDIY     = 5253 // 建筑自定义保存
	OpCodeBuildingSaveDIYResp = 5254 // 保存返回
	OpCodeBuildingUpdateResp  = 5255 //建筑新加/更新(广播)
	OpCodeBuildingRemove      = 5256 //建筑拆除
	OpCodeBuildingRemoveResp  = 5257 //
	OpCodeBuildingMove        = 5258 //建筑移动性拆除(返还重建图纸)
	OpCodeBuildingMoveResp    = 5259 //

	//Game1客户端模拟server 5260 -> 5279
	OpCodeClientSimControllerResp   = 5260 // 通知客户端其是否为客户端Server模拟器的主控端
	OpCodeClientSimulatorMsg        = 5261 // 主控客户端模拟器Server发送该消息到服务器
	OpCodeCliSimNpcStartReq         = 5262 // 主控模拟服务器收到开始与npc对话请求
	OpCodeCliSimNpcStopReq          = 5263 // 主控模拟服务器收到停止与npc对话请求
	OpcodeCliSimPlayerEnterMapResp  = 5264 // 通知主控端有玩家进入地图
	OpcodeCliSimSaveNpcSnapshotReq  = 5265 // 主控客户端同步Npc快照数据到服务器端
	OpcodeCliSimSyncNpcSnapshotResp = 5266 // 服务器端同步Npc数据到主控客户端
	OpcodeCliSimCloseNpcResp        = 5267 // 服务器通知主控端关闭npc
	OpcodeCliSimSyncNpcSnapshotReq  = 5268 // 获取服务器端Npc数据到主控客户端
	OpcodeCliSimControllerPingReq   = 5269 // 主控客户端发ping，无需pong
	OpcodeCliSimAliveReq            = 5270 // 客户端是否alive请求[server -> client]
	OpcodeCliSimAliveResp           = 5271 // 客户端是否alive应答[client -> server]
	OpcodeCliSimSendAction          = 5272 // 服务器给客户端发送action动作
	OpcodeCliSimSendActionResp      = 5273
	OpcodeCliSimRefreshAgentNpcs    = 5274 //如果有新增代理npc更新主控

	OpCodeCliSimOutOfControlNpcResp = 5275 //
	OpCodeCliSimAddControlNpcResp   = 5278 //

	OpCodeCliSimNpcTakeUpDownResp = 5276
	OpCodeCliSimNpcPlayActionResp = 5277

	OpcodePetHatchAddEgg         = 5280 // 孵化添加宠物蛋
	OpcodePetHatchAddEggResp     = 5281 // 孵化添加宠物蛋返回
	OpcodePetHatchGetPet         = 5282 // 孵化完成，获取宠物
	OpcodePetHatchGetPetResp     = 5283 // 孵化完成，获取宠物返回
	OpcodePetHatchAccelerate     = 5284 // 加速孵化，获取宠物
	OpcodePetHatchAccelerateResp = 5285 // 加速孵化，获取宠物返回

	//Game1房间区域操作
	OpcodeChangeWallpaper      = 5290 // 请求修改墙纸
	OpcodeChangeWallpaperResp  = 5291 //
	OpcodeChangeFloorboard     = 5292 // 请求地板贴图
	OpcodeChangeFloorboardResp = 5293 //

	//Game1矿石
	OpcodeOreListResp         = 5300 // 矿砂数据返回
	OpcodeOreGetPos           = 5301 // 获取客户端位置，广播发出，使用第一个
	OpcodeOreGetPosResp       = 5302 // 位置信息返回
	OpcodeOreGetPosErrMsgResp = 5303 // 给客户端返回错误信息，用于查错
	OpcodeOreInteract         = 5304 // 矿石交互
	OpcodeOreInteractResp     = 5305 // 矿石交互返回

	//Game1建造/建筑/桥梁/斜坡
	OpCodeApplyBuildPrint     = 5331 //领取建筑建造图纸
	OpCodeApplyBuildPrintResp = 5332 //
	OpCodeBuildCancel         = 5333 //建造取消
	OpCodeBuildCancelResp     = 5334 //OpCodeBridgeUpdate(广播)
	OpCodeBuildDonate         = 5335 //建造捐资
	OpCodeBuildDonateResp     = 5336 //OpCodeBridgeUpdate(广播)
	OpCodeBuildSpeedUp        = 5338 //建造加速立即完成
	OpCodeBuildSpeedUpResp    = 5339 //OpCodeBridgeUpdate(广播)
	OpCodeBuildComplete       = 5340 //建造竣工剪彩
	OpCodeBuildCompleteResp   = 5341 //(广播)

	OpCodeBuildingUpgrade       = 5342 //建筑升级
	OpCodeBuildingUpgradeResp   = 5343 //建筑升级
	OpCodeBuildingUpSpeedUp     = 5348 //建筑升级加速立即完成
	OpCodeBuildingUpSpeedUpResp = 5349 //(广播)
	OpcodeBuildOpenUpdateResp   = 5350 //建筑解锁

)
