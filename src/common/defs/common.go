package defs

const (
	ToolHoe        = 1 // 锄头
	ToolShovel     = 2 // 铲子
	ToolAxe        = 3 // 斧子
	ToolPickaxe    = 4 // 镐
	ToolKettle     = 5 // 水壶
	ToolFishingRod = 6 // 鱼竿
	ToolMagicStick = 7 // 魔法棒
)

const (
	ToolParamIdxLevel  = 1 // 工具等级在ItemInfo.ItemParameter中的索引号
	ToolParamIdxDamage = 2 // 工具伤害在ItemInfo.ItemParameter中的索引号
)

const (
	UnitBase        = 0  //空姬
	UnitPlayer      = 1  //玩家
	UnitNpc         = 2  //NPC
	UnitBullet      = 3  //子弹
	UnitLaser       = 4  //闪光(后端无.仅通知前端)
	UnitBuilding    = 5  //建筑
	UnitDropItem    = 6  //掉落物
	UnitPlant       = 7  //植物
	UnitFurniture   = 8  //摆放物（原家具）
	UnitBridge      = 9  //桥或坡
	UnitOre         = 10 //矿石
	UnitHouseRegion = 11 //房间 区域
	UnitPet         = 12 //宠物
	UnitSummon      = 13 //NPC
	UnitFishCrowd   = 14 //鱼群
)

const (
	//地图类型

	Game1MapTypeHomestead = 1 //家园私人地图(岛及室内)
	Game1MapTypeSquare    = 2 //家园广场
	Game1MapTypeDungeon   = 3 //家园副本

)

const (
	MaxCellBits     = 36
	TotalIdBits     = 53
	HostIdBits      = 19
	MaxHostId       = 1<<HostIdBits - 1
	HostIdShiftBits = TotalIdBits - HostIdBits
	MaxCellId       = 1 << MaxCellBits
	MaxCellMask     = MaxCellId - 1
)

const (
	FrameTypeServer = 1
	FrameTypeClient = 2
)

const (
	//GameId GameName
	GameNone    = iota //占位
	HomeSteadId        //家园
	MaxGames

	HomeSteadName        = "clubkoala"
	HomeSteadAssetPrefix = "homestead" //家园
)
