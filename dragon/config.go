package main

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/text"
)

const (
	DEBUG = false

	// TotalBorderStyle 游戏区域边界样式
	TotalBorderStyle = linestyle.Double
	// TotalBorderColor 游戏区域边界颜色
	TotalBorderColor = cell.ColorYellow
	// TotalBorderTitle 游戏区域边界标题
	TotalBorderTitle = "按Ctrl + W退出"

	// OperateAreaBorderStyle 操作提示区边界样式
	OperateAreaBorderStyle = linestyle.Round
	// OperateAreaBorderColor 操作提示区边界颜色
	OperateAreaBorderColor = cell.ColorYellow
	// OperateAreaBorderTitle 操作提示区边界标题
	OperateAreaBorderTitle = "操作提示"

	// InputAreaBorderStyle 输入区边界样式
	InputAreaBorderStyle = linestyle.Round
	// InputAreaBorderColor 输入区边界颜色
	InputAreaBorderColor = cell.ColorYellow
	// InputAreaBorderTitle 输入区边界标题
	InputAreaBorderTitle = "按回车完成输入"

	// ValuesAreaBorderStyle 数值区边界样式
	ValuesAreaBorderStyle = linestyle.Round
	// ValuesAreaBorderColor 数值区边界颜色
	ValuesAreaBorderColor = cell.ColorYellow
	// ValuesAreaBorderTitle 数值区边界标题
	ValuesAreaBorderTitle = "龙の数值"

	// HistoryAreaBorderStyle 历史记录区边界样式
	HistoryAreaBorderStyle = linestyle.Round
	// HistoryAreaBorderColor 历史记录区边界颜色
	HistoryAreaBorderColor = cell.ColorYellow
	// HistoryAreaBorderTitle 历史记录区边界标题
	HistoryAreaBorderTitle = "龙生经历"

	// RankAreaBorderStyle 排行榜区边界样式
	RankAreaBorderStyle = linestyle.Round
	// RankAreaBorderColor 排行榜区边界颜色
	RankAreaBorderColor = cell.ColorYellow
	// RankAreaBorderTitle 排行榜区边界标题
	RankAreaBorderTitle = "龙の排行榜 TOP 10"
)

var (
	TextOptionBold          = text.WriteCellOpts(cell.Bold())
	TextOptionUnderline     = text.WriteCellOpts(cell.Underline())
	TextOptionItalic        = text.WriteCellOpts(cell.Italic())
	TextOptionBlink         = text.WriteCellOpts(cell.Blink())
	TextOptionInverse       = text.WriteCellOpts(cell.Inverse())
	TextOptionStrikethrough = text.WriteCellOpts(cell.Strikethrough())

	TextOptionRed     = text.WriteCellOpts(cell.FgColor(cell.ColorRed))
	TextOptionBlue    = text.WriteCellOpts(cell.FgColor(cell.ColorBlue))
	TextOptionCyan    = text.WriteCellOpts(cell.FgColor(cell.ColorCyan))
	TextOptionGray    = text.WriteCellOpts(cell.FgColor(cell.ColorGray))
	TextOptionGreen   = text.WriteCellOpts(cell.FgColor(cell.ColorGreen))
	TextOptionMagenta = text.WriteCellOpts(cell.FgColor(cell.ColorMagenta))
	TextOptionYellow  = text.WriteCellOpts(cell.FgColor(cell.ColorYellow))
	TextOptionWhite   = text.WriteCellOpts(cell.FgColor(cell.ColorWhite))
	TextOptionBlack   = text.WriteCellOpts(cell.FgColor(cell.ColorBlack))
)
