package qimen

// 奇门三元
var _Yuan3Name = []string{"", "上元", "中元", "下元"}

// 节气索引
var _JieQiIndex = map[string]int{
	"冬至": 1, "小寒": 2, "大寒": 3,
	"立春": 4, "雨水": 5, "惊蛰": 6,
	"春分": 7, "清明": 8, "谷雨": 9,
	"立夏": 10, "小满": 11, "芒种": 12,
	"夏至": 13, "小暑": 14, "大暑": 15,
	"立秋": 16, "处暑": 17, "白露": 18,
	"秋分": 19, "寒露": 20, "霜降": 21,
	"立冬": 22, "小雪": 23, "大雪": 24,
}

// 奇门局数
var _QiMenJu = [][]int{{0, 0, 0},
	{+1, +7, +4}, {+2, +8, +5}, {+3, +9, +6}, //坎宫 冬至{上元,中元,下元},小寒{上元,中元,下元},大寒{上元,中元,下元},
	{+8, +5, +2}, {+9, +6, +3}, {+1, +7, +4}, //艮宫 立春...
	{+3, +9, +6}, {+4, +1, +7}, {+5, +2, +8}, //震...
	{+4, +1, +7}, {+5, +2, +8}, {+6, +3, +9}, //巽
	{-9, -3, -6}, {-8, -2, -5}, {-7, -1, -4}, //离
	{-2, -5, -8}, {-1, -4, -7}, {-9, -3, -6}, //坤
	{-7, -1, -4}, {-6, -9, -3}, {-5, -8, -2}, //兑
	{-6, -9, -3}, {-5, -8, -2}, {-4, -7, -1}, //乾
	//{1, 7, 4},
}

// 盘式
var _QMType = []string{"转盘", "飞盘", "鸣法"}

const (
	QMTypeRollDoor = 0
	QMTypeFlyDoor  = 1
	QMTypeAmaze    = 2
)

// 飞盘九星飞宫 "鸣法"=="阴阳皆顺"
var _QMFlyType = []string{"阴阳皆顺", "阳顺阴逆?"}

const (
	QMFlyTypeAllOrder   = 0
	QMFlyTypeYinReverse = 1
)

// 转盘寄宫
var _QMRollHostingType = []string{"中宫寄坤?", "阳艮阴坤", "长夏寄四库?"}

const (
	QMRollHostingType2    = 0
	QMRollHostingType28   = 1
	QMRollHostingType2846 = 3
)

// 起局方式
var _QMStartType = []string{"拆补", "茅山", "置闰", "自选"}

// 暗干起法
var _QMHideGanType = []string{"值使门起", "门地盘起"}

