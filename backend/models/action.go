package models

type Action interface {
	Run() error
}
