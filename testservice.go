package main

type TestService struct {
}

func (t TestService) GetStatus(credentials *Credentials) (output int, err error) {
	return 1, nil
}

func (t TestService) SetCredentials(credentials *Credentials) error {
	return nil
}

func (t TestService) Insert(credentials *Credentials, value string) (err error) {
	return nil
}

func (t TestService) Exists(credentials *Credentials, value string) (exists bool, err error) {
	return true, nil
}
func (t TestService) Delete(credentials *Credentials, value string) (err error) {
	return nil
}
