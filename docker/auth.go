package docker

import (
	"encoding/base64"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
)

type AuthConfig struct {
	Auths AuthCredentials `yaml:"auths"`
}

type AuthCredentials struct {
	Registry string `yaml:"registry"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func LoadAuth(filePath string) (*AuthCredentials, error) {
	authConfig := &AuthConfig{}
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, authConfig)
	if err != nil {
		return nil, err
	}
	return &authConfig.Auths, nil
}

// EncodeAuthToBase64 encodes Docker auth credentials to Base64
func EncodeAuthToBase64(auth *AuthCredentials) (string, error) {
	authData := map[string]string{
		"username": auth.Username,
		"password": auth.Password,
	}
	authBytes, err := json.Marshal(authData)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(authBytes), nil
}
