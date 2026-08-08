package workers

import "time"

type Worker struct {
	UserId         int64     `json:"user_id"`
	OutletID       int64     `json:"outlet_id"`
	WorkerPresent  string    `json:"worker_present"`
	SignInTime     time.Time `json:"sign_in_time"`
	ShiftTime      time.Time `json:"shift_time"`
	SignInLocation string    `json:"sign_in_location"`
}
