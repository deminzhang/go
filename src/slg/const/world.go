package Const

//世界常量
const (
	MAIN_WORLD  = 0    //主世界ID
	WORLD_WIDTH = 1200 //世界宽
	//区域用于分区刷怪/建号
	AREA_LINE     = 8                       //区域行列数
	AREA_ROWCOL   = AREA_LINE * AREA_LINE   //区域数
	AREA_WIDTH    = WORLD_WIDTH / AREA_LINE //区域宽
	AREA_TILE_NUM = AREA_WIDTH * AREA_WIDTH //区域格数
	//用于视野同步
	SIGHT_WIDTH  = 10 //视块宽
	SIGHT_ROWCOL = WORLD_WIDTH / SIGHT_WIDTH
)
