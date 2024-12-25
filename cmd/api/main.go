package main

import (
	"GopherNetwork/internal/db"
	"GopherNetwork/internal/env"
	"GopherNetwork/internal/mailer"
	"GopherNetwork/internal/store"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const version = "0.0.1"

//	@title			GopherNetwork API
//	@description	API for GopherNetwork, a social network for gophers
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description

func main() {
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	cfg := config{
		addr:        env.GetString("ADDR", ":8080"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:8080"),
		frontendURL: env.GetString("FRONTEND_URL", "localhost:4000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:1234@localhost/gophernetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp: time.Hour * 24 * 3, // 3 days,
			mailtrap: mailTrapConfig{
				apiKey: env.GetString("MAILTRAP_API_KEY", ""),
			},
			fromEmail: env.GetString("FROM_EMAIL", ""),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("DB connection pool established")

	store := store.NewStorage(db)

	mailtrap, err := mailer.NewMailTrapClient(cfg.mail.mailtrap.apiKey, cfg.mail.fromEmail)
	if err != nil {
		logger.Fatal(err)
	}
	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
		mailer: mailtrap,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
