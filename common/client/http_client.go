package client

import (
	"github.com/gogf/gf/v2/frame/g"
)

func InitDefaultHttpClient() {
	
	s := g.Server()
	
	register(s)
	
	s.SetPort(8000)
	s.Run()
}
