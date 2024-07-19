package file

import (
	"fmt"
	"github.com/rogelioConsejo/golibs/helpers"
	"os"
	"testing"
)

func TestPersistence(t *testing.T) {
	const routines = 10
	t.Cleanup(func() {
		err := makePersistence().(*persistence).destroy()
		if err != nil {
			fmt.Printf("could not destroy: %s", err.Error())
		}
	})
	/* make parallel tests with the defined amount of routines */
	t.Run("parallel", func(t *testing.T) {
		t.Parallel()
		for i := 0; i < routines; i++ {
			t.Run(fmt.Sprintf("routine-%d", i), testPersistence)
		}
	})

}

func testPersistence(t *testing.T) {
	t.Helper()
	p := makePersistence()
	p.Lock()
	defer p.Unlock()
	var text = helpers.MakeRandomString(10)
	var dto = DTO{}
	err := p.Get(&dto)
	if err != nil {
		t.Fatalf("could not get: %s", err.Error())
	}
	if dto.Data == nil {
		dto.Data = make(map[string]interface{})
	}
	dto.Data[text] = true
	err = p.Save(dto)
	if err != nil {
		t.Fatalf("could not save: %s", err.Error())
	}

	var dto2 = DTO{}
	err = p.Get(&dto2)
	if err != nil {
		t.Fatalf("could not get: %s", err.Error())
	}
	if _, ok := dto2.Data[text]; !ok {
		t.Fatalf("could not find the text in the dto")
	}
}

func makePersistence() Persistence {
	var fileName Name = "filename.test"
	var p = GetPersistence(fileName)
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
	Data map[string]interface{}
}
