## 一、前言

新年就要到了，祝大家新的一年：🐲 龙行龘龘，🔥 前程朤朤！

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

### 1. 初始化

- 起一个你喜欢的名字！然后分配100点能力值到三维属性。
- 攻击力：当冒险遇到 NPC 时，会与其发生回合制战斗，每回合你对它可以造成：你的攻击力 - NPC 防御力的伤害。
- 防御力：当冒险遇到 NPC 时，会与其发生回合制战斗，每回合它对你可以造成：NPC 的攻击力 - 你防御力的伤害。
- 生命值：战斗和触发事件造成生命值为0均无法继续冒险，必须返回休养生息（每秒恢复最大生命值10%）。
- 输入整个游戏的轮次：这里输入200轮。

> 🐲 这里也向大家发起挑战：200轮次内，挑战可以获取的最高修为值！欢迎加q群分享：622383022，或者评论区展示你的战绩！

![image-20240119102637116](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240119102637116.png)

### 2. 外出冒险

选择2则出发外出冒险（输入冒险轮次，200轮挑战就是限制这个），冒险有大概率遇到 NPC 与其战斗，小概率触发随机事件，战斗与事件都会对你的各项属性造成影响，**但是生命值为0会丢失此次冒险的剩余轮次**。

🐲 **每 0.5 秒进行一轮冒险，打印日志。**

![image-20240119104025273](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240119104025273.png)

🐲 **冒险终止条件**：轮次消耗完毕或者生命为0。

🐲 **丢失轮次**：如果你选择外出冒险50轮，但是在第25轮的时候，生命为0，则会丢失剩余25轮次。

![image-20240119104240812](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240119104240812.png)

🐲 **游戏结束**：所有轮次消耗完毕。

### 3. 修为进阶

```go
// 修为进阶表
1, 2, 4, 8, 16, 32, 64, 128...
// 当修为到达2^x的时候，会触发进阶判断，询问你是否进阶（y/n）
1. 进阶成功：随机一项能力*2
2. 进阶失败：随机一项能力/2，修为/2
3. 放弃进阶：只有进入下一阶才能继续增长修为
// 当前成为2的5次方龙之前，默认自动选择进阶
```

![image-20240119105108313](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240119105108313.png)

### 4. 修养生息

![image-20240119105201498](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240119105201498.png)

🐲 生命回复、修为提升：

![image-20240119105602631](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240119105602631.png)

### 5. 游戏结算

🐲 这里当我还剩下60轮游戏的时候，我选择了直接进行60轮冒险，但是直接在冒险中暴毙了...（损失60轮，有时候运气才是最重要的）

最终修为：236！成为了2的7次方龙，好弱...

![image-20240119105911030](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20240119105911030.png)

## 三、游戏设计

### 3.1 事件库与 NPC 库

🐲 Key 是当前修为阶段的幂，如果是2^14，则会遇到14列表的 NPC 或者事件。

```go
var (
   // NPCs NPC库
   NPCs = map[int][]NPC{
      // 省略...
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
      // 省略...
   }
)
```

### 3.2 巨龙的战斗方法

🐲 回合制战斗，你死我活！

```go
// Fight 与 NPC 战斗
func (d *Dragon) Fight(n *NPC) {
   // 回合制战斗，你死我活！
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
      appendExperience(d, -n.Experience)
   } else {
      fmt.Printf("你打败了%s\n", n.Name)
      appendExperience(d, n.Experience)
   }
}
```

### 3.3 巨龙遭遇事件

```go
// Process 处理偶发事件
func (d *Dragon) Process(e *Event) {
   fmt.Println(e.Name)
   d.Attack += e.Attack
   d.Defense += e.Defense
   // 增加生命
   appendLife(d, e.Life)
   // 增加修为
   appendExperience(d, e.Experience)
}
```

### 3.4 修为进阶函数

🐲 进阶的成功概率通过 const 常量 `SuccessRate` 进行控制，这里是75%。

```go
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
```

### 3.5 修养生息

🐲 恢复生命值，并增长修为。

```go
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
```

### 3.6 外出冒险

🐲 小概率遭遇事件，大概率遭遇 NPC。

```go
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
```

## 四、小结

你长达200轮的一生真是波澜壮阔，你达成了以下成就：
姓名：白泽，修为：236，称号：2的7次方龙，攻击力：70，防御力：2225，生命值：0，剩余寿命：0轮



> 🐲 游戏代码在我的开源 [Go 学习仓库](https://github.com/BaiZe1998/go-learning/tree/main/dragon)的`dragon`路径下，可以直接运行可执行程序体验。
>
> 这里也向大家发起挑战：200轮次内，挑战可以获取的最高修为值！欢迎加q群分享：622383022，或者评论区展示你的战绩！

