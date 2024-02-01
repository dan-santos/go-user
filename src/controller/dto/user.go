package dto

type UserRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*_-"`
	Name string `json:"name" binding:"required,min=3,max=50"`
	Age int8 `json:"age" binding:"required,min=1,max=100"`
}

type UserUpdateRequest struct {
	Name string `json:"name" binding:"omitempty,min=3,max=50"`
	Age int8 `json:"age" binding:"omitempty,min=1,max=100"`
}

type UserLoginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*_-"`
}

type UserResponse struct {
	ID string `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Age int8 `json:"age"`
}