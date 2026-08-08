package queries

var WorkerQueries = struct {
	CreateWorker           string
	GetWorkers             string
	GetWorkersByOutlet     string
	SignInWorker            string
	UpdateShiftTime        string
	UpdateSignInLocation   string
	UpdateOutletID         string
}{
	CreateWorker: `
		INSERT INTO workers (
			user_id,
			outlet_id,
			worker_present,
			shift_time,
			sign_in_location
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING
			user_id,
			outlet_id,
			worker_present,
			sign_in_time,
			shift_time,
			sign_in_location
	`,

	GetWorkers: `
		SELECT
			user_id,
			outlet_id,
			worker_present,
			sign_in_time,
			shift_time,
			sign_in_location
		FROM workers
		ORDER BY shift_time ASC
	`,

	GetWorkersByOutlet: `
		SELECT
			user_id,
			outlet_id,
			worker_present,
			sign_in_time,
			shift_time,
			sign_in_location
		FROM workers
		WHERE outlet_id = $1
		ORDER BY shift_time ASC
	`,

	SignInWorker: `
		UPDATE workers
		SET
			worker_present = 'Present',
			sign_in_time = NOW()
		WHERE user_id = $1
			AND worker_present = 'Absent'
			AND NOW() >= shift_time - INTERVAL '5 minutes'
	`,

	UpdateShiftTime: `
		UPDATE workers
		SET shift_time = $1
		WHERE user_id = $2
	`,

	UpdateSignInLocation: `
		UPDATE workers
		SET sign_in_location = $1
		WHERE user_id = $2
	`,

	UpdateOutletID: `
		UPDATE workers
		SET outlet_id = $1
		WHERE user_id = $2
	`,
}