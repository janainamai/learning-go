package usecase

import (
	"context"

	"github.com/janainamai/learning-go/8-uow/internal/entity"
	"github.com/janainamai/learning-go/8-uow/internal/repository"
	"github.com/janainamai/learning-go/8-uow/pkg/uow"
)

type (
	InputUseCaseUow struct {
		CategoryName     string
		CourseName       string
		CourseCategoryID int
	}

	AddCourseUseCaseUow struct {
		Uow uow.Uow
	}
)

func NewAddCourseUseCaseUow(uow uow.Uow) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{
		Uow: uow,
	}
}

func (a *AddCourseUseCaseUow) Execute(ctx context.Context, input InputUseCase) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		// tudo o que colocarmos aqui, será feito um begin e um commit/rollback
		categoryRepo := a.getCategoryRepository(ctx)
		category := entity.Category{
			Name: input.CategoryName,
		}
		err := categoryRepo.Insert(ctx, category)
		if err != nil {
			return err
		}

		courseRepo := a.getCourseRepository(ctx)
		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}
		err = courseRepo.Insert(ctx, course)
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *AddCourseUseCaseUow) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}

	return repo.(repository.CategoryRepositoryInterface)
}

func (a *AddCourseUseCaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}

	return repo.(repository.CourseRepositoryInterface)
}
