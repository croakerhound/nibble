package main

func main() {

	rdb, ctx, _ := connectToRedis()
	keys := getAllKeys(rdb, ctx)
	data, _ := getAllKeysAndValues(rdb, ctx)
	printKeys(keys)
	printKeysAndValues(data)
}
