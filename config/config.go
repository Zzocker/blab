package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// C represents loaded config from yaml file
type C struct {
	Port int `yaml:"port"`
	//
	Core Core `yaml:"core"`
}

type Core struct {
	User    UserCoreConf    `yaml:"user"`
	OAuth   OAuthCoreConf   `yaml:"oauth"`
	Book    BookCoreConf    `yaml:"book"`
	Comment CommentCoreConf `yaml:"coment"`
}

type CommentCoreConf struct {
	CommentStoreConf DatastoreConf `yaml:"commentStoreConf"`
}
type OAuthCoreConf struct {
	TokenStoreConf DatastoreConf `yaml:"tokenStoreConf"`
}
type UserCoreConf struct {
	UserStoreConf DatastoreConf `yaml:"userStoreConf"`
}

type BookCoreConf struct {
	BookStoreConf DatastoreConf `yaml:"bookStoreConf"`
}

// DatastoreConf : config for connecting database
type DatastoreConf struct {
	URL        string `yaml:"url"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
}

// Init load config file from confPath
// os.Exit(1) is called if any error occurred while reading config file
func Init(confPath string) *C {
	f, err := os.Open(confPath)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	var conf C
	if err = yaml.NewDecoder(f).Decode(&conf); err != nil {
		os.Exit(1)
	}
	return &conf
}
