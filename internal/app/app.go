package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	db "github.com/rem-aster/goth-template/internal/database"
	"github.com/rem-aster/goth-template/internal/handlers"
)

type App struct {
	Env *Env
}

type Env struct {
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_DB       string
}

func New() *App {
	return &App{}
}

func (a *App) Init() {
	a.SetupEnv()
	a.SetupServer()
}

func (a *App) SetupEnv() {
	a.Env = &Env{
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
	}
}

func (a *App) SetupServer() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		a.Env.POSTGRES_USER,
		a.Env.POSTGRES_PASSWORD,
		a.Env.POSTGRES_HOST,
		a.Env.POSTGRES_PORT,
		a.Env.POSTGRES_DB,
	))
	if err != nil {
		log.Fatalf("Error Connecting to database:\nMessage:\n%v", err.Error())
	}
	db := db.New(conn)

	h := handlers.New(db)
	h.SetupRoutes()

	err = h.StartServer()
	if err != nil {
		log.Printf("Error Starting Server:\nMessage:\n%v", err.Error())
	}
}
