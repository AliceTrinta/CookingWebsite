package model

type Recipe struct {
	ID 	 	 	int    `json:"id" db:"id"`
	UserID		int    `json:"userId" db:"userId"`
	Name 		string `json:"name" db:"name" form:"name"`
	Type 		string `json:"type" db:"type" form:"type"`
	Description string `json:"description" db:"description" form:"description"`
}
