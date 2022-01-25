package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

func Init(stringTime string, machineId int64) (err error) {
	var st time.Time

	st, err = time.Parse("2006-01-02", stringTime)
	if err != nil {
		return err
	}

	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineId)
	return
}

func GenId() int64 {
	return node.Generate().Int64()
}

//func main() {
//	if err := Init("2022-01-25",1);err != nil{
//		fmt.Println("初始化获取id值失败 err=",err)
//		return
//	}
//	id := GenId()
//
//	fmt.Println("id值 ：",id)
//}
