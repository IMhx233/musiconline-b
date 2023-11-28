package model

type Music struct {
	ID       int
	Name     string
	Path     string
	Tags     string
	Lrcpath  string
	Artist   string
	Album    string
	Length   int
	isdelete bool
	hot      int
}
