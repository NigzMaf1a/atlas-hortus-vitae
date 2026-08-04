package auth

type User struct {
	UserId    int    `json:"user_id"`
	SectorId  int    `json:"sector_id"`
	RoleId    int    `json:"role_id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	AccStatus string `json:"acc_status"`
	RegType   string `json:"reg_type"`
	Location  string `json:"location"`
}

type LoginCred struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type HortusVirtaeCred struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	OutletID int64  `json:"outlet_id"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type HortusLoginResponse struct {
	Token      string `json:"token"`
	User       User   `json:"user"`
	OutletName string `json:"outlet_name"`
	OutletID   string `json:"outlet_id"`
}
