package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ServiceWeaver/weaver"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmsgprefix)
	log.SetPrefix(fmt.Sprintf("[%d] [weaver] ", os.Getpid()))
}

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

func serve(ctx context.Context, app *App) error {
	server := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	server.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	server.Get("/users", func(c *fiber.Ctx) error {
		users, err := app.userRepository.Get().FindAll(c.Context())

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(users)
	})

	server.Get("/users/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		user, err := app.userRepository.Get().FindById(c.Context(), uint64(id))

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(user)
	})

	server.Post("/users", func(c *fiber.Ctx) error {
		var body struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		user, err := app.userRepository.Get().Create(c.Context(), User{
			Name:  body.Name,
			Email: body.Email,
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(user)
	})

	server.Put("/users/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		var body struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		user, err := app.userRepository.Get().Update(c.Context(), User{
			Id:    uint64(id),
			Name:  body.Name,
			Email: body.Email,
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(user)
	})

	server.Delete("/users/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		if err := app.userRepository.Get().Delete(c.Context(), uint64(id)); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.SendStatus(fiber.StatusNoContent)
	})

	log.Printf("Server is listening on: http://%s\n", app.listener.Addr())
	log.Printf("Check health status at: http://%s/health\n", app.listener.Addr())
	return server.Listener(app.listener)
}

type App struct {
	weaver.Implements[weaver.Main]
	listener       weaver.Listener
	userRepository weaver.Ref[UserRepository]
}

type User struct {
	weaver.AutoMarshal
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserRepository interface {
	FindAll(context.Context) ([]User, error)
	FindById(context.Context, uint64) (User, error)
	Create(context.Context, User) (User, error)
	Update(context.Context, User) (User, error)
	Delete(context.Context, uint64) error
}

type userRepositoryImpl struct {
	weaver.Implements[UserRepository]
	db *sql.DB
}

func (r *userRepositoryImpl) Init(context.Context) error {
	db, err := sql.Open("sqlite3", "./dev.db?_journal_mode=WAL&mode=rwc&cache=shared")

	if err != nil {
		return err
	}

	r.db = db

	return nil
}

func (r *userRepositoryImpl) FindAll(ctx context.Context) ([]User, error) {
	log.Println("UserRepository.FindAll()")
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepositoryImpl) FindById(ctx context.Context, id uint64) (User, error) {
	log.Printf("UserRepository.FindById(%d)\n", id)
	var user User
	if err := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *userRepositoryImpl) Create(ctx context.Context, user User) (User, error) {
	result, err := r.db.ExecContext(ctx, "INSERT INTO users(name, email) VALUES(?, ?)", user.Name, user.Email)

	if err != nil {
		return User{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return User{}, err
	}

	return r.FindById(ctx, uint64(id))
}

func (r *userRepositoryImpl) Update(ctx context.Context, user User) (User, error) {
	log.Println("UserRepository.Update()")
	_, err := r.db.ExecContext(ctx, "UPDATE users SET name = ?, email = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?", user.Name, user.Email, user.Id)

	if err != nil {
		return User{}, err
	}

	return r.FindById(ctx, user.Id)
}

func (r *userRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	log.Printf("UserRepository.Delete(%d)\n", id)
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
