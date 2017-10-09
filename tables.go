package rockets

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
	fuel         string
	oxidizer     string
	manufacturer string
}
