package cfg

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type File struct {
	Type string
	Name string
	File string
}

func FromFile(f File) (*viper.Viper, error) {
	v := viper.New()
	if f.File != "" {
		v.SetConfigFile(f.File)
	} else {
		v.SetConfigName(f.Name)
		v.SetConfigType(f.Type)
	}
	return v, v.ReadInConfig()
}

func ReadFile(file string, out any) (err error) {
	v := viper.New()
	v.SetConfigFile(file)
	if err = v.ReadInConfig(); err != nil {
		return
	}
	if err = mapstructure.Decode(v.AllSettings(), out); err != nil {
		return
	}
	return
}
