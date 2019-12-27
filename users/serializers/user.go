package serializers

type RegisterDefaultUserRequest struct {
	Email     string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password1 string `form:"password1" json:"password1" xml:"password1" binding:"required"`
	Password2 string `form:"password2" json:"password2" xml:"password2" binding:"required"`
}

type RefreshTokenRequest struct {
	Token        string `form:"token" json:"token" xml:"token"  binding:"required"`
	RefreshToken string `form:"refresh_token" json:"refresh_token" xml:"refresh_token"  binding:"required"`
}
