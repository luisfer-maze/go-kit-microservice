package service

type Service interface {
	ProcessSalute(salute string) string
}
