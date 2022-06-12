package cmd

import (
	"fmt"

	"os"

	"github.com/ferealqq/cup-util/pkg/api"
	"github.com/ferealqq/cup-util/pkg/store"
	"github.com/spf13/cobra"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var ID = 187356618

var p *tea.Program
var statusList []list.Item

var getTask = &cobra.Command{
	Use:   "task",
	Short: "get tasks from list",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		statusList = []list.Item{}

		tasks, err := store.FetchTasks(api.TaskQuery{List: ID}, false)

		if err == nil {
			for _, v := range tasks {
				statusList = append(statusList, item{title: v.Status.Status, desc: v.Status.Type})
			}

			m := status{list: list.New(statusList, list.NewDefaultDelegate(), 0, 0)}
			m.list.Title = "Statuses"

			p = tea.NewProgram(m)

			if err := p.Start(); err != nil {
				fmt.Println("Error running program:", err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Error occured whilst fetching items")
		}
	},
}

// s is the status name
func createTaskList(s string) ([]list.Item, error) {
	items := []list.Item{
		item{title: "Raspberry Pi's", desc: "I have em all over my house", status: s},
	}

	tasks, err := store.FetchTasks(api.TaskQuery{List: ID}, false)
	if err == nil {
		for _, v := range tasks {
			if s == v.Status.Status {
				items = append(items, item{title: v.Name, desc: v.Description, status: v.Status.Status})
			}
		}

		return items, nil
	} else {
		return nil, err
	}
}

var docStyle = lipgloss.NewStyle().Margin(1, 3)

type item struct {
	title, desc, status string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) Status() string      { return i.status }
func (i item) FilterValue() string { return i.title }

type status struct {
	list list.Model
}

func (m status) Init() tea.Cmd {
	return nil
}

func (m status) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		var title string

		if i, ok := m.list.SelectedItem().(item); ok {
			title = i.Title()
		}

		switch msg.String() {
		case "enter":
			if l, err := createTaskList(title); err == nil {

				return m, m.list.SetItems(l)
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m status) View() string {
	return docStyle.Render(m.list.View())
}
