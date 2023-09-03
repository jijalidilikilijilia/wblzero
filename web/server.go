package web

import (
	"net/http"

	"github.com/jijalidilikilijilia/wblzero/handlers"
	"github.com/jijalidilikilijilia/wblzero/internal/cache"
	"github.com/jijalidilikilijilia/wblzero/internal/database"
	"github.com/jijalidilikilijilia/wblzero/internal/nats"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router      *gin.Engine
	cache       *cache.CacheHandler //Redis - cache
	db          *database.DBHandler // Db - Postgres/gorm
	natsHandler *nats.NatsHandler   // Nats
}

func NewServer(cache *cache.CacheHandler, db *database.DBHandler, natsHandler *nats.NatsHandler) *Server {
	router := gin.Default()

	s := &Server{
		router:      router,
		cache:       cache,
		db:          db,
		natsHandler: natsHandler,
	}

	s.router.LoadHTMLGlob("../templates/*")
	s.setupRoutes()
	return s
}

func (s *Server) Run(address string) error {
	return s.router.Run(address)
}

func (s *Server) setupRoutes() {
	s.router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	s.router.POST("/get-order", func(c *gin.Context) {
		uid := c.PostForm("uid")
		c.Redirect(http.StatusFound, "order/"+uid)
	})
	s.router.GET("/order/:uid", handlers.GetOrderHandler(s.db, s.cache))
}
