package defs

// requests
type UserCreadential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

//

// response
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

// session
type SimpleSession struct {
	Username string  // login name
	TTL		 int64
}