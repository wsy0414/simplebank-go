package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"simplebank/api"
	"simplebank/config"
	"simplebank/db/sqlc"
	_ "simplebank/flags"
	"simplebank/gapi"
	"simplebank/pb"
	"syscall"

	_ "simplebank/docs"

	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

//	@title			Swagger SimpleBank API
//	@version		1.0
//	@description	This is a simple bank server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost:8080

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	config.LoadConfig("./config")
	db := connDB()

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()
	waitGroup, ctx := errgroup.WithContext(ctx)

	runHttpServer(ctx, waitGroup, &config.ConfigVal, sqlc.NewStore(db))
	runGrpcServer(ctx, waitGroup, &config.ConfigVal, sqlc.NewStore(db))

	err := waitGroup.Wait()
	if err != nil {
		log.Println("error from wait group")
	}
}

// connDB return a sql package's DB implement
func connDB() *sql.DB {
	db, err := sql.Open(config.ConfigVal.Database.Driver, config.ConfigVal.Database.Source)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func runHttpServer(ctx context.Context, waitGroup *errgroup.Group, config *config.Config, store sqlc.Store) {
	server := api.NewServer(store)

	httpServer := http.Server{
		Handler: server,
		Addr:    config.Server.Port,
	}

	waitGroup.Go(func() error {
		err := httpServer.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				return nil
			}
			log.Printf("start server error: %s", err.Error())
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Println("Gracefully Shutdown http server")
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatalf("failed to shutdown http server: %s", err.Error())
			return err
		}
		return nil
	})

}

func runGrpcServer(ctx context.Context, waitGroup *errgroup.Group, config *config.Config, store sqlc.Store) {
	server := gapi.NewServer(store)

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.Grpc.Port)
	if err != nil {
		log.Fatal("cannot create listener", err)
	}

	log.Println("grpc server listen address", listener.Addr().String())
	waitGroup.Go(func() error {
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatal("cannot connect server")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Println("Gracefully Stop Grpc Server")
		grpcServer.GracefulStop()

		return nil
	})

}
