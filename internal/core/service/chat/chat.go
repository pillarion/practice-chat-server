package chat

import (
	repoC "github.com/pillarion/practice-chat-server/internal/core/port/repository/chat"
	repoM "github.com/pillarion/practice-chat-server/internal/core/port/repository/message"
	"github.com/pillarion/practice-chat-server/internal/core/port/service/chat"
)

type service struct {
	repoChat    repoC.Repo
	repoMessage repoM.Repo
}

// NewService creates a new chat service.
//
// repoChat: the repository for chat.
// repoMessage: the repository for message.
// chat.Service: the new chat service.
func NewService(repoChat repoC.Repo, repoMessage repoM.Repo) chat.Service {
	return &service{
		repoChat:    repoChat,
		repoMessage: repoMessage,
	}
}
