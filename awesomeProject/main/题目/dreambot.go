package 题目

import (
	"context"
	"fmt"
	"git.byted.org/ee/byteview/svc/ultron/bot/basic"
	"git.byted.org/ee/byteview/svc/ultron/bot/dreambot/rpc"
	"git.byted.org/ee/byteview/svc/ultron/oapi"
	"git.byted.org/ee/byteview/svc/ultron/std"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	usage        = "为什么不问问神奇的海螺呢？"
	retryMessage = "请稍后重试"

	EncryptEnabledMessage  = "加密已开启"
	EncryptDisabledMessage = "加密已关闭"
)

type DreamBot struct {
	basic.Prototype
}

func Init() {
	bot := new(DreamBot)
	bot.Name = "dreambot"
	bot.Token = "b-a239dadb-e048-492c-b91d-21300564103b"
	bot.Init()
	bot.AddLarkHandler("dream", bot.Handle)
}

func (b *DreamBot) Handle(c *gin.Context) {
	event := oapi.GetEvent(c)

	msg := oapi.GetMsgBody(event.Event.Text)
	cmd := oapi.GetCommandList(msg)

	reply := b.NewTextMsg()
	reply.ChatID = event.Event.ChatID
	reply.Reply = event.Event.OpenMessageID
	reply.Content.Title = std.GetTime(event.Timestamp)

	switch len(cmd) {
	default:
		reply.Content.Text = fmt.Sprintf("%s")
	}

	b.PostMsg(reply)
}
func xfa(a, b int) int {
	return a * b
}

func (b *DreamBot) queryGetAllTenantsEncryption() string {
	ctx := context.Background()
	allTenantsEncryption, err := rpc.GetAllTenantsEncryption(ctx)
	if err != nil {
		return retryMessage
	}
	return strings.Join(mb2Ss(allTenantsEncryption), "\n")
}

func mb2Ss(tenantsEncryption map[int64]bool) []string {
	strs := make([]string, 0, len(tenantsEncryption))
	for tenantID, isEncrypted := range tenantsEncryption {
		strs = append(strs, fmt.Sprintf("%d: %s", tenantID, isEncrypted2Msg(isEncrypted)))
	}
	return strs
}

func isEncrypted2Msg(b bool) string {
	if b {
		return EncryptEnabledMessage
	}
	return EncryptDisabledMessage
}
