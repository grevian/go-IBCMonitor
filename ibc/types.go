package ibc

import "time"

const (
	HEATING = "Heating"
	STANDBY = "Standby"
	PURGING = "Purging"
	IGNITING = "Igniting"
	INITIALIZING = "Initializing"
)

const (
	IBC_MASTER_BOILER_DATA = 2
	IBC_BOILER_STATUS_DATA = 3
	IBC_BOILER_RUN_PROFILE_DATA = 5
	IBC_BOILER_LOG_DATA = 6
	IBC_BOILER_BOILER_DATA = 11
	IBC_BOILER_STANDARD_DATA = 13
	IBC_BOILER_ADV_SETTINGS_DATA = 15
	IBC_BOILER_EXT_DETAIL_DATA = 19
	IBC_CLOCK_DATA = 24
	IBC_BOILER_SITE_DATA = 34

	// --- Advanced settings, things like load pairing information not useful to most people
	IBC_LOAD_PAIRING_DATA = 25
	// TODO will this tell me when I need to run a cleaning cycle?, what does that entail?
	IBC_BOILER_CLEANING_SETTINGS_DATA = 18
	IBC_BOILER_MULTI_SETTINGS_DATA = 17

	// --- These fail when hit with the parameters I was using, more discovery needed
	// Example response { "object_no": 201, "fail_code": 1, "operation": 7, "object_index": -1 }
	IBC_BOILER_ERRLOG_DATA = 7
	IBC_BOILER_SETBACK_DATA = 14
	IBC_BOILER_LOAD_SETTINGS_DATA = 16
	IBC_SITELOG_DATA = 23
	IBC_BOILER_CAPTURE_DATA = 26

	// --- DO NOT USE --- (These are either for setting values, which is not supported with this class, or look
	// dangerous enough I didn't want to hit them to find out
	// IBC_PASSWORD_DATA = 99
	// IBC_ALERT_DATA = 31
	// IBC_BOILER_BACKUP_ADVANCED = 27 # Discovered by me, http://192.168.2.13/js/adv.js:4
	// IBC_BOILER_RESTORE_ADVANCED = 28 # Discovered by me, http://192.168.2.13/js/adv.js:5
	// IBC_BOILER_RESTORE = 29
	// IBC_BOILER_FACTORY_DATA = 20
	// IBC_BOILER_FACTORY_SETTINGS_DATA = 21

	_UNSET_SENTINEL_VALUE = 32766
)

const TARGET_URL = "cgi-bin/bc2-cgi"

const MINIMUM_DELAY = time.Second * 9

type IBCRequest struct {
	Object_no int `json:"object_no"`
	Object_request int `json:"object_request"`
	Boiler_no int `json:"boiler_no"`
}

type IBCResponseExtDetail struct {
	ObjectNumber int `json:"object_no"`
	Status string
	Warnings string
	Errors string
	MBH int
	SupplyT int
	ReturnT int
	TargetT int
	StackT int
	AirT int
	IndoorT int
	OutdoorT int
	SecondaryT int
	TankT int
	InletPressure float32
	OutletPressure float32
	DeltaPressure float32
	Servicing int
	MajorError int
	MinorError int
	SystemError int
	WarnFlags int
	OpStatus int
}