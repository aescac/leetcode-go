package main

/**
https://leetcode.cn/problems/nim-game/
思路：只需要保证石头数量非 4 的整数倍即可，这样自己第一次拿的时候就可让剩余的石头数量为 4 的整数倍，然后可以把 朋友拿-自己拿 看成一个回合，
     这样无论朋友拿多少个，自己只需要保证每个回合拿的总数为4，所有回合结束自己便能拿到最后的石头，4=（3+1）一次拿的最大值+一次拿的最小值
	 通用：n % (max + min) != 0
*/
func canWinNim(n int) bool {
	if n%4 != 0 {
		return true
	}
	return false
}
// 测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试------------------------------------------

// git 测试
func main() {
	println(canWinNim(5))
	println(canWinNim(6))
	println(canWinNim(7))
	println(canWinNim(8))
	println(canWinNim(9))
}
