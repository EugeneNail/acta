package resource

import "github.com/EugeneNail/acta/journal/internal/domain"

type Habit struct {
	Uuid string      `json:"uuid"`
	Icon domain.Icon `json:"icon"`
	Name string      `json:"name"`
}

// NewHabit constructs a habit resource from the domain entity.
func NewHabit(habit domain.Habit) Habit {
	return Habit{
		Uuid: habit.Uuid.String(),
		Icon: habit.Icon,
		Name: habit.Name,
	}
}
