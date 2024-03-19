package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	grpcChatController "github.com/pillarion/practice-chat-server/internal/adapter/controller/grpc"
	"github.com/pillarion/practice-chat-server/internal/adapter/controller/interceptor"
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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	closer "github.com/pillarion/practice-platform/pkg/closer"
	pgClient "github.com/pillarion/practice-platform/pkg/dbclient"
	txManager "github.com/pillarion/practice-platform/pkg/pgtxmanager"

	accessClientDriver "github.com/pillarion/practice-chat-server/internal/core/tools/access_v1"
	accessClient "github.com/pillarion/practice-chat-server/internal/core/tools/client/access"
)

type serviceProvider struct {
	config *config.Config

	dbDriver          pgClient.DB
	dbClient          pgClient.Client
	txManager         txManager.TxManager
	chatRepository    chatRepoPort.Repo
	messageRepository messageRepoPort.Repo
	journalRepository journalRepoPort.Repo

	chatService chatServicePort.Service

	chatServer *grpcChatController.Server

	accessClientDriver accessClientDriver.AccessV1Client
	accessClient       accessClient.V1Client

	interceptor *interceptor.ChatServerInterceptor
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

func (s *serviceProvider) DBDriver(ctx context.Context) pgClient.DB {
	if s.dbDriver == nil {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			s.Config().Database.Host,
			s.Config().Database.Port,
			s.Config().Database.User,
			s.Config().Database.Db,
			s.Config().Database.Pass,
		)
		db, err := pgClient.NewDB(ctx, dsn)
		if err != nil {
			log.Fatalf("failed to create db driver: %v", err)
		}

		s.dbDriver = db
	}

	return s.dbDriver
}

func (s *serviceProvider) DBClient(ctx context.Context) pgClient.Client {
	if s.dbClient == nil {
		cl, err := pgClient.New(s.DBDriver(ctx))
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) txManager.TxManager {
	if s.txManager == nil {
		s.txManager = txManager.NewTransactionManager(s.DBClient(ctx).DB())
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

// AccessClient returns an access client
func (s *serviceProvider) AccessClientDriver(_ context.Context) accessClientDriver.AccessV1Client {
	if s.accessClientDriver == nil {
		pemServerCA, err := os.ReadFile(s.Config().Access.CAcert)
		if err != nil {
			log.Fatal("failed to load server CA's certificate")
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(pemServerCA) {
			log.Fatal("failed to add server CA's certificate")
		}

		// Create the credentials and return it
		tlsConfig := &tls.Config{
			RootCAs:    certPool,
			MinVersion: tls.VersionTLS12,
		}

		conn, err := grpc.Dial(
			s.Config().Access.Address,
			grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		)
		if err != nil {
			log.Fatalf("failed to create access client: %v", err)
		}
		closer.Add(conn.Close)

		cl := accessClientDriver.NewAccessV1Client(conn)

		s.accessClientDriver = cl
	}

	return s.accessClientDriver
}

func (s *serviceProvider) AccessClient(ctx context.Context) accessClient.V1Client {
	if s.accessClient == nil {
		s.accessClient = accessClient.NewV1Client(s.AccessClientDriver(ctx))
	}

	return s.accessClient
}

func (s *serviceProvider) Interceptor(ctx context.Context) *interceptor.ChatServerInterceptor {
	if s.interceptor == nil {
		s.interceptor = interceptor.NewChatServerInterceptor(s.AccessClient(ctx))
	}

	if s.interceptor == nil {
		log.Fatalf("failed to create interceptor")
	}

	return s.interceptor
}
