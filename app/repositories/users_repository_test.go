package repositories

import (
	"os"
	"testing"

	"github.com/Tulenien/Avito-backend-trainee-assignment-2023/app/models"
	"github.com/Tulenien/Avito-backend-trainee-assignment-2023/app/database"
)

func TestUsersRepository(t *testing.T) {
	database.ConnectDB()
	cursor, err := database.GetConnection()
	if err != nil {
		os.Exit(1)
	}
	ur := NewUsersRepository(cursor)

	//segment_temp := models.Segments{Name: "AVITO_VOICE_MESSAGES"}
	user_add := models.Users{Segments: nil}
	var userID uint

	t.Run("Add user to db", func (t *testing.T) {
		user_add, err := ur.Create(user_add)

		if err != nil {
			t.Error(err)
		} else {
			userID = user_add.ID
			t.Logf("User %v created", user_add)
		}
	})

	t.Run("Find user by id", func(t *testing.T) {
		_, err := ur.FindByID(userID)

		if err != nil {
			t.Errorf("User not found")
			t.Error(err)
		}
	})
}