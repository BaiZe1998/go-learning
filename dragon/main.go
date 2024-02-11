package main

import (
	"fmt"
	"github.com/mum4k/termdash"
	"math"
	"math/rand"
	"strconv"
	"strings"
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
	// StageMax 世界等级上限
	StageMax = 18
	// ShowRankSize 展示的排行榜大小
	ShowRankSize = 10
)

var (
	p *printer

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
	originExperience := dragon.Experience
	dragon.Experience += value
	if dragon.Experience >= int(math.Pow(AdvanceThreshold, float64(dragon.ExperienceStage))) {
		result := handleAdvance(dragon)
		switch result {
		case 0:
			dragon.Experience = int(math.Pow(AdvanceThreshold, float64(dragon.ExperienceStage)))
			dragon.ExperienceStage++
			p.addHistory(newHistoryInfo(fmt.Sprintf("恭喜，修为增加了 %d，进阶为2的%d次方龙！\n", dragon.Experience-originExperience, dragon.ExperienceStage)))
		case 1:
			dragon.Experience = originExperience / 2
			p.addHistory(newHistoryInfo(fmt.Sprintf("修为减半了！, 还剩余 %d\n", dragon.Experience)))
		case 2:
			dragon.Experience = originExperience
		}
	} else {
		if value < 0 {
			if dragon.Experience < 0 {
				dragon.Experience = 0
			}
			p.addHistory(newHistoryInfo(fmt.Sprintf("先前修为 %d，本次修为 %d, 当前修为 %d\n", originExperience, value, dragon.Experience)))
		} else {
			p.addHistory(newHistoryInfo(fmt.Sprintf("修为增加了 %d\n", value)))
		}
	}
}

// 增加生命
func appendLife(dragon *Dragon, value int) {
	dragon.basic.life += value
	if dragon.basic.life > dragon.basic.maxLife {
		dragon.basic.life = dragon.basic.maxLife
	} else if dragon.basic.life < 0 {
		dragon.basic.life = 0
	}
}

// 修为进阶
func handleAdvance(dragon *Dragon) int {
	p.addHistoryLn(newHistoryInfo("\n修为达到了瓶颈，是否进阶？(y/n)"))
	var choice string
	if dragon.ExperienceStage <= 12 {
		p.addHistory(newHistoryInfo("修为低于2的12次方龙，默认自动进阶\n"))
		choice = "y"
	} else {
		p.addOperateHint("修为达到了瓶颈，是否进阶？(y/n)")
		choice = <-p.scanned
	}

	if choice == "y" {
		if rand.Float64() <= SuccessRate {
			p.addHistory(newHistoryInfo("恭喜，修为成功进阶！\n"))
			randomIncreaseState(dragon)
			return 0
		} else {
			p.addHistory(newHistoryInfo("很遗憾，修为进阶失败。\n"))
			randomDecreaseState(dragon)
			return 1
		}
	}
	p.addHistory(newHistoryInfo("你选择了放弃进阶。\n"))
	return 2
}

// 是否游戏结束
func isGameOver(dragon *Dragon) bool {
	return dragon.Remaining <= 0
}

// 创建龙
func createDragon() Dragon {
	dragon := Dragon{
		basic: &basic{},
	}

	p.addOperateHint("请在下方输入龙的名称：")
	name := <-p.scanned
	dragon.Name = name

	for {
		p.addOperateHint("分配生命、攻击力、防御力的能力值（总和为100，以空格分隔）: ")
		valueString := <-p.scanned
		values := strings.Split(valueString, " ")
		if len(values) != 3 {
			p.addOperateHint("输入格式错误，请重新输入")
			time.Sleep(1 * time.Second)
			continue
		}
		var valuesInt []int
		for i := 0; i < 3; i++ {
			i, _ := strconv.ParseInt(values[i], 10, 64)
			valuesInt = append(valuesInt, int(i))
		}
		if valuesInt[0]+valuesInt[1]+valuesInt[2] != 100 {
			p.addOperateHint("总和不为100，请重新输入")
			time.Sleep(1 * time.Second)
			continue
		}
		dragon.basic.life = valuesInt[0]
		dragon.basic.attack = valuesInt[1]
		dragon.basic.defense = valuesInt[2]
		break
	}

	p.addOperateHint("请输入初始寿命（轮）: ")
	remaining := <-p.scanned
	remainingInt, _ := strconv.ParseInt(remaining, 10, 64)

	dragon.Remaining = int(remainingInt)
	dragon.MaxRemaining = int(remainingInt)
	dragon.basic.maxLife = dragon.basic.life

	return dragon
}

// 随机增加属性
func randomIncreaseState(dragon *Dragon) {
	stat := rand.Intn(3) // 0: Attack, 1: Defense, 2: Life
	switch stat {
	case 0:
		dragon.basic.attack *= 2
		p.addHistory(newHistoryInfo(fmt.Sprintf("攻击力翻倍了！, 现在是 %d\n", dragon.basic.attack)))
	case 1:
		dragon.basic.defense *= 2
		p.addHistory(newHistoryInfo(fmt.Sprintf("防御力翻倍了！, 现在是 %d\n", dragon.basic.defense)))
	case 2:
		dragon.basic.maxLife *= 2
		dragon.basic.life = dragon.basic.maxLife
		p.addHistory(newHistoryInfo(fmt.Sprintf("生命值翻倍了！, 现在是 %d\n", dragon.basic.life)))
	}
}

