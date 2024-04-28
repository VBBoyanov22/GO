package main
import ("fmt")
import("strings")

func main() {
  
    fmt.Println("Please enter something:")


    var inputs []string
	
	for i:=0;i < 4;i++{

		var input string

		fmt.Printf("Please enter something (%d/4): ", i+1)

		_, err := fmt.Scanln(&input)
        if err != nil {
            fmt.Println("Error reading input:", err)
            return
        }
	
		inputs = append(inputs, input)
	}


	fmt.Println("You entered:")
    fmt.Println(strings.Join(inputs, " "))
}