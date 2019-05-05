package main

import "fmt"

// PersonInfo 是一个包含详细个人信息的类型
type PersonInfo struct {
	ID string
	Name string
	Address string
}

func main() {
	//map[KeyType]ValueType
	var personMap map[string] PersonInfo
	personMap = make(map[string] PersonInfo)
	
	// 往这个 map 里插入几条数据
	personMap["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personMap["1"] = PersonInfo{"1", "Jack", "Room 101,..."}
	
	// 从这个 map 查找 key 为"1234"的信息
	person, ok := personMap["12345"]
	
	// ok 是一个返回的 bool 型,返回 true 表示找到了对应的数据
	if ok {
		fmt.Println("Found person", person.Name, "with ID 12345.")
	} else {
		fmt.Println("cannot find person with ID 12345.")
	}
	//删除map
	delete(personMap,"12345")
	person2, ok := personMap["12345"]
	if ok {
		fmt.Println("after delete 12345 Found person", person2.Name, "with ID 12345.")
	} else {
		fmt.Println("after delete ,cannot not find person with ID 12345.")

	}
}