// 宫序环
var _GongIdx = []int{9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// 九宫八卦
var _Gua8In9 = []string{"", "坎", "坤", "震", "巽", "中", "乾", "兑", "艮", "离"}

// 三奇六仪
var _Qm3Q6Y = []string{"戊", "己", "庚", "辛", "壬", "癸", "丁", "丙", "乙"}
var _Qm3Q6YCircle = []string{"戊", "己", "庚", "辛", "壬", "癸", "丁", "丙", "乙", "戊", "己", "庚", "辛", "壬", "癸", "丁", "丙"}

// 旬首遁甲
var _HideJia = map[string]string{
	"甲子": "戊",
	"甲戌": "己",
	"甲申": "庚",
	"甲午": "辛",
	"甲辰": "壬",
	"甲寅": "癸",
}
var _XunIdx = map[string]string{
	"甲子": "甲子乙",
	"甲戌": "甲戌",
	"甲申": "庚",
	"甲午": "辛",
	"甲辰": "壬",
	"甲寅": "癸",
}

// 奇门九星
var _QiMenStar9 = []string{"", "天蓬", "天芮", "天冲", "天辅", "天禽", "天心", "天柱", "天任", "天英"}
var _QiMenStar9Circle = []string{"", "天蓬", "天芮", "天冲", "天辅", "天禽", "天心", "天柱", "天任", "天英",
	"天蓬", "天芮", "天冲", "天辅", "天禽", "天心", "天柱", "天任"}

// 奇门九神
// 九神飞盘阳遁
var _QiMenGod9S = []string{"值符", "腾蛇", "太阴", "六合", "勾陈", "太常", "朱雀", "九地", "九天"}
var _QiMenGod9SCircle = []string{
	"值符", "腾蛇", "太阴", "六合", "勾陈", "太常", "朱雀", "九地", "九天",
	"值符", "腾蛇", "太阴", "六合", "勾陈", "太常", "朱雀", "九地", "九天"}

// 九神飞盘阴遁
var _QiMenGod9L = []string{"值符", "腾蛇", "太阴", "六合", "白虎", "太常", "玄武", "九地", "九天"}
var _QiMenGod9LCircle = []string{
	"值符", "腾蛇", "太阴", "六合", "白虎", "太常", "玄武", "九地", "九天",
	"值符", "腾蛇", "太阴", "六合", "白虎", "太常", "玄武", "九地", "九天"}

// 奇门八门
var _QiMenDoor9 = []string{"", "休门", "死门", "伤门", "杜门", "中门", "开门", "惊门", "生门", "景门"}

// 农历日期信息
// 阴历1900年到2100年每年中的月天数信息
// 阴历每月只能是29或30天，一年用12（或13）个二进制位表示，对应位为1 代表30天，否则为29天
// 闰月不会出现在正月、冬月或腊月,不会连续两年闰月
//var lunarMonthData = [201]int{
//	// 0xf   =0000 0000 0000 1111
//	//       =1234 5678 9ABC 1000闰月月数8
//	//       =1234 5678 9ABC 去年闰月大小 0小 1111大
//	0x4bd8, //1900:0100 1011 1101 1000(小大小小 大小大大 大大小大 闰八月)
//	0x4ae0, //1901:0100 1010 1110 0000(去年闰八月小)
//	0xa570,
//	0x54d5, //1903:0101 0100 1101 0101(闰五月)
//	0xd260, //1904:1101 0010 0110 0000(去年闰五月小)
//	0xd950,
//	0x5554, //1906:0101 0101 0101 0100(闰四月)
//	0x56af, //1907:0101 0110 1010 1111(去年闰四月大)
//	0x9ad0, 0x55d2,
//	0x4ae0, 0xa5b6, 0xa4d0, 0xd250, 0xd255, 0xb54f, 0xd6a0, 0xada2, 0x95b0, 0x4977,
//	0x497f, 0xa4b0, 0xb4b5, 0x6a50, 0x6d40, 0xab54, 0x2b6f, 0x9570, 0x52f2, 0x4970,
//	0x6566, 0xd4a0, 0xea50, 0x6a95, 0x5adf, 0x2b60, 0x86e3, 0x92ef, 0xc8d7, 0xc95f,
//	0xd4a0, 0xd8a6, 0xb55f, 0x56a0, 0xa5b4, 0x25df, 0x92d0, 0xd2b2, 0xa950, 0xb557,
//	0x6ca0, 0xb550, 0x5355, 0x4daf, 0xa5b0, 0x4573, 0x52bf, 0xa9a8, 0xe950, 0x6aa0,
//	0xaea6, 0xab50, 0x4b60, 0xaae4, 0xa570, 0x5260, 0xf263, 0xd950, 0x5b57, 0x56a0,
//	0x96d0, 0x4dd5, 0x4ad0, 0xa4d0, 0xd4d4, 0xd250, 0xd558, 0xb540, 0xb6a0, 0x95a6,
//	0x95bf, 0x49b0, 0xa974, 0xa4b0, 0xb27a, 0x6a50, 0x6d40, 0xaf46, 0xab60, 0x9570,
//	0x4af5, 0x4970, 0x64b0, 0x74a3, 0xea50, 0x6b58, 0x5ac0, 0xab60, 0x96d5, 0x92e0,
//	0xc960, 0xd954, 0xd4a0, 0xda50, 0x7552, 0x56a0, 0xabb7, 0x25d0, 0x92d0, 0xcab5,
//	0xa950, 0xb4a0, 0xbaa4, 0xad50, 0x55d9, 0x4ba0, 0xa5b0, 0x5176, 0x52bf, 0xa930,
//	0x7954, 0x6aa0, 0xad50, 0x5b52, 0x4b60, 0xa6e6, 0xa4e0, 0xd260, 0xea65, 0xd530,
//	0x5aa0, 0x76a3, 0x96d0, 0x4afb, 0x4ad0, 0xa4d0, 0xd0b6, 0xd25f, 0xd520, 0xdd45,
//	0xb5a0, 0x56d0, 0x55b2, 0x49b0, 0xa577, 0xa4b0, 0xaa50, 0xb255, 0x6d2f, 0xada0,
//	0x4b63, 0x937f, 0x49f8, 0x4970, 0x64b0, 0x68a6, 0xea5f, 0x6b20, 0xa6c4, 0xaaef,
//	0x92e0, 0xd2e3, 0xc960, 0xd557, 0xd4a0, 0xda50, 0x5d55, 0x56a0, 0xa6d0, 0x55d4,
//	0x52d0, 0xa9b8, 0xa950, 0xb4a0, 0xb6a6, 0xad50, 0x55a0, 0xaba4, 0xa5b0, 0x52b0,
//	0xb273, 0x6930, 0x7337, 0x6aa0, 0xad50, 0x4b55, 0x4b6f, 0xa570, 0x54e4, 0xd260,
//	0xe968, 0xd520, 0xdaa0, 0x6aa6, 0x56df, 0x4ae0, 0xa9d4, 0xa4d0, 0xd150, 0xf252,
//	0xd520, //2100:1101 0101 0010 0000
//}

// 廿四节气信息 (0小寒)
// 从第0个节气的分钟数
var termData = []int{
	0, 21208, 42467, 63836, 85337, 107014, 128867, 150921, 173149, 195551, 218072, 240693,
	263343, 285989, 308563, 331033, 353350, 375494, 397447, 419210, 440795, 462224, 483532,
	504758,
}

// ShowChars
const (
	HS     = "甲乙丙丁戊已庚辛壬癸"
	EB     = "子丑寅卯辰巳午末申酉戌亥"
	Term24 = "小寒大寒立春雨水惊蛰春分清明谷雨立夏小满芒种夏至小暑大暑立秋处暑白露秋分寒露霜降立冬小雪大雪冬至"

	Diagrams8   = "坎坤震巽乾竞艮离"                 //八卦
	Diagrams9   = "坎坤震巽  乾竞艮离"               //九宫八卦
	QMStars     = "蓬芮冲辅禽心柱任英"                //奇门九星
	QMGodsRoll  = "值符腾蛇太阴六合白虎玄武九地九天"         //奇门转盘八神
	QMGodsFly   = "值符腾蛇太阴六合勾陈太常朱雀九地九天"       //奇门飞盘九神
	QMDoorsRoll = "休生伤杜景死惊开"                 //转盘八门
	QMDoorsFly  = "休死伤杜中开惊生景"                //飞盘九门
	BuildStar12 = "建除满平定执破危成收开闭"             //十二建星
	God12       = "青龙明堂天刑朱雀金匮天德白虎玉堂天牢玄武司命勾陈" //十二原神
)
