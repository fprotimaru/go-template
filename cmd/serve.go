package cmd

import (
	"log"

	"project/internal/config"
	"project/internal/service/auth/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	Run: func(c *cobra.Command, args []string) {
		cfg, err := config.New(cmd.cfg)
		if err != nil {
			log.Fatalln(err)
		}

		authUC := authusecase.NewUseCase([]byte(cfg.SecretKey))

		app := fiber.New(fiber.Config{
			JSONEncoder: jsoniter.Marshal,
			JSONDecoder: jsoniter.Unmarshal,
		})
		app.Use(cors.New(cors.ConfigDefault))
		app.Use(logger.New(logger.ConfigDefault))

		app.Use(authUC.Middleware)

		if err := app.Listen(cmd.addr); err != nil {
			log.Fatalln(err)
		}
	},
}
