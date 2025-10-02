package pulid

import (
	"database/sql/driver"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
)

// ID implements a PULID - a prefixed ULID.
type ID string

// The default entropy source.
var defaultEntropySource *ulid.LockedMonotonicReader

func init() {
	// Seed the default entropy source.
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	defaultEntropySource = &ulid.LockedMonotonicReader{MonotonicReader: ulid.Monotonic(rng, 0)}
}

// newULID returns a new ULID for time.Now() using the default entropy source.
func newULID() ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(time.Now()), defaultEntropySource)
}

// MustNew returns a new PULID for time.Now() given a prefix. This uses the default entropy source.
func MustNew(prefix string) ID { return ID(prefix + fmt.Sprint(newULID())) }

func (u *ID) String() string {
	return string(*u)
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *ID) UnmarshalGQL(v interface{}) error {
	return u.Scan(v)
}

// MarshalGQL implements the graphql.Marshaler interface
func (u ID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(u)))
}

// Scan implements the Scanner interface.
func (u *ID) Scan(src interface{}) error {
	if src == nil {
		return fmt.Errorf("pulid: expected a value")
	}

	q, ok := src.(ID)
	if ok {
		*u = ID(q)
		return nil
	}

	s, ok := src.(string)
	if !ok {
		return fmt.Errorf("pulid: expected a string")
	}
	*u = ID(s)
	return nil
}

// Value implements the driver Valuer interface.
func (u ID) Value() (driver.Value, error) {
	return string(u), nil
}
