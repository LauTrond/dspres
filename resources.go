package main

import (
	"fmt"
	"sort"
)

type Resource struct {
	Name         string
	Formula      int     //配方编号
	Facility     string
	Duration     float64 //生产时长
	Num          float64 //每次产出数量
	Requirements []ResourceRequirement
}

func (r *Resource) Key() string {
	return fmt.Sprintf("%s-%d", r.Name, r.Formula)
}

type ResourceRequirement struct {
	ResourceName string
	Num          float64 //每次需求数量
}

type ManufactureParameters struct {
	FacilityRate       map[string]float64
	Formula            map[string]int
	ImportingResources map[string]bool
}

var AllResources = []*Resource{
	//基础资源
	{"铁矿", 0, "采矿机", 1, 1, nil},
	{"铜矿", 0, "采矿机", 1, 1, nil},
	{"石矿", 0, "采矿机", 1, 1, nil},
	{"煤矿", 0, "采矿机", 1, 1, nil},
	{"钛矿", 0, "采矿机", 1, 1, nil},
	{"硅石", 0, "采矿机", 1, 1, nil},
	{"水", 0, "抽水机", 1, 1, nil},
	{"氢", 0, "原油精炼厂", 8, 3, []ResourceRequirement{
		{"原油", 2},
		{"高能石墨", -2},
	}},
	{"氢", 1, "行星采集器", 1, 20, nil},
	{"重氢", 0, "分馏器", 100.0 / 30, 1, []ResourceRequirement{
		{"氢", 1},
	}},
	{"重氢", 1, "行星采集器", 1, 1, nil},

	//珍奇
	{"原油", 0, "原油萃取机", 1, 1, nil},
	{"硫酸",1,  "抽水机", 1, 1, nil},
	{"可燃冰", 0, "行星采集器", 1, 1, nil},
	{"有机晶体",1,  "采矿机", 1, 1, nil},
	{"刺笋结晶", 0, "采矿机", 1, 1, nil},
	{"光栅石",0,  "采矿机", 1, 1, nil},
	{"金伯利矿石", 0, "采矿机", 1, 1, nil},
	{"分形硅",0,  "采矿机", 1, 1, nil},
	{"单极磁石",0,  "采矿机", 1, 1, nil},

	//
	{"临界光子", 0, "微波接收器", 12, 1, nil},
	{"反物质", 0, "粒子对撞机", 2, 2, []ResourceRequirement{
		{"临界光子", 2},
		{"氢", -2},
	}},

	//
	{"铁块", 0, "电弧熔炉", 1, 1, []ResourceRequirement{
		{"铁矿", 1},
	}},
	{"磁铁", 0, "电弧熔炉", 1.5, 1, []ResourceRequirement{
		{"铁矿", 1},
	}},
	{"铜块", 0, "电弧熔炉", 1, 1, []ResourceRequirement{
		{"铜矿", 1},
	}},
	{"高纯硅块", 0, "电弧熔炉", 2, 1, []ResourceRequirement{
		{"硅石", 2},
	}},
	{"钛块", 0, "电弧熔炉", 2, 1, []ResourceRequirement{
		{"钛矿", 2},
	}},
	{"石材", 0, "电弧熔炉", 1, 1, []ResourceRequirement{
		{"石矿", 1},
	}},
	{"玻璃",0,  "电弧熔炉", 2, 1, []ResourceRequirement{
		{"石矿", 2},
	}},
	{"高能石墨", 0, "电弧熔炉", 2, 1, []ResourceRequirement{
		{"煤矿", 2},
	}},
	{"金刚石", 0, "电弧熔炉", 2, 1, []ResourceRequirement{
		{"高能石墨", 1},
	}},
	{"金刚石",1,  "电弧熔炉", 1.5, 2, []ResourceRequirement{
		{"金伯利矿石", 1},
	}},
	{"精炼油",0,  "原油精炼厂", 4, 2, []ResourceRequirement{
		{"原油", 2},
		{"氢", -1},
	}},
	{"晶格硅", 0, "电弧熔炉", 2, 1, []ResourceRequirement{
		{"高纯硅块", 1},
	}},
	{"晶格硅", 1, "制造台", 1.5, 2, []ResourceRequirement{
		{"分形硅", 1},
	}},
	{"钢材",0,  "电弧熔炉", 3, 1, []ResourceRequirement{
		{"铁块", 3},
	}},

	{"硫酸",0,  "化工厂", 6, 4, []ResourceRequirement{
		{"精炼油", 6},
		{"石矿", 8},
		{"水", 4},
	}},
	{"塑料",0,  "化工厂", 3, 1, []ResourceRequirement{
		{"精炼油", 2},
		{"高能石墨", 1},
	}},
	{"有机晶体",0,  "化工厂", 6, 1, []ResourceRequirement{
		{"塑料", 2},
		{"精炼油", 1},
		{"水", 1},
	}},
	{"石墨烯", 0, "化工厂", 3, 2, []ResourceRequirement{
		{"高能石墨", 3},
		{"硫酸", 1},
	}},
	{"石墨烯", 1, "化工厂", 2, 2, []ResourceRequirement{
		{"可燃冰", 2},
		{"氢", -1},
	}},
	{"碳纳米管", 0, "化工厂", 4, 2, []ResourceRequirement{
		{"石墨烯", 3},
		{"钛块", 1},
	}},
	{"碳纳米管", 1, "化工厂", 4, 2, []ResourceRequirement{
		{"刺笋结晶", 2},
	}},



	{"电路板",0,  "制造台", 1, 2, []ResourceRequirement{
		{"铁块", 2},
		{"铜块", 1},
	}},
	{"磁线圈", 0, "制造台", 1, 2, []ResourceRequirement{
		{"磁铁", 2},
		{"铜块", 1},
	}},
	{"齿轮", 0, "制造台", 1, 1, []ResourceRequirement{
		{"铁块", 1},
	}},
	{"电动机",0,  "制造台", 2, 1, []ResourceRequirement{
		{"铁块", 2},
		{"齿轮", 1},
		{"磁线圈", 1},
	}},
	{"电磁涡轮", 0, "制造台", 2, 1, []ResourceRequirement{
		{"电动机", 2},
		{"磁线圈", 2},
	}},
	{"超级磁场环", 0, "制造台", 3, 1, []ResourceRequirement{
		{"电磁涡轮", 2},
		{"磁铁", 3},
		{"高能石墨", 1},
	}},
	{"粒子容器", 0, "制造台", 4, 1, []ResourceRequirement{
		{"电磁涡轮", 2},
		{"铜块", 2},
		{"石墨烯", 2},
	}},
	{"粒子容器", 1, "制造台", 4, 1, []ResourceRequirement{
		{"单极磁石", 10},
		{"铜块", 2},
	}},

	//==============================

	{"棱镜",0,  "制造台", 2, 2, []ResourceRequirement{
		{"玻璃", 3},
	}},
	{"电浆激发器", 0, "制造台", 2, 1, []ResourceRequirement{
		{"磁线圈", 4},
		{"棱镜", 2},
	}},
	{"钛晶石",0,  "制造台", 4, 1, []ResourceRequirement{
		{"有机晶体", 1},
		{"钛块", 3},
	}},
	{"钛合金",0,  "电弧熔炉", 12, 4, []ResourceRequirement{
		{"钛块", 4},
		{"钢材", 4},
		{"硫酸", 8},
	}},

	//==============================

	{"粒子宽带", 0, "制造台", 8, 1, []ResourceRequirement{
		{"碳纳米管", 2},
		{"晶格硅", 2},
		{"塑料", 1},
	}},
	{"微晶元件",0,  "制造台", 2, 1, []ResourceRequirement{
		{"高纯硅块", 2},
		{"铜块", 1},
	}},
	{"处理器", 0, "制造台", 3, 1, []ResourceRequirement{
		{"电路板", 2},
		{"微晶元件", 2},
	}},

	//==============================

	{"钛化玻璃",0,  "制造台", 5, 2, []ResourceRequirement{
		{"玻璃", 2},
		{"钛块", 2},
		{"水", 2},
	}},
	{"卡西米尔晶体", 0, "制造台", 4, 1, []ResourceRequirement{
		{"钛晶石", 1},
		{"石墨烯", 2},
		{"氢", 12},
	}},
	{"卡西米尔晶体", 1, "制造台", 4, 1, []ResourceRequirement{
		{"光栅石", 4},
		{"石墨烯", 2},
		{"氢", 12},
	}},
	{"位面过滤器",0,  "制造台", 12, 1, []ResourceRequirement{
		{"卡西米尔晶体", 1},
		{"钛化玻璃", 2},
	}},
	{"量子芯片", 0, "制造台", 6, 1, []ResourceRequirement{
		{"处理器", 2},
		{"位面过滤器", 2},
	}},

	{"奇异物质",0,  "粒子对撞机", 8, 1, []ResourceRequirement{
		{"粒子容器", 2},
		{"铁块", 2},
		{"重氢", 10},
	}},

	{"引力透镜",0,  "制造台", 6, 1, []ResourceRequirement{
		{"金刚石", 4},
		{"奇异物质", 1},
	}},

	//==============================

	{"液氢燃料棒",0,  "制造台", 6, 2, []ResourceRequirement{
		{"钛块", 1},
		{"氢", 10},
	}},
	{"氘核燃料棒",0,  "制造台", 12, 2, []ResourceRequirement{
		{"重氢", 20},
		{"钛合金", 1},
		{"超级磁场环", 1},
	}},
	{"湮灭约束球",0,  "制造台", 20, 1, []ResourceRequirement{
		{"粒子容器", 1},
		{"处理器", 1},
	}},
	{"反物质燃料棒",0,  "制造台", 24, 2, []ResourceRequirement{
		{"反物质", 12},
		{"氢", 12},
		{"湮灭约束球", 1},
		{"钛合金", 1},
	}},

	//==============================

	{"推进器",0,  "制造台", 4, 1, []ResourceRequirement{
		{"钢材", 2},
		{"铜块", 3},
	}},
	{"物流运输机",0,  "制造台", 4, 1, []ResourceRequirement{
		{"铁块", 5},
		{"处理器", 2},
		{"推进器", 2},
	}},
	{"加力推进器",0,  "制造台", 6, 1, []ResourceRequirement{
		{"钛合金", 5},
		{"电磁涡轮", 5},
	}},
	{"星际物流运输船", 0, "制造台", 6, 1, []ResourceRequirement{
		{"钛合金", 10},
		{"处理器", 10},
		{"加力推进器", 2},
	}},
	{"空间翘曲器",0,  "制造台", 10, 8, []ResourceRequirement{
		{"引力矩阵", 1},
	}},

	//==============================

	{"光子合并器", 0, "制造台", 3, 1, []ResourceRequirement{
		{"棱镜", 2},
		{"电路板", 1},
	}},
	{"光子合并器",1,  "制造台", 3, 1, []ResourceRequirement{
		{"光栅石", 1},
		{"电路板", 1},
	}},
	{"太阳帆",0,  "制造台", 4, 2, []ResourceRequirement{
		{"石墨烯", 1},
		{"光子合并器", 1},
	}},
	{"框架材料", 0, "制造台", 6, 1, []ResourceRequirement{
		{"碳纳米管", 4},
		{"钛合金", 1},
		{"高纯硅块", 1},
	}},
	{"戴森球组件", 0, "制造台", 8, 1, []ResourceRequirement{
		{"框架材料", 3},
		{"太阳帆", 3},
		{"处理器", 3},
	}},
	{"小型运输火箭",0,  "制造台", 6, 1, []ResourceRequirement{
		{"戴森球组件", 2},
		{"氘核燃料棒", 4},
		{"量子芯片", 2},
	}},

	//==============================

	{"电磁矩阵", 0, "研究所", 3, 1, []ResourceRequirement{
		{"磁线圈", 1},
		{"电路板", 1},
	}},
	{"能量矩阵", 0, "研究所", 6, 1, []ResourceRequirement{
		{"高能石墨", 2},
		{"氢", 2},
	}},
	{"结构矩阵", 0, "研究所", 8, 1, []ResourceRequirement{
		{"金刚石", 1},
		{"钛晶石", 1},
	}},
	{"信息矩阵", 0, "研究所", 10, 1, []ResourceRequirement{
		{"处理器", 2},
		{"粒子宽带", 1},
	}},
	{"引力矩阵", 0, "研究所", 24, 2, []ResourceRequirement{
		{"引力透镜", 1},
		{"量子芯片", 1},
	}},
	{"宇宙矩阵", 0, "研究所", 15, 1, []ResourceRequirement{
		{"电磁矩阵", 1},
		{"能量矩阵", 1},
		{"结构矩阵", 1},
		{"信息矩阵", 1},
		{"引力矩阵", 1},
		{"反物质", 1},
	}},

	//==============================

	{"行星内物流运输站", 0, "制造台", 20, 1, []ResourceRequirement{
		{"钢材", 40},
		{"钛块", 40},
		{"处理器", 40},
		{"粒子容器", 20},
	}},
	{"星际物流运输站", 0, "制造台", 30, 1, []ResourceRequirement{
		{"行星内物流运输站", 1},
		{"钛合金", 40},
		{"粒子容器", 20},
	}},
	{"推进器", 0, "制造台", 4, 1, []ResourceRequirement{
		{"钢材", 2},
		{"铜块", 3},
	}},
	{"物流运输机", 0, "制造台", 4, 1, []ResourceRequirement{
		{"铁块", 5},
		{"处理器", 2},
		{"推进器", 2},
	}},
	{"加力推进器", 0, "制造台", 6, 1, []ResourceRequirement{
		{"钛合金", 5},
		{"电磁涡轮", 5},
	}},
	{"星际物流运输船", 0, "制造台", 6, 1, []ResourceRequirement{
		{"钛合金", 10},
		{"处理器", 10},
		{"加力推进器", 2},
	}},

}

