[TOC]

## 一、前言

新年快乐，祝大家新的一年：🐲 龙行龘龘，🔥 前程朤朤！

白泽花了点时间，用 Go 写了一个控制台的小游戏：《模拟龙生》，在游戏中你将模拟一条新生的巨龙，开始无尽的冒险！

> Tips：运气很重要！不然会抓狂！还有游戏可能有些小 BUG，你那么帅，一定不好意思说我（欢迎 pr）。

游戏玩法图解：

![image-20240119100821996](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240119100821996.png)

> 🐲 游戏代码在我的开源 [Go 学习仓库](https://github.com/BaiZe1998/go-learning/tree/main/dragon)的`dragon`路径下，可以直接运行可执行程序体验。
>
> 仓库里还包含 Go 各阶段学习文章、读书笔记、电子书、简历模板等，欢迎 star。

![image-20240119102429937](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240119102429937.png)

## 二、游戏玩法

🐲 《模拟龙生》大部分游戏时间都是挂机冒险，不占用大家时间～

**游戏的目的是在指定轮次的游戏回合内，想方设法获得最高的修为！**

![dragon](dragon.gif)

### 1. 初始化

- 起一个你喜欢的名字！然后分配100点能力值到三维属性。
- 攻击力：当冒险遇到 NPC 时，会与其发生回合制战斗，每回合你对它可以造成：你的攻击力 - NPC 防御力的伤害。**（无法击穿对方护甲则会逃跑，但是必须承受一次对方的攻击）**
- 防御力：当冒险遇到 NPC 时，会与其发生回合制战斗，每回合它对你可以造成：NPC 的攻击力 - 你防御力的伤害。
- 生命值：战斗和触发事件造成生命值为0均无法继续冒险，必须返回休养生息（每秒恢复最大生命值10%）。
- 外出冒险有几率遭遇 NPC 发生战斗，遭遇随机事件获得效果（增/减），或者其他...
- 输入整个游戏的轮次：这里输入200轮。

> 🐲 这里也向大家发起挑战：200轮次内，挑战可以获取的最高修为值！欢迎加q群分享：622383022，或者评论区展示你的战绩！

![image-20240212124854835](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240212124854835.png)

### 2. 外出冒险

选择2则出发外出冒险（输入冒险轮次，200轮挑战就是限制这个），冒险有大概率遇到 NPC 与其战斗，小概率触发随机事件，战斗与事件都会对你的各项属性造成影响，**但是生命值为0会丢失此次冒险的剩余轮次**。

🐲 **每 0.5 秒进行一轮冒险，打印日志。**

![image-20240212125225636](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240212125225636.png)

🐲 **冒险终止条件**：轮次消耗完毕或者生命为0。

🐲 **丢失轮次**：如果你选择外出冒险50轮，但是在第25轮的时候，生命为0，则会丢失剩余25轮次，此时必须输入1返回修养生息（每个轮次回复 10% 生命值）。

🐲 **游戏结束**：所有轮次消耗完毕。

### 3. 修为进阶

```go
// 修为进阶表
1, 2, 4, 8, 16, 32, 64, 128...
// 当修为到达2^x的时候，会触发进阶判断，询问你是否进阶（y/n）
1. 进阶成功：随机一项能力*2
2. 进阶失败：随机一项能力/2，修为/2
3. 放弃进阶：只有进入下一阶才能继续增长修为（🌟但是主动放弃进阶会使自己随机一项属性值*1.25）
// 当前成为2的12次方龙之前，默认自动选择进阶
```

![image-20240212125609172](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240212125609172.png)

### 4. 修养生息

🐲 生命回复，修为每次恢复10%。

![image-20240212125710033](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240212125710033.png)

### 5. 游戏结算

🐲 当剩余最后45轮的时候，输入冒险45轮，但是在生命剩余23轮的时候，遇到不可击败的敌人（无法击穿护甲），尝试逃跑，但是逃跑失败（无法承受一次攻击），直接死亡，丢失剩余所有轮数。

最后进入排行榜第二名，排行榜根据经验值（修为）进行排序，取前十名。

![image-20240212125927688](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240212125927688.png)

## 三、游戏设计

### 3.1 游戏 UI 展示

[termdash](https://github.com/mum4k/termdash)：Termdash 是一款基于终端的跨平台定制仪表盘，《模拟龙生》这款游戏中，所有的终端页面都是借助这个开源库实现的。

![image-20240212132156626](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240212132156626.png)

### 3.2 使用 channel 传递消息



### 3.3 事件库与 NPC 库

🐲 Key 是当前修为阶段的幂，如果是2^14，则会遇到14列表的 NPC 或者事件。

```go
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
    // ...
)
```

### 3.4 巨龙的战斗方法

🐲 回合制战斗，你死我活！

```go
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
	randomDecreaseState(d)
}
```

### 3.5 巨龙遭遇事件

```go
// Process 处理偶发事件
func (d *Dragon) Process(e *Event) {
	p.addHistoryLn(newHistoryInfo(e.Name))
	d.basic.attack += e.Attack
	d.basic.defense += e.Defense
	appendLife(d, e.Life)
	appendExperience(d, e.Experience)
}
```

### 3.6 修为进阶函数

🐲 进阶的成功概率通过 const 常量 `SuccessRate` 进行控制，这里是75%。

```go
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
```

### 3.7 修养生息

🐲 恢复生命值，并增长修为。

```go
// 修养生息
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
```

### 3.8 外出冒险

🐲 小概率遭遇事件，大概率遭遇 NPC。

```go
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
```

## 四、排行榜

当游戏结算之后，按照所获得的经验值，录入排行榜，当前展示 Top10。

```go
func (r *Rank) save() {
	// 打开数据库连接。如果数据库文件不存在，将会创建它。
	db, err := sql.Open("sqlite3", "rank.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查并创建表
	if err := createTableIfNotExists(db); err != nil {
		log.Fatal(err)
	}

	// 插入数据
	insertSQL := `
	INSERT INTO ranks (name, experience, experience_stage, attack, defense, life) VALUES (?, ?, ?, ?, ?, ?);
	`
	_, err = db.Exec(insertSQL, r.Name, r.Experience, r.ExperienceStage, r.Attack, r.Defense, r.Life)
	if err != nil {
		log.Fatal(err)
	}
}
```

## 五、小结

> 🐲 游戏代码在我的开源 [Go 学习仓库](https://github.com/BaiZe1998/go-learning/tree/main/dragon)的`dragon`路径下，可以直接运行可执行程序体验。
>
> 这里也向大家发起挑战：200轮次内，挑战可以获取的最高修为值！欢迎加q群分享：622383022，或者评论区展示你的战绩！

