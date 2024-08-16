package redis

import (
	"github.com/redis/go-redis/v9"

	"app/config"
)

func Connect(db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisAddress,
		Password: config.AppConfig.RedisPassword,
		DB:       db,
	})
}

/* -------------------------------------------------------------------------- */
/*                                   Set Key                                  */
/* -------------------------------------------------------------------------- */
// rdb := redis.Connect(0)
// ctx := context.Background()
// err := rdb.Set(ctx, "Key", "Value", 0).Err()
// if err != nil {
// 	fmt.Println("error on set redis", err)
// }
/* -------------------------------------------------------------------------- */

/* -------------------------------------------------------------------------- */
/*                                   Get Key                                  */
/* -------------------------------------------------------------------------- */
// rdb := redis.Connect(0)
// ctx := context.Background()
// value, err := rdb.Get(ctx, "Key").Result()
// if err != nil {
// 	fmt.Println("error on get redis", err)
// }
// fmt.Println(value)
/* -------------------------------------------------------------------------- */

/* -------------------------------------------------------------------------- */
/*                                 Delete Key                                 */
/* -------------------------------------------------------------------------- */
// rdb := redis.Connect(0)
// ctx := context.Background()
// err := rdb.Del(ctx, "Key").Err()
// if err != nil {
// 	fmt.Println("error on get redis", err)
// }
/* -------------------------------------------------------------------------- */
