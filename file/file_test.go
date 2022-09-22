package file

import (
	"fmt"
	"os"
	"testing"
)

func TestPersistence(t *testing.T) {
	p := makePersistence(t)
	const text = "some text"
	var smt DTO = DTO{SomeText: text}
	err := p.Save(smt)
	if err != nil {
		t.Fatalf(err.Error())
	}
	var retrievedSomething DTO
	p2 := makePersistence(t)
	err = p2.Get(&retrievedSomething)

	if savedText, retrievedText := smt.Text(), retrievedSomething.Text(); savedText != retrievedText {
		t.Errorf("did not retrieve 'DTO' correctly: expected (%s) -- retrieved (%s)", savedText, retrievedText)
	}

	err = p.(persistence).destroy()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func makePersistence(t *testing.T) Persistence {
	var fileName Name = "filename.test"
	var p, err = New(fileName)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return p
}

func (p persistence) destroy() error {
	fileName := p.file.Name()
	err := p.file.Close()
	if err != nil {
		return fmt.Errorf("could not close file: %w", err)
	}
	err = os.Remove(fileName)
	if err != nil {
		return fmt.Errorf("could not remove file: %w", err)
	}
	return nil
}

type DTO struct {
	SomeText string `json:"some_text"`
}

func (s DTO) Text() string {
	return s.SomeText
}
