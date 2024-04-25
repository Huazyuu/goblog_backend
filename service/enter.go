package service

import "gvb_server/service/img_service"

type ServiceGroup struct {
	ImageService img_service.ImageService
}

var ServiceApp = new(ServiceGroup)
