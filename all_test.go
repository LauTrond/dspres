package main

import (
	"testing"
)

var StartContext = &ManufactureParameters{
	FacilityRate: map[string]float64{
		"制造台": 0.75,
		"电弧熔炉": 1,
	},
	Formula: map[string]int{
		"氢": 1,
	},
}

var SecondContext = &ManufactureParameters{
	FacilityRate: map[string]float64{
		"制造台": 0.75,
		"电弧熔炉": 1,
	},
	Formula: map[string]int{
		"氢": 1,
		"石墨烯": 1,
		"有机晶体": 1,
	},
}

var DefaultContext = &ManufactureParameters{
	FacilityRate: map[string]float64{
		"制造台": 1.5,
		"电弧熔炉": 2,
	},
	Formula: map[string]int{
		"石墨烯": 1,
		"有机晶体": 1,
		"晶格硅": 1,
		"金刚石": 1,
		"碳纳米管": 1,
		"卡西米尔晶体": 0,
		"粒子容器": 0,
		"光子合并器": 0,
		"氢": 1,
		"重氢": 1,
	},
}

var AllResContext = &ManufactureParameters{
	FacilityRate: map[string]float64{
		"制造台": 1.5,
		"电弧熔炉": 2,
	},
	Formula: map[string]int{
		"石墨烯": 1,
		"有机晶体": 1,
		"晶格硅": 1,
		"金刚石": 1,
		"碳纳米管": 1,
		"卡西米尔晶体": 1,
		"粒子容器": 1,
		"光子合并器": 1,
		"氢": 1,
		"重氢": 1,
	},
}

func TestStart(t *testing.T) {
	p := &ManufactureParameters{
		FacilityRate: map[string]float64{
			"制造台": 0.75,
			"电弧熔炉": 1,
		},
		Formula: map[string]int{
			"氢": 1,
		},
		ImportingResources: map[string]bool{
			"氢": true,
		},
	}

	p.ShowRequirement([]ResourceRate{
		{"硫酸", 6.0},
	})
	p.ImportingResources["硫酸"] = true

	p.ShowRequirement([]ResourceRate{
		{"处理器", 3.0},
	})
	p.ImportingResources["处理器"] = true


	p.ShowRequirement([]ResourceRate{
		{"粒子容器", 0.375},
	})
	p.ImportingResources["粒子容器"] = true

	p.ShowRequirement([]ResourceRate{
		{"电磁涡轮", 0.375},
	})
	p.ImportingResources["电磁涡轮"] = true

	p.ShowRequirement([]ResourceRate{
		{"钛块", 1.0},
		{"钛合金", 3.0},
	})
	p.ImportingResources["钛块"] = true
	p.ImportingResources["钛合金"] = true

	p.ShowRequirement([]ResourceRate{
		{"行星内物流运输站", 1.0/60},
		{"星际物流运输站", 0.5/60},
		{"物流运输机", 22.5/60},
		{"星际物流运输船", 7.5/60},
	})
}

func TestRocketBox(t *testing.T) {

	(&ManufactureParameters{
		FacilityRate: map[string]float64{
			"制造台": 1.5,
			"电弧熔炉": 2,
		},
		Formula: map[string]int{
			"石墨烯": 1,
			"有机晶体": 1,
			"晶格硅": 1,
			"金刚石": 1,
			"碳纳米管": 1,
			"氢": 1,
			"重氢": 1,
		},
	}).ShowRequirement([]ResourceRate{
		{"小型运输火箭",0.4},
	})
}