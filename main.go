package main

import (
	"context"
	"fmt"
	"gofiber/handler"
	"gofiber/repository"
	"gofiber/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	gormLogger "gorm.io/gorm/logger"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

type SqlLogger struct {
	gormLogger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n ================================== \n", sql)
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: &SqlLogger{},
	})
	if err != nil {
		panic("failed to connect database")
	}

	userRepository := repository.NewUserSqlliteRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	db.AutoMigrate(&User{})
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

	app.Post("/signup", userHandler.SignUp)
	app.Post("/login", Login)
	app.Post("/hello", Login)

	app.Listen(":3000")
}

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func SignUp(c *fiber.Ctx) error {
	request := SignupRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	if request.Username == "" || request.Email == "" {
		return fiber.ErrUnprocessableEntity
	}

	return nil
}

func Login(c *fiber.Ctx) error {
	return nil
}

func Hello(c *fiber.Ctx) error {
	return nil
}
