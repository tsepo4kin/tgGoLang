package storage

import (
	"errors"
	"crypto/sha1"
	"fmt"
	"io"
	"tgGoLang/lib/error"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

var ErrNoSavedPages = errors.New("no saved page")

type Page struct {
	URL string
	UserName string
}

func (p *Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", errorWrapper.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", errorWrapper.Wrap("can't calculate hash", err)
	}

	return fmt.Sprintf("%X", h.Sum(nil)), nil
}