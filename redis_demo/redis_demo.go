package redis_demo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var rdb *redis.Client

/**
*@author: 廖理
*@date:2022/12/23
**/
func initRedisDb() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
}

// doCommand go-redis基本使用示例
func doCommand() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	key := "test"
	// 执行命令获取结果
	val, err := rdb.Get(ctx, key).Result()
	fmt.Println(val, err)

	// 先获取到命令对象
	cmder := rdb.Get(ctx, key)
	fmt.Println(cmder.Val()) // 获取值
	fmt.Println(cmder.Err()) // 获取错误

	// 直接执行命令获取错误
	err = rdb.Set(ctx, key, 100, time.Hour).Err()

	// 直接执行命令获取值
	value := rdb.Get(ctx, key).Val()
	fmt.Println(value)
}

// doDemo rdb.Do 方法使用示例
func doDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 直接执行命令获取错误
	//err := rdb.Do(ctx, "set", "haha", 10, "EX", 10).Err()
	//fmt.Println(err)

	// 执行命令获取结果
	val, err := rdb.Do(ctx, "get", "haha").Result()
	fmt.Println(val, err)
}

// getValueFromRedis redis.Nil判断
func getValueFromRedis(key, defaultValue string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			return defaultValue, nil
		}
		// 出其他错了
		return "", err
	}
	return val, nil
}

//zset示例
//下面的示例代码演示了如何使用 go-redis 库操作 zset。

// zsetDemo 操作zset示例
func zsetDemo() {
	// key
	zsetKey := "language_rank"
	// value
	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// ZADD
	err := rdb.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

//执行上面的函数将得到如下输出结果。

type testZ struct {
	Name string
	Age  int64
}

func (z testZ) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(z)

	if err != nil {
		return nil, err
	}
	return
}

