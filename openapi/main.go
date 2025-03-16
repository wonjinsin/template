package main

import (
	"fmt"
	"net/http"

	"github.com/wonjinsin/template/openapi/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	// 여기에 필요한 의존성을 추가하세요
	// 예: db *sql.DB
}

// ServerInterface 구현
func (s *Server) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"resultCode": "200",
		"resultData": []map[string]interface{}{
			{
				"id":    "1",
				"name":  "Test User",
				"email": "test@example.com",
			},
		},
	})
}

func (s *Server) PostUsers(c echo.Context) error {
	var reqBody api.PostUsersJSONRequestBody
	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"resultCode": "200",
		"resultData": map[string]interface{}{
			"id":    "2",
			"name":  reqBody.Name,
			"email": reqBody.Email,
		},
	})
}

func (s *Server) GetUsersId(c echo.Context, id uint64) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"resultCode": "200",
		"resultData": map[string]interface{}{
			"id":    fmt.Sprintf("%d", id),
			"name":  "Test User",
			"email": "test@example.com",
		},
	})
}

func (s *Server) PutUsersId(c echo.Context, id uint64) error {
	var reqBody api.PutUsersIdJSONRequestBody
	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"resultCode": "200",
		"resultData": map[string]interface{}{
			"id":    fmt.Sprintf("%d", id),
			"name":  reqBody.Name,
			"email": "test@example.com",
		},
	})
}

func (s *Server) DeleteUsersId(c echo.Context, id uint64) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"resultCode": "200",
		"resultMsg":  "User deleted successfully",
	})
}

func main() {
	e := echo.New()

	// 미들웨어 설정
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 서버 인스턴스 생성
	server := &Server{
		// 여기에 필요한 의존성 초기화
		// 예: db: database.Connect(),
	}

	// API 핸들러 등록
	api.RegisterHandlers(e, server)

	// 서버 시작
	e.Logger.Fatal(e.Start(":8080"))
}
