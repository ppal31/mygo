package types

type ResourceType string

const (
	BOOK  ResourceType = "BOOK"
	BLOG               = "BLOG"
	VIDEO              = "VIDEO"
)

type Resource struct {
	Id          int64        `db:"id"           json:"id"`
	Rtype       ResourceType `db:"rtype"        json:"rtype"`
	Name        string       `db:"name"         json:"name"`
	DisplayName string       `db:"display_name" json:"display_name"`
	Url         string       `db:"url"          json:"url"`
	AuthorId    string       `db:"author_id"    json:"author_id"`
}

type Author struct {
	Id   int64  `db:"id"   json:"id"`
	Name string `db:"name" json:"name"`
}

// Params stores query parameters.
type Params struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func (p Params) Offset() int {
	return p.Page * p.SizeOrDefault()
}

func (p Params) SizeOrDefault() int {
	if p.Size == 0 {
		return 10
	}
	return p.Size
}
