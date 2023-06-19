package repository

import (
	"fmt"
	"ge-rest-api/src/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, id uint, userId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, id uint, userId uint) error
	DeleteTask(id uint, userId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).Find(tasks).Error; err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) GetTaskById(task *model.Task, id uint, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).First(task, id).Error; err != nil {
		return err
	}

	return nil

}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, id uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", id, userId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *taskRepository) DeleteTask(userId uint, id uint) error {
	result := tr.db.Where("id=? AND user_id=?", id, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
