//go泛型的学习  
//endless实现优雅启动
//channel学习

func (test TestController) Delete(c *gin.Context) {
fmt.Println("delete")

	urlMap := map[string]string{
		"url1": "http://127.0.0.1:9090/update",
		"url2": "http://127.0.0.1:9090/update",
	}

	// 使用channel收集结果
	resChan := make(chan int, len(urlMap))
	var wg sync.WaitGroup

	// 并发发送请求
	for _, url := range urlMap {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			resChan <- test.SendHttp(u)
		}(url)
	}

	// 等待所有goroutine完成
	go func() {
		wg.Wait()
		close(resChan)
	}()

	// 收集结果
	resArr := []int{}
	for res := range resChan {
		resArr = append(resArr, res)
	}

	fmt.Printf("并发请求成功返回: %d\n", resArr)
}