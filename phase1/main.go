package main

import (
	"fmt"
	"phase1/go_basic"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	goBasic()
}

func goBasic() {
	// 只出现一次的数字
	fmt.Println(go_basic.SingleNumber([]int{2, 2, 1}))
	// 回文数
	fmt.Println(go_basic.IsPalindrome(123321))
}
