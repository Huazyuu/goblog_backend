package service

import (
	"gvb_server/service/imgServer"
	"gvb_server/service/userServer"
)

type ServiceGroup struct {
	ImageService imgServer.ImageService
	UserService  userServer.UserService
}

var ServiceApp = new(ServiceGroup)
