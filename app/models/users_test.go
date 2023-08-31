package models

import (
	"testing"
)

func TestUsersRepository(t *testing.T) {
	user := Users{Segments: nil}

	name1 := "AVITO_VOICE_MESSAGES"
	segment1_add := Segments{Name: name1}


	t.Run("Add segment to user", func (t *testing.T) {
		err := AddSegment(&user, segment1_add)

		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Segment %v added", segment1_add)
		}
	})
}