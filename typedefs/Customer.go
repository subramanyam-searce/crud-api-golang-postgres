package typedefs

type Customer struct {
	Name      string `json:"name"`
	Id        int    `json:"id"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
}
