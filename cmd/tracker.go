package cmd 

type Expense struct{
  ID int `json:"id"`
  Description string `json:"description"`
  Amount float32 `json:"amount"`
  Date string `json:"date"`
}