package services

import "gin-generate-framework/app/models"

type TestService struct {
	BaseService[models.Test]
}
