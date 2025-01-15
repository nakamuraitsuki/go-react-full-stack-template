package handler

import (
	"database/sql"
	"errors"
	"log"
	"server/model"

	"github.com/go-playground/validator"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)


type TodoListHandler struct {
	db			*sqlx.DB
	validator 	*validator.Validate
}

//インスタンス生成・依存性注入
func NewTodoListHandler(db *sqlx.DB) *TodoListHandler {
	return &TodoListHandler{db: db, validator: validator.New()}
}

func (h *TodoListHandler) Register(g *echo.Group) {
	g.GET("/todo-lists", h.GetTodoListsByUser)
	g.GET("/todo-lists/:id", h.GetTodoListsByID)
	g.POST("/todo-lists", h.CreateTodoList)
	g.PUT("/todo-lists/:id", h.UpdateTodoList)
}

type GetTodoListsResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

//user_idからその人の作成したTODOリストを取得
func (h *TodoListHandler) GetTodoListsByUser(c echo.Context) error {
	//リクエストからuser_idパラメータを取得
	userID := c.Get("user_id").(int)

	var todoLists []model.TodoList
	err := h.db.Select(&todoLists, "SELECT * FROM todo_lists WHERE user_id = ?", userID)
	if err != nil {
		log.Println(err)
		//0件取得でのエラーの場合
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(200, []GetTodoListsResponse{})
		}
		return c.JSON(500, map[string]string{"message": "Internal Server Error"})
	}

	res := make([]GetTodoListsResponse, len(todoLists))
	for i, todolist := range todoLists {
		res[i] = GetTodoListsResponse{
			ID:		todolist.ID,
			Name: 	todolist.Name,
		}
	}

	return c.JSON(200, res)
}

func (h *TodoListHandler) GetTodoListsByID(c echo.Context) error {
    id := c.Param("id")
    if id == "" {
        return c.JSON(400, map[string]string{"message": "user_id is required"})
    }

	var todoList model.TodoList
	err := h.db.Get(&todoList, "SELECT * FROM todo_lists WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return c.JSON(500, map[string]string{"message": "Internal Server Error"})
	}

	return c.JSON(200, todoList)
}

type CreateTodoListRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateTodoListResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

//TODOリスト作成
func (h *TodoListHandler) CreateTodoList(c echo.Context) error {
	userID := c.Get("user_id").(int)

	req := new(CreateTodoListRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, map[string]string{"message": "Bad Request"})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.JSON(400, map[string]string{"message": "Bad Request"})
	}

	res, err := h.db.Exec("INSERT INTO todo_lists (user_id, name) VALUES (?, ?)", userID, req.Name)
	if err != nil {
		return c.JSON(500, map[string]string{"message": "Internal Server Error"})
	}

	id, _ := res.LastInsertId()
	return c.JSON(201, CreateTodoListResponse{
		ID: 	int(id),
		Name: 	req.Name,
	})
}

type UpdateTodoListRequest struct {
	Name string `json:"name"`
}

//TodoListを編集する
func (h *TodoListHandler) UpdateTodoList(c echo.Context) error {
	id := c.Param("id")
    if id == "" {
        return c.JSON(400, map[string]string{"message": "user_id is required"})
    }

	req := new(UpdateTodoListRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, map[string]string{"message": "Bad Request"})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.JSON(400, map[string]string{"message": "Bad Request"})
	}

	_, err := h.db.Exec("UPDATE todo_lists SET name = ? WHERE id = ?", req.Name, id)
	if err != nil {
		return c.JSON(500, map[string]string{"message": "Internal Server Error"})
	}

	return c.JSON(200, nil)
}