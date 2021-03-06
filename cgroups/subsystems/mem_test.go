package subsystems

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMemoryCgroup(t *testing.T) {
	arr := []int{}
	for i := 0; i < 1024*1024*102; i++ {
		arr = append(arr, 0)
	}
	fmt.Println(len(arr), os.Getpid())

	memSubSys := MemorySubSystem{}
	resConfig := ResourceConfig{
		MemoryMax: "1024m",
	}
	testCgroup := "testmemorylimit"

	if err := memSubSys.Set(testCgroup, &resConfig); err != nil {
		t.Fatalf("cgroup fail %v", err)
	}

	if err := memSubSys.Apply(testCgroup, os.Getpid()); err != nil {
		t.Fatalf("cgroup Apply %v", err)
	}

	//将进程移回到根Cgroup节点
	if err := memSubSys.Apply("", os.Getpid()); err != nil {
		t.Fatalf("cgroup Apply %v", err)
	}

	time.Sleep(5 * time.Second)

	if err := memSubSys.Remove(testCgroup); err != nil {
		t.Fatalf("cgroup remove %v", err)
	}
	for {
		select {}
	}
}
