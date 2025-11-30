package main

func main() {

	rdb, ctx, _ := connectToRedis()
	keys := getAllKeys(rdb, ctx)
	printKeys(keys)
}
