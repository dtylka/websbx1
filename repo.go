package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func RepoListTodo() []Todo {
	var todos []Todo

	session := repoHelper.getSessionCopy()
	defer session.Close()

	if err := session.DB("sbx1").C("todo").Find(bson.M{}).All(&todos); err != nil {
		panic(err)
	}

	return todos
}

func RepoFindTodo(id bson.ObjectId) (bool, Todo) {
	t := Todo{}

	session := repoHelper.getSessionCopy()
	defer session.Close()

	mgo.SetStats(true)

	if err := session.DB("sbx1").C("todo").FindId(id).One(&t); err != nil {
		return false, t
	}

	stats := mgo.GetStats()
	log.Printf("Stats: %+v\n", stats)

	return true, t
}

func RepoCreateTodo(t Todo) Todo {
	t.Id = bson.NewObjectId()

	session := repoHelper.getSessionCopy()
	defer session.Close()

	session.DB("sbx1").C("todo").Insert(t)

	return t
}

func RepoDestroyTodo(name string) {
	session := repoHelper.getSessionCopy()
	defer session.Close()

	session.DB("sbx1").C("todo").Remove(bson.M{"name": name})
}
