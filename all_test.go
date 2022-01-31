package main

import (
	"testing"
)

var StartContext = &ManufactureParameters{
	FacilityRate: map[string]float64{
		"制造台":  0.75,
		"电弧熔炉": 1,
	},
	FormulaNums: map[string]int{
		"氢": 1,
	},
}

var SecondContext = &ManufactureParameters{
	FacilityRate: map[string]float64{
		"制造台":  0.75,
		"电弧熔炉": 1,
	},
	FormulaNums: map[string]int{
		"氢":    1,
		"石墨烯":  1,
		"有机晶体": 1,
	},
}

var BlackboxContext = &ManufactureParameters{
	FacilityRate: map[string]float64{
		"制造台":  1.5,
		"电弧熔炉": 2,
	},
	FormulaNums: map[string]int{
		"硫酸":     1,
		"石墨烯":    1,
		"有机晶体":   1,
		"晶格硅":    1,
		"金刚石":    1,
		"碳纳米管":   1,
		"卡西米尔晶体": 1,
		"粒子容器":   1,
		"光子合并器":  1,
		"氢":      1,
		"重氢":     1,
	},
	SprayingResources: map[string]SprayType{
		"增产剂MK1": SprayDisabled,
		"增产剂MK2": SprayDisabled,
		"增产剂MK3": SprayDisabled,
	},
	ImportingResources: map[string]bool{
		"粒子容器": true,
		"铁矿": true,
		"铜矿": true,
		"石矿": true,
		"煤矿": true,
		"钛矿": true,
		"硅石": true,
		"水": true,
		"氢": true,
		"重氢": true,
		"原油": true,
		"可燃冰": true,
		"有机晶体": true,
		"刺笋结晶": true,
		"光栅石": true,
		"金伯利矿石": true,
		"分形硅": true,
		"临界光子": true,
	},
	SprayingClass: SprayingMark3,
}

func TestMatrix(t *testing.T) {
	c := BlackboxContext.Clone()
	//c.SprayingResources["磁线圈"] = SpraySpeedUp
	//c.SprayingResources["电路板"] = SpraySpeedUp
	c.SprayingResources["卡西米尔晶体"] = SpraySpeedUp
	c.SprayingResources["钛晶石"] = SpraySpeedUp
	c.SprayingResources["奇异物质"] = SpraySpeedUp
	c.SprayingResources["能量矩阵"] = SpraySpeedUp
	c.SprayingResources["晶格硅"] = SpraySpeedUp

	c.ShowRequirement(map[string]float64{
		"宇宙矩阵": 10,
	}, templateMatrix)
}

const templateMatrix = `
粒子容器

铁矿
铜矿
石矿
煤矿
钛矿
硅石
水
氢
重氢
原油
可燃冰
有机晶体
刺笋结晶
光栅石
金伯利矿石
分形硅
临界光子

增产剂MK1
增产剂MK2
增产剂MK3

铁块
铜块
电路板
磁铁
磁线圈
电磁矩阵

高能石墨
能量矩阵

金刚石
钛块
钛晶石
结构矩阵

高纯硅块
精炼油
塑料
碳纳米管
晶格硅
粒子宽带
微晶元件
处理器
信息矩阵

石墨烯
卡西米尔晶体
玻璃
钛化玻璃
位面过滤器
量子芯片
奇异物质
引力透镜
引力矩阵

反物质
宇宙矩阵
`