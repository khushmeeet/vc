package vc

type Commit struct {
	Tree    string
	Parent  string
	Message string
}

type RefValue struct {
	symbolic bool
	value    string
}
