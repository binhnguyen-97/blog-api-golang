package types

type InsertArticleRequest struct {
	Title            string `json:"title" binding:"required"`
	ShortDescription string `json:"shortDescription" binding:"required"`
	Author           string `json:"author" binding:"required"`
	Content          string `json:"content" binding:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateAccountRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}
