package db

import (
	"context"
	"demo_echo/models"
)

func GetAllStudent() (*[]models.Student, error) {
	students := []models.Student{}
	filter := map[string]interface{}{}
	cursor, err := Client.Database(DBName).Collection(Col).Find(
		context.TODO(),
		filter,
	)
	if err != nil {
		return nil, err
	}

	cursor.All(context.TODO(), &students)

	return &students, nil
}
