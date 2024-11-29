package users

type UserRegistrationRequestDTO struct {
	TelegramID int    `json:"telegram_id"`
	Username   string `json:"username"`
	FullName   string `json:"full_name"`
}

type UserDetailOutDTO struct {
	TelegramID    int    `json:"telegram_id"`
	Username      string `json:"username"`
	FullName      string `json:"full_name"`
	BirthdayCount int    `json:"birthday_count"`
	IsPremium     bool   `json:"is_premium"`
}
