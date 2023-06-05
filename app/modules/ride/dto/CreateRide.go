package rideDto

type CreateRide struct {
	DriverID uint `json:"driver_id"`
	RiderID  uint `json:"rider_id"`
}

type CreateRideResponse struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}


