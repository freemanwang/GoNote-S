package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["alice"] = Vertex{
		18, 2000,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["alice"])

	var n = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(n)

	//若顶级类型只是一个类型名，你可以在文法的元素中省略它。
	// 注意，和上块文法比起来，内部少了 Vertex
	var o = map[string]Vertex{
		"Taobao": {200, -200.25},
		"Tencent":    {300, -350.50},
	}
	fmt.Println(o)



}

