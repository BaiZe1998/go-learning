package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	// HealingRate 生命回复比例
	HealingRate = 0.1
	// AdvanceThreshold 修为阶段的底数（2次方龙）
	AdvanceThreshold = 2
	// SuccessRate 修为进阶成功几率
	SuccessRate = 0.75
	// EventChance 遇到事件几率
	EventChance = 0.2
	// NPCChance 遇到NPC几率
	NPCChance = 0.6
)

var (
	// NPCs NPC库
	NPCs = map[int][]NPC{
		0: []NPC{
			NPC{"小妖", 10, 1, 1, 1},
		},
		1: []NPC{
			NPC{"小妖", 100, 1, 1, 1},
		},
		2: []NPC{
			NPC{"小妖", 100, 1, 1, 1},
		},
		3: []NPC{
			NPC{"小妖", 100, 1, 1, 1},
			NPC{"中妖", 200, 2, 2, 2},
			NPC{"大妖", 300, 3, 3, 3},
		},
		4: []NPC{
			NPC{"小妖", 100, 1, 1, 1},
			NPC{"中妖", 200, 2, 2, 2},
			NPC{"大妖", 300, 3, 3, 3},
		},
		5: []NPC{
			NPC{"小妖", 100, 1, 1, 1},
			NPC{"中妖", 200, 2, 2, 2},
			NPC{"大妖", 300, 3, 3, 3},
			NPC{"哪吒", 500, 5, 5, 5},
			NPC{"二郎神", 1000, 10, 10, 10},
		},
		6: []NPC{
			NPC{"小妖", 100, 1, 1, 1},
			NPC{"中妖", 200, 2, 2, 2},
			NPC{"大妖", 300, 3, 3, 3},
			NPC{"哪吒", 500, 5, 5, 5},
			NPC{"二郎神", 1000, 10, 10, 10},
		},
		7: []NPC{
			NPC{"小妖", 100, 1, 1, 1},
			NPC{"中妖", 200, 2, 2, 2},
			NPC{"大妖", 300, 3, 3, 3},
			NPC{"哪吒", 500, 5, 5, 5},
			NPC{"葫芦娃", 1000, 10, 10, 10},
			NPC{"托塔天王", 2000, 20, 20, 20},
			NPC{"牛魔王", 5000, 50, 50, 50},
		},
		8: []NPC{
			NPC{"小妖", 100, 1, 1, 1},
			NPC{"中妖", 200, 2, 2, 2},
			NPC{"大妖", 300, 3, 3, 3},
			NPC{"哪吒", 500, 5, 5, 5},
			NPC{"葫芦娃", 1000, 10, 10, 10},
			NPC{"托塔天王", 2000, 20, 20, 20},
			NPC{"牛魔王", 5000, 50, 50, 50},
		},
		9: []NPC{
			NPC{"小妖", 100, 1, 1, 1},
			NPC{"中妖", 200, 2, 2, 2},
			NPC{"大妖", 300, 3, 3, 3},
			NPC{"哪吒", 500, 5, 5, 5},
			NPC{"葫芦娃", 1000, 10, 10, 10},
			NPC{"托塔天王", 2000, 20, 20, 20},
			NPC{"牛魔王", 5000, 50, 50, 50},
			NPC{"二郎神", 10000, 100, 100, 100},
		},
		10: []NPC{
			NPC{"小妖", 100, 1, 1, 1},
			NPC{"中妖", 200, 2, 2, 2},
			NPC{"大妖", 300, 3, 3, 3},
			NPC{"哪吒", 500, 5, 5, 5},
			NPC{"葫芦娃", 1000, 10, 10, 10},
			NPC{"托塔天王", 2000, 20, 20, 20},
			NPC{"牛魔王", 5000, 50, 50, 50},
			NPC{"二郎神", 10000, 100, 100, 100},
		},
		11: []NPC{
			NPC{"小妖", 100, 1, 1, 1},
			NPC{"中妖", 200, 2, 2, 2},
			NPC{"大妖", 300, 3, 3, 3},
			NPC{"哪吒", 500, 5, 5, 5},
			NPC{"葫芦娃", 1000, 10, 10, 10},
			NPC{"托塔天王", 2000, 20, 20, 20},
			NPC{"牛魔王", 5000, 50, 50, 50},
			NPC{"二郎神", 10000, 100, 100, 100},
			NPC{"狐尼克", 20000, 200, 200, 200},
			NPC{"朱迪警官", 50000, 500, 500, 500},
		},
		12: []NPC{
			NPC{"小妖", 100, 1, 1, 1},
			NPC{"中妖", 200, 2, 2, 2},
			NPC{"大妖", 300, 3, 3, 3},
			NPC{"哪吒", 500, 5, 5, 5},
			NPC{"葫芦娃", 1000, 10, 10, 10},
			NPC{"托塔天王", 2000, 20, 20, 20},
			NPC{"牛魔王", 5000, 50, 50, 50},
			NPC{"二郎神", 10000, 100, 100, 100},
			NPC{"狐尼克", 20000, 200, 200, 200},
			NPC{"朱迪警官", 50000, 500, 500, 500},
		},
		13: []NPC{
			NPC{"小妖", 100, 1, 1, 1},
			NPC{"中妖", 200, 2, 2, 2},
			NPC{"大妖", 300, 3, 3, 3},
			NPC{"哪吒", 500, 5, 5, 5},
			NPC{"葫芦娃", 1000, 10, 10, 10},
			NPC{"托塔天王", 2000, 20, 20, 20},
			NPC{"牛魔王", 5000, 50, 50, 50},
			NPC{"二郎神", 10000, 100, 100, 100},
			NPC{"狐尼克", 20000, 200, 200, 200},
			NPC{"朱迪警官", 50000, 500, 500, 500},
			NPC{"狮子王", 100000, 1000, 1000, 1000},
		},
		14: []NPC{
			NPC{"小妖", 100, 1, 1, 1},
			NPC{"中妖", 200, 2, 2, 2},
			NPC{"大妖", 300, 3, 3, 3},
			NPC{"哪吒", 500, 5, 5, 5},
			NPC{"葫芦娃", 1000, 10, 10, 10},
			NPC{"托塔天王", 2000, 20, 20, 20},
			NPC{"牛魔王", 5000, 50, 50, 50},
			NPC{"二郎神", 10000, 100, 100, 100},
			NPC{"狐尼克", 20000, 200, 200, 200},
			NPC{"朱迪警官", 50000, 500, 500, 500},
			NPC{"狮子王", 100000, 1000, 1000, 1000},
		},
		15: []NPC{
			NPC{"雅典娜", 500000, 1000, 1000, 10000},
		},
		16: []NPC{
			NPC{"雅典娜", 500000, 2000, 2000, 20000},
		},
		17: []NPC{
			NPC{"宙斯", 5000000, 5000, 5000, 50000},
		},
		18: []NPC{
			NPC{"宙斯", 5000000, 10000, 10000, 100000},
		},
	}
	// Events 事件库
	Events = map[int][]Event{
		0: []Event{
			Event{"从龙穴中出生，获得了一本《龙族秘典》", 10, 10, 10, 5},
		},
		1: []Event{
			Event{"从龙穴中出生，获得了一本《龙族秘典》", 10, 10, 10, 5},
		},
		2: []Event{
			Event{"从龙穴中出生，获得了一本《龙族秘典》", 0, 0, 0, 100},
			Event{"遇到了一位古老的龙族长者，传授古老龙法", 0, 0, 0, 100},
			Event{"迷失在龙穴中，生命值减少，攻击力提升", -5, 5, 0, 20},
			Event{"发现龙草，恢复一些生命值", 10, 0, 0, 10},
		},
		3: []Event{
			Event{"从龙穴中出生，获得了一本《龙族秘典》", 0, 0, 0, 100},
			Event{"遇到了一位古老的龙族长者，传授古老龙法", 0, 0, 0, 100},
			Event{"迷失在龙穴中，生命值减少，攻击力提升", -5, 5, 0, 20},
			Event{"发现龙草，恢复一些生命值", 10, 0, 0, 10},
			Event{"遭遇龙兽，生命值减少，经验值提升", -10, 0, 0, 30},
			Event{"发现宝箱，获得一把龙之利爪", 0, 15, 0, 10},
		},
		4: []Event{
			Event{"从龙穴中出生，获得了一本《龙族秘典》", 0, 0, 0, 100},
			Event{"遇到了一位古老的龙族长者，传授古老龙法", 0, 0, 0, 100},
			Event{"迷失在龙穴中，生命值减少，攻击力提升", -5, 5, 0, 20},
			Event{"发现龙草，恢复一些生命值", 10, 0, 0, 10},
			Event{"遭遇龙兽，生命值减少，经验值提升", -10, 0, 0, 30},
			Event{"发现宝箱，获得一把龙之利爪", 0, 15, 0, 10},
			Event{"遇到了神秘的三体生物，与之交流，经验值大幅提升", -20, 10, -5, 80},
		},
		5: []Event{
			Event{"从龙穴中出生，获得了一本《龙族秘典》", 0, 0, 0, 100},
			Event{"遇到了一位古老的龙族长者，传授古老龙法", 0, 0, 0, 100},
			Event{"迷失在龙穴中，生命值减少，攻击力提升", -5, 5, 0, 20},
			Event{"发现龙草，恢复一些生命值", 10, 0, 0, 10},
			Event{"遭遇龙兽，生命值减少，经验值提升", -10, 0, 0, 30},
			Event{"发现宝箱，获得一把龙之利爪", 0, 15, 0, 10},
		},
		6: []Event{
			Event{"龙眼泪降临，龙族秘法激发，获得潜在能力", 5, 0, 0, 10},
			Event{"穿越时空之门，与古代龙族互动，经验值大幅提升", -15, 10, -5, 70},
			Event{"遇到了修仙者，学习修仙之术，修为大幅提升", 0, 0, 0, 120},
			Event{"在龙潭中修炼，生命值大幅提升，攻击力、防御力提升", 20, 15, 15, 40},
			Event{"发现神秘遗迹，获得未知科技的力量", 0, 20, 10, 10},
		},
		7: []Event{
			Event{"龙眼泪降临，龙族秘法激发，获得潜在能力", 5, 0, 0, 10},
			Event{"穿越时空之门，与古代龙族互动，经验值大幅提升", -15, 10, -5, 70},
			Event{"遇到了修仙者，学习修仙之术，修为大幅提升", 0, 0, 0, 120},
			Event{"在龙潭中修炼，生命值大幅提升，攻击力、防御力提升", 20, 15, 15, 40},
			Event{"发现神秘遗迹，获得未知科技的力量", 0, 20, 10, 10},
			Event{"遇到了三体文明的代表，进行文明交流", -10, 10, -5, 80},
			Event{"龙吟之音传遍四海，引来无数追随者", 0, 10, 0, 30},
		},
		8: []Event{
			Event{"龙眼泪降临，龙族秘法激发，获得潜在能力", 5, 0, 0, 10},
			Event{"穿越时空之门，与古代龙族互动，经验值大幅提升", -15, 10, -5, 70},
			Event{"遇到了修仙者，学习修仙之术，修为大幅提升", 0, 0, 0, 120},
			Event{"在龙潭中修炼，生命值大幅提升，攻击力、防御力提升", 20, 15, 15, 40},
			Event{"发现神秘遗迹，获得未知科技的力量", 0, 20, 10, 10},
			Event{"遇到了三体文明的代表，进行文明交流", -10, 10, -5, 80},
			Event{"龙吟之音传遍四海，引来无数追随者", 0, 10, 0, 30},
			Event{"遭遇时空风暴，生命值减少，但获得未知力量", -10, 0, 0, 50},
		},
		9: []Event{
			Event{"龙眼泪降临，龙族秘法激发，获得潜在能力", 5, 0, 0, 10},
			Event{"穿越时空之门，与古代龙族互动，经验值大幅提升", -15, 10, -5, 70},
			Event{"遇到了修仙者，学习修仙之术，修为大幅提升", 0, 0, 0, 120},
			Event{"在龙潭中修炼，生命值大幅提升，攻击力、防御力提升", 20, 15, 15, 40},
			Event{"发现神秘遗迹，获得未知科技的力量", 0, 20, 10, 10},
			Event{"遇到了三体文明的代表，进行文明交流", -10, 10, -5, 80},
			Event{"龙吟之音传遍四海，引来无数追随者", 0, 10, 0, 30},
			Event{"遭遇时空风暴，生命值减少，但获得未知力量", -10, 0, 0, 50},
			Event{"掌握了时空穿梭的能力，增加了冒险机会", 0, 0, 0, 10},
		},
		10: []Event{
			Event{"龙眼泪降临，龙族秘法激发，获得潜在能力", 5, 0, 0, 10},
			Event{"穿越时空之门，与古代龙族互动，经验值大幅提升", -15, 10, -5, 70},
			Event{"遇到了修仙者，学习修仙之术，修为大幅提升", 0, 0, 0, 120},
			Event{"在龙潭中修炼，生命值大幅提升，攻击力、防御力提升", 20, 15, 15, 40},
			Event{"发现神秘遗迹，获得未知科技的力量", 0, 20, 10, 10},
			Event{"遇到了三体文明的代表，进行文明交流", -10, 10, -5, 80},
			Event{"龙吟之音传遍四海，引来无数追随者", 0, 10, 0, 30},
			Event{"遭遇时空风暴，生命值减少，但获得未知力量", -10, 0, 0, 50},
			Event{"掌握了时空穿梭的能力，增加了冒险机会", 0, 0, 0, 10},
			Event{"科技文明的代表寻求与你合作，经验值大幅提升", 0, 0, 0, 60},
		},
		11: []Event{
			Event{"龙眼泪降临，龙族秘法激发，获得潜在能力", 5, 0, 0, 10},
			Event{"穿越时空之门，与古代龙族互动，经验值大幅提升", -15, 10, -5, 70},
			Event{"遇到了修仙者，学习修仙之术，修为大幅提升", 0, 0, 0, 120},
			Event{"在龙潭中修炼，生命值大幅提升，攻击力、防御力提升", 20, 15, 15, 40},
			Event{"发现神秘遗迹，获得未知科技的力量", 0, 20, 10, 10},
			Event{"遇到了三体文明的代表，进行文明交流", -10, 10, -5, 80},
			Event{"龙吟之音传遍四海，引来无数追随者", 0, 10, 0, 30},
			Event{"遭遇时空风暴，生命值减少，但获得未知力量", -10, 0, 0, 50},
			Event{"掌握了时空穿梭的能力，增加了冒险机会", 0, 0, 0, 10},
			Event{"科技文明的代表寻求与你合作，经验值大幅提升", 0, 0, 0, 60},
			Event{"成功阻止一场时空灾难，名望大增", 0, 0, 0, 100},
		},
		12: []Event{
			Event{"在龙穴中发现古老的传送阵，传送到异次元空间", 0, 0, 0, 50},
			Event{"遇到修仙者，学到了龙族法术，攻击力大幅提升", 0, 20, 0, 20},
			Event{"触发龙族神迹，生命值、攻击力、防御力全部提升", 20, 20, 20, 30},
			Event{"遭遇宇宙异兽，与之激战，经验值大幅提升", -30, 15, -10, 100},
			Event{"踏入龙门，感悟龙之奥义，修为提升", 0, 0, 0, 50},
			Event{"遭遇时空扭曲，生命值降低，但修为提升", -15, 0, 0, 30},
			Event{"龙吟之夜，获得神秘力量，全属性提升", 30, 30, 30, 50},
		},
		13: []Event{
			Event{"在龙穴中发现古老的传送阵，传送到异次元空间", 0, 0, 0, 50},
			Event{"遇到修仙者，学到了龙族法术，攻击力大幅提升", 0, 20, 0, 20},
			Event{"触发龙族神迹，生命值、攻击力、防御力全部提升", 20, 20, 20, 30},
			Event{"遭遇宇宙异兽，与之激战，经验值大幅提升", -30, 15, -10, 100},
			Event{"踏入龙门，感悟龙之奥义，修为提升", 0, 0, 0, 50},
			Event{"遭遇时空扭曲，生命值降低，但修为提升", -15, 0, 0, 30},
			Event{"龙吟之夜，获得神秘力量，全属性提升", 30, 30, 30, 50},
			Event{"遭遇龙魂幻影，与之对话，修为大幅提升", 0, 0, 0, 80},
		},
		14: []Event{
			Event{"在龙穴中发现古老的传送阵，传送到异次元空间", 0, 0, 0, 50},
			Event{"遇到修仙者，学到了龙族法术，攻击力大幅提升", 0, 20, 0, 20},
			Event{"触发龙族神迹，生命值、攻击力、防御力全部提升", 20, 20, 20, 30},
			Event{"遭遇宇宙异兽，与之激战，经验值大幅提升", -30, 15, -10, 100},
			Event{"踏入龙门，感悟龙之奥义，修为提升", -10, 0, 0, 50},
			Event{"遭遇时空扭曲，生命值降低，但修为提升", -15, 0, 0, 30},
			Event{"龙吟之夜，获得神秘力量，全属性提升", 30, 30, 30, 50},
			Event{"遭遇龙魂幻影，与之对话，修为大幅提升", -10, 0, 0, 80},
			Event{"修炼龙舞之术，全属性提升，生命值大幅回复", 50, 50, 50, 50},
		},
		15: []Event{
			Event{"遭遇龙魂幻影，与之对话，修为大幅提升", -100, 0, 0, 800},
			Event{"修炼龙舞之术，全属性提升，生命值大幅回复", 500, 500, 500, 500},
		},
		16: []Event{
			Event{"遭遇龙魂幻影，与之对话，修为大幅提升", -200, 0, 0, 800},
			Event{"修炼龙舞之术，全属性提升，生命值大幅回复", 500, 500, 500, 500},
		},
		17: []Event{
			Event{"遭遇龙魂幻影，与之对话，修为大幅提升", -500, 0, 0, 8000},
			Event{"修炼龙舞之术，全属性提升，生命值大幅回复", 5000, 5000, 5000, 5000},
		},
		18: []Event{
			Event{"遭遇龙魂幻影，与之对话，修为大幅提升", -1000, 0, 0, 8000},
			Event{"修炼龙舞之术，全属性提升，生命值大幅回复", 5000, 5000, 5000, 5000},
		},
	}
)

