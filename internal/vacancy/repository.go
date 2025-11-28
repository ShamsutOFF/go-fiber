package vacancy

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type VacancyRepository struct {
	Dbpool       *pgxpool.Pool
	CustomLogger *zerolog.Logger
}

func NewVacancyRepository(
	dbpool *pgxpool.Pool,
	customLogger *zerolog.Logger,
) *VacancyRepository {
	return &VacancyRepository{
		Dbpool:       dbpool,
		CustomLogger: customLogger,
	}
}

func (repo *VacancyRepository) addVacancy(form *VacancyCreateForm) error {
	query := `INSERT INTO vacancies (email, role, company, salary, type, location, createdat) VALUES(@email, @role, @company, @salary, @type, @location, @createdat)`
	args := pgx.NamedArgs{
		"email":     form.Email,
		"role":      form.Role,
		"company":   form.Company,
		"salary":    form.Salary,
		"type":      form.Type,
		"location":  form.Location,
		"createdat": time.Now(),
	}
	_, err := repo.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("Невозможно создать вакансию %w", err)
	}
	return nil
}

func (repo *VacancyRepository) GetCountAll() int {
	query := `SELECT COUNT(*) FROM vacancies`
	var count int
	repo.Dbpool.QueryRow(context.Background(), query).Scan(&count)
	return count
}

func (repo *VacancyRepository) GetAllVacancies(limit, offset int) ([]Vacancy, error) {
	query := `SELECT * FROM vacancies ORDER BY createdat DESC LIMIT @limit OFFSET @offset`
	args := pgx.NamedArgs{
		"limit":  limit,
		"offset": offset,
	}
	if limit == 0 {
		query = `SELECT * FROM vacancies ORDER BY createdat DESC`
	}
	rows, err := repo.Dbpool.Query(context.Background(), query, args)
	if err != nil {
		return nil, fmt.Errorf("Невозможно получить вакансии %w", err)
	}
	vacancies, err := pgx.CollectRows(rows, pgx.RowToStructByName[Vacancy])
	if err != nil {
		return nil, fmt.Errorf("Невозможно получить вакансии %w", err)
	}
	defer rows.Close()
	return vacancies, nil
}
