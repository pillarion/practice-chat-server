package chat

import (
	repoC "github.com/pillarion/practice-chat-server/internal/core/port/repository/chat"
	repoJ "github.com/pillarion/practice-chat-server/internal/core/port/repository/journal"
	repoM "github.com/pillarion/practice-chat-server/internal/core/port/repository/message"
	"github.com/pillarion/practice-chat-server/internal/core/port/service/chat"
	txManager "github.com/pillarion/practice-platform/pkg/pgtxmanager"
)

type service struct {
	repoChat    repoC.Repo
	repoMessage repoM.Repo
	repoJournal repoJ.Repo
	txManager   txManager.TxManager
}

// NewService creates a new chat service.
//
// repoChat: the repository for chat.
// repoMessage: the repository for message.
// chat.Service: the new chat service.
func NewService(repoChat repoC.Repo, repoMessage repoM.Repo, repoJournal repoJ.Repo, txManager txManager.TxManager) chat.Service {
	return &service{
		repoChat:    repoChat,
		repoMessage: repoMessage,
		repoJournal: repoJournal,
		txManager:   txManager,
	}
}
