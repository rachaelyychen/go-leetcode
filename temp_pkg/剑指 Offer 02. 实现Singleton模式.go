package temp_pkg

import (
	"fmt"
	"sync"
)

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/17/21 10:43 AM
**/

/**
懒惰模式的单例模式，考虑多线程安全，使用sync包下的once.Do(f)方法初始化对象，保证多线程环境下只会执行一次。
这是如何做到的呢？对于一个Once对象来说，只会在Do(f)方法第一次被调用时调用它，也即调用f函数，实现方式也是双重加锁检查：
1.Once对象有两个字段：uint32类型的done、mutex类型的m
2.第一次检查通过atomic包的原子操作完成done判断（有没有调用过，没有才进行调用），第二次检查通过m加锁解锁再次完成done判断
3.原子操作使用的是cas机制乐观锁，mutex是一种悲观锁，相对来说一个效率高些（fast）、一个效率低些（slow）
那为什么不使用一个乐观锁解决呢：
if atomic.CompareAndSwapUint32(&o.done, 0, 1) {
	f()
}
因为要保证Do(f)方法返回时f调用结束，假设两个线程抢占乐观锁，失败的线程会立刻返回，而不会等待成功的线程调用f结束；
如果再加一个悲观锁，那么另一个线程会被阻塞，等待f调用结束。
 */

type Singleton struct{}

var singleton *Singleton
var once sync.Once

func NewSingleton() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}

func main() {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	const cnt = 100
	wg.Add(cnt)
	singletons := [cnt]*Singleton{}
	for i := 0; i < cnt; i++ {
		go func(i int) {
			<-start
			singletons[i] = NewSingleton()
			wg.Done()
		}(i)
	}
	close(start)
	wg.Wait()
	for i := 0; i < cnt; i++ {
		if singletons[i] != singleton {
			fmt.Println("singleton not equal")
			return
		}
	}
}
