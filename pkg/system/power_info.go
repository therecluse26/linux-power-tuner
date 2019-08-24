package system

import (
	"github.com/distatus/battery"
	"github.com/therecluse26/uranium/pkg/utils"
	"strconv"
)

type BatteryInfo []*battery.Battery

func GetBatteries() BatteryInfo {
	batts, err := battery.GetAll()
	if err != nil {
		utils.HandleError(err, 0, true, true)
	}
	return batts
}

/*
 * Returns current battery level percent (float64)
 */
func (b *BatteryInfo) GetBatteryPercent() float64 {
	var batCount float64 = 0
	var percent float64
	for _, bat := range *b {
		percent += (bat.Current / bat.Design) * 100
		batCount++
	}
	return percent / batCount
}

/*
 * Returns boolean charging status result
 */
func (b *BatteryInfo) isCharging() bool {
	for _, bat := range *b {
		if bat.State == battery.Charging {
			return true
		}
	}
	return false
}

/*
 * Returns charging rate
 *
 * Formatted as string in Watts by default (indicating charge/discharge status)
 * If optional "raw" parameter is supplied a true value, raw rate in mW is returned
 */
func (b *BatteryInfo) GetChargingRate(raw ...bool) interface{} {
	var batCount int = 0
	var rate int
	var rateFmt string
	for _, bat := range *b {
		rate += int(bat.ChargeRate)
		batCount++
	}
	// If a "true" argument is passed for "raw" parameter, return raw value
	for _, r := range raw {
		if r == true {
			return rate / batCount
		}
	}
	// Format as string
	if b.isCharging() == true {
		rateFmt = "Charging ~"
	} else if b.isCharging() == false {
		rateFmt = "Discharging ~"
	}
	rateFmt += strconv.Itoa((rate / batCount) / 1000) + "W"

	if rate == 0 {
		rateFmt = "Idle"
	}

	return rateFmt
}