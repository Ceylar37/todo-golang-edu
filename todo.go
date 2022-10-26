package todo

type Todo struct {
	Id          int    `json:"id" db:"id"`
	Description string `json:"description" db:"description"`
	IsDone      bool   `json:"isDone" db:"is_done"`
	IsFavourite bool   `json:"isFavourite" db:"is_favourite"`
}

type ChangeDoneStatusDto struct {
	Id     int  `json:"-"`
	IsDone bool `json:"isDone"`
}

type ChangeFavouriteStatusDto struct {
	Id          int  `json:"-"`
	IsFavourite bool `json:"isFavourite"`
}
