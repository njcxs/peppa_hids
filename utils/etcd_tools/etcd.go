package etcd_tools


import (
"github.com/etcd-io/etcd/clientv3"
"fmt"
"time"
"context"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		// Endpoints:   []string{"localhost:4001"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")

	defer cli.Close()
	//设置1秒超时，访问etcd有超时控制
	t1:=time.Now()
	ctx, _ := context.WithCancel(context.TODO())
	//操作etcd
	_, err = cli.Put(ctx, "key", "v")
	//操作完毕，取消etcd
	// cancel()

	t2 :=time.Now()
	fmt.Println("put耗时",t2.Sub(t1))
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}
	//取值，设置超时为1秒
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	t1= time.Now()
	resp, err := cli.Get(ctx, "key")
	fmt.Println("get 耗时:",time.Now().Sub(t1))
	// 	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}

	//测试redis
}
