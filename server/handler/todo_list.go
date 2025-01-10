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
	g.GET("/todolists", h.GetTodoListsByUser)
	g.POST("/todolists", h.CreateTodoList)
}

type GetTodoListsByUserResponse struct {
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
			return c.JSON(200, []GetTodoListsByUserResponse{})
		}
		return c.JSON(500, map[string]string{"message": "Internal Server Error"})
	}

	res := make([]GetTodoListsByUserResponse, len(todoLists))
	for i, todolist := range todoLists {
		res[i] = GetTodoListsByUserResponse{
			ID:		todolist.ID,
			Name: 	todolist.Name,
		}
	}

	return c.JSON(200, res)
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