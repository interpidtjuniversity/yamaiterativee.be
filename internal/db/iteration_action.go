package db

type IterationAction struct {
	ID             int64    `xorm:"id autoincr pk"`
	FinallyPass    bool     `xorm:"finally_pass" json:"-"`
	PipeLineId     int64	`xorm:"pipeline_id"`
	EnvGroup       string   `xorm:"env_group"`
	State          string   `xorm:"state"`
	PipeLineExecId []int64	`xorm:"pipeline_exec_id"`
}

type IterationMergeRequest struct {
	IterationAction
	SponsorPassId []int64  `xorm:"sponsor_pass_id" json:"-"`
	SponsorID     int64    `xorm:"sponsor_id"`
	SponsorName   string   `xorm:"sponsor_name"`
	ReviewersID   []int64  `xorm:"reviewers_id" json:"-"`
	ReviewersName []string `xorm:"reviewers_name" json:"-"`
}

type IterationServerApply struct {
	IterationAction
}

type IterationResourceRelease struct {
	IterationAction
}

