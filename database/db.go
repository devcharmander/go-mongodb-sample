package database

// Actions performed on the database
type Actions interface {
	Create(data interface{}) (*[]Habit, error)
	Retrieve(data interface{}) ([]*Habit, error)
	//	Update(data interface{}) (*Habit, error)
	//	Delete(data interface{}) (*Habit, error)
}