// 随机减少属性
func randomDecreaseState(dragon *Dragon) {
	stat := rand.Intn(3) // 0: Attack, 1: Defense, 2: Life
	switch stat {
	case 0:
		dragon.basic.attack /= 2
		p.addHistory(newHistoryInfo(fmt.Sprintf("攻击力减半了！, 还剩余 %d\n", dragon.basic.attack)))
	case 1:
		dragon.basic.defense /= 2
		p.addHistory(newHistoryInfo(fmt.Sprintf("防御力减半了！, 还剩余 %d\n", dragon.basic.defense)))
	case 2:
		dragon.basic.maxLife /= 2
		dragon.basic.life /= 2
		p.addHistory(newHistoryInfo(fmt.Sprintf("生命值减半了！, 还剩余 %d\n", dragon.basic.life)))
	}
}

// 休养生息
func toHeal(dragon *Dragon, turn int) {
	p.addHistory(newHistoryInfo("你开始休养生息，恢复生命值，并增长修为\n"))
	for turn > 0 {
		p.flush()
		turn--
		dragon.Remaining--
		p.addHistory(newHistoryInfo(fmt.Sprintf("休养中ing...\n剩余寿命 %d 轮\n", dragon.Remaining)))
		appendLife(dragon, int(float64(dragon.basic.maxLife)*HealingRate)+1)
		appendExperience(dragon, int(float64(dragon.Experience)*HealingRate)+1)

		if isGameOver(dragon) {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// 外出冒险
func toAdventure(dragon *Dragon, turn int) {
	p.addHistory(newHistoryInfo("你开始外出冒险，增长修为\n"))
	for turn > 0 {
		p.flush()
		if dragon.basic.life <= 0 {
			dragon.Remaining -= turn
			p.addHistoryLn(newHistoryInfo(fmt.Sprintf("你已经死亡，无法继续冒险!丢失%d冒险回合，请按1休养生息！！！！！！！！", turn)))
			break
		}
		turn--
		dragon.Remaining--
		p.addHistory(newHistoryInfo(fmt.Sprintf("剩余寿命 %d 轮 ", dragon.Remaining)))
		rad := rand.Float64()
		if rad <= NPCChance {
			npc := NPCs.get(dragon.ExperienceStage)
			dragon.Fight(npc)
		} else if rad <= NPCChance+EventChance {
			// 当龙的修为大于等于阶段最大值时，触发阶段最大值的事件（系统制裁）
			if dragon.ExperienceStage > StageMax {
				event := Events[StageMax][rand.Intn(len(Events[StageMax]))]
				dragon.Process(&event)
			} else {
				event := Events[dragon.ExperienceStage][rand.Intn(len(Events[dragon.ExperienceStage]))]
				dragon.Process(&event)
			}
		} else {
			p.addHistoryLn(newHistoryInfo("你踏入了一片宁静的山林，潜心修炼"))
			appendExperience(dragon, int(float64(dragon.Experience)*HealingRate)*2+1)
		}

		if isGameOver(dragon) {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func getInputTurn() int {
	for {
		turn := <-p.scanned
		turnI64, _ := strconv.ParseInt(turn, 10, 64)
		if int(turnI64) <= p.dragon.Remaining {
			return int(turnI64)
		}
		p.addOperateHint("输入轮数超过剩余寿命，请重新输入")
	}
}

func init() {
	initNPCs()
}

// 主函数
func main() {
	fmt.Println("\033[H\033[2J") // Clear screen

	p = newPrinter()
	// need to be called after newPrinter(), this func depends on p
	go showRank(getRanks(), nil)

	go func() {
		if err := termdash.Run(p.ctx, p.terminal, p.container, termdash.KeyboardSubscriber(p.keyBinding)); err != nil {
			panic(err)
		}
	}()

	dragon := createDragon()
	p.setDragon(&dragon)

	for !isGameOver(&dragon) {
		p.flush()
		p.addOperateHint("请选择操作:\n1. 休养生息\n2. 外出冒险")

		choice := <-p.scanned

		switch choice {
		case "1":
			p.addOperateHint("请输入休养的轮数: ")
			toHeal(&dragon, getInputTurn())
		case "2":
			p.addOperateHint("请输入冒险的轮数: ")
			toAdventure(&dragon, getInputTurn())
		default:
			p.addOperateHint("无效的选择，请重新输入")
			time.Sleep(1 * time.Second)
		}
	}
	p.addHistoryLn(newHistoryInfo("龙生结束，您的一生真是波澜壮阔，不虚此行！按 CRTL + W 退出游戏"))

	rank := newRank(&dragon)
	rank.save()
	go showRank(getRanks(), rank)

	<-p.scanned
}
