package main
import "fmt"
type Yangcun struct{
	 members1 []string//创建一个空字符串切片用于储存成员
} //构建一个羊村结构体

type Langbao struct{
	 members2 []string
} //构建一个狼堡结构体

func (yang *Yangcun) add(name string){//改变结构体的值，通过结构体指针来处理
	yang.members1 = append(yang.members1,name)
}//向羊村添加角色的方法

func (lang *Langbao) add(name string){
	lang.members2 = append(lang.members2,name)
}//向狼堡添加角色的方法

func main(){
	yangcun := Yangcun{}
	langbao := Langbao{}
	var flag int
	var name string
	for {
		fmt.Println("请输入数字来完成对应的操作:")
		fmt.Println("1.添加羊村角色")
		fmt.Println("2.添加狼堡角色")
		fmt.Println("3.退出")
		fmt.Scanln(&flag)
		if flag == 1 {
			fmt.Println("请输入名字:")
			fmt.Scanln(&name)
			yangcun.add(name)
			fmt.Printf("恭喜 %s 加入羊村\n\n",name)
		}else if flag == 2 {
			fmt.Println("请输入名字:")
			fmt.Scanln(&name)
			langbao.add(name)
			fmt.Printf("恭喜 %s 加入狼堡\n\n",name)
		}else if flag == 3 {
				break
		}else {fmt.Println("输入有误，请重新输入\n\n")}
		
	}//调用相应的方法添加成员
	fmt.Println("羊村成员:",yangcun.members1)
	fmt.Println("狼堡成员:",langbao.members2)//最后输出成员列表
}