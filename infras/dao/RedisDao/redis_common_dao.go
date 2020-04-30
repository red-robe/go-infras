package RedisDao

import (
	"GoWebScaffold/infras/utils"
	"errors"
)

// 单次执行命令的R函数,执行完命令自动关闭连接
func (p *RedisPool) R(command string, args ...interface{}) (reply interface{}, err error) {
	conn := p.GetRedisConn()
	defer func() {
		conn.Close()
	}()
	return conn.Do(command, args...)
}

// pipeline 串行命令，减少网络开销
// e.g.: {{"SET","name","ginger"},{"SET","key","value"},{"GET","key"}}
type CommandPipe [][]interface{}
type ReplysPipe []interface{}

func (p *RedisPool) P(commands CommandPipe) (ReplysPipe, error) {
	conn := p.GetRedisConn()
	defer func() {
		conn.Close()
	}()

	var err error
	var replys ReplysPipe

	for _, cp := range commands {
		if cmd, ok := cp[0].(string); ok {
			params := cp[1:]
			err = conn.Send(cmd, params...)
			if utils.ErrorHandler(err) {
				return nil, err
			}
		} else {
			return nil, errors.New("commandPipe type error")
		}
	}

	err = conn.Flush()
	if utils.ErrorHandler(err) {
		return nil, err
	}

	cmdCount := len(commands)
	replys = make(ReplysPipe, cmdCount)

	for i := 0; i < cmdCount; i++ {
		rs, err := conn.Receive()
		if utils.ErrorHandler(err) {
			return nil, err
		}
		replys = append(replys, rs)
	}

	return replys, nil

}
