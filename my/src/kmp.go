package main 
func main(){
	patter := "aaaccefg"
	arr := "-1,0,1"

	dp :=  make( []int , len(patter) )
	dp[0] =-1
	dp[0] = 0 
	k := 0 
	for i:= 2 ;i<len (patter) ;i++{
		 if k == -1 ||  patter[i-1] == patter[k] {
			dp[ i] = k+1 
		 }else{
			 k = next[k]
		 }
		
	}
}

func search( str string , pt string , dp [] int ){
	i , k := 0
	for i<len(str) && k <len(pt){
		 if   k = -1 || str[i] == pt[k]{
			i++;
			k++;
		 }else{
			k=dp[k]
		 }
	}
}
