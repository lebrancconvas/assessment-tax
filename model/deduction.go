package model

import "github.com/lebrancconvas/assessment-tax/db"

func CreatePersonalDeduction(amount float64) error {
	db := db.GetDB()

	stmt, err := db.Prepare(`
		INSERT INTO deductions(personal_deduction)
		VALUES($1)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(amount)
	if err != nil {
		return err
	}

	return nil
}

func GetPersonalDeduction() (float64, error) {
	db := db.GetDB()

	stmt, err := db.Prepare(`
		SELECT personal_deduction
		FROM deductions
		ORDER BY id DESC
		LIMIT 1
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var amount float64
	err = stmt.QueryRow().Scan(&amount)
	if err != nil {
		return 0, err
	}

	return amount, nil
}
