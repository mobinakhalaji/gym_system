package services

import (
	"context"
	"errors"

	"book.api/train4/internal/model"
	"book.api/train4/internal/repositories"
)

func CreateMember(member model.Member) (model.Member, error) {
	if member.Age < 18 {
		return model.Member{},
			errors.New("member must be at least 18 years old")
	}
	newMember, err := repositories.CreateMember(member)
	if err != nil {
		return model.Member{}, err
	}
	return newMember, nil
}

func GetMembers() ([]model.Member, error) {
	return repositories.GetMembers()
}
func GetMember(id int) (model.Member, error) {
	return repositories.GetMember(id)
}

func ChangeMember(id int, member model.Member) (model.Member, error) {
	if member.Age < 18 {
		return model.Member{}, errors.New("member must be at least 18 years old")
	}
	return repositories.ChangeMember(context.Background(), int64(id), member)
}

func DeleteMember(id int) error {
	return repositories.DeleteMember(context.Background(), int64(id))
}

func CreateTrainer(trainer model.Trainer) (model.Trainer, error) {
	if trainer.Specialty == "body buillding" {
		return model.Trainer{},
			errors.New("trainer is for body billding")
	}
	newTrainer, err := repositories.CreateTrainer(trainer)
	if err != nil {
		return model.Trainer{}, err
	}
	return newTrainer, nil
}
func GetTrainers() ([]model.Trainer, error) {
	return repositories.GetTrainers()
}
func GetTrainer(id int) (model.Trainer, error) {
	return repositories.GetTrainer(id)
}

func ChangeTrainer(id int, trainer model.Trainer) (model.Trainer, error) {
	return repositories.ChangeTrainer(context.Background(), int64(id), trainer)
}

func DeleteTrainer(id int) error {
	return repositories.DeleteTrainer(context.Background(), int64(id))
}

func CreateClass(ctx context.Context, class model.Class) (model.Class, error) {
	exists, err := repositories.TrainerExists(ctx, class.TrainerID)

	if err != nil {
		return model.Class{}, err
	}

	if !exists {
		return model.Class{}, errors.New("trainer not found")
	}

	newClass, err := repositories.CreateClassRepository(ctx, class)

	if err != nil {
		return model.Class{}, err
	}

	return newClass, nil
}
func ShowClasses() ([]model.Class, error) {
	return repositories.ShowClasses()
}
func ShowClass(id int) (model.Class, error) {
	return repositories.ShowClass(id)
}

func Register(memberID int, classID int) (model.Register, error) {
	memberExists, err := repositories.MemberExists(context.Background(), memberID)
	if err != nil {
		return model.Register{}, err
	}
	if !memberExists {
		return model.Register{}, errors.New("member not found")
	}

	classExists, err := repositories.ClassExists(context.Background(), classID)
	if err != nil {
		return model.Register{}, err
	}
	if !classExists {
		return model.Register{}, errors.New("class not found")
	}

	class, err := repositories.ShowClass(classID)
	if err != nil {
		return model.Register{}, err
	}

	members, err := repositories.MemberOfthisClass(classID)
	if err != nil {
		return model.Register{}, err
	}

	if len(members) >= class.Capacity {
		return model.Register{}, errors.New("class is full")
	}

	register := model.Register{
		MemberID: memberID,
		ClassID:  classID,
	}

	newRegister, err := repositories.CreateRegister(context.Background(), register)
	if err != nil {
		return model.Register{}, err
	}

	return newRegister, nil
}

func ClassesOfthisMember(id int) ([]model.Class, error) {
	return repositories.ClassesOfthisMember(id)
}

func MemberOfthisClass(id int) ([]model.Member, error) {
	class, err := repositories.ShowClass(id)
	if err != nil {
		return nil, err
	}

	members, err := repositories.MemberOfthisClass(id)
	if err != nil {
		return nil, err
	}

	if len(members) > class.Capacity {
		return nil, errors.New("class capacity exceeded")
	}

	return members, nil
}

func SearchMembers(ctx context.Context, name string) ([]model.Class, error) {
	return repositories.SearchMembers(name)
}