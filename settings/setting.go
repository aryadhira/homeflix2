package settings

import "github.com/joho/godotenv"

type Setting struct{}

func (setting *Setting) GetConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
