package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	dgrpc "github.com/pillarion/practice-chat-server/internal/adapter/controller/grpc"
	config "github.com/pillarion/practice-chat-server/internal/adapter/driver/config/env"
	dchat "github.com/pillarion/practice-chat-server/internal/adapter/driver/db/postgresql/chat"
	dmsg "github.com/pillarion/practice-chat-server/internal/adapter/driver/db/postgresql/message"
	chat "github.com/pillarion/practice-chat-server/internal/core/service/chat"
	desc "github.com/pillarion/practice-chat-server/pkg/chat_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Get()
	if err != nil {
		slog.Warn("failed to get config", "Error", err)
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Db, cfg.Database.Pass)

	pgx, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		slog.Warn("failed to listen", "Error", err)
	}
	repoC, err := dchat.New(pgx)
	if err != nil {
		slog.Warn("failed to create repo", "Error", err)
	}
	repoM, err := dmsg.New(pgx)
	if err != nil {
		slog.Warn("failed to create repo", "Error", err)
	}
	cs := chat.NewService(repoC, repoM)

	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterChatV1Server(s, dgrpc.NewServer(cs))

	slog.Info("server listening at", "address", lis.Addr().String())

	if err = s.Serve(lis); err != nil {
		slog.Warn("failed to serve", "Error", err)
	}
}
