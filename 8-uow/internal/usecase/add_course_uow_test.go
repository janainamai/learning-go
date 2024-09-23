package usecase

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/janainamai/learning-go/8-uow/internal/db"
	"github.com/janainamai/learning-go/8-uow/internal/repository"
	"github.com/janainamai/learning-go/8-uow/pkg/uow"
	"github.com/stretchr/testify/assert"
)

func TestAddCourseUow(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbt.Exec("drop table if exists `courses`;")
	dbt.Exec("drop table if exists `categories`;")

	dbt.Exec("create table if not exists `categories` (id int primary key auto_increment, name varchar(255) not null);")
	dbt.Exec("create table if not exists `courses` (id int primary key auto_increment, name varchar(255) not null, category_id integer not null, foreign key (category_id) references categories(id));")

	ctx := context.Background()
	uow := uow.NewUow(dbt)
	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})
	uow.Register("CourseRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})

	input := InputUseCase{
		CategoryName:     "Category 1",
		CourseName:       "Course 1",
		CourseCategoryID: 2, //simulando erro para verificar se o banco de dados ficou com dados inconsistentes
	}

	useCase := NewAddCourseUseCaseUow(*uow)
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
