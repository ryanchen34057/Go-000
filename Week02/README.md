## Homeworl problems
### 1. dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么?
应该，因为这样调用dao层的人才可以得到堆栈讯息，而且也不用每个地方都打印一次日志
