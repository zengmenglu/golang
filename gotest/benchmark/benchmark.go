package benchmark

func Add(){
	var n int
	var nums []int
	for i:=0;i<100;i++{
		n+=i
		nums = append(nums,i)
	}
}

func AddArray(arr []int){
	n:=0
	for i := range arr{
		n+=arr[i]
	}
}

func Parallel()  {
	
}