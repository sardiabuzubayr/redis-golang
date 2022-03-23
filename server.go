package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	con := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	})

	ctx := context.Background()
	err := con.Set(ctx, "token", "101", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := con.Get(ctx, "token").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	err = con.Del(ctx, "token").Err()
	if err != nil {
		panic(err)
	}

	val, err = con.Get(ctx, "token").Result()
	if err != nil {
		fmt.Println("Data kosong")
	}

	fmt.Println("Data dengan key token setelah dihapus ", val)

	fmt.Println("\nTes set angka/numeric di redis")
	err = con.Set(ctx, "angka", 0, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err = con.Get(ctx, "angka").Result()
	if err != nil {
		fmt.Println("Data dengan key angka kosong")
	} else {
		fmt.Println("Angka sebelum increment ", val)
	}

	err = con.Incr(ctx, "angka").Err()
	if err != nil {
		panic(err)
	}

	val, err = con.Get(ctx, "angka").Result()
	if err != nil {
		fmt.Println("Data dengan key angka kosong")
	} else {
		fmt.Println("Angka setelah increment ", val)
	}

	err = con.IncrBy(ctx, "angka", 10).Err()
	if err != nil {
		panic(err)
	}

	val, err = con.Get(ctx, "angka").Result()
	if err != nil {
		fmt.Println("Data dengan key angka kosong")
	} else {
		fmt.Println("Angka setelah increment 10 = ", val)
	}

	err = con.Decr(ctx, "angka").Err()
	if err != nil {
		panic(err)
	}

	val, err = con.Get(ctx, "angka").Result()
	if err != nil {
		fmt.Println("Data dengan key angka kosong")
	} else {
		fmt.Println("Angka setelah decrement = ", val)
	}

	err = con.DecrBy(ctx, "angka", 4).Err()
	if err != nil {
		panic(err)
	}

	val, err = con.Get(ctx, "angka").Result()
	if err != nil {
		fmt.Println("Data dengan key angka kosong")
	} else {
		fmt.Println("Angka setelah decrement 4 = ", val)
	}
}
