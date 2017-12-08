package main

type IDataService interface {
	GetStatus(*Credentials) (int, error)
	Insert(*Credentials, string) error
	Exists(*Credentials, string) (bool, error)
	Delete(*Credentials, string) error
}
