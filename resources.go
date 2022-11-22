package main

import (
	"fmt"
	"sort"
	"strings"
)

type SprayType int

var (
	SprayDisabled = SprayType(0)
	SpraySpeedUp  = SprayType(1)
	SprayBonus    = SprayType(2)
)

type SprayingClass struct {
	Resource    string
	SpeedUpRate float64
	BonusRate   float64
	PowerRate   float64
	Count       float64
}

type Formula struct {
	ResourceName string  //目标资源名称
	FormulaNum   int     //配方编号
	Facility     string  //生产设施
	Duration     float64 //生产时长
	OutputNum    float64 //每次产出数量
	CanSpeedUp   bool    //支持增产剂加速生产
	CanBonus     bool    //支持增产剂额外产出
	Materials    []ResourceNum
}

func (f *Formula) Key() string {
	return fmt.Sprintf("%s-%d", f.ResourceName, f.FormulaNum)
}

type Produce struct {
	*Formula
	OutputRes    string
	OutputRate   float64
	SprayType    SprayType
	SprayRate    float64
	Inputs       map[string]float64
	FacilityName string
	FacilityNum  float64
}

type ResourceNum struct {
	ResourceName string
	Num          float64 //每次需求数量
}

type ManufactureParameters struct {
	FacilityRate       map[string]float64
	FormulaNums        map[string]int
	//资源从外部运入，不计算其生产，计算其运输需要的物流塔数量
	//值是距离（光年）
	ImportingResources map[string]float64
	SprayingResources  map[string]SprayType
	SprayingClass      *SprayingClass
	//计算增产剂用量时，默认使用增产剂喷增产剂自身，
	//可以提高单个增产剂喷涂次数，减少增产剂消耗
	//设为true禁用这个计算
	NoRecurseSpraying bool
}

func (mp *ManufactureParameters) Clone() *ManufactureParameters {
	r := &ManufactureParameters{
		FacilityRate:       map[string]float64{},
		FormulaNums:        map[string]int{},
		ImportingResources: map[string]float64{},
		SprayingResources:  map[string]SprayType{},
	}
	for k, v := range mp.FacilityRate {
		r.FacilityRate[k] = v
	}
	for k, v := range mp.FormulaNums {
		r.FormulaNums[k] = v
	}
	for k, v := range mp.ImportingResources {
		r.ImportingResources[k] = v
	}
	for k, v := range mp.SprayingResources {
		r.SprayingResources[k] = v
	}
	r.SprayingClass = mp.SprayingClass
	return r
}

func (mp *ManufactureParameters) SetImportingResources(r map[string]float64) {
	for k, v := range r {
		mp.ImportingResources[k]=v
	}
}

var mappingResources = func() map[string]*Formula {
	m := map[string]*Formula{}
	for _, r := range AllFormulas {
		m[r.Key()] = r
	}
	return m
}()

var resourceOrder = func() map[string]int {
	m := map[string]int{}
	for i, r := range AllFormulas {
		m[r.Key()] = i
	}
	return m
}()

func (mp *ManufactureParameters) getFormula(resName string) *Formula {
	if distance,importing := mp.ImportingResources[resName]; importing {
		return &Formula{
			ResourceName: resName,
			FormulaNum:   -1,
			Facility:     "星际物流站",
			Duration:     (distance * 60 + 646) / 290.0,
			OutputNum:    2000,
		}
	}
	resKey := fmt.Sprintf("%s-%d", resName, mp.FormulaNums[resName])
	return mappingResources[resKey]
}

