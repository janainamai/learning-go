package entity

type (
	Category struct {
		ID       int
		Name     string
		CourseID []int
	}

	Course struct {
		ID         int
		Name       string
		CategoryID int
	}
)

func (c *Category) AddCourse(id int) {
	c.CourseID = append(c.CourseID, id)
}
