package storage

type Storage interface {
	Save(p *Page) error
	PickRandom(userName) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL string
	UserName string
}

func (p *Page) Hash() (string, err) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", error.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", error.Wrap("can't calculate hash", err)
	}

	return fmt.Sprintf("%X", h.Sum(nil)), nil
}