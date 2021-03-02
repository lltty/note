1 rpc深入理解
  rpc的通信协议一般是使用tcp或者udp的。当然也可以使用http 
4 runtime的深入理解
    runtime实现了golang的并发调度，内存分配，并发垃圾回收。
5 make和new的区别

6 context包

7 三大装逼利器: 内存分配，线程调取， 垃圾回收

    1 内存分配
    栈中一般存储函数的参数，临时变量。栈的内存读取速度快，但是缺乏灵活性，栈中数据的声明周期是固定的，与函数共存亡
    堆上一般存储引用数据，比如对象，数组等等。优势是可以动态分配内存大小，适合不可预知大小的内存分配，声明周期享有主动控制权，缺点是复杂的内存分配管理会占用资源、速度慢、而且会出现内存碎片已经内存泄漏的问题。
    栈内存可能会逃逸到堆内存。
    
    Golang语言内置运行时，抛弃了传统的内存分配方式，改为自主管理。Golang运行时的内存分配算法主要源自于Google为C语言开发的TCMalloc算法。核心思想就是把内存分为多级管理，从而降低锁的粒度。它将可用的堆内存采用二级分配的方式进行管理。每个线程都会自行维护一个独立的内存池，进行内存分配时，优先从该内存池中分配，当内存不足时才会想全局内存申请，以避免不同线程对全局内存的池的频繁竞争。
    
    Go程序在启动的时候，会先向操作系统申请一块虚拟内存，切成小块自己进行管理。申请到的内存被分配成三个快，在64位机器上分别是spans(512M),bitmap(16G),arena(512G)的大小。
    
    arena区域就是我们所谓的堆区，Go动态分配的内存都是在这个区域，它把内存分割成8KM大小的页，一些页组合起来称为mspan.
    
    bitmap区域标识arena区域哪些地址保存了对象，并且保存了对象是否包含指针，GC信息。bitmap中的一个字节的内存对应arean区域中4个指针大小的内存，一个指针是8个字节，所以bitmap的大小是512G/(4*8B)=16GB
    
    bitmap是由高地址向低地址增长的，所以bitmap的高地址是指向arean地址的低地址的。
    
    spans区域存放mspan(也就是arena分割的页面组合起来的内存管理基本单元)的指针，每个指针指向一页，所以span区域的大小是 512G/(8KB*8B)=512MB.创建mspan的时候，按页填充对应的span区域，在回收Object时，根据地址很容易就能找到它所属的mspan.
    
    1.1 内存管理单元
    
    mspan可以理解为是一个包含起始位置、mspn规格、页的数量等内容的双端链表，每个msapn按照它自身的属性Size的大小分割成若干个object,每个object可存储一个对象。并且会使用位图来标记尚未使用的obejct.属性size决定obejct的大小，而mspan只会分配尺寸大小接近的对象。
    
    Size class = Span class / 2
   
  这是因为其实每个size class 有两个mspan,也就是两个Span class.其中一个分配给含有指针的对象，另一个分配给不含有指针的对象。这个主要是为了垃圾回收。
  
  mspan根据size class可以得到它划分的Object的大小，分配器就可以将内存接近的对象保存到object里，对于微小对象，分配器会将其合并，将几个对象分配到同一个object中。
  
  超过32KB,此对象就是大对象了，它会被特别对待。类型Size class为0大对象，它是直接由堆内存分配，而小对象是通过mspan来分配。
  
  1.2 内存管理组件
  内存分配由内存分配器完成。分配器由3种组件构成：mcache, mcentral, mheap.
  
  1.2.1 mcache
  每个工作线程都会绑定一个mcache,本地缓存可用的mspan资源，这样就可以直接给Goroutine分配内存，因为不存在多个线程竞争的情况，所以不会消耗锁资源。
  
  mcache用Span Classes 作为索引管理多个用于分配的的mspan,它包含所有规格的mspan。它是_NumSizeClasses的2倍。这里是2倍的原因是：为了加速之后内存的回收速度，数组里一半的mspan中分配的对象不包含指针，另一半则包含指针。
  
  对于无指针对象的mspan在进行垃圾回收的时候无需进一步扫描它是否引用了其他活跃的对象，可以直接回收。
  
  mcache在初始化的时候是没有任何mspan资源的，在使用的过程中会动态的从mcentral申请，之后会缓存下来。当对象小于等于32KB大小时，使用mcache的相应规格的mspan进行分配。
  
  1.2.2 mcentral
  为所有mcache提供切分好的mspn资源。每个central保存一种特定大小的全局mspan列表,包括已分配出去的和还未分配出去的。每个mcentral对应一种mspan,而mspan的种类导致它分割的object大小不同。当工作线程mcache中没合合适的mspan时就会从mcentral获取。
  
  mcentral被所有的工作线程共同享有，存在多个线程竞争的情况，因此会消耗资源。结构体定义：
  
  type mcentral struct {
       lock mutex          //互斥锁
       sizeclass int 32    //规格
       nonempty mSpanList  //尚有空间Object的mspan链表
       empty  mSpanList    //没有空闲Object的mspan链表，或者已被mcache取走的mspan链表
       nmalloc uint64      //已经累计分配的对象个数
  }
  
  empty表示这条链表里的mspan都被分配的obejct，或者都已经被mcache取走了mspan,这个mspan被那个工作线程独占了。而nonempty则表示有空闲对象的mspn列表。每个mcentral结构体都在mhap中维护。
  
  简单说下mcache从mcentral中获取和归还mspan的流程：
  * 获取加锁:从nonempty列表中找到一个可用的mspan;并将其从nonempty链表中删除；将取出的mspan加入到empty链表；将mspan返回给工作线程；解锁。
  * 归还解锁：将mspan从empty链表删除；将mspan加入到nonempty;解锁
  
  1.2.3 mheap
  代表Go程序持有的所有堆空间，Go程序使用一个mheap的全局对象 _mheap来管理堆内存。
  
  当mcentral没有足够的mspan时，会向mheap申请。而mheap没有资源时，会向操作系统申请内存。mheap主要用于大对象的内存分配，以及管理未切割的mspan,用于给mcentral切割成小对象。
    
   mheap中含有所有规格的mcentral,所以，当一个mcache从mcentral申请mspan时，只需要独立的mcentral中使用锁，并不会影响申请其他规格的mspan.
   
   mheap的结构体定义：
   type mheap struct {
        lock mutex
        spans []*mspan   //指向mspan区域，用于映射mspan和page的关系
        bitmap uintprt   //指向bitmap首地址，bitmap是从高到低增长的
        area_start uintptr //指向arena区首地址
        area_used uintptr  //指向arena区中已使用地址位置
        area_end uintptr   //指向arena区末地址
        central [67*2]struct {
            mcentral mcentral
            pad [sys.CacheLineSize - unsafe.Sizeof(mcentral{})%sys.CacheLineSize]byte
        }
   }
   
   1.2.4 分配流程
   变量是在栈上还是在堆上分配，是由逃逸分析的结果决定的。
   Go的内存分配器在分配对象时，根据对象的大小，分成三类：微对象(<=16B),小对象(大于16B，小于等于32K )，大对象（>32KB）.
   
   微对象直接在mcache的tiny分配器分配，大对象直接在mheap分配，小对象首先计算对象的规格，然后使用mcache中相应规格的mspan分配。
   如果mcache没有对应规格大小的mspan,则向mcentral申请
   如果mcentral也没有，向mheap申请
   如果mheap中也没有，则向操作系统申请。
   
   1.2.5 总结
   * Go在程序启动时，会向操作系统申请一大块内存，之后进行管理
   * Go内存管理的基本单元是mspan,它由于若干个页组成，每种mspan可以分配特定大小的Object,多个变量也可以组合保存到一个Object里。
   * mcache,mcentral，mheap是Go内存管理的三大组件，层层递进。mcache管理线程在本地缓存的mspan;mcentral管理全局的mspan供所有线程使用;mheap管理Go所有的动态分配内存。
   * 极小对象对分配在同一个Object中，以节省资源，使用tiny分配器分配；一般小对象通过mspan分配内存；大对象则直接由mheap分配内存。
   
   1.2.6 补充
   
   知乎链接地址:[https://zhuanlan.zhihu.com/p/59125443](https://zhuanlan.zhihu.com/p/59125443)
   
   大佬博客:[https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-memory-allocator/](https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-memory-allocator/) 
  
  
8 google三架马车:GFS, BigTable, MapReduce

9 部分包的理解:sync.map

10 反射

11 select

12 defer三要素
   1 defer传递的参数是在声明时候立即求值的
   2 多个defer的声明是先声明，后执行
   3 return,defer,返回值的执行顺序
   return 最先执行，return负责将计算结果赋值给返回值;接着defer做一些收尾工作；最后函数携带当前返回值退出
   
   
13 控制请求超时的方式

14 context的作用

    在并发编程中，由于超时、取消操作或者一些异常情况，往往需要进行抢占操作或者中断后续操作。熟悉channel的，就会想到，可以定义一个channel作为开关通知，如果主协程需要在某个时刻发送消息通知子协程中断任务，那么就可以让子协程监听这个信号，一旦主协程关闭，那么子协程也就可以退出了。但是这个是有限的。如果主协程有多个子任务，子任务又有多个子任务，这个子任务，既要感知主协程的信号，又要感知父任务的信号，这个就需要定义多个信号来处理，嵌套的越深，这个就麻烦。所以就可以用context来控制就很容易了。
    
    上层任务取消后，所有的下层任务都会被取消
    中间某一层的任务取消后，只会将当前任务的子任务取消，不会影响其他层级的任务。


    1 可以控制不同节点的执行时间,主节点超时，子节点结束，或者子节点超时，通知其他节点结束
    2 多个线程间数据共享，一般控制在一个请求周期内的所有线程共享数据
    3 context的设计是不可变的,包括超时时间，context.Value()。
    4 context适合保存request范畴的值
    5 context.Background() 和 context.Todo() 有何区别
    实际上，没有任何区别，只是语意上的区别
    context.Background: 是上下文的默认值，所有其他的上下文都应该从它衍生出来
    context.Todo: 是在不确定使用哪种上下文的时候，使用。
    
15 互斥锁的正常模式于饥饿模式
    在正常模式下，锁的等待会按照先进先出的顺讯获取锁。但是刚刚被唤起的线程与新创建的线程竞争时，大概率是获取不到的锁的，为了解决这种情况，引入了饥饿模式。一旦线程超过1ms没有获取到锁，当前互斥锁会被切换为饥饿模式，为了防止部分线程被饿死。
     饥饿模式可以保证公平性，在饥饿模式中，互斥锁会直接交给队列最前面的线程。新的线程在该状态下是不能获取到锁的、也不会进入自旋状态，它们只会在队列的最末尾等待。如果一个线程获取到的锁，并且它在队列的最末尾或者获取锁的时间少于1ms,那么当前的互斥锁就会切换为正常模式。
     相比饥饿模式，正常模式下的互斥锁能够提供更好的性能，饥饿模式能避免线程由于陷入等待无法获取到锁而造成的高尾延时。
 
  获取锁的步骤
  
16 方法的结构指针接受者和值结构值接收者
   指针接收者和值接收者使用起来都是一样的，区别在于指针接收者，可以在方法内部改变结构内部的变量的值。
   对于指针接收者，如果调用的是值方法，并不会改变结构内的变量值
   对于值接收者，如果调用的是指针方法，也会改变结构内的变量值
   但是实现接收时候，就会有些不同，如果要把一个变量赋值给接口，那么值方法实现的接口，就一定要赋值值接收者，如果是指针方法实现的接口，那就一定要赋值指针接收者才可以。
   
17 channel的实现原理

18 网络轮询器
    网络轮询器是Go语言用来处理I/O操作的关键组件，它使用了操作系统的I/O多路复用机制增强程序的并发处理能力。
    
19 unsafe.pointer

20 golang使用何种io模型
    IO多路复用模型

20 pipeline

21 runtime

    runtime/debug.SetMaxThreads设置系统创建的go程序可用的最大的线程(M)数量
    runtime.GOMAXPROCS(n); // n > 0 设置可同时执行的CPU的最大数目；n =0 获取设置的数量
    runtime.NumGorouting; //获取当前正在执行Gorouting数量
22 线上内存和cpu突然飙高，如何排查
