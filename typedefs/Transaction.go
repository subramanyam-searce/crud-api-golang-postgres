package typedefs

type Transaction struct {
	From      int    `json:"from"`
	To        int    `json:"to"`
	Amount    int    `json:"amount"`
	TimeStamp string `json:"time_stamp"`
}
