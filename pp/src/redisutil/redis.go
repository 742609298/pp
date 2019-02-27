package redisutil

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379",)

	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()




	//存值
	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	//取值
	mykeyvalue, err := redis.String(c.Do("GET", "mykey"))

	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		log.Print(mykeyvalue)
	}

	//EX seconds -- 指定过期时间，单位为秒.
	//PX milliseconds -- 指定过期时间，单位为毫秒.
	// 设置过期时间5秒
	_, err = c.Do("SET", "mykey", "superWang", "EX", "5")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	// 设置过期时间为24小时
	n, _ := c.Do("EXPIRE", "mykey", 24*3600)
	if n == int64(1) {
		fmt.Println("success")
	}

	//删除
	_, err = c.Do("DEL", "mykey")
	if err != nil {
		fmt.Println("redis delelte failed:", err)
	}

	//批量操作
	_, err = c.Do("MSet", "name", "nick", "age", "18")
	if err != nil {
		fmt.Println("MSet error: ", err)
		return
	}

	r2, err := redis.Strings(c.Do("MGet", "name", "age"))
	if err != nil {
		fmt.Println("MGet error: ", err)
		return
	}
	fmt.Println(r2)

	//hash操作
	_, err = c.Do("HSet", "names", "nick", "suoning")
	if err != nil {
		fmt.Println("hset error: ", err)
		return
	}

	r3,err := redis.String(c.Do("HGet", "names", "nick"))
	if err != nil {
		fmt.Println("hget error: ", err)
		return
	}
	fmt.Println(r3)


	// 队列
	_, err = c.Do("lpush", "Queue", "nick", "dawn", 9)
	if err != nil {
		fmt.Println("lpush error: ", err)
		return
	}
	for {
		r4, err := redis.String(c.Do("lpop", "Queue"))
		if err != nil {
			fmt.Println("lpop error: ", err)
			break
		}
		fmt.Println(r4)
	}
	r5, err := redis.Int(c.Do("llen", "Queue"))
	if err != nil {
		fmt.Println("llen error: ", err)
		return
	}
	fmt.Println(r5)

}

//使用连接池
func getRedis(_tcp string,_ipport string) (redis.Conn) {
	c, err := redis.Dial(_tcp, "ipport")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		//return
	}
	defer c.Close()
	return c

}

//使用连接池
func GetRedisPool() (*redis.Pool) {
	// 从配置文件获取redis的ip以及db
	//host := beego.AppConfig.String("redis.host")
	//dbIndex, _ := beego.AppConfig.Int("redis.db")
	host := "localhost:6379"
	dbIndex := 0
	//log.Println(host)
	//log.Println(db)
	RedisClient := &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", dbIndex)
			return c, nil
		},
	}

	return RedisClient
	//c:=pool.Get()//获取连接
	//defer c.Close()//关闭连接

}



