package main

import (
	"gopkg.in/mgo.v2"
)

type RepoHelper struct {
	session *mgo.Session
}

func NewRepoHelper(s *mgo.Session) *RepoHelper {
	return &RepoHelper{s}
}

func (rh *RepoHelper) getSessionCopy() *mgo.Session {
	return rh.session.Copy()
}