// Dragon 龙的结构体
type Dragon struct {
	Name            string
	Life            int
	MaxLife         int
	Attack          int
	Defense         int
	Experience      int
	ExperienceStage int
	Remaining       int
	MaxRemaining    int
}

// Fight 与 NPC 战斗
func (d *Dragon) Fight(n *NPC) {
	for d.Life > 0 && n.Life > 0 {
		if n.Attack > d.Defense {
			d.Life -= n.Attack - d.Defense
		}
		if d.Attack > n.Defense {
			n.Life -= d.Attack - n.Defense
		}
	}
	if d.Life <= 0 {
		fmt.Printf("你被%s打败了\n", n.Name)
		d.Life = 0
		appendExperience(d, -d.Experience/2)
		randomDecrease(d)
	} else {
		fmt.Printf("你打败了%s\n", n.Name)
		appendExperience(d, n.Experience)
	}
}

// Process 处理偶发事件
func (d *Dragon) Process(e *Event) {
	fmt.Println(e.Name)
	d.Attack += e.Attack
	d.Defense += e.Defense
	appendLife(d, e.Life)
	appendExperience(d, e.Experience)
}

// NPC 的结构体
type NPC struct {
	Name       string
	Life       int
	Attack     int
	Defense    int
	Experience int
}

// Event 的结构体
type Event struct {
	Name       string
	Life       int
	Attack     int
	Defense    int
	Experience int
}

