package file

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestPersistence(t *testing.T) {
	p := makePersistence()
	const text = "some text"
	var smt DTO = DTO{SomeText: text}
	err := p.Save(smt)
	if err != nil {
		t.Fatalf(err.Error())
	}
	var retrievedSomething DTO
	p2 := makePersistence()
	err = p2.Get(&retrievedSomething)

	if savedText, retrievedText := smt.Text(), retrievedSomething.Text(); savedText != retrievedText {
		t.Errorf("did not retrieve 'DTO' correctly: expected (%s) -- retrieved (%s)", savedText, retrievedText)
	}

	const otherText = "some other text"
	smt.SomeText = otherText
	err = p.Save(smt)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = p2.Get(&retrievedSomething)
	if err != nil && err != io.EOF {
		t.Fatalf(err.Error())
	}
	if retrievedSomething.Text() != otherText {
		t.Fatalf("did not update 'something'")
	}

	err = p.(*persistence).destroy()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func makePersistence() Persistence {
	var fileName Name = "filename.test"
	var p = New(fileName)
	return p
}

func (p *persistence) destroy() error {
	err := os.Remove(string(p.fileName))
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
