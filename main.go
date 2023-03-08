package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"example.com/lbcexercise/controllers"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	Error string `json:"error"`
}

type fizzbuzzResponse struct {
	Result []string `json:"result"`
}

type fizzbuzzRequest struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

func setHandlers(r *echo.Echo) {
	r.POST("/fizzbuzz", func(c echo.Context) error {
		req := new(fizzbuzzRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusInternalServerError, errorResponse{Error: "Failed to parse request: " + err.Error()})
		}

		result, err := controllers.Fizzbuzzcustom(req.Int1, req.Int2, req.Limit, req.Str1, req.Str2)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		}

		return c.JSON(http.StatusOK, fizzbuzzResponse{Result: result})
	})
}

func main() {
	port := flag.Int("p", 8090, "Provide a port number")
	flag.Parse()
	e := echo.New()
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	setHandlers(e)
	if err := e.Start(fmt.Sprintf(":%d", *port)); err != http.ErrServerClosed {
		log.Println("server closed unexpectedly: ", err.Error())
	}
}
