package group

import (
	"my_arch/services/contact/internal/domain/contact"
	"my_arch/services/contact/internal/domain/group"
	"my_arch/services/contact/internal/repository"
	usecase "my_arch/services/contact/internal/useCase"
)

type GroupUsecase struct {
	gropRepo repository.InterfaceGroupRepository
}

func NewGroupUseCase(repo repository.InterfaceGroupRepository) usecase.InterfaceGroupUsecase {
	return &GroupUsecase{
		gropRepo: repo,
	}
}

func (gs *GroupUsecase) CreateGroup(group.Group) (int, error) {
	return 0, nil
}

func (gs *GroupUsecase) GetGroup() (group.Group, error) {
	return group.Group{}, nil
}

func (gs *GroupUsecase) InsertContact(contact contact.Contact, id int) error {
	return nil
}
