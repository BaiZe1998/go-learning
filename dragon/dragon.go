package main

import "fmt"

// Dragon 龙的结构体
type Dragon struct {
	basic           *basic
	Name            string
	Experience      int
	ExperienceStage int
	Remaining       int
	MaxRemaining    int
}

func (d *Dragon) isAlive() bool {
	return d.basic.life > 0
}

// Fight 与 NPC 战斗
func (d *Dragon) Fight(n *NPC) {
	// 根本打不过直接润
	if d.basic.attack <= n.basic.defense {
		fmt.Printf("绝无可能击败的敌人%s\n", n.Name)
		deduct := d.basic.attacked(n.basic.attack)
		if d.basic.isAlive() {
			fmt.Printf("逃跑成功 耗费%d点血量\n", deduct)
		} else {
			fmt.Printf("逃跑失败\n")
		}
		return
	}

	for d.basic.isAlive() {
		n.basic.attacked(d.basic.attack)
		if !n.basic.isAlive() {
			fmt.Printf("你打败了%s\n", n.Name)
			appendExperience(d, n.Experience)
			return
		}
		d.basic.attacked(n.basic.attack)
	}

	fmt.Printf("你被%s打败了\n", n.Name)
	appendExperience(d, -d.Experience/2)
	randomDecrease(d)
}

// Process 处理偶发事件
func (d *Dragon) Process(e *Event) {
	fmt.Println(e.Name)
	d.basic.attack += e.Attack
	d.basic.defense += e.Defense
	appendLife(d, e.Life)
	appendExperience(d, e.Experience)
}
