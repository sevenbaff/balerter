package manager

import (
	"github.com/balerter/balerter/internal/alert/message"
	"go.uber.org/zap"
)

type SendOption func()

func (m *Manager) Send(alertName, text string, channels []string, fields []string, image string) {
	chs := make(map[string]alertChannel)

	if len(channels) > 0 {
		for _, channelName := range channels {
			ch, ok := m.channels[channelName]
			if !ok {
				m.logger.Error("channel not found", zap.String("channel name", channelName))
				continue
			}
			chs[channelName] = ch
		}
	} else {
		chs = m.channels
	}

	if len(chs) == 0 {
		m.logger.Error("empty channels")
		return
	}

	for name, module := range chs {
		if err := module.Send(message.New(alertName, text, fields, image)); err != nil {
			m.logger.Error("error send message to channel", zap.String("channel name", name), zap.Error(err))
		}
	}
}
