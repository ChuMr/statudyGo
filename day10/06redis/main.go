package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
	// "github.com/go-redis/redis" // 注意导入的是新版本
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.10.175:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func V8Example() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		return
	}

	err := rdb.Set(ctx, "key", "value123", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func ZsetExample() {
	if err := initClient(); err != nil {
		return
	}
	zsetKey := "language_rank"

	// rdb.ZAdd(ctx, zsetKey, &redis.Z{
	// 	Score:  90,
	// 	Member: "Golang",
	// })
	// rdb.ZAdd(ctx, zsetKey, &redis.Z{
	// 	Score:  91,
	// 	Member: "Java",
	// })
	// rdb.ZAdd(ctx, zsetKey, &redis.Z{
	// 	Score:  92,
	// 	Member: "Python",
	// })
	// rdb.ZAdd(ctx, zsetKey, &redis.Z{
	// 	Score:  97,
	// 	Member: "JavaScript",
	// })
	// rdb.ZAdd(ctx, zsetKey, &redis.Z{
	// 	Score:  99,
	// 	Member: "C/C++",
	// })
	// // 把Golang的分数加10
	// newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	// if err != nil {
	// 	fmt.Printf("zincrby failed, err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("Golang's score is %f now.\n", newScore)

	// // 增加一个PHP
	// rdb.ZAdd(ctx, zsetKey, &redis.Z{
	// 	Score:  95,
	// 	Member: "PHP",
	// })

	// 取分数最高的3个
	// ret, err := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Result()
	// if err != nil {
	// 	fmt.Printf("zrevrange failed, err:%v\n", err)
	// 	return
	// }
	// for _, z := range ret {
	// 	fmt.Println(z.Member, z.Score)
	// }

	// 取分数最低的3个
	ret, err := rdb.ZRangeWithScores(ctx, zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()

	// if err := initClient(); err != nil {
	// 	return
	// }

	// zsetKey := "language_rank"
	// languages := []redis.Z{
	// 	redis.Z{Score: 90.0, Member: "Golang"},
	// 	redis.Z{Score: 98.0, Member: "Java"},
	// 	redis.Z{Score: 95.0, Member: "Python"},
	// 	redis.Z{Score: 97.0, Member: "JavaScript"},
	// 	redis.Z{Score: 99.0, Member: "C/C++"},
	// }
	// // ZADD
	// num, err := rdb.ZAdd(ctx, zsetKey, languages...).Result()
	// if err != nil {
	// 	fmt.Printf("zadd failed, err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("zadd %d succ.\n", num)

	// // 把Golang的分数加10
	// newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	// if err != nil {
	// 	fmt.Printf("zincrby failed, err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("Golang's score is %f now.\n", newScore)

	// // 取分数最高的3个
	// ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	// if err != nil {
	// 	fmt.Printf("zrevrange failed, err:%v\n", err)
	// 	return
	// }
	// for _, z := range ret {
	// 	fmt.Println(z.Member, z.Score)
	// }

	// // 取95~100分的
	// op := redis.ZRangeBy{
	// 	Min: "95",
	// 	Max: "100",
	// }
	// ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	// if err != nil {
	// 	fmt.Printf("zrangebyscore failed, err:%v\n", err)
	// 	return
	// }
	// for _, z := range ret {
	// 	fmt.Println(z.Member, z.Score)
	// }

}

func main() {
	// V8Example()
	ZsetExample()
	// testConn()
}
