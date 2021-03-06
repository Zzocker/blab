package adapters

import (
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/core/ports"
	"github.com/Zzocker/blab/pkg/datastore"
	"github.com/Zzocker/blab/pkg/errors"
)

// CreateUserStore create new userstore adapter to be used by userCore
func CreateUserStore(conf config.DatastoreConf) (ports.UserStorePort, errors.E) {
	db, err := datastore.NewSmartDS(conf)
	if err != nil {
		return nil, err
	}
	return &userStore{db: db}, nil
}

// CreateBookStore creates a new bookstore adapter to be used by bookcore
func CreateBookStore(conf config.DatastoreConf) (ports.BookStorePort, errors.E) {
	db, err := datastore.NewSmartDS(conf)
	if err != nil {
		return nil, err
	}
	return &bookStore{db: db}, nil
}

// CreateCommentStore creates a new commentStore adapter to be used by commentCore
func CreateCommentStore(conf config.DatastoreConf) (ports.CommentStore, errors.E) {
	db, err := datastore.NewSmartDS(conf)
	if err != nil {
		return nil, err
	}
	return &commentStore{db: db}, nil
}

// CreateOAuthStore creates a new oauthStore adapter to be used by oauthCore
func CreateOAuthStore(conf config.DatastoreConf) (ports.OAuthStore, errors.E) {
	db, err := datastore.NewDumbDS(conf)
	if err != nil {
		return nil, err
	}
	return &oauthStore{db: db}, nil
}
