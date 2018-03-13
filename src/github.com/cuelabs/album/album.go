package album

import "net/url"

type Album struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
	Url url.URL `json:"url"`
}

func (a *Album) getName(s string) (string, error) {
	return a.Name, nil
}

