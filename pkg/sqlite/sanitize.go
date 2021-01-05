package sqlite

import (
	"fmt"

	"github.com/sirupsen/logrus"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/operator-framework/operator-registry/pkg/registry"
)

type SQLSanitizer interface {
	SanitizeDanglingRecords() error
}

// Sanitizer removes a package from the database
type Sanitizer struct {
	store registry.Load
}

var _ SQLSanitizer = &Sanitizer{}

func NewSQLSanitizerForDanglingRecords(store registry.Load) *Sanitizer {
	return &Sanitizer{
		store: store,
	}
}

func (d *Sanitizer) SanitizeDanglingRecords() error {
	log := logrus.WithField("sanitize", "table data")

	log.Info("sanitizing tables")

	var errs []error

	//errs = append(errs, fmt.Errorf("error removing operator package"))

	if err := d.store.Sanitize(); err != nil {
		errs = append(errs, fmt.Errorf("error sanitizing: %s", err))
	}

	return utilerrors.NewAggregate(errs)
}
