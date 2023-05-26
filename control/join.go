package control

import (
	"elf-talk/config"
	"elf-talk/utils"
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func Join(UUID string, Pwd string) (string, error) {
	if Pwd != config.Conf.Server.Password {
		return "", errors.New("P")
	}

	randId, _ := gonanoid.New(7)
	payload := map[string]interface{}{
		"uuid": UUID,
		"rand": randId,
	}
	token, err := utils.Generate(payload)
	if err != nil {
		return "", err
	}

	return token, nil
}
