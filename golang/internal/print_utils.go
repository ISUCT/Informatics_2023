package internal

import "fmt"

func PrintSlice(slice []ResultRecord) {
	for index, record := range slice {
		fmt.Printf("%v. (f(%2.2f) = %2.3f)", index+1, record.X, record.Y)

		if index != len(slice)-1 {
			fmt.Print(", ")

			if index%3 == 2 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
}