func zsetDemo2() {

	// key
	zsetKey := "language_rank"
	// value
	languages := []*redis.Z{
		{Score: 90.0, Member: &testZ{
			Name: "liq",
			Age:  33,
		}},
		{Score: 98.0, Member: &testZ{
			Name: "hc",
			Age:  31,
		}},
		{Score: 95.0, Member: &testZ{
			Name: "how",
			Age:  32,
		}},
		{Score: 97.0, Member: &testZ{
			Name: "qq",
			Age:  43,
		}},
		{Score: 99.0, Member: &testZ{
			Name: "koala",
			Age:  53,
		}},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// ZADD
	err := rdb.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")
}

//扫描或遍历所有key
//你可以使用KEYS prefix:* 命令按前缀获取所有 key。
//
//vals, err := rdb.Keys(ctx, "prefix*").Result()
//但是如果需要扫描数百万的 key ，那速度就会比较慢。这种场景下你可以使用Scan 命令来遍历所有符合要求的 key。

// scanKeysDemo1 按前缀查找所有key示例
func scanKeysDemo1() {
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()

	var cursor uint64
	for {
		var keys []string
		var err error
		// 按前缀扫描key
		keys, cursor, err = rdb.Scan(ctx, cursor, "*", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key", key)
		}

		if cursor == 0 { // no more keys
			break
		}
	}
}

//Go-redis 允许将上面的代码简化为如下示例。

// scanKeysDemo2 按前缀扫描key示例
func scanKeysDemo2() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	// 按前缀扫描key
	iter := rdb.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

//例如，我们可以写出一个将所有匹配指定模式的 key 删除的示例。

// delKeysByMatch 按match格式扫描所有key并删除
func delKeysByMatch(match string, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	iter := rdb.Scan(ctx, 0, match, 0).Iterator()
	for iter.Next(ctx) {
		err := rdb.Del(ctx, iter.Val()).Err()
		if err != nil {
			panic(err)
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

//此外，对于 Redis 中的 set、hash、zset 数据类型，go-redis 也支持类似的遍历方法。
//
//iter := rdb.SScan(ctx, "set-key", 0, "prefix:*", 0).Iterator()
//iter := rdb.HScan(ctx, "hash-key", 0, "prefix:*", 0).Iterator()
//iter := rdb.ZScan(ctx, "sorted-hash-key", 0, "prefix:*", 0).Iterator()

//
//Redis Pipeline 允许通过使用单个 client-server-client 往返执行多个命令来提高性能。区别于一个接一个地执行100个命令，你可以将这些命令放入 pipeline 中，然后使用1次读写操作像执行单个命令一样执行它们。这样做的好处是节省了执行命令的网络往返时间（RTT）。
//
//y在下面的示例代码中演示了使用 pipeline 通过一个 write + read 操作来执行多个命令。
//
//pipe := rdb.Pipeline()
//
//incr := pipe.Incr(ctx, "pipeline_counter")
//pipe.Expire(ctx, "pipeline_counter", time.Hour)
//
//cmds, err := pipe.Exec(ctx)
//if err != nil {
//panic(err)
//}
//
//// 在执行pipe.Exec之后才能获取到结果
//fmt.Println(incr.Val())
//上面的代码相当于将以下两个命令一次发给 Redis Server 端执行，与不使用 Pipeline 相比能减少一次RTT。
//
//INCR pipeline_counter
//EXPIRE pipeline_counts 3600
//或者，你也可以使用Pipelined 方法，它会在函数退出时调用 Exec。
//
//var incr *redis.IntCmd
//
//cmds, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
//	incr = pipe.Incr(ctx, "pipelined_counter")
//	pipe.Expire(ctx, "pipelined_counter", time.Hour)
//	return nil
//})
//if err != nil {
//panic(err)
//}

// 在pipeline执行后获取到结果
//fmt.Println(incr.Val())

//我们可以遍历 pipeline 命令的返回值依次获取每个命令的结果。下方的示例代码中使用pipiline一次执行了100个 Get 命令，在pipeline 执行后遍历取出100个命令的执行结果。

func PipelinedDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	cmds, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 1; i < 2; i++ {
			pipe.Get(ctx, fmt.Sprintf("test%d", i))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, cmd := range cmds {
		fmt.Println(cmd.(*redis.StringCmd).Val())
	}
}

//在那些我们需要一次性执行多个命令的场景下，就可以考虑使用 pipeline 来优化。

//事务
//Redis 是单线程执行命令的，因此单个命令始终是原子的，但是来自不同客户端的两个给定命令可以依次执行，
//例如在它们之间交替执行。但是，Multi/exec能够确保在multi/exec两个语句之间的命令之间没有其他客户端正在执行命令。
//
//在这种场景我们需要使用 TxPipeline 或 TxPipelined 方法将 pipeline 命令使用 MULTI 和EXEC包裹起来。

// TxPipeline demo
func TxPipelineDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	pipe := rdb.TxPipeline()
	incr := pipe.Incr(ctx, "tx_pipeline_counter")
	pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
	_, err := pipe.Exec(ctx)
	fmt.Println(incr.Val(), err)

	// TxPipelined demo
	var incr2 *redis.IntCmd
	_, err = rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		incr2 = pipe.Incr(ctx, "tx_pipeline_counter")
		pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
		return nil
	})
	fmt.Println(incr2.Val(), err)
}

//上面代码相当于在一个RTT下执行了下面的redis命令：
//
//MULTI
//INCR pipeline_counter
//EXPIRE pipeline_counts 3600
//EXEC

//我们通常搭配 WATCH命令来执行事务操作。从使用WATCH命令监视某个 key 开始，直到执行EXEC命令的这段时间里，
//如果有其他用户抢先对被监视的 key 进行了替换、更新、删除等操作，那么当用户尝试执行EXEC的时候，
//事务将失败并返回一个错误，用户可以根据这个错误选择重试事务或者放弃事务。
//
//Watch方法接收一个函数和一个或多个key作为参数。
//
//Watch(fn func(*Tx) error, keys ...string) error
//下面的代码片段演示了 Watch 方法搭配 TxPipelined 的使用示例。

// watchDemo 在key值不变的情况下将其值+1
func watchDemo(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Millisecond)
	defer cancel()
	return rdb.Watch(ctx, func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		// 假设操作耗时5秒
		// 5秒内我们通过其他的客户端修改key，当前事务就会失败
		time.Sleep(5 * time.Second)
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n+1, time.Hour)
			return nil
		})
		return err
	}, key)
}

//将上面的函数执行并打印其返回值，如果我们在程序运行后的5秒内修改了被 watch 的 key 的值，
//那么该事务操作失败，返回redis: transaction failed错误。
const routineCount = 100

func watchINCR() {

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	increment := func(key string) error {
		txf := func(tx *redis.Tx) error {
			// 获得当前值或零值
			n, err := tx.Get(ctx, key).Int()
			if err != nil && err != redis.Nil {
				return err
			}

			// 实际操作（乐观锁定中的本地操作）
			n++

			// 仅在监视的Key保持不变的情况下运行
			_, err = tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
				// pipe 处理错误情况
				pipe.Set(ctx, key, n, 0)
				return nil
			})
			return err
		}

		for retries := routineCount; retries > 0; retries-- {
			err := rdb.Watch(ctx, txf, key)
			if err != redis.TxFailedErr {
				return err
			}
			// 乐观锁丢失
		}
		return errors.New("increment reached maximum number of retries")
	}

	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			defer wg.Done()

			if err := increment("counter3"); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := rdb.Get(ctx, "counter3").Int()
	fmt.Println("ended with", n, err)
}

func RedisDemo() {
	initRedisDb()
	//doCommand()
	//doDemo()
	//r, err := getValueFromRedis("test1", "no value")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(r)

	//zsetDemo()
	//zsetDemo2()

	//scanKeysDemo1()
	//scanKeysDemo2()
	//PipelinedDemo()

	//err := watchDemo("test2")
	//if err != nil {
	//
	//	fmt.Println(err)
	//}
	watchINCR()
}
