package main

func main() {

	rdb, ctx, _ := connectToRedis()
	getAllKeys(rdb, ctx)
}
