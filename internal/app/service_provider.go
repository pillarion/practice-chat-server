package app

import (
	"context"
	"fmt"
	"log"

	grpcChatController "github.com/pillarion/practice-chat-server/internal/adapter/controller/grpc"
	configDriver "github.com/pillarion/practice-chat-server/internal/adapter/driver/config/env"
	pgChatDriver "github.com/pillarion/practice-chat-server/internal/adapter/driver/db/postgresql/chat"
	pgJournalDriver "github.com/pillarion/practice-chat-server/internal/adapter/driver/db/postgresql/journal"
	pgMessageDriver "github.com/pillarion/practice-chat-server/internal/adapter/driver/db/postgresql/message"
	config "github.com/pillarion/practice-chat-server/internal/core/entity/config"
	chatRepoPort "github.com/pillarion/practice-chat-server/internal/core/port/repository/chat"
	journalRepoPort "github.com/pillarion/practice-chat-server/internal/core/port/repository/journal"
	messageRepoPort "github.com/pillarion/practice-chat-server/internal/core/port/repository/message"
	chatServicePort "github.com/pillarion/practice-chat-server/internal/core/port/service/chat"
	chatService "github.com/pillarion/practice-chat-server/internal/core/service/chat"
	pgClientDriver "github.com/pillarion/practice-chat-server/internal/core/tools/dbclient/adapter/pgclient"
	txManagerDriver "github.com/pillarion/practice-chat-server/internal/core/tools/dbclient/adapter/pgtxmanager"
	pgClientRepoPort "github.com/pillarion/practice-chat-server/internal/core/tools/dbclient/port/pgclient"
	txManagerRepoPort "github.com/pillarion/practice-chat-server/internal/core/tools/dbclient/port/pgtxmanager"
	pgClientService "github.com/pillarion/practice-chat-server/internal/core/tools/dbclient/service/pgclient"
)

type serviceProvider struct {
	config *config.Config

	dbDriver          pgClientRepoPort.DB
	dbClient          pgClientRepoPort.Client
	txManager         txManagerRepoPort.TxManager
	chatRepository    chatRepoPort.Repo
	messageRepository messageRepoPort.Repo
	journalRepository journalRepoPort.Repo

	chatService chatServicePort.Service

	chatServer *grpcChatController.Server
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Config() *config.Config {
	if s.config == nil {
		cfg, err := configDriver.Get()
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) DBDriver(ctx context.Context) pgClientRepoPort.DB {
	if s.dbDriver == nil {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			s.Config().Database.Host,
			s.Config().Database.Port,
			s.Config().Database.User,
			s.Config().Database.Db,
			s.Config().Database.Pass,
		)
		db, err := pgClientDriver.NewDB(ctx, dsn)
		if err != nil {
			log.Fatalf("failed to create db driver: %v", err)
		}

		s.dbDriver = db
	}

	return s.dbDriver
}

func (s *serviceProvider) DBClient(ctx context.Context) pgClientRepoPort.Client {
	if s.dbClient == nil {
		cl, err := pgClientService.New(s.DBDriver(ctx))
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) txManagerRepoPort.TxManager {
	if s.txManager == nil {
		s.txManager = txManagerDriver.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) chatRepoPort.Repo {
	if s.chatRepository == nil {
		repo, err := pgChatDriver.New(s.DBClient(ctx))
		if err != nil {
			log.Fatalf("failed to create user repository: %v", err)
		}

		s.chatRepository = repo
	}

	return s.chatRepository
}

func (s *serviceProvider) MessageRepository(ctx context.Context) messageRepoPort.Repo {
	if s.messageRepository == nil {
		repo, err := pgMessageDriver.New(s.DBClient(ctx))
		if err != nil {
			log.Fatalf("failed to create user repository: %v", err)
		}

		s.messageRepository = repo
	}

	return s.messageRepository
}

func (s *serviceProvider) JournalRepository(ctx context.Context) journalRepoPort.Repo {
	if s.journalRepository == nil {
		repo, err := pgJournalDriver.New(s.DBClient(ctx))
		if err != nil {
			log.Fatalf("failed to create user repository: %v", err)
		}

		s.journalRepository = repo
	}

	return s.journalRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) chatServicePort.Service {
	if s.chatService == nil {
		service := chatService.NewService(
			s.ChatRepository(ctx),
			s.MessageRepository(ctx),
			s.JournalRepository(ctx),
			s.TxManager(ctx),
		)

		s.chatService = service
	}

	return s.chatService
}

func (s *serviceProvider) ChatServer(ctx context.Context) *grpcChatController.Server {
	if s.chatServer == nil {
		server := grpcChatController.NewServer(s.ChatService(ctx))

		s.chatServer = server
	}

	return s.chatServer
}