var mappingResources = func() map[string]*Resource {
	m := map[string]*Resource{}
	for _, r := range AllResources {
		m[r.Key()] = r
	}
	return m
}()

var resourceOrder = func() map[string]int {
	m := map[string]int{}
	for i, r := range AllResources {
	m[r.Key()] = i
	}
	return m
}()

var facilityOrder = map[string]int{
	"采矿机":    1,
	"原油萃取机":  2,
	"抽水机":    3,
	"行星采集器":  4,
	"电弧熔炉":   5,
	"原油精炼厂":  6,
	"化工厂":    7,
	"分馏器":    8,
	"制造台": 9,
	"粒子对撞机":  10,
	"微波接收器":  11,
	"研究所": 12,
}

type ResourceRate struct {
	ResourceName string
	Rate         float64
}

func (ctx *ManufactureParameters) getResource(resName string) *Resource {
	if ctx.ImportingResources[resName] {
		return &Resource{
			Name: resName,
			Formula: -1,
			Facility: "外部输入",
			Duration: 1,
			Num: 1,
		}
	}
	resKey := fmt.Sprintf("%s-%d", resName, ctx.Formula[resName])
	return mappingResources[resKey]
}

//递归计算所有资源公式
//返回 资源名称 - 制造设备数量
func (ctx *ManufactureParameters) calculateRequirement(resRate ResourceRate) map[string]float64 {
	result := map[string]float64{}

	res := ctx.getResource(resRate.ResourceName)
	if res == nil {
		panic(fmt.Errorf("配方 %s 不存在", resRate.ResourceName))
	}
	for _, resReq := range res.Requirements {
		subResReq := ctx.calculateRequirement(ResourceRate{
			ResourceName: resReq.ResourceName,
			Rate:         resRate.Rate / res.Num * resReq.Num,
		})
		for subRes, subNum := range subResReq {
			result[subRes] += subNum
		}
	}
	facilityRate, rateExists := ctx.FacilityRate[res.Facility]
	if !rateExists {
		facilityRate = 1
	}
	result[resRate.ResourceName] += resRate.Rate * res.Duration / res.Num / facilityRate
	return result
}

