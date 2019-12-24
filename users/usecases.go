package users

func RegisterUserUsecase(email string, password1 string, password2 string) {

}

func GoogleLoginUsecase(code string) (string, string) {
	return "a", "b"
}

func KakaoLoginUsecase(code string) (string, string) {
	return "a", "b"
}

func RefreshTokenUsecase(token string, refreshToken string) string {
	return "a"
}

func VerifyTokenUsecase(token string) bool {
	return true
}
