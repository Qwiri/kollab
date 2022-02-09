package main

import (
	"fmt"
	"github.com/Qwiri/gobby/pkg/gobby"
	"github.com/Qwiri/gobby/pkg/validate"
	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

const (
	RoleEditor gobby.Role = 1 << iota
)

func main() {
	app := fiber.New(fiber.Config{})

	g := gobby.New(app)
	g.AppVersion = "1.1.1"

	g.MustOn(func(event *gobby.LobbyCreate) {
		log.WithField("lobby", event.Lobby).Info("Lobby created. Making lobby password-less.")
		event.Lobby.Secret = ""
	})

	g.MustOn(func(event *gobby.MessageReceive) {
		log.Debugf("[->] (%s): %s [%+v]", event.Sender.Name, event.Message.Command, event.Message.Args)
	})

	g.MustOn(func(event *gobby.Join) {
		log.Infof("[+] %s joined lobby %s.", event.Client.Name, event.Lobby.ID)
		// set default role
		event.Client.Role = RoleEditor
	})

	g.Handle("foo", &gobby.Handler{
		Validation: validate.Schemes{
			// foo bar 1
			validate.String().Min(3).Max(16).As("user"),
			validate.Number().RangeClosed(1, 10).As("duration"),
		},
		Handler: func(event *gobby.Handle) error { // TODO: add replyTo to gobby.Handle
			user := event.Args.String("user")
			dura := event.Args.Number("duration")

			return event.Message.ReplyTo(event.Client,
				gobby.NewBasicMessage("FOO", fmt.Sprintf("user: %s, v: %v", user, dura)))
		},
	})

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":8082"); err != nil {
			log.WithError(err).Error("cannot listen on port 8082")
			sc <- syscall.SIGTERM
		}
	}()

	<-sc

	log.Info("Stopping webserver")
	if err := app.Shutdown(); err != nil {
		log.WithError(err).Warn("Cannot gracefully shutdown webserver")
	}
}
