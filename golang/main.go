// gurev ivan 1-279 formula 6
package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/structs"
)

func main() {
	id1 := structs.CreateStruct(89019998833, "MTS", "Evlampiy")
	id2 := structs.CreateStruct(89991124823, "", "Robert")
	fmt.Printf("Номер данного пользователя: %d\n", id1.GetNumber())
	fmt.Printf("Имя данного пользователя: %s\n", id1.GetName())
	fmt.Printf("Провайдер данного пользователя %s\n", id1.GetProvider())
	id2.SetNumber(8999777665)
	id2.SetProvider("TELE 2")
	fmt.Printf("Номер данного пользователя: %d\n", id2.GetNumber())
	fmt.Printf("Провайдер данного пользователя  %s\n", id2.GetProvider())
}