type FacilityRequirement struct {
	Resource   string
	Facility   string
	Num        float64
	Inputs     []ResourceRate
	OutputRate float64
}

func (ctx *ManufactureParameters) FormatReq(fr *FacilityRequirement) string {
	facilityRate, exists := ctx.FacilityRate[fr.Facility]
	if !exists {
		facilityRate = 1
	}

	s := fmt.Sprintf("%s %.2f/s <- (%s:%.2f) <- (", fr.Resource, fr.OutputRate*facilityRate, fr.Facility, fr.Num)
	for i, input := range fr.Inputs {
		if i > 0 {
			s += " "
		}
		s += fmt.Sprintf("%s: %.2f/s", input.ResourceName, input.Rate*facilityRate)
	}
	s += ")"
	return s
}

func (ctx *ManufactureParameters) ShowRequirement(reqs []ResourceRate) {
	resultMap := map[string]float64{}
	fmt.Println("====目标====")
	for _, req := range reqs {
		m := ctx.calculateRequirement(req)
		for k, v := range m {
			resultMap[k] += v
		}
		fmt.Printf("%s %.2f/s\n", req.ResourceName, req.Rate)
	}

	fmt.Println("====生产线====")
	result := make([]*FacilityRequirement, 0)
	facilitiesCount := map[string]float64{}
	for resourceName, facilityNum := range resultMap {
		resource := ctx.getResource(resourceName)
		fr := &FacilityRequirement{
			Resource:   resourceName,
			Facility:   resource.Facility,
			Num:        facilityNum,
			OutputRate: resource.Num / resource.Duration * facilityNum,
		}
		for _, subRes := range resource.Requirements {
			fr.Inputs = append(fr.Inputs, ResourceRate{
				ResourceName: subRes.ResourceName,
				Rate:         subRes.Num / resource.Duration * facilityNum,
			})
		}
		result = append(result, fr)
		facilitiesCount[fr.Facility] += facilityNum
	}

	sort.Slice(result, func(i, j int) bool {
		res1 := ctx.getResource(result[i].Resource)
		ro1 := resourceOrder[res1.Key()]

		res2 := ctx.getResource(result[j].Resource)
		ro2 := resourceOrder[res2.Key()]

		if ro1 != ro2 {
			return ro1 < ro2
		}
		return res1.Name < res2.Name
	})
	for _, item := range result {
		if item.Num != 0 {
			fmt.Println(ctx.FormatReq(item))
		}
	}

	fmt.Println("====设施====")
	for fac, num := range facilitiesCount {
		if num != 0 {
			fmt.Printf("%s: %.2f\n", fac, num)
		}
	}
}