// 增加经验
func appendExperience(dragon *Dragon, value int) {
	tmp := dragon.Experience
	dragon.Experience += value
	if dragon.Experience >= int(math.Pow(AdvanceThreshold, float64(dragon.ExperienceStage))) {
		result := handleAdvance(dragon)
		switch result {
		case 0:
			dragon.Experience = int(math.Pow(AdvanceThreshold, float64(dragon.ExperienceStage)))
			dragon.ExperienceStage++
			fmt.Printf("恭喜，修为增加了 %d，进阶为2的%d次方龙！\n", dragon.Experience-tmp, dragon.ExperienceStage-1)
		case 1:
			dragon.Experience = tmp / 2
			fmt.Printf("修为减半了！, 还剩余 %d\n", dragon.Experience)
		case 2:
			dragon.Experience = tmp
		}
	} else {
		if value < 0 {
			dragon.Experience += value
			if dragon.Experience < 0 {
				dragon.Experience = 0
			}
			fmt.Printf("修为减少了 %d\n", value)
		} else {
			fmt.Printf("修为增加了 %d\n", value)
		}
	}
}

// 增加生命
func appendLife(dragon *Dragon, value int) {
	dragon.Life += value
	if dragon.Life > dragon.MaxLife {
		dragon.Life = dragon.MaxLife
	} else if dragon.Life < 0 {
		dragon.Life = 0
	}
}

