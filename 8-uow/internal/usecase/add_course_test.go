package usecase

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/janainamai/learning-go/8-uow/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestAddCourse(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbt.Exec("drop table if exists `courses`;")
	dbt.Exec("drop table if exists `categories`;")

	dbt.Exec("create table if not exists `categories` (id int primary key auto_increment, name varchar(255) not null);")
	dbt.Exec("create table if not exists `courses` (id int primary key auto_increment, name varchar(255) not null, category_id integer not null, foreign key (category_id) references categories(id));")

	input := InputUseCase{
		CategoryName:     "Category 1",
		CourseName:       "Course 1",
		CourseCategoryID: 1,
	}

	ctx := context.Background()
	useCase := NewAddCourseUseCase(repository.NewCourseRepository(dbt), *repository.NewCategoryRepository(dbt))
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
