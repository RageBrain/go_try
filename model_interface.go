package main

type Model interface {
  documentsList() []string
  documentExist() bool
}
