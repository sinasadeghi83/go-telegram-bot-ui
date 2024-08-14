package dialog

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Button struct {
	Text            string
	NodeID          string
	URL             string
	CallbackHandler bot.HandlerFunc
}

type Node struct {
	ID       string
	Text     string
	Keyboard [][]Button
}

func (n Node) buildKB(prefix, nodePrefix, callbackPrefix, parentNodeID string) models.ReplyMarkup {
	if len(n.Keyboard) == 0 {
		return nil
	}

	var kb [][]models.InlineKeyboardButton

	for _, row := range n.Keyboard {
		var kbRow []models.InlineKeyboardButton
		for _, btn := range row {
			b := models.InlineKeyboardButton{
				Text: btn.Text,
			}
			switch {
			case btn.URL != "":
				b.URL = btn.URL
			case btn.CallbackHandler != nil:
				b.CallbackData = prefix + callbackPrefix + parentNodeID + "_" + btn.Text
			default:
				b.CallbackData = prefix + nodePrefix + btn.NodeID

			}
			kbRow = append(kbRow, b)
		}
		kb = append(kb, kbRow)
	}

	return models.InlineKeyboardMarkup{
		InlineKeyboard: kb,
	}
}
