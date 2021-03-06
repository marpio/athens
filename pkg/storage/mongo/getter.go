package mongo

import (
	"bytes"
	"io/ioutil"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/gomods/athens/pkg/storage"
)

// Get a specific version of a module
func (s *ModuleStore) Get(module, vsn string) (*storage.Version, error) {
	c := s.s.DB(s.d).C(s.c)
	result := &storage.Module{}
	err := c.Find(bson.M{"module": module, "version": vsn}).One(result)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err = storage.ErrVersionNotFound{Module: module, Version: vsn}
		}
		return nil, err
	}
	return &storage.Version{
		RevInfo: storage.RevInfo{
			Version: result.Version,
			Name:    result.Version,
			Short:   result.Version,
			Time:    time.Now(),
		},
		Mod: result.Mod,
		Zip: ioutil.NopCloser(bytes.NewReader(result.Zip)),
	}, nil
}
