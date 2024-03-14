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
		"硫酸":   1,
	},
	SprayingResources: map[string]SprayType{
		"增产剂MK1": SprayDisabled,
		"增产剂MK2": SprayDisabled,
		"增产剂MK3": SprayDisabled,
		"晶格硅":    SpraySpeedUp,
		"钛合金":    SprayExtra,
		"钢材":     SprayExtra,
	},
	ImportingResources: map[string]float64{},
}

var BlackboxContext = &ManufactureParameters{
	FacilityRate: map[string]float64{
		"制造台":  1.5,
		"电弧熔炉": 2,
		"化工厂":  2,
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
		"碳纳米管": SprayExtra,
		"晶格硅":  SpraySpeedUp,
		"钛合金":  SprayExtra,
		"钢材":   SprayExtra,
	},
	ImportingResources: map[string]float64{
		"粒子容器":  16,
		"铁矿":    1,
		"铜矿":    1,
		"石矿":    1,
		"煤矿":    8,
		"钛矿":    1,
		"硅石":    1,
		"水":     6,
		"氢":     1,
		"重氢":    6,
		"原油":    6,
		"可燃冰":   6,
		"有机晶体":  6,
		"刺笋结晶":  12,
		"光栅石":   12,
		"金伯利矿石": 8,
		"分形硅":   10,
		"临界光子":  16,
	},
	SprayingClass: SprayingMark3,
}

var DarkForkBlackboxContext = &ManufactureParameters{
	FacilityRate: map[string]float64{
		"制造台":  3,
		"电弧熔炉": 3,
		"化工厂":  2,
		"研究所":  3,
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
		"碳纳米管":   SpraySpeedUp,
		"晶格硅":    SpraySpeedUp,
		"钛合金":    SprayExtra,
		"奇异物质":   SpraySpeedUp,
		"卡西米尔晶体": SprayDisabled,
		"增产剂MK1": SprayDisabled,
		"增产剂MK2": SprayDisabled,
		"增产剂MK3": SprayDisabled,
	},
	ImportingResources: map[string]float64{
		"粒子容器":  16,
		"铁矿":    1,
		"铜矿":    1,
		"石矿":    1,
		"煤矿":    8,
		"钛矿":    1,
		"硅石":    1,
		"水":     6,
		"氢":     1,
		"重氢":    6,
		"原油":    6,
		"可燃冰":   6,
		"有机晶体":  6,
		"刺笋结晶":  12,
		"光栅石":   12,
		"金伯利矿石": 8,
		"分形硅":   10,
		"临界光子":  16,
	},
	SprayingClass: SprayingMark3,
}

func TestMatrix(t *testing.T) {
	c := BlackboxContext.Clone()
	c.SetImportingResources(map[string]float64{
		"煤矿":    16,
		"水":     16,
		"原油":    16,
		"可燃冰":   6,
		"有机晶体":  16,
		"刺笋结晶":  16,
		"光栅石":   16,
		"金伯利矿石": 16,
		"分形硅":   16,
		"临界光子":  16,
		"粒子容器":  40,
	})
	c.ShowRequirement(map[string]float64{
		"宇宙矩阵": 10,
	}, templateMatrix)
}

func TestDarkForkMatrix(t *testing.T) {
	c := DarkForkBlackboxContext.Clone()
	c.SetImportingResources(map[string]float64{
		"煤矿":    16,
		"水":     16,
		"原油":    16,
		"可燃冰":   6,
		"有机晶体":  16,
		"刺笋结晶":  16,
		"光栅石":   16,
		"金伯利矿石": 16,
		"分形硅":   16,
		"临界光子":  16,
		"粒子容器":  40,
	})
	c.ShowRequirement(map[string]float64{
		"宇宙矩阵": 1,
	}, templateMatrix)
}

func TestRocket(t *testing.T) {
	c := BlackboxContext.Clone()
	c.SetImportingResources(map[string]float64{
		"煤矿":    16,
		"水":     16,
		"原油":    16,
		"可燃冰":   6,
		"有机晶体":  16,
		"刺笋结晶":  16,
		"光栅石":   16,
		"金伯利矿石": 16,
		"分形硅":   16,
		"临界光子":  16,
	})
	c.ShowRequirement(map[string]float64{
		"小型运输火箭": 3.125,
	}, templateRocket)
}

func TestEng(t *testing.T) {
	c := BlackboxContext.Clone()
	c.SetImportingResources(map[string]float64{
		"增产剂MK3": 16,
		"粒子容器":   30,
		"反物质燃料棒": 1,
	})
	c.SprayingResources = map[string]SprayType{
		"铁块":      SprayDisabled,
		"钢材":      SprayDisabled,
		"钛块":      SprayDisabled,
		"钛合金":     SprayDisabled,
		"高纯硅块":    SprayDisabled,
		"碳纳米管":    SprayDisabled,
		"奇异物质":    SprayDisabled,
		"奇异湮灭燃料棒": SpraySpeedUp,
	}
	c.ShowRequirement(map[string]float64{
		"奇异湮灭燃料棒": 0.75,
	}, templateEng)
}

