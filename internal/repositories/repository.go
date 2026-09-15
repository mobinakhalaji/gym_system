package repositories

import (
	"context"

	"book.api/train4/internal/database"
	"book.api/train4/internal/model"
	"github.com/jackc/pgx/v5"
)

func CreateMember(member model.Member) (model.Member, error) {

	query := `
		INSERT INTO members(name, age, active)
		VALUES($1, $2, $3)
		RETURNING id
	`

	err := database.DB.QueryRow(
		context.Background(),
		query,
		member.Name,
		member.Age,
		member.Active,
	).Scan(&member.ID)

	if err != nil {
		return model.Member{}, err
	}

	return member, nil
}

func GetMembers() ([]model.Member, error) {
	rows, err := database.DB.Query(
		context.Background(),
		`SELECT id, name, phone, age, active FROM members ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	members := make([]model.Member, 0)
	for rows.Next() {
		var member model.Member
		if err := rows.Scan(&member.ID, &member.Name, &member.Phone, &member.Age, &member.Active); err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}

func GetMember(id int) (model.Member, error) {
	var member model.Member

	err := database.DB.QueryRow(
		context.Background(),
		`SELECT id, name, phone, age, active
         FROM members
         WHERE id = $1`,
		id,
	).Scan(
		&member.ID,
		&member.Name,
		&member.Phone,
		&member.Age,
		&member.Active,
	)

	if err != nil {
		return model.Member{}, err
	}

	return member, nil
}

func ChangeMember(
	ctx context.Context,
	id int64,
	member model.Member,
) (model.Member, error) {

	var updatedMember model.Member

	err := database.DB.QueryRow(ctx,
		`UPDATE members
		SET
			name = $1,
			phone = $2,
			age = $3,
			active = $4
		WHERE id = $5
		RETURNING id, name, phone, age, active
		`,
		member.Name,
		member.Phone,
		member.Age,
		member.Active,
		id,
	).Scan(
		&updatedMember.ID,
		&updatedMember.Name,
		&updatedMember.Phone,
		&updatedMember.Age,
		&updatedMember.Active,
	)

	if err != nil {
		return model.Member{}, err
	}

	return updatedMember, nil
}

func DeleteMember(ctx context.Context, id int64) error {
	result, err := database.DB.Exec(ctx, `DELETE FROM members WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func CreateTrainer(trainer model.Trainer) (model.Trainer, error) {

	query := `
		INSERT INTO trainers(name,Specialty,phone)
		VALUES($1, $2, $3)
		RETURNING id
	`

	err := database.DB.QueryRow(
		context.Background(),
		query,
		trainer.Name,
		trainer.Specialty,
		trainer.Phone,
	).Scan(&trainer.ID)

	if err != nil {
		return model.Trainer{}, err
	}

	return trainer, nil
}
func GetTrainers() ([]model.Trainer, error) {
	rows, err := database.DB.Query(
		context.Background(),
		`SELECT * FROM trainers ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	trainers := make([]model.Trainer, 0)
	for rows.Next() {
		var trainer model.Trainer
		if err := rows.Scan(&trainer.ID, &trainer.Name, &trainer.Specialty, &trainer.Phone); err != nil {
			return nil, err
		}
		trainers = append(trainers, trainer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return trainers, nil
}

func GetTrainer(id int) (model.Trainer, error) {
	var trainer model.Trainer

	err := database.DB.QueryRow(
		context.Background(),
		`SELECT *
         FROM trainers
         WHERE id = $1`,
		id,
	).Scan(
		&trainer.ID,
		&trainer.Name,
		&trainer.Specialty,
		&trainer.Phone,
	)

	if err != nil {
		return model.Trainer{}, err
	}

	return trainer, nil
}
func ChangeTrainer(
	ctx context.Context,
	id int64,
	trainer model.Trainer,
) (model.Trainer, error) {

	var updatedTrainer model.Trainer

	err := database.DB.QueryRow(ctx,
		`UPDATE trainers
		SET
			name = $1,
			Specialty = $2,
			phone = $3
		WHERE id = $4
		RETURNING id, name,Specialty, phone
		`,
		trainer.Name,
		trainer.Specialty,
		trainer.Phone,

		id,
	).Scan(
		&updatedTrainer.ID,
		&updatedTrainer.Name,
		&updatedTrainer.Specialty,
		&updatedTrainer.Phone,
	)

	if err != nil {
		return model.Trainer{}, err
	}

	return updatedTrainer, nil
}

func DeleteTrainer(ctx context.Context, id int64) error {
	result, err := database.DB.Exec(ctx, `DELETE FROM trainers WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func MemberExists(ctx context.Context, memberID int) (bool, error) {
	var exists bool

	err := database.DB.QueryRow(
		ctx,
		`SELECT EXISTS(
			SELECT 1
			FROM members
			WHERE id = $1
		)`,
		memberID,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func ClassExists(ctx context.Context, classID int) (bool, error) {
	var exists bool

	err := database.DB.QueryRow(
		ctx,
		`SELECT EXISTS(
			SELECT 1
			FROM classes
			WHERE id = $1
		)`,
		classID,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func TrainerExists(ctx context.Context, trainerId int) (bool, error) {
	var exists bool

	err := database.DB.QueryRow(
		ctx,
		`SELECT EXISTS(
			SELECT 1
			FROM trainers
			WHERE id = $1
		)`,
		trainerId,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func CreateRegister(ctx context.Context, register model.Register) (model.Register, error) {
	err := database.DB.QueryRow(
		ctx,
		`INSERT INTO registers (member_id, class_id)
		 VALUES ($1, $2)
		 RETURNING id`,
		register.MemberID,
		register.ClassID,
	).Scan(&register.ID)

	if err != nil {
		return model.Register{}, err
	}

	return register, nil
}
func CreateClassRepository(ctx context.Context, class model.Class) (model.Class, error) {

	err := database.DB.QueryRow(
		ctx,
		`INSERT INTO classes (name, trainer_id, capacity, price)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id`,
		class.Name,
		class.TrainerID,
		class.Capacity,
		class.Price,
	).Scan(&class.ID)

	if err != nil {
		return model.Class{}, err
	}

	return class, nil
}
func ShowClasses() ([]model.Class, error) {
	rows, err := database.DB.Query(
		context.Background(),
		`SELECT c.id, c.name, c.trainer_id, t.name, c.capacity, c.price
		 FROM classes c
		 JOIN trainers t ON c.trainer_id = t.id
		 ORDER BY c.id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	classes := make([]model.Class, 0)
	for rows.Next() {
		var class model.Class
		if err := rows.Scan(
			&class.ID,
			&class.Name,
			&class.TrainerID,
			&class.TrainerName,
			&class.Capacity,
			&class.Price,
		); err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return classes, nil
}
func ShowClass(id int) (model.Class, error) {
	var class model.Class
	err := database.DB.QueryRow(
		context.Background(),
		`SELECT 
			c.id,
			c.name,
			c.trainer_id,
			t.name,
			c.capacity,
			c.price
		 FROM classes c
		 JOIN trainers t ON c.trainer_id = t.id
		 WHERE c.id = $1`,
		id,
	).Scan(
		&class.ID,
		&class.Name,
		&class.TrainerID,
		&class.TrainerName,
		&class.Capacity,
		&class.Price,
	)
	if err != nil {
		return model.Class{}, nil
	}
	return class, nil
}
func ClassesOfthisMember(id int) ([]model.Class, error) {
	rows, err := database.DB.Query(
		context.Background(),
		`SELECT c.id, c.name, c.trainer_id, t.name, c.capacity, c.price
		 FROM registers r
		 JOIN classes c ON r.class_id = c.id
		 JOIN trainers t ON c.trainer_id = t.id
		 WHERE r.member_id = $1
		 ORDER BY c.id`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	classes := make([]model.Class, 0)
	for rows.Next() {
		var class model.Class
		if err := rows.Scan(
			&class.ID,
			&class.Name,
			&class.TrainerID,
			&class.TrainerName,
			&class.Capacity,
			&class.Price,
		); err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return classes, nil
}

func MemberOfthisClass(id int) ([]model.Member, error) {
	rows, err := database.DB.Query(
		context.Background(),
		`SELECT m.id, m.name, m.phone, m.age, m.active
		 FROM registers r
		 JOIN members m ON r.member_id = m.id
		 WHERE r.class_id = $1
		 ORDER BY m.id`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	members := make([]model.Member, 0)
	for rows.Next() {
		var member model.Member
		if err := rows.Scan(
			&member.ID,
			&member.Name,
			&member.Phone,
			&member.Age,
			&member.Active,
		); err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}

func SearchMembers(name string) ([]model.Class, error) {
	rows, err := database.DB.Query(
		context.Background(),
		`SELECT c.id, c.name, c.trainer_id, t.name AS trainer_name, c.capacity, c.price
		 FROM registers r
		 JOIN members m ON r.member_id = m.id
		 JOIN classes c ON r.class_id = c.id
		 JOIN trainers t ON c.trainer_id = t.id
		 WHERE m.name = $1
		 ORDER BY c.id`,
		name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	classes := make([]model.Class, 0)
	for rows.Next() {
		var class model.Class
		if err := rows.Scan(
			&class.ID,
			&class.Name,
			&class.TrainerID,
			&class.TrainerName,
			&class.Capacity,
			&class.Price,
		); err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return classes, nil
}
