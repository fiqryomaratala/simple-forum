package memberships

import "time"

//Membuat UserModel dan SignUp Request

//Membuat SignUp request untuk menerima request saat ada yang melakukan SignUp dari Front-End

type (
	SignUp struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

//Membuat UserModel

type (
	UserModel struct {
		ID        int64     `db:"id"`
		Email     string    `db:"email"`
		Username  string    `db:"username"`
		Password  string    `db:"password"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedBy string    `db:"updated_by"`
	}
)
