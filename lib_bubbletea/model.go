package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	todos    []string // todo列表
	cursor   int      // 对应的todo的游标
	selected []bool   // 如果已经完成则置true
}

// Init 程序初始化时会调用
func (m model) Init() tea.Cmd {
	return nil
}

// Update 每次的按键时间进行响应，并且返回对应的cmd让程序进行执行
// msg 来自按键
// tea.Cmd 给予程序的命令（让其执行）：返回一些tea内置命令，让程序执行
// 如：tea.qui则让程序退出
// `ctrl+c`, `q`：程序退出
// `up`：游标上移
// `down`：游标下移
// `enter： 选择selected
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// 断言msg的类型，断言其为tea.KeyMsg类型
	// keyword.String()字符串代表按键输入的字符串
	// 如在键盘上按下`K`键，则其.String()输出的字符串是"k"
	// `enter`键同理，输出"enter"
	switch keyword := msg.(type) {
	case tea.KeyMsg:
		switch keyword.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor == 0 {
				break
			}
			m.cursor--
		case "down":
			if m.cursor == len(m.todos)-1 {
				break
			}
			m.cursor++
		case "enter":
			value := m.selected[m.cursor]
			m.selected[m.cursor] = !value
		}
	}
	return m, nil
}

func (m model) View() string {
	str := "todo：\n"
	for i, todo := range m.todos {
		cursorStr := " "
		selectedStr := " "
		if m.cursor == i {
			cursorStr = ">"
		}
		if m.selected[i] {
			selectedStr = "o"
		}
		str += fmt.Sprintf("%s [%s] %s\n", cursorStr, selectedStr, todo)
	}
	return str + "\n"
}

func NewModel(todos []string) *model {
	selected := make([]bool, len(todos))
	for i := 0; i < len(todos); i++ {
		selected = append(selected, false)
	}
	return &model{
		todos:    todos,
		cursor:   0,
		selected: selected,
	}
}
