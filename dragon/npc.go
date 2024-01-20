package main

import "math/rand"

// NPC 的结构体
type NPC struct {
	Name       string
	basic      *basic
	Experience int
}

func (n *NPC) strengthen(attack, defense, life, experience float64) *NPC {
	return &NPC{
		Name: n.Name,
		basic: &basic{
			life:    int(float64(n.basic.life) * life),
			attack:  int(float64(n.basic.attack) * attack),
			defense: int(float64(n.basic.defense) * defense),
		},
		Experience: int(float64(n.Experience) * experience),
	}
}

var (
	smallMonster = &NPC{
		Name: "小妖",
		basic: &basic{
			life:    100,
			attack:  1,
			defense: 1,
		},
		Experience: 1,
	}

	middleMonster = &NPC{
		Name: "中妖",
		basic: &basic{
			life:    200,
			attack:  2,
			defense: 2,
		},
		Experience: 2,
	}

	bigMonster = &NPC{
		Name: "大妖",
		basic: &basic{
			life:    300,
			attack:  3,
			defense: 3,
		},
		Experience: 3,
	}

	nezha = &NPC{
		Name: "哪吒",
		basic: &basic{
			life:    500,
			attack:  5,
			defense: 5,
		},
		Experience: 5,
	}

	erlang = &NPC{
		Name: "二郎神",
		basic: &basic{
			life:    1000,
			attack:  10,
			defense: 10,
		},
		Experience: 10,
	}

	gourd = &NPC{
		Name: "葫芦娃",
		basic: &basic{
			life:    1000,
			attack:  10,
			defense: 10,
		},
		Experience: 10,
	}

	towerer = &NPC{
		Name: "托塔天王",
		basic: &basic{
			life:    2000,
			attack:  20,
			defense: 20,
		},
		Experience: 20,
	}

	bullMonster = &NPC{
		Name: "牛魔王",
		basic: &basic{
			life:    5000,
			attack:  50,
			defense: 50,
		},
		Experience: 50,
	}

	superErlang = erlang.strengthen(
		10, 10, 10, 10,
	)

	nick = &NPC{
		Name: "狐尼克",
		basic: &basic{
			life:    20000,
			attack:  200,
			defense: 200,
		},
		Experience: 200,
	}

	judy = &NPC{
		Name: "朱迪警官",
		basic: &basic{
			life:    50000,
			attack:  500,
			defense: 500,
		},
		Experience: 500,
	}

	lionKing = &NPC{
		Name: "狮子王",
		basic: &basic{
			life:    100000,
			attack:  1000,
			defense: 1000,
		},
		Experience: 1000,
	}

	athena = &NPC{
		Name: "雅典娜",
		basic: &basic{
			life:    500000,
			attack:  1000,
			defense: 1000,
		},
		Experience: 10000,
	}

	superAthena = athena.strengthen(
		10, 10, 1, 10,
	)

	zeus = &NPC{
		Name: "宙斯",
		basic: &basic{
			life:    5000000,
			attack:  5000,
			defense: 5000,
		},
		Experience: 50000,
	}

	superZeus = zeus.strengthen(
		2, 2, 1, 2,
	)
)

type NPCList []*NPC
type NPCIndex struct {
	index map[int]NPCList
}

func newNPCIndex() *NPCIndex {
	return &NPCIndex{
		index: make(map[int]NPCList),
	}
}

func (n *NPCIndex) add(stage int, npc *NPC, num int) {
	if _, ok := n.index[stage]; !ok {
		n.index[stage] = make(NPCList, 0)
	}

	for i := 0; i < num; i++ {
		n.index[stage] = append(n.index[stage], npc)
	}
}

func (n *NPCIndex) get(stage int) *NPC {
	if stage > StageMax {
		stage = StageMax
	}

	return n.index[stage][rand.Intn(len(n.index[stage]))]
}

var NPCs = newNPCIndex()

