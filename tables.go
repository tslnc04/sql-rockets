package rockets

import (
	"fmt"
)

// Rocket is a struct for the rockets table used for easy queries
type Rocket struct {
	ID           int
	Name         string
	Height       float32
	Diameter     float32
	Manufacturer string
}

// Stage is a struct for the stages table used for easy queries
type Stage struct {
	StageID      int
	RocketID     int
	EngineID     int
	EngineAmount int
}

// Engine is a struct for the engines table used for easy queries
type Engine struct {
	ID           int
	Name         string
	ThrustSL     float32
	ThrustVac    float32
	IspSL        float32
	IspVac       float32
	Fuel         string
	Oxidizer     string
	Manufacturer string
}

func (r Rocket) String() string {
	var output string

	output += fmt.Sprintf("ID: %d\n", r.ID)
	output += fmt.Sprintf("Name: %s\n", r.Name)
	output += fmt.Sprintf("Height: %f\n", r.Height)
	output += fmt.Sprintf("Diameter: %f\n", r.Diameter)
	output += fmt.Sprintf("Manufacturer: %s\n", r.Manufacturer)

	return output
}

func (s Stage) String() string {
	var output string

	output += fmt.Sprintf("Stage ID: %d\n", s.StageID)
	output += fmt.Sprintf("Rocket ID: %d\n", s.RocketID)
	output += fmt.Sprintf("Engine ID: %d\n", s.EngineID)
	output += fmt.Sprintf("Engine Amount: %d\n", s.EngineAmount)

	return output
}

func (e Engine) String() string {
	var output string

	output += fmt.Sprintf("ID: %d\n", e.ID)
	output += fmt.Sprintf("Name: %d\n", e.Name)
	output += fmt.Sprintf("Sea Level Thrust: %f\n", e.ThrustSL)
	output += fmt.Sprintf("Vacuum Thrust: %f\n", e.ThrustVac)
	output += fmt.Sprintf("Sea Level ISP: %f\n", e.IspSL)
	output += fmt.Sprintf("Vacuum ISP: %f\n", e.IspVac)
	output += fmt.Sprintf("Fuel: %s\n", e.Fuel)
	output += fmt.Sprintf("Oxidizer: %s\n", e.Oxidizer)
	output += fmt.Sprintf("Manufacturer: %s\n", e.Manufacturer)

	return output
}
