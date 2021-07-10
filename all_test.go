package main

import (
	"fmt"
	"testing"
)

var StartContext = &ManufactureParameters{
	FacilityRate: map[string]float64{
		"制造台": 0.75,
		"电弧熔炉": 1,
	},
	Formula: map[string]int{
		"氢": 1,
		"石墨烯": 1,
		"硫酸": 1,
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
		"光子合并器": 1,
		"氢": 1,
		"重氢": 1,
	},
}

func TestResearch(t *testing.T) {
	DestRate := 60.0
	fmt.Println("========")
	fmt.Println("科研二所")
	DefaultContext.ShowRequirement([]ResourceRate{
		{"宇宙矩阵",DestRate},
		{"奇异物质",-DestRate/2},
		{"磁线圈", -DestRate},
		{"处理器",-DestRate * 3},
		{"电路板",-DestRate},
		{"钛晶石",-DestRate * 2},
		{"钛化玻璃",-DestRate * 2},
	})
	fmt.Println("========")
	fmt.Println("科研二所Fe")
	DefaultContext.ShowRequirement([]ResourceRate{
		{"奇异物质",DestRate /2},
		{"磁线圈", DestRate},
	})
	fmt.Println("========")
	fmt.Println("科研二所Si")
	DefaultContext.ShowRequirement([]ResourceRate{
		{"处理器",DestRate * 3},
		{"电路板",DestRate},
	})
	fmt.Println("========")
	fmt.Println("科研二所Ti")
	DefaultContext.ShowRequirement([]ResourceRate{
		{"钛晶石",DestRate * 2},
		{"钛化玻璃",DestRate * 2},
	})
}

func TestResearch180(t *testing.T) {
	DestRate := 3.0
	fmt.Println("白糖黑盒")
	DefaultContext.ShowRequirement([]ResourceRate{
		{"宇宙矩阵",DestRate},
	})
}

func TestN(t *testing.T) {
	fmt.Println("太阳帆工厂")
	DefaultContext.ShowRequirement([]ResourceRate{
		{"太阳帆",240},
	})
}
