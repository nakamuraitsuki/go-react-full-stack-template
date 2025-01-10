package handler

import (
	"database/sql"
	"errors"
	"log"
	"server/model"
	"sort"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	db        *sqlx.DB
	validator *validator.Validate
}

func NewTodoHandler(db *sqlx.DB) *TodoHandler {
	return &TodoHandler{db: db, validator: validator.New()}
}

func (h *TodoHandler) Register(g *echo.Group) {
	g.GET("/todos", h.GetTodos)
	g.POST("/todos", h.CreateTodo)
	g.PUT("/todos/:id", h.UpdateTodo)
	g.DELETE("/todos/:id", h.DeleteTodo)
}

type GetTodosResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (h *TodoHandler) GetTodos(c echo.Context) error {
	todoListID := c.QueryParam("todo_list_id")
    if todoListID == "" {
        return c.JSON(400, map[string]string{"message": "user_id is required"})
    }

	if _, err := strconv.Atoi(todoListID); err != nil {
        return c.JSON(400, map[string]string{"message": "user_id must be a valid integer"})
    }

	var todos []model.Todo
	err := h.db.Select(&todos, "SELECT * FROM todos WHERE todo_list_id = ?", todoListID)
	if err != nil {
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(200, []GetTodosResponse{})
		}
		return c.JSON(500, map[string]string{"message": "Internal Server Error"})
	}

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].CreatedAt.After(todos[j].CreatedAt)
	})

	res := make([]GetTodosResponse, len(todos))
	for i, todo := range todos {
		res[i] = GetTodosResponse{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
		}
	}

	return c.JSON(200, res)
}

type CreateTodoRequest struct {
	Title string `json:"title" validate:"required"`
}

type CreateTodoResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
	todoListID := c.QueryParam("todo_list_id")
    if todoListID == "" {
        return c.JSON(400, map[string]string{"message": "user_id is required"})
    }

	if _, err := strconv.Atoi(todoListID); err != nil {
        return c.JSON(400, map[string]string{"message": "user_id must be a valid integer"})
    }

	req := new(CreateTodoRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, map[string]string{"message": "Bad Request"})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.JSON(400, map[string]string{"message": "Bad Request"})
	}

	res, err := h.db.Exec("INSERT INTO todos (todo_list_id, title) VALUES (?, ?)", todoListID, req.Title)
	if err != nil {
		return c.JSON(500, map[string]string{"message": "Internal Server Error"})
	}

	id, _ := res.LastInsertId()
	return c.JSON(201, CreateTodoResponse{
		ID:        int(id),
		Title:     req.Title,
		Completed: false,
	})
}

type UpdateTodoRequest struct {
	Id        int    `param:"id" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Completed bool   `json:"completed"` // bool型の場合はrequiredを指定しない(ゼロ値==falseが入る)
}

func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	todoListID := c.QueryParam("todo_list_id")
    if todoListID == "" {
        return c.JSON(400, map[string]string{"message": "user_id is required"})
    }

	if _, err := strconv.Atoi(todoListID); err != nil {
        return c.JSON(400, map[string]string{"message": "user_id must be a valid integer"})
    }

	req := new(UpdateTodoRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, map[string]string{"message": "Bad Request"})
	}

	if err := h.validator.Struct(req); err != nil {
		log.Println("validator error: ", err)
		return c.JSON(400, map[string]string{"message": "Bad Request"})
	}

	_, err := h.db.Exec("UPDATE todos SET title = ?, completed = ? WHERE id = ? AND todo_list_id = ?", req.Title, req.Completed, req.Id, todoListID)
	if err != nil {
		return c.JSON(500, map[string]string{"message": "Internal Server Error"})
	}

	return c.JSON(200, nil)
}

type DeleteTodoRequest struct {
	ID int `param:"id" validate:"required"`
}

func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	todoListID := c.QueryParam("todo_list_id")
    if todoListID == "" {
        return c.JSON(400, map[string]string{"message": "user_id is required"})
    }

	if _, err := strconv.Atoi(todoListID); err != nil {
        return c.JSON(400, map[string]string{"message": "user_id must be a valid integer"})
    }

	req := new(DeleteTodoRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, map[string]string{"message": "Bad Request"})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.JSON(400, map[string]string{"message": "Bad Request"})
	}

	_, err := h.db.Exec("DELETE FROM todos WHERE id = ? AND todo_list_id = ?", req.ID, todoListID)
	if err != nil {
		return c.JSON(500, map[string]string{"message": "Internal Server Error"})
	}

	return c.JSON(200, nil)
}
