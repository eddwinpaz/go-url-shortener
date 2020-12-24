package shortener

import (
	"errors"
	"time"

	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
)

var (
	// ErrRedirectNotFound Return new Error notifying that redirect was not able to be found.
	ErrRedirectNotFound = errors.New("Redirect not found")
	// ErrRedirectInvalid Return new Error notifying that redirect was invalid.
	ErrRedirectInvalid = errors.New("Redirect invalid")
)

type redirectService struct {
	redirectRepo RedirectRepository
}

// NewRedirectService return repo interface
func NewRedirectService(redirectRepo RedirectRepository) RedirectService {
	return &redirectService{redirectRepo}
}

func (r *redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.redirectRepo.Store(redirect)
}