// 修为进阶
func handleAdvance(dragon *Dragon) int {
	fmt.Println("修为达到了瓶颈，是否进阶？(y/n)")
	var choice string
	if dragon.ExperienceStage <= 5 {
		fmt.Printf("修为低于2的5次方龙，默认自动进阶")
		choice = "y"
	} else {
		fmt.Scanln(&choice)
	}

	if choice == "y" {
		if rand.Float64() <= SuccessRate {
			fmt.Println("恭喜，修为成功进阶！")
			randomIncrease(dragon)
			return 0
		} else {
			fmt.Println("很遗憾，修为进阶失败。")
			randomDecrease(dragon)
			return 1
		}
	}
	fmt.Println("你选择了放弃进阶。")
	return 2
}

// 是否游戏结束
func isGameOver(dragon *Dragon) bool {
	return dragon.Remaining <= 0
}

// 创建龙
func createDragon() Dragon {
	var dragon Dragon

	fmt.Print("请输入龙的名称: ")
	fmt.Scanln(&dragon.Name)

	for {
		fmt.Print("分配生命、攻击力、防御力的能力值（总和为100，以空格分隔）: ")
		fmt.Scanln(&dragon.Life, &dragon.Attack, &dragon.Defense)
		if dragon.Life+dragon.Attack+dragon.Defense == 100 {
			break
		} else {
			fmt.Println("总和不为100，请重新输入")
		}
	}

	fmt.Print("请输入初始寿命（轮）: ")
	fmt.Scanln(&dragon.MaxRemaining)

	dragon.Remaining = dragon.MaxRemaining
	dragon.MaxLife = dragon.Life

	return dragon
}