//迭代计算所有资源公式
func (mp *ManufactureParameters) calculateProduce(requirements map[string]float64) map[string]*Produce {
	producesMap := map[string]*Produce{}
	for changed := true; changed; {
		changed = false
		for resName, resRate := range requirements {
			if resRate < 0.000001 {
				continue
			}
			f := mp.getFormula(resName)

			//生产速度加成
			rateRatio := 1.0
			//产出加成
			outputRatio := 1.0

			//设施生产速度
			facilityRatio, ok := mp.FacilityRate[f.Facility]
			if !ok {
				facilityRatio = 1.0
			}

			//计算喷涂加速
			st := mp.getSprayType(f)
			switch st {
			case SprayDisabled:
			case SpraySpeedUp:
				rateRatio = mp.SprayingClass.SpeedUpRate
			case SprayBonus:
				outputRatio = mp.SprayingClass.BonusRate
			}

			//每秒生产（公式应用）次数
			produceRate := resRate / f.OutputNum / outputRatio
			//设施数量
			facilityNum := produceRate * f.Duration / facilityRatio / rateRatio

			materials := f.Materials
			var sprayRate float64
			if st != SprayDisabled {
				//计算喷涂开销
				materials = make([]ResourceNum, len(f.Materials), len(f.Materials)+1)
				copy(materials, f.Materials)
				sumSprayCount := 0.0
				for _, m := range f.Materials {
					sumSprayCount += m.Num
				}
				sprayingCoung := mp.SprayingClass.Count
				if !mp.NoRecurseSpraying {
					sprayingCoung = sprayingCoung * mp.SprayingClass.BonusRate - 1
				}
				sprayRate = sumSprayCount / sprayingCoung
				materials = append(materials, ResourceNum{
					ResourceName: mp.SprayingClass.Resource,
					Num:          sprayRate,
				})
			}

			//先从需求中减去这个资源
			delete(requirements, resName)

			//创建 Produce 生产详情项
			p, ok := producesMap[resName]
			if !ok {
				p = &Produce{
					Formula:     f,
					OutputRes:   resName,
					SprayType:   st,
					Inputs:      map[string]float64{},
				}
				producesMap[resName] = p
			}
			p.OutputRate += resRate
			p.FacilityNum += facilityNum
			p.SprayRate += sprayRate * produceRate
			for _, m := range f.Materials {
				p.Inputs[m.ResourceName] += m.Num * produceRate
			}
			for _, m := range materials {
				requirements[m.ResourceName] += m.Num * produceRate
			}
			changed = true
		}
	}
	return producesMap
}

func (mp *ManufactureParameters) getSprayType(f *Formula) SprayType {
	if mp.SprayingClass == nil {
		return SprayDisabled
	}

	t, ok := mp.SprayingResources[f.ResourceName]
	if !ok {
		if defaultIncreaseFacilities[f.Facility] {
			t = SprayBonus
		} else {
			t = SpraySpeedUp
		}
	}
	if !f.CanBonus && t == SprayBonus {
		t = SpraySpeedUp
	}
	if !f.CanSpeedUp && t == SpraySpeedUp {
		t = SprayDisabled
	}
	return t
}

func (mp *ManufactureParameters) ShowRequirement(reqs map[string]float64, outputTemplate string) {
	fmt.Println("====目标====")
	for res, rate := range reqs {
		fmt.Printf("%s %.3f/s\n", res, rate)
	}

	producesMap := mp.calculateProduce(reqs)
	produces := make([]*Produce, 0, len(producesMap))
	for _, p := range producesMap {
		produces = append(produces, p)
	}

	sort.Slice(produces, func(i, j int) bool {
		res1 := produces[i]
		ro1 := resourceOrder[res1.Key()]

		res2 := produces[j]
		ro2 := resourceOrder[res2.Key()]

		if ro1 != ro2 {
			return ro1 < ro2
		}
		return res1.ResourceName < res2.ResourceName
	})

	fmt.Println("====生产线====")
	outputTemplate = strings.Trim(outputTemplate, "\n")
	if outputTemplate != "" {
		for _, line := range strings.Split(outputTemplate, "\n") {
			if p, ok := producesMap[line]; ok {
				fmt.Println(mp.FormatReq(p))
				delete(producesMap, line)
			} else {
				fmt.Println(line)
			}
		}
		fmt.Println()
	}
	for _, p := range produces {
		if p, ok := producesMap[p.ResourceName]; ok {
			fmt.Println(mp.FormatReq(p))
		}
	}

	fmt.Println("====设施====")
	facilitiesCount := map[string]float64{}
	for _, p := range produces {
		facilitiesCount[p.Facility] += p.FacilityNum
	}
	for fac, num := range facilitiesCount {
		fmt.Printf("%s: %.3f\n", fac, num)
	}
}

func (mp *ManufactureParameters) FormatReq(p *Produce) string {
	s := fmt.Sprintf("%s %.3f/s <- (%s:%.3f)", p.OutputRes, p.OutputRate, p.Facility, p.FacilityNum)
	inputNames := make([]string, 0, len(p.Inputs))
	for inputName := range p.Inputs {
		inputNames = append(inputNames, inputName)
	}

	switch p.SprayType {
	case SpraySpeedUp:
		s += " <- 加速"
	case SprayBonus:
		s += " <- 增产"
	}
	if p.SprayType != SprayDisabled {
		s += fmt.Sprintf("(%s: %.3f/s)", mp.SprayingClass.Resource, p.SprayRate)
	}

	if len(p.Inputs) > 0 {
		sort.Slice(inputNames, func(i, j int) bool {
			order1 := resourceOrder[inputNames[i]]
			order2 := resourceOrder[inputNames[j]]
			return order1 < order2
		})
		s += " <- ( "
		for _, inputRes := range inputNames {
			inputRate := p.Inputs[inputRes]
			s += fmt.Sprintf("%s: %.2f/s", inputRes, inputRate)
			s += " "
		}
		s += ")"
	}
	return s
}
