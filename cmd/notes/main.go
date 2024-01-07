package main

import (
	"context"
	"database/sql"
	"errors"
	"net"
	"net/http"
	"os"

	"github.com/IEP/sqlite-wasm/database/migrations"
	"github.com/IEP/sqlite-wasm/database/migrations/sqlite3"
	pb "github.com/IEP/sqlite-wasm/gen/go/protos/notes/v1"
	"github.com/IEP/sqlite-wasm/internal/notes"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Info().Msg("starting notes service")

	// start - config
	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = "file:data.sqlite?_txlock=immediate"
	}
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "8001"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	// end - config

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open DB")
	}
	db.SetMaxOpenConns(1)
	dbMigrate(db)

	// start - init services
	svc := notes.NewNotesGRPCService(db)
	// end - init services

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		log.Info().Msgf("starting grpc server on port %s", grpcPort)

		l, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			return err
		}
		grpcServer := grpc.NewServer()
		pb.RegisterNoteServiceServer(grpcServer, svc)
		if err := grpcServer.Serve(l); err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		log.Info().Msgf("starting http server on port %s", port)

		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}
		endpoint := net.JoinHostPort("localhost", grpcPort)
		err := pb.RegisterNoteServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
		if err != nil {
			return err
		}

		if err := http.ListenAndServe(":"+port, mux); err != nil {
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatal().Err(err).Msg("failed to run server")
	}
}

func dbMigrate(db *sql.DB) {
	sd, err := iofs.New(migrations.FS, ".")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create fs")
	}
	dd, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create db instance")
	}

	m, err := migrate.NewWithInstance("iofs", sd, "", dd)
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal().Err(err).Msg("failed to perform db migration")
	}
}