// 随机增加属性
func randomIncrease(dragon *Dragon) {
	stat := rand.Intn(3) // 0: Attack, 1: Defense, 2: Life
	switch stat {
	case 0:
		dragon.Attack *= 2
		fmt.Printf("攻击力翻倍了！, 现在是 %d\n", dragon.Attack)
	case 1:
		dragon.Defense *= 2
		fmt.Printf("防御力翻倍了！, 现在是 %d\n", dragon.Defense)
	case 2:
		dragon.MaxLife *= 2
		dragon.Life = dragon.MaxLife
		fmt.Printf("生命值翻倍了！, 现在是 %d\n", dragon.Life)
	}
}

// 随机减少属性
func randomDecrease(dragon *Dragon) {
	stat := rand.Intn(3) // 0: Attack, 1: Defense, 2: Life
	switch stat {
	case 0:
		dragon.Attack /= 2
		fmt.Printf("攻击力减半了！, 还剩余 %d\n", dragon.Attack)
	case 1:
		dragon.Defense /= 2
		fmt.Printf("防御力减半了！, 还剩余 %d\n", dragon.Defense)
	case 2:
		dragon.MaxLife /= 2
		dragon.Life /= 2
		fmt.Printf("生命值减半了！, 还剩余 %d\n", dragon.Life)
	}
}

