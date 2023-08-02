package entity

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"required"`
	Email     string `json:"email" binding:"gte: 1, lte: 100"`
}

type Video struct {
	Title       string `json:"title" binding:"min=2"`
	Description string `json:"description" binding:"required"`
	URL         Person `json:"author" binding:"required"`
}
