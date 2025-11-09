package program

var count int

func TowerOfHanoi(disk int, start, end, temp string) int {
	// Total moves to transfer disk from source to destination using auxiliary rod == 2^n-1
	if disk < 1 {
		return 0
	}
	count++
	TowerOfHanoi(disk-1, start, temp, end)
	TowerOfHanoi(disk-1, temp, end, start)
	return count
}