func TestContainer(t *testing.T) {
	c := &ManufactureParameters{
		FacilityRate: map[string]float64{
			"制造台":  1.5,
			"电弧熔炉": 2,
			"化工厂":  2,
		},
		FormulaNums: map[string]int{
			"硫酸":     1,
			"石墨烯":    1,
			"有机晶体":   1,
			"晶格硅":    1,
			"金刚石":    1,
			"碳纳米管":   1,
			"卡西米尔晶体": 1,
			"光子合并器":  1,
			"氢":      1,
			"重氢":     1,
		},
		SprayingResources: map[string]SprayType{},
		ImportingResources: map[string]float64{
			"增产剂MK3": 16,
		},
		SprayingClass: SprayingMark3,
	}
	c.ShowRequirement(map[string]float64{
		"粒子容器": 22.5,
	}, "")
}

func TestNuclear(t *testing.T) {
	c := &ManufactureParameters{
		FacilityRate: map[string]float64{
			"制造台":  0.75,
			"电弧熔炉": 1,
		},
		FormulaNums: map[string]int{
			"氢":    1,
			"石墨烯":  0,
			"有机晶体": 1,
			"硫酸":   1,
		},
		SprayingResources: map[string]SprayType{
			"增产剂MK1": SprayDisabled,
			"增产剂MK2": SprayDisabled,
			"增产剂MK3": SprayDisabled,
			"铁块":     SprayDisabled,
			"磁铁":     SprayDisabled,
			"铜块":     SprayDisabled,
			"钛块":     SprayDisabled,
			"高能石墨":   SprayDisabled,
			"金刚石":    SprayDisabled,
			"石墨烯":    SprayDisabled,
			"碳纳米管":   SprayDisabled,
			"钛合金":    SprayExtra,
			"钢材":     SprayExtra,
		},
		SprayingClass: SprayingMark3,
		ImportingResources: map[string]float64{
			"重氢": 1,
		},
	}
	c.ShowRequirement(map[string]float64{
		"氘核燃料棒": 0.625,
	}, "")
}

func TestMissiles(t *testing.T) {
	c := SecondContext.Clone()
	c.SprayingClass = SprayingMark3
	c.SetImportingResources(map[string]float64{
		"增产剂MK3": 16,
		"金刚石":    1,
		"精炼油":    1,
		"钢材":     1,
		"铜块":     1,
		"硫酸":     1,
		"水":      1,
		"磁线圈":    1,
		"电路板":    1,
		"塑料":     1,
		"处理器":    1,
		"钛晶石":    1,
		"卡西米尔晶体": 1,
		"粒子容器":   30,
		"奇异物质":   1,
	})
	c.SprayingResources = map[string]SprayType{
		"爆破单元":   SprayExtra,
		"晶石爆破单元": SprayExtra,
		//"导弹组":    SprayExtra,
		//"超音速导弹组": SprayExtra,
		//"引力导弹组":  SprayExtra,
	}
	c.ShowRequirement(map[string]float64{
		"引力导弹组": 3,
	}, "")

}

func TestAny(t *testing.T) {
	c := BlackboxContext.Clone()
	c.SetImportingResources(map[string]float64{
		"增产剂MK3": 16,
		"粒子容器":   30,
		"临界光子":   1,
	})
	c.SprayingResources = map[string]SprayType{
		"磁铁":    SprayDisabled,
		"石墨烯":   SprayDisabled,
		"等离子胶囊": SprayDisabled,
		"反物质胶囊": SprayDisabled,
	}
	c.ShowRequirement(map[string]float64{
		"反物质胶囊": 2.25,
	}, "")
}

const templateMatrix = `
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
粒子容器
奇异物质
引力透镜
引力矩阵

反物质
宇宙矩阵
`

const templateRocket = `
铁矿
铜矿
石矿
硅石
钛矿
氢
重氢
水
煤矿
硫酸
金伯利矿石
光栅石
刺笋结晶
可燃冰

#### 基础材料 ####

铁块
铜块
高纯硅块
钛块

#### 增产剂 ####

金刚石
#碳纳米管->框架材料
碳纳米管
增产剂MK1
增产剂MK2
增产剂MK3

#### 氘棒 ####

磁铁
磁线圈
齿轮
电动机
电磁涡轮
高能石墨
超级磁场环
氘核燃料棒

#### 芯片 ####

#电路板->光子合并器
电路板
微晶元件
#处理器->戴森球组件
处理器
#石墨烯->太阳帆
石墨烯
卡西米尔晶体
玻璃
钛化玻璃
位面过滤器
量子芯片

#### 球体 ####

光子合并器
太阳帆
钢材
钛合金
框架材料
戴森球组件
小型运输火箭
`

const templateEng = `
反物质燃料棒
核心素
粒子容器
增产剂MK3
刺笋结晶
铁矿
硅石
钛矿
重氢
硫酸
碳纳米管
高纯硅块
铁块
钢材
钛块
钛合金
框架材料
奇异物质
奇异湮灭燃料棒
`
