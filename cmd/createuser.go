package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"syscall"

	"github.com/fprotimaru/go-template/internal/config"
	"github.com/fprotimaru/go-template/internal/entity"
	"github.com/fprotimaru/go-template/internal/service/auth/usecase"
	"github.com/fprotimaru/go-template/internal/service/user/repository"
	"github.com/fprotimaru/go-template/internal/service/user/usecase"

	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"golang.org/x/crypto/ssh/terminal"
)

var createUserCmd = &cobra.Command{
	Use:   "createuser",
	Short: "create a new user",
	Run: func(c *cobra.Command, args []string) {
		cfg, err := config.New(cmd.cfg)
		if err != nil {
			log.Fatalln(err)
		}

		var username string
		fmt.Printf("Username: ")
		_, err = fmt.Scan(&username)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Password: ")
		passwordBytes, err := terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println()

		sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.PostgresUrl)))
		db := bun.NewDB(sqlDB, pgdialect.New())

		userRepo := userrepository.NewRepository(db)
		userUC := userusecase.NewUseCase(userRepo)

		authUC := authusecase.NewUseCase([]byte(cfg.SecretKey))
		password := authUC.HashPassword(passwordBytes)

		var user = entity.User{
			Username: username,
			Password: password,
		}

		var ctx = context.Background()
		_, err = userUC.Create(ctx, &user)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("user successfully created")
	},
}
