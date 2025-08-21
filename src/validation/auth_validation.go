package validation

type Register struct {
	Name     string `json:"name" validate:"required,max=50" example:"fake name"`
	Email    string `json:"email" validate:"required,email,max=50" example:"fake@example.com"`
	Password string `json:"password" validate:"required,min=8,max=20,password" example:"password1"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email,max=50" example:"fake@example.com"`
	Password string `json:"password" validate:"required,min=8,max=20,password" example:"password1"`
}

type GoogleLogin struct {
	Name          string `json:"name" validate:"required,max=50"`
	Email         string `json:"email" validate:"required,email,max=50"`
	VerifiedEmail bool   `json:"verified_email" validate:"required"`
}

type Logout struct {
	RefreshToken string `json:"refresh_token" validate:"required,max=255"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required,max=255"`
}

type ForgotPassword struct {
	Email string `json:"email" validate:"required,email,max=50" example:"fake@example.com"`
}

type Token struct {
	Token string `json:"token" validate:"required,max=255"`
}

type QueryAuth struct {
	Page   int    `validate:"omitempty,number,max=50"`
	Limit  int    `validate:"omitempty,number,max=50"`
	Search string `validate:"omitempty,max=50"`
}
type UpdateAuth2 struct {
	PhoneNumber string `json:"phoneNumber" validate:"omitempty,e164"`
	FullName    string `json:"fullName" validate:"omitempty,min=2,max=100"`
	Password    string `json:"password" validate:"omitempty,min=6"`
	UserAgent   string `json:"userAgent" validate:"omitempty"`
	IPAddress   string `json:"ipAddress" validate:"omitempty,ip"`
}
