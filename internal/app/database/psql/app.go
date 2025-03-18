package psqlapp

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log/slog"
	"runtime"
)

type PsqlApp struct {
	log *slog.Logger
	db  *sqlx.DB
}

func (psql *PsqlApp) Db() *sqlx.DB {
	return psql.db
}

func NewPsqlApp(
	log *slog.Logger,
	user string,
	dbname string,
	password string,
	host string,
	port string,
) *PsqlApp {
	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"user=%s dbname=%s sslmode=disable password=%s host=%s port=%s",
			user,
			dbname,
			password,
			host,
			port,
		),
	)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	return &PsqlApp{
		log: log,
		db:  db,
	}
}

func (psql *PsqlApp) MustRun() {
	if err := psql.Run(); err != nil {
		panic(err)
	}
}

func (psql *PsqlApp) Run() error {
	psql.log.With(slog.String("operation", psql.MethodForLog())).
		Info(
			"Starting PSQL server",
			slog.String("conn", psql.db.DriverName()),
		)

	if err := psql.db.Ping(); err != nil {
		return fmt.Errorf("%s: %w", psql.MethodForLog(), err)
	}

	return nil
}

func (psql *PsqlApp) Stop() {
	psql.log.With(slog.String("operation", psql.MethodForLog())).
		Info(
			"Stopping PSQL server",
			slog.String("conn", psql.db.DriverName()),
		)

	if err := psql.db.Close(); err != nil {
		panic(err)
	}
}

func (psql *PsqlApp) MethodForLog() string {
	pc, _, _, _ := runtime.Caller(1)

	return runtime.FuncForPC(pc).Name()
}
