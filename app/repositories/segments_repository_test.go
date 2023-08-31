package repositories

import (
	"os"
	"testing"

	"github.com/Tulenien/Avito-backend-trainee-assignment-2023/app/models"
	"github.com/Tulenien/Avito-backend-trainee-assignment-2023/app/database"
)

func TestSegmentsRepository(t *testing.T) {
	database.ConnectDB()
	cursor, err := database.GetConnection()
	if err != nil {
		os.Exit(1)
	}
	ur := NewSegmentsRepository(cursor)

	name := "AVITO_VOICE_MESSAGES"
	segment_add := models.Segments{Name: name}
	var segmentID uint

	t.Run("Add segment to db", func (t *testing.T) {
		segment_add, err := ur.Create(segment_add)

		if err != nil {
			t.Error(err)
		} else {
			segmentID = segment_add.ID
			t.Logf("Segment %v created", segment_add)
		}
	})

	t.Run("Find segment by id", func(t *testing.T) {
		_, err := ur.FindByID(segmentID)

		if err != nil {
			t.Errorf("Segment not found")
			t.Error(err)
		}
	})

	t.Run("Find segment by name", func(t *testing.T) {
		_, err := ur.FindByName(name)

		if err != nil {
			t.Errorf("Segment not found")
			t.Error(err)
		}
	})
}