package main

import "sync/atomic"

func main() {
	var reqNUm uint32
	atomic.StoreUint32(&reqNUm, 100)               //原子的给变量写入值
	_ = atomic.LoadUint32(&reqNUm)                 //原子的从变量获取值
	atomic.CompareAndSwapUint32(&reqNUm, 100, 200) //原子的比较，如果新值和旧值不一样，则设置为新值
	atomic.AddUint32(&reqNUm, 1)                   //原子的加1
	atomic.AddUint32(&reqNUm, ^uint32(1-1))        //原子的减1
	atomic.SwapUint32(&reqNUm, 10)                 //原子的保存新值，且返回旧值
}
