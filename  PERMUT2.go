package main 
import "fmt"
func main() {
	for ; ; {
		var n int
		var arr[100000] int
		fmt.Scan(&n)
		if(n==0){
			break
		}
		for i := 0; i < n; i++ {
				fmt.Scan(&arr[i])
			}
		tmp :=0
		for i := 0; i < n; i++ {
			if(i !=arr[arr[i]]){
				tmp = 1
				break
			}
		}
		if(tmp ==1){
			fmt.Println("ambigous")
		}else{
			fmt.Println("unambigous")
		}

	
	}
}