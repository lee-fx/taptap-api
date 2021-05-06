package taskrunner

const (
	READY_TO_DISPATCH = "d" // producer/dispatch 生产者
	READY_TO_EXECUTE = "e"  // consume/execute 消费者
	CLOSE = "c"             // 关闭

	VIDEO_PATH = "./videos/"
)

// contro channel
type controlChan chan string

// data channel
type dataChan chan interface{}

// dispather/executor
type fn func(dc dataChan) error
