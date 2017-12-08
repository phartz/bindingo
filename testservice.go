package main

type TestService struct {
}

func (t TestService) GetStatus(id string, credentials string) (output int, err error) {
	return 1, nil
}

func (t TestService) SetCredentials(id string, credentials string) error {
	return nil
}

func (t TestService) Insert(id string, credentials string, value string) (err error) {
	return nil
}

func (t TestService) Exists(id string, credentials string, value string) (exists bool, err error) {
	return true, nil
}
func (t TestService) Delete(id string, credentials string, value string) (err error) {
	return nil
}
