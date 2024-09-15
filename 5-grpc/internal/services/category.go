package services

import (
	"context"
	"io"

	"github.com/janainamai/learning-go/5-grpc/internal/database"
	"github.com/janainamai/learning-go/5-grpc/internal/pb"
	"google.golang.org/grpc"
)

type (
	categoryService struct {
		pb.UnimplementedCategoryServiceServer // fornece implementação padrão para os métodos ainda não implementados
		categoryDB                            database.Category
	}
)

func NewCategoryService(categoryDB database.Category) *categoryService {
	return &categoryService{categoryDB: categoryDB}
}

func (c *categoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.categoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *categoryService) CreateCategoryStream(stream grpc.ClientStreamingServer[pb.CreateCategoryRequest, pb.CategoryList]) error {
	categories := &pb.CategoryList{}

	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}

		categoryResponse, err := c.categoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResponse.ID,
			Name:        categoryResponse.Name,
			Description: categoryResponse.Description,
		})
	}
}

func (c *categoryService) CreateCategoryStreamBidirectional(stream grpc.BidiStreamingServer[pb.CreateCategoryRequest, pb.Category]) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		categoryResponse, err := c.categoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.Category{
			Id:          categoryResponse.ID,
			Name:        categoryResponse.Name,
			Description: categoryResponse.Description,
		})
		if err != nil {
			return err
		}
	}
}

func (c *categoryService) ListCategories(context.Context, *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.categoryDB.FindAll()
	if err != nil {
		panic(err)
	}

	var categoriesResponse []*pb.Category
	for _, category := range categories {
		categoryResponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return &pb.CategoryList{Categories: categoriesResponse}, nil
}

func (c *categoryService) GetCategoryByID(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.categoryDB.FindByID(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}
