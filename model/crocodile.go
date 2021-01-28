package model

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

const (
	layout = "2006-01-02"
)

// Crocodile main crocodile model definition
type Crocodile struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Sex       Gender     `json:"sex"`
	BirthDate *Timestamp `json:"date_of_birth"`
	Age       int        `json:"age"`
}

// CrocodileInput DTO used to create new/edit crocs
type CrocodileInput struct {
	Name      string     `json:"name"`
	Sex       Gender     `json:"sex"`
	BirthDate *Timestamp `json:"date_of_birth"`
	Age       int        `json:"age"`
}

// Timestamp holder used to add custom date marshalling/unmarshalling
type Timestamp struct{ *time.Time }

// UnmarshalJSON custom timestamp unmarshaller
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	d, _ := time.Parse(layout, strings.ReplaceAll(string(b), "\"", ""))
	t.Time = &d
	return nil
}

// MarshalJSON custom timestamp marshaller
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format(layout) + `"`), nil
}

// Gender constant used to define available gender types
type Gender string

const (
	// GenderM male crocodile identifier
	GenderM Gender = "M"
	// GenderF female crocodile identifier
	GenderF Gender = "F"
)

// AllGender retrieve all available gender identifiers
var AllGender = []Gender{
	GenderM,
	GenderF,
}

// IsValid define whether the provided gender is a valid constant
func (e Gender) IsValid() bool {
	switch e {
	case GenderM, GenderF:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

// UnmarshalGQL custom gender GraphQL unmarshalling
func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

// MarshalGQL custom gender GraphQL marshalling
func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// MarshalTimestamp custom graphql marshalling for Timestamps
func MarshalTimestamp(t Timestamp) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		res, err := t.MarshalJSON()
		if err != nil {
			panic(err)
		}
		io.WriteString(w, string(res))
	})
}

// UnmarshalTimestamp custom graphql unmarshalling for Timestamps
func UnmarshalTimestamp(v interface{}) (Timestamp, error) {
	if tmpStr, ok := v.(string); ok {
		t, err := time.Parse(layout, tmpStr)
		return Timestamp{&t}, err
	}
	return Timestamp{&time.Time{}}, fmt.Errorf("Unable to parse time")
}
