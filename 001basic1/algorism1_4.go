
package main

import "fmt"
import "time"
import "math/rand"

//go run algorism1_4.go
//1〜Nの値を1回使ってできるランダムな順列を作る

func main() {
	//乱数の元になるシードを与える。randomにこのコードが必要。
	rand.Seed(time.Now().UnixNano())
	const N = 20

	var (
        i int = 0
		j int = 0
		flag int = 1
		a[N+1] int
    )
	//1個目の値をセット
	a[0] = irnd(N);
	fmt.Print("random ", irnd(N));

	for i = 1; i <= N; i++ {
		fmt.Print(i, j)
		fmt.Print(" flag ", flag, "\n");
		for flag == 1 {
			//新しいランダム値をセット
			a[i] = irnd(N);
			flag = 0;
			for j = 0; j <= i; j++ {
				fmt.Print(" ai ", a[i], "\n");
				fmt.Print(" aj ", a[j], "\n");
				if (a[i] == a[j]){
					
					fmt.Printf("%d\n", a) 
					fmt.Print("return flag ", flag, "\n");
					break;
				}
			}
			fmt.Print("");
		}
		flag = 1;
	}
	fmt.Printf("%d\n", a) 
	fmt.Print("\n");

}

func irnd(n int) int {
    return rand.Intn(n);
}


