package models //models/PlanModel.go

// CycleParam type
type CycleParam struct {
	ID         uint `json:"-"`
	PriceRenew string
	PriceOrder string
	Months     int
}

// CyclePeriod type
type CyclePeriod struct {
	ID           uint       `json:"-"`
	Monthly      CycleParam `json:"monthly"`
	Semiannually CycleParam `json:"semiannually"`
	Biennially   CycleParam `json:"biennially"`
	Triennially  CycleParam `json:"triennially"`
	Quarterly    CycleParam `json:"quarterly"`
	Annually     CycleParam `json:"annually"`	
}

//Products type
type Products struct {
	Name         string      `json:"name"`
	ID           uint        `json:"id"`
	CyclePeriods CyclePeriod `json:"cycle"`
}


