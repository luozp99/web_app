package modles

type UserSignUp struct {
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type LoginUser struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDO struct {
	Id       int64  `db:"id""`
	Name     string `db:"string"`
	Age      int    `db:"age"`
	Phone    string `db:"phone"`
	Password string `db:"password"`
}
