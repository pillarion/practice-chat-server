package app

import (
	"context"
	"fmt"
	"net"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	desc "github.com/pillarion/practice-chat-server/pkg/chat_v1"
)

// App is the main application struct.
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

// NewApp initializes a new App.
//
// Takes a context as a parameter.
// Returns a pointer to App and an error.
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	desc.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatServer(ctx))

	return nil
}

// Run runs the App.
//
// No parameters.
// Returns an error.
func (a *App) Run() error {
	defer func() {
		CloseAll()
		Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) runGRPCServer() error {
	lAddress := fmt.Sprintf(":%d", a.serviceProvider.Config().GRPC.Port)
	list, err := net.Listen("tcp", lAddress)
	if err != nil {
		return err
	}
	slog.Info("GRPC server is running", "ListenAddress", lAddress)

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
