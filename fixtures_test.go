package sheru

type Post struct {
	Model
	Title string
}

type Blog struct {
	Model
}

func (Blog) TableName() string {
	return "thoughts"
}

type User struct {
	Model
}

func (User) TableName() string {
	return "blog.users"
}
