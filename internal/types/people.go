package types

type UriID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type People struct {
	ID        string `json:"id"`
	Nickname  string `json:"nickname"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	Stack     string `json:"stack"`
}

func (p *People) SearchValue() string {
	return p.Name + p.Nickname + p.Birthdate
}
