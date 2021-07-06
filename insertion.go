package main
import("fmt")

func main(){
	
	var arr =[]int{8,4,1,5,9,2}
	
	fmt.Println(arr)
	fmt.Println(insertion(arr))

}


func insertion(arr []int) []int{
	for i:=0;i<len(arr)-1;i++{
		temp:=arr[i]
		j:=i-1
		for j>0 && a[j]>temp{
			a[j+1]=a[j]
			j--
		}
		a[j+1]=temp
	}
	return arr
}