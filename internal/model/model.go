package model

// Kratos hello kratos.
type Kratos struct {
	Hello string
}

//Customer
type Customer struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	NickName        string `json:"nickName"`
	Passwd          string `json:"passwd"`
	NewPasswd       string `json:"newPasswd"` //新密码,用户改密时使用
	Salt            int32  `json:"salt"`
	Phone           string `json:"phone"`
	AccountType     int16  `json:"accountType"`
	Status          int16  `json:"status"`
	TimeCreate      int64  `json:"timeCreate"`
	TimeUpdate      int64  `json:"timeUpdate"`
	TimeLatestLogin int64  `json:"timeLatestLogin"`
}

type Blog struct {
	Id         int64  `json:"id"`
	CustomerId int64  `json:"customerId"`
	TypeId     int64  `json:"typeId"`
	BlogTitle  string `json:"blogTitle"`
	Content    string `json:"content"`
	ReadNum    int64  `json:"readNum"`
	Status     int16  `json:"status"`
	TimeCreate int64  `json:"timeCreate"`
	TimeUpdate int64  `json:"timeUpdate"`
	Token      string `json:"token"`
}
