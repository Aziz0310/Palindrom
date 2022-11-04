package palindrom

func CheckPalindrome(num int) bool {
	input_num := num
	var remainder int
	res := 0
	for num > 0 {
		remainder = num % 10
		res = (res * 10) + remainder
		num = num / 10
	}
	if input_num == res {
		return true
	}
	{
		return false
	}
}
