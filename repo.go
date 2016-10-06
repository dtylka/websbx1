package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func RepoListTodo() []Todo {
	var todos []Todo

	session := getSession()
	if err := session.DB("sbx1").C("todo").Find(bson.M{}).All(&todos); err != nil {
		panic(err)
	}

	return todos
}

func RepoFindTodo(id bson.ObjectId) (bool, Todo) {
	t := Todo{}

	session := getSession()
	if err := session.DB("sbx1").C("todo").FindId(id).One(&t); err != nil {
		return false, t
	}

	return true, t
}

func RepoCreateTodo(t Todo) Todo {
	t.Id = bson.NewObjectId()

	session := getSession()
	session.DB("sbx1").C("todo").Insert(t)

	return t
}

func RepoDestroyTodo(name string) {
	session := getSession()
	session.DB("sbx1").C("todo").Remove(bson.M{"name": name})
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s
}
