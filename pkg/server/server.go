package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Address string
	Handler *echo.Echo
}

func NewServer(address string, handler *echo.Echo) Server {
	return Server{
		Address: address,
		Handler: handler,
	}
}

func (s Server) ListenAndServe() {
	log.Println("This server is deployed using CI CD")
	go func() {
		err := s.Handler.Start(s.Address)
		if err != nil {
			s.Handler.Logger.Fatalf("ERROR (server): %v\n", err.Error())
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("")
	log.Println("(SERVER): shutting down server...")
}
