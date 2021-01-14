package gpool

var (
	PoolGO *Pool
)

type Pool struct {
	EntryChannel chan *Task
	JobsChannel  chan *Task
	workerNum    int
}


func init() {
	PoolGO = NewPool(20)
	go PoolGO.Run()
}

func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		JobsChannel:  make(chan *Task),
		workerNum:    cap,
	}

	return &p
}

func (p *Pool) do(workerID int) {
	for task := range p.JobsChannel {
		task.Execute()
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.workerNum; i++ {
		go p.do(i)
	}

	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}
}