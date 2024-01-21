package toast

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/remshams/common/tui/styles"
)

type Toast interface {
	Message() string
}

type WarningToast struct {
	message string
}

func (warningToast WarningToast) Message() string {
	return styles.TextWarningColor.Render(warningToast.message)
}

func CreateWarningToastAction(message string) tea.Cmd {
	return func() tea.Msg {
		return WarningToast{message}
	}
}

type ErrorToast struct {
	message string
}

func (errorToast ErrorToast) Message() string {
	return styles.TextErrorColor.Render(errorToast.message)
}

func CreateErrorToastAction(message string) tea.Cmd {
	return func() tea.Msg {
		return ErrorToast{message}
	}
}

type InfoToast struct {
	message string
}

func (infoToast InfoToast) Message() string {
	return styles.TextInfoColor.Render(infoToast.message)
}

func CreateInfoToastAction(message string) tea.Cmd {
	return func() tea.Msg {
		return InfoToast{message}
	}
}

type SuccessToast struct {
	message string
}

func (successToast SuccessToast) Message() string {
	return styles.TextSuccessColor.Render(successToast.message)
}

func CreateSuccessToastAction(message string) tea.Cmd {
	return func() tea.Msg {
		return SuccessToast{message}
	}
}