// 修养生息
func toHeal(dragon *Dragon, turn int) {
	fmt.Printf("你开始修养生息，恢复生命值，并增长修为\n")
	for turn > 0 {
		turn--
		dragon.Remaining--
		fmt.Printf("修养中ing...\n剩余寿命 %d 轮\n", dragon.Remaining)
		appendLife(dragon, int(float64(dragon.MaxLife)*HealingRate)+1)
		appendExperience(dragon, int(float64(dragon.Experience)*HealingRate)+1)
		printStatus(dragon)

		if isGameOver(dragon) {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// 外出冒险
func toAdventure(dragon *Dragon, turn int) {
	fmt.Printf("你开始外出冒险，增长修为\n")
	for turn > 0 {
		if dragon.Life <= 0 {
			dragon.Remaining -= turn
			fmt.Printf("你已经死亡，无法继续冒险!丢失%d冒险回合，请按1修养生息！！！！！！！！\n", turn)
			break
		}
		turn--
		dragon.Remaining--
		fmt.Printf("\n剩余寿命 %d 轮\n", dragon.Remaining)
		rad := rand.Float64()
		if rad <= NPCChance {
			event := Events[dragon.ExperienceStage][rand.Intn(len(Events[dragon.ExperienceStage]))]
			dragon.Process(&event)
		} else if rad <= NPCChance+EventChance {
			npc := NPCs[dragon.ExperienceStage][rand.Intn(len(NPCs[dragon.ExperienceStage]))]
			dragon.Fight(&npc)
		} else {
			fmt.Println("你踏入了一片宁静的山林，潜心修炼")
			appendExperience(dragon, int(float64(dragon.Experience)*HealingRate)*2+1)
		}
		printStatus(dragon)

		if isGameOver(dragon) {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// 打印龙的各项属性
func printStatus(dragon *Dragon) {
	fmt.Printf("姓名：%s，修为：%d，称号：2的%d次方龙，攻击力：%d，防御力：%d，生命值：%d，剩余寿命：%d轮\n",
		dragon.Name, dragon.Experience, dragon.ExperienceStage-1, dragon.Attack, dragon.Defense, dragon.Life, dragon.Remaining)
}

// 游戏结束成就打印
func gameOver(dragon *Dragon) {
	fmt.Printf("你长达%d轮的一生真是波澜壮阔，你达成了以下成就：\n", dragon.MaxRemaining-dragon.Remaining)
	printStatus(dragon)
}

// 主函数
func main() {
	fmt.Printf("\033[H\033[2J") // Clear screen

	dragon := createDragon()

	for !isGameOver(&dragon) {
		//fmt.Printf("\033[H\033[2J") // Clear screen
		printStatus(&dragon)

		fmt.Println("请选择操作:")
		fmt.Println("1. 修养生息")
		fmt.Println("2. 外出冒险")
		fmt.Println("z. 结束游戏")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			fmt.Println("请输入修养的轮数: ")
			var turn int
			fmt.Scanln(&turn)
			toHeal(&dragon, turn)
		case "2":
			fmt.Println("请输入冒险的轮数: ")
			var turn int
			fmt.Scanln(&turn)
			toAdventure(&dragon, turn)
		case "z":
			fmt.Println("结束游戏")
			gameOver(&dragon)
			return
		default:
			fmt.Println("无效的选择，请重新输入")
		}
	}
	gameOver(&dragon)
}
