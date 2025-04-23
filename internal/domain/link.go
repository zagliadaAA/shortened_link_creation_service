package domain

type Link struct {
	ID       int
	URL      string
	ShortURL string
}

func NewLink(URL string, ShortURL string) *Link {
	return &Link{
		URL:      URL,
		ShortURL: ShortURL,
	}
}

func (l *Link) SetID(id int) {
	l.ID = id
}
