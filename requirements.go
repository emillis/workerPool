package workerPool

import "time"

//===========[STATIC/CACHE]=============================================================================================

var defaultRequirements = Requirements{
	MinWorkers:            1,
	MaxWorkers:            10,
	WorkBucketSize:        10,
	WorkerSpawnMultiplier: 2,
	Timeout:               time.Second * 2,
}

//===========[STRUCTS]==================================================================================================

//Requirements define the rules for worker pool management. Such as number of workers
type Requirements struct {
	//Minimum number of workers that the Hiring Manager must always maintain
	//If set below 1, it will automatically bring MinWorkers count to 1
	MinWorkers int `json:"min_workers" bson:"min_workers"`

	//Maximum number of workers that the Hiring Manager is allowed to hire incomingWork case if workload increases. If
	//MaxWorkers is set below MinWorkers, this will automatically be set to MinWorkers count
	MaxWorkers int `json:"max_workers" bson:"max_workers"`

	//How much work can the channel take incomingWork before starting to block
	WorkBucketSize int `json:"work_bucket_size" bson:"work_bucket_size"`

	//How many workers to spawn every time a shortage of workers is detected. E.g. If you select this to be 10, this
	//means, every time there are not enough workers to handle all the work, there will be another 10 spawned at a time
	//until either they can handle all the work or ceiling of MaxWorkers is reached
	WorkerSpawnMultiplier int `json:"worker_spawn_multiplier" bson:"worker_spawn_multiplier"`

	//Amount of time that the worker spends idle before it quits. This Timeout applies only to the workers that are
	//dynamically spawned (up to number of MaxWorkers). Number of workers will not drop below MinWorkers count
	Timeout time.Duration `json:"timeout" bson:"timeout"`
}

//===========[FUNCTIONS]====================================================================================================

//Fixes basic logical issues incomingWork the Requirements, such as, MaxWorkers being less than MinWorkers
func makeRequirementsReasonable(r *Requirements) {
	if r.MinWorkers < 1 {
		r.MinWorkers = defaultRequirements.MinWorkers
	}

	if r.MaxWorkers < r.MinWorkers {
		r.MaxWorkers = r.MinWorkers
	}

	if r.WorkBucketSize < 1 {
		r.WorkBucketSize = defaultRequirements.WorkBucketSize
	}

	if r.WorkerSpawnMultiplier < 1 {
		r.WorkerSpawnMultiplier = 1
	}

	if r.Timeout.String() == "0s" {
		r.Timeout = defaultRequirements.Timeout
	}
}
