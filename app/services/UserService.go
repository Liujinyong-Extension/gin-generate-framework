package services

import "gin-generate-framework/app/models"

type UserService struct {
	BaseService[models.User]
}
