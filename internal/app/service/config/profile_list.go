package config

import (
	"github.com/BurntSushi/toml"

	"github.com/nitschmann/cfdns/internal/pkg/model"
)

type ProfileListService interface {
	Get() (*map[string]*model.ConfigProfile, error)
}

type ProfileListServiceObj struct {
	configFilepath string
}

func NewProfileListService(configFilepath string) *ProfileListServiceObj {
	return &ProfileListServiceObj{configFilepath: configFilepath}
}

func (serv *ProfileListServiceObj) Get() (map[string]*model.ConfigProfile, error) {
	list := make(map[string]*model.ConfigProfile)
	_, err := toml.DecodeFile(serv.configFilepath, &list)

	return list, err
}
