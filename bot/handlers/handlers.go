package handlers

import (
	"fmt"
	"strings"

	"kheft/bot"
	"kheft/bot/languages"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func NonMemberStart(b *gotgbot.Bot, ctx *ext.Context) error {
	response := fmt.Sprintf(strings.Join(languages.Response.Messages.NonMember.Response, "\n"),
		ctx.Message.Chat.FirstName, bot.Configs.ChannelUsername)

	var keyboards [][]gotgbot.KeyboardButton
	btns := languages.Response.Messages.NonMember.Btns
	keyboard := make([]gotgbot.KeyboardButton, len(btns))

	for i, data := range btns {
		keyboard[i].Text = data
	}
	keyboards = append(keyboards, keyboard[:])
	markup := gotgbot.ReplyKeyboardMarkup{
		Keyboard:        keyboards,
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}

	_, err := ctx.EffectiveMessage.Reply(b, response,
		&gotgbot.SendMessageOpts{
			ReplyMarkup: markup,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to send reply: %w", err)
	}
	return nil
}

func MemberStart(b *gotgbot.Bot, ctx *ext.Context) error {
	var keyboards [][]gotgbot.KeyboardButton
	btns := languages.Response.Messages.Member.Btns
	keyboard := make([]gotgbot.KeyboardButton, len(btns))

	for i, data := range btns {
		keyboard[i].Text = data
	}
	keyboards = append(keyboards, keyboard[:])
	markup := gotgbot.ReplyKeyboardMarkup{
		Keyboard:        keyboards,
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}
	_, err := ctx.EffectiveMessage.Reply(b,
		strings.Join(languages.Response.Messages.Member.Response, "\n"),
		&gotgbot.SendMessageOpts{
			ReplyMarkup: markup,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to reply nonmemberchecking: %s", err)
	}
	return nil
}

func NonMemberChecking(b *gotgbot.Bot, ctx *ext.Context) error {
	status := (&bot.CheckMembershipOpts{}).CheckMessage(b)(ctx.EffectiveMessage)
	if status {
		err := MemberStart(b, ctx)
		if err != nil {
			return fmt.Errorf("failed to start nonmembercheking: %s", err)
		}
	} else {
		_, err := ctx.EffectiveMessage.Reply(b, languages.Response.Messages.NonMember.Failed, nil)
		if err != nil {
			return fmt.Errorf("failed to reply nonmemberchecking: %s", err)
		}
	}
	return nil
}