// 初始化NPC列表，每一级别可以通过添加的NPC的数量控制遇到NPC的频率
func initNPCs() {
	NPCs.add(0, smallMonster, 1)

	NPCs.add(1, smallMonster, 10)

	NPCs.add(2, smallMonster, 10)
	NPCs.add(2, middleMonster, 10)

	NPCs.add(3, smallMonster, 10)
	NPCs.add(3, middleMonster, 10)
	NPCs.add(3, bigMonster, 10)

	NPCs.add(4, smallMonster, 20)
	NPCs.add(4, middleMonster, 20)
	NPCs.add(4, bigMonster, 20)

	NPCs.add(5, smallMonster, 20)
	NPCs.add(5, middleMonster, 20)
	NPCs.add(5, bigMonster, 20)
	NPCs.add(5, nezha, 2)
	NPCs.add(5, erlang, 2)

	NPCs.add(6, smallMonster, 20)
	NPCs.add(6, middleMonster, 20)
	NPCs.add(6, bigMonster, 20)
	NPCs.add(6, nezha, 2)
	NPCs.add(6, erlang, 2)

	NPCs.add(7, smallMonster, 20)
	NPCs.add(7, middleMonster, 20)
	NPCs.add(7, bigMonster, 20)
	NPCs.add(7, nezha, 2)
	NPCs.add(7, erlang, 2)
	NPCs.add(7, gourd, 2)
	NPCs.add(7, towerer, 2)
	NPCs.add(7, bullMonster, 2)

	NPCs.add(8, smallMonster, 20)
	NPCs.add(8, middleMonster, 20)
	NPCs.add(8, bigMonster, 20)
	NPCs.add(8, nezha, 2)
	NPCs.add(8, erlang, 2)
	NPCs.add(8, gourd, 2)
	NPCs.add(8, towerer, 2)
	NPCs.add(8, bullMonster, 2)

	NPCs.add(9, smallMonster, 20)
	NPCs.add(9, middleMonster, 20)
	NPCs.add(9, bigMonster, 20)
	NPCs.add(9, nezha, 2)
	NPCs.add(9, superErlang, 2)
	NPCs.add(9, gourd, 2)
	NPCs.add(9, towerer, 2)
	NPCs.add(9, bullMonster, 2)

	NPCs.add(10, smallMonster, 20)
	NPCs.add(10, middleMonster, 20)
	NPCs.add(10, bigMonster, 20)
	NPCs.add(10, nezha, 2)
	NPCs.add(10, superErlang, 2)
	NPCs.add(10, gourd, 2)
	NPCs.add(10, towerer, 2)
	NPCs.add(10, bullMonster, 2)
	NPCs.add(10, superErlang, 2)

	NPCs.add(11, smallMonster, 20)
	NPCs.add(11, middleMonster, 20)
	NPCs.add(11, bigMonster, 20)
	NPCs.add(11, nezha, 2)
	NPCs.add(11, superErlang, 2)
	NPCs.add(11, gourd, 2)
	NPCs.add(11, towerer, 2)
	NPCs.add(11, bullMonster, 2)
	NPCs.add(11, nick, 2)
	NPCs.add(11, judy, 2)

	NPCs.add(12, smallMonster, 20)
	NPCs.add(12, middleMonster, 20)
	NPCs.add(12, bigMonster, 20)
	NPCs.add(12, nezha, 2)
	NPCs.add(12, superErlang, 2)
	NPCs.add(12, gourd, 2)
	NPCs.add(12, towerer, 2)
	NPCs.add(12, bullMonster, 2)
	NPCs.add(12, nick, 2)
	NPCs.add(12, judy, 2)

	NPCs.add(13, smallMonster, 20)
	NPCs.add(13, middleMonster, 20)
	NPCs.add(13, bigMonster, 20)
	NPCs.add(13, nezha, 2)
	NPCs.add(13, superErlang, 2)
	NPCs.add(13, gourd, 2)
	NPCs.add(13, towerer, 2)
	NPCs.add(13, bullMonster, 2)
	NPCs.add(13, nick, 2)
	NPCs.add(13, judy, 2)
	NPCs.add(13, lionKing, 2)

	NPCs.add(14, smallMonster, 20)
	NPCs.add(14, middleMonster, 20)
	NPCs.add(14, bigMonster, 20)
	NPCs.add(14, nezha, 2)
	NPCs.add(14, superErlang, 2)
	NPCs.add(14, gourd, 2)

	NPCs.add(15, athena, 2)

	NPCs.add(16, superAthena, 2)

	NPCs.add(17, zeus, 2)

	NPCs.add(18, superZeus, 2)
}
