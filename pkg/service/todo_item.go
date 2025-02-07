package service

import (
	todo "Notes_GoRest"
	"Notes_GoRest/pkg/repository"
	"errors"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTotoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId int, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, errors.New("failed to get user list")
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId int, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId int, itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId int, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) UpdateItem(userId int, itemId int, input todo.UpdateItemInput) error {
	return s.repo.UpdateItem(userId, itemId, input)
}
