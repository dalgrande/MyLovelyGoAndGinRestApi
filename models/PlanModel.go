package models //models/PlanModel.go

// CycleParam type
type CycleParam struct {
	// CycleParamID uint `json:"-"`
	ID         uint `json:"-"`
	PriceRenew string
	PriceOrder string
	Months     int
}

// CyclePeriod type
type CyclePeriod struct {
	// CyclePeriodID uint       `json:"-"`
	ID           uint       `json:"-"`
	Monthly      CycleParam `json:"monthly"`
	Semiannually CycleParam `json:"semiannually"`
	Biennially   CycleParam `json:"biennially"`
	Triennially  CycleParam `json:"triennially"`
	Quarterly    CycleParam `json:"quarterly"`
	Annually     CycleParam `json:"annually"`
}

//Product type
type Product struct {
	Name         string        `json:"name"`
	ID           uint          `json:"id"`
	CyclePeriods []CyclePeriod `json:"cycle"`
}
