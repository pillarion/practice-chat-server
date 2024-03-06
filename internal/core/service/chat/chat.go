package chat

import (
	repoC "github.com/pillarion/practice-chat-server/internal/core/port/repository/chat"
	repoJ "github.com/pillarion/practice-chat-server/internal/core/port/repository/journal"
	repoM "github.com/pillarion/practice-chat-server/internal/core/port/repository/message"
	"github.com/pillarion/practice-chat-server/internal/core/port/service/chat"
	repoTxM "github.com/pillarion/practice-chat-server/internal/core/tools/dbclient/port/pgtxmanager"
)

type service struct {
	repoChat    repoC.Repo
	repoMessage repoM.Repo
	repoJournal repoJ.Repo
	txManager   repoTxM.TxManager
}

// NewService creates a new chat service.
//
// repoChat: the repository for chat.
// repoMessage: the repository for message.
// chat.Service: the new chat service.
func NewService(repoChat repoC.Repo, repoMessage repoM.Repo, repoJournal repoJ.Repo, txManager repoTxM.TxManager) chat.Service {
	return &service{
		repoChat:    repoChat,
		repoMessage: repoMessage,
		repoJournal: repoJournal,
		txManager:   txManager,
	}
}
