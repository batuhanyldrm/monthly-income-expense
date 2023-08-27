package main

type Api struct {
	service *Service
}

func NewApi(service *Service) Api {
	return Api{
		service: service,
	}
}
