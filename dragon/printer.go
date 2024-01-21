package main

import (
	"context"
	"fmt"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/mum4k/termdash/widgets/textinput"
	"math"
	"os"
)

const (
	CLEAR = "\033[H\033[2J"
)

type historyInfo struct {
	info    string
	options []text.WriteOption
}

func newHistoryInfo(info string, options ...text.WriteOption) historyInfo {
	return historyInfo{
		info:    info,
		options: options,
	}
}

type printer struct {
	dragon          *Dragon
	ctx             context.Context
	terminal        *tcell.Terminal
	container       *container.Container
	historyText     *text.Text
	history         chan historyInfo
	operateHintText *text.Text
	operateHint     chan string
	scanned         chan string
	flushChannel    chan struct{}
	values          *text.Text
	experienceBar   *gauge.Gauge
	hpBar           *gauge.Gauge
	keyBinding      func(*terminalapi.Keyboard)
}

func newPrinter() *printer {
	terminal, err := tcell.New()
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())

	// 历史记录区
	historyPanel, _ := text.New(text.RollContent(), text.WrapAtWords())

	// 数值区
	values, _ := text.New(text.RollContent(), text.WrapAtRunes())

	// 操作提示区
	operationHint, _ := text.New(text.RollContent(), text.WrapAtRunes())

	// 输入区
	inputs, _ := textinput.New(
		textinput.DefaultText(""),
		textinput.MaxWidthCells(30),
		textinput.ExclusiveKeyboardOnFocus(),
		textinput.FillColor(cell.ColorYellow),
	)

	// 经验条
	experience, _ := gauge.New(
		gauge.Color(cell.ColorBlue),
		gauge.TextLabel("龙の经验条"),
		gauge.Height(1),
	)

	hpBar, _ := gauge.New(
		gauge.Color(cell.ColorRed),
		gauge.TextLabel("龙の血条"),
		gauge.Height(1),
	)

	// 创建界面
	c, err := container.New(
		terminal,
		container.Border(TotalBorderStyle),
		container.TitleColor(TotalBorderColor),
		container.BorderTitle(TotalBorderTitle),
		container.SplitVertical(
			container.Left(
				container.SplitVertical(
					container.Left(
						container.SplitHorizontal(
							container.Top(
								container.PlaceWidget(operationHint),
								container.Border(OperateAreaBorderStyle),
								container.BorderColor(OperateAreaBorderColor),
								container.BorderTitle(OperateAreaBorderTitle),
							),
							container.Bottom(
								container.Focused(),
								container.PlaceWidget(inputs),
								container.BorderTitle(InputAreaBorderTitle),
								container.Border(InputAreaBorderStyle),
								container.BorderColor(InputAreaBorderColor),
							),
							container.SplitPercent(80),
						),
					),
					container.Right(
						container.SplitHorizontal(
							container.Top(
								container.SplitHorizontal(
									container.Top(
										container.PlaceWidget(experience),
									),
									container.Bottom(
										container.PlaceWidget(hpBar),
									),
								),
							),
							container.Bottom(
								container.PlaceWidget(values),
							),
							container.SplitPercent(30),
						),
						container.BorderTitle(ValuesAreaBorderTitle),
						container.Border(ValuesAreaBorderStyle),
						container.BorderColor(ValuesAreaBorderColor),
					),
					container.SplitPercent(40),
				),
			),
			container.Right(
				container.PlaceWidget(historyPanel),
				container.BorderTitle(HistoryAreaBorderTitle),
				container.Border(HistoryAreaBorderStyle),
				container.BorderColor(HistoryAreaBorderColor),
				container.KeyFocusSkip(),
			),
			container.SplitPercent(30),
		),
	)

	p := &printer{
		terminal:        terminal,
		ctx:             ctx,
		container:       c,
		history:         make(chan historyInfo),
		historyText:     historyPanel,
		operateHintText: operationHint,
		operateHint:     make(chan string),
		scanned:         make(chan string),
		flushChannel:    make(chan struct{}),
		values:          values,
		experienceBar:   experience,
		hpBar:           hpBar,
		keyBinding: func(k *terminalapi.Keyboard) {
			// Ctrl + W 退出
			if k.Key == keyboard.KeyCtrlW {
				cancel()
				os.Exit(0)
			}

			// Enter 完成输入
			if k.Key == keyboard.KeyEnter {
				value := inputs.ReadAndClear()
				p.scanned <- value
			}
		},
	}
	go p.updateValuesPanel()
	go p.receiveHistory()
	go p.receiveOperateHint()

	return p
}

// 更新数值区域
func (p *printer) updateValuesPanel() {
	for {
		select {
		case <-p.flushChannel:
			if p.dragon == nil {
				return
			}

			p.values.Reset()
			p.values.Write(fmt.Sprint("龙の名字: ", p.dragon.Name, "\n"))
			p.values.Write(fmt.Sprint("龙の剩余寿命: ", p.dragon.Remaining, "\n"))
			p.values.Write(fmt.Sprintf("龙の级别: 2的%d次方龙\n", p.dragon.ExperienceStage))
			p.values.Write(fmt.Sprint("龙の攻击力: ", p.dragon.basic.attack, "\n"))
			p.values.Write(fmt.Sprint("龙の防御力: ", p.dragon.basic.defense, "\n"))

			p.experienceBar.Absolute(p.dragon.Experience, int(math.Pow(2, float64(p.dragon.ExperienceStage))))

			p.hpBar.Absolute(p.dragon.basic.life, p.dragon.basic.maxLife)
		}
	}
}

func (p *printer) setDragon(d *Dragon) {
	p.dragon = d
}

// 接收操作提示后的处理方法
func (p *printer) receiveOperateHint() {
	go func() {
		for {
			info := <-p.operateHint
			p.operateHintText.Reset()
			p.operateHintText.Write(info)
		}
	}()
}

// 接收历史数据处理方法
func (p *printer) receiveHistory() {
	go func() {
		for {
			select {
			case info := <-p.history:
				p.historyText.Write(info.info, info.options...)
			}
		}
	}()
}

// 显示操作提示
func (p *printer) addOperateHint(msg string) {
	p.operateHint <- msg
}

// 接收历史数据
func (p *printer) addHistory(info historyInfo) {
	p.history <- info
}

// 接收历史数据，并换行
func (p *printer) addHistoryLn(info historyInfo) {
	info.info += "\n"
	p.history <- info
}

// 接收信号，用于刷新数值区
func (p *printer) flush() {
	p.flushChannel <- struct{}{}
}
