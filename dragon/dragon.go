package main

import (
	"fmt"
)

// Dragon 龙的结构体
type Dragon struct {
	basic           *basic
	ID              int
	Name            string
	Experience      int
	ExperienceStage int
	Remaining       int
	MaxRemaining    int
}

func (d *Dragon) isAlive() bool {
	return d.basic.life > 0
}

func (d *Dragon) decreaseRemaining() {
	d.Remaining--
}

// Fight 与 NPC 战斗
func (d *Dragon) Fight(n *NPC) {
	// 根本打不过直接润
	if d.basic.attack <= n.basic.defense {
		p.addHistory(newHistoryInfo(fmt.Sprintf("绝无可能击败的敌人%s\n", n.Name)))
		deduct := d.basic.attacked(n.basic.attack)
		if d.basic.isAlive() {
			p.addHistory(newHistoryInfo(fmt.Sprintf("逃跑成功 耗费%d点血量\n", deduct)))
		} else {
			p.addHistory(newHistoryInfo("逃跑失败\n"))
		}
		return
	}

	for d.basic.isAlive() {
		n.basic.attacked(d.basic.attack)
		if !n.basic.isAlive() {
			p.addHistory(newHistoryInfo("你打败了"))
			p.addHistory(newHistoryInfo(n.Name, TextOptionUnderline))
			appendExperience(d, n.Experience)
			return
		}
		d.basic.attacked(n.basic.attack)
	}

	p.addHistory(newHistoryInfo("你被"))
	p.addHistory(newHistoryInfo(n.Name, TextOptionUnderline))
	p.addHistoryLn(newHistoryInfo("打败了"))
	appendExperience(d, -d.Experience/2)
	randomDecrease(d)
}

// Process 处理偶发事件
func (d *Dragon) Process(e *Event) {
	p.addHistoryLn(newHistoryInfo(e.Name))
	d.basic.attack += e.Attack
	d.basic.defense += e.Defense
	appendLife(d, e.Life)
	appendExperience(d, e.Experience)
}
