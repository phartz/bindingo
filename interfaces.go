package main

type IDataService interface {
	GetStatus(string, string) (int, error)
	Insert(string, string, string) error
	Exists(string, string, string) (bool, error)
	Delete(string, string, string) error
}
