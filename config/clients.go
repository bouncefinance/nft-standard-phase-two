package config

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"sync"
)

type ProIDFlag struct {
	task       int
	PorID      string
	HttpClient *ethclient.Client
}

func (f *ProIDFlag) Increase() {
	f.task++
}
func (f *ProIDFlag) Decrease() {
	f.task--
}
func NewProIDFlag(proID string, task int) *ProIDFlag {
	client, err := ethclient.Dial(RPCHttpURL + proID)
	if err != nil {
		log.Fatal("dial http client error: ", err)
	}
	return &ProIDFlag{
		task:       task,
		PorID:      proID,
		HttpClient: client,
	}
}

type ProIDFlags struct {
	Ps     []*ProIDFlag
	locker sync.Mutex
}

/**
 * @parameter:
 * @return:
 * @Description: 返回节点中，任务最少的client；当前的任务数量;client所在的下标
 * @author: shalom
 * @date: 2020/12/29 15:08
 */
func (p *ProIDFlags) GetMin() (*ProIDFlag, int, int) {
	p.locker.Lock()
	defer p.locker.Unlock()
	var (
		j     = INT_MAX
		index int
		s     *ProIDFlag
	)
	for i := 0; i < len(p.Ps); i++ {
		if p.Ps[i].task < j {
			j = p.Ps[i].task
			s = p.Ps[i]
			index = i
		}
	}
	p.Ps[index].Increase()
	return s, j, index
}

func NewProIDFlags(ps []*ProIDFlag) *ProIDFlags {
	return &ProIDFlags{
		Ps:     ps,
		locker: sync.Mutex{},
	}
}
