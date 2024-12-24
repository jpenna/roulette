package simulation

var List = []int{21, 4, 30, 8, 20, 5, 27, 9, 32, 18, 20, 30, 24, 0, 13, 30, 21, 4, 18, 36, 9, 34, 10, 29, 3, 16, 11, 33, 16, 13, 13, 8, 10, 21, 1, 16, 25, 30, 25, 12, 4, 32, 4, 24, 34, 5, 12, 17, 34, 33, 11, 31, 13, 22, 0, 27, 3, 31, 24, 4, 17, 7, 7, 31, 25, 26, 30, 16, 17, 10, 13, 9, 25, 33, 15, 3, 5, 35, 8, 10, 35, 15, 34, 33, 9, 0, 9, 19, 18, 28, 23, 8, 21, 14, 17, 17, 27, 5, 22, 16, 33, 22, 4, 21, 3, 18, 26, 12, 31, 29, 1, 31, 15, 5, 8, 13, 24, 3, 21, 18, 12, 27, 10, 12, 24, 32, 18, 22, 23, 36, 10, 35, 11, 0, 33, 5, 5, 10, 17, 26, 18, 36, 18, 4, 32, 26, 5, 19, 30, 0, 7, 19, 21, 29, 14, 29, 22, 35, 20, 21, 31, 10, 17, 16, 5, 0, 2, 9, 10, 14, 4, 35, 5, 5, 36, 34, 18, 23, 33, 7, 9, 1, 24, 12, 23, 21, 22, 24, 30, 36, 0, 27, 22, 1, 0, 17, 35, 11, 32, 30, 28, 27, 28, 35, 1, 2, 7, 22, 13, 11, 12, 6, 4, 0, 7, 26, 20, 23, 25, 21, 33, 31, 11, 18, 18, 11, 32, 15, 17, 4, 5, 23, 1, 33, 27, 19, 4, 18, 4, 11, 1, 23, 10, 5, 14, 9, 30, 12, 12, 33, 22, 13, 10, 3, 31, 31, 6, 6, 32, 33, 11, 14, 11, 3, 3, 27, 12, 17, 5, 1, 25, 18, 8, 18, 23, 30, 20, 17, 15, 12, 5, 7, 0, 24, 1, 9, 2, 0, 15, 29, 10, 15, 32, 22, 32, 18, 29, 23, 23, 17, 3, 32, 19, 26, 8, 23, 20, 17, 29, 32, 21, 36, 23, 27, 13, 27, 33, 34, 10, 6, 24, 30, 8, 14, 24, 15, 26, 22, 15, 16, 31, 33, 28, 32, 18, 1, 3, 21, 7, 23, 14, 23, 10, 11, 1, 26, 33, 29, 28, 30, 16, 1, 29, 28, 27, 26, 1, 23, 1, 36, 7, 8, 21, 8, 33, 35, 28, 24, 21, 36, 30, 28, 32, 29, 19, 31, 17, 18, 10, 16, 30, 1, 32, 5, 12, 18, 14, 16, 4, 18, 20, 10, 2, 28, 34, 26, 27, 28, 15, 1, 35, 0, 34, 27, 15, 32, 35, 12, 16, 33, 18, 32, 18, 11, 18, 27, 2, 1, 10, 21, 6, 29, 25, 4, 8, 7, 27, 10, 32, 23, 30, 8, 27, 20, 5, 13, 12, 31, 30, 5, 19, 9, 1, 18, 12, 8, 19, 24, 16, 21, 9, 30, 9, 18, 36, 10, 33, 18, 5, 27, 17, 4, 15, 16, 12, 13, 26, 1, 10, 22, 36, 21, 2, 6, 7, 15, 10, 27, 13, 6, 28, 3, 14, 32, 34, 32, 5, 11, 24, 12, 30, 22, 16, 28, 30, 4, 36, 5, 4, 35}

var DiffList = []int{6, 36, 20, 33, 34, 13, 28, 22, 32, 30, 33, 10, 33, 17, 32, 31, 21, 15, 30, 22, 9, 28, 1, 7, 21, 25, 20, 21, 24, 29, 32, 19, 33, 29, 12, 22, 21, 25, 35, 24, 32, 27, 25, 0, 18, 24, 14, 34, 16, 25, 9, 15, 14, 2, 7, 9, 25, 18, 31, 5, 9, 5, 0, 2, 1, 24, 21, 6, 36, 35, 23, 11, 7, 14, 31, 10, 24, 10, 23, 9, 19, 19, 13, 29, 33, 6, 5, 17, 27, 21, 10, 7, 2, 18, 7, 11, 25, 28, 8, 4, 29, 17, 27, 3, 31, 20, 2, 12, 7, 36, 5, 0, 24, 6, 24, 36, 30, 35, 19, 8, 13, 6, 29, 9, 35, 0, 12, 18, 32, 6, 19, 14, 23, 22, 25, 36, 30, 8, 11, 13, 18, 4, 30, 0, 16, 31, 23, 27, 21, 0, 17, 27, 34, 0, 36, 29}

var List2 = []int{6, 36, 20, 33, 34, 13, 28, 22, 32, 30, 33, 10, 33, 17, 32, 31, 21, 15, 30, 22, 9, 28, 1, 7, 21, 25, 20, 21, 24, 29, 32, 19, 33, 29, 12, 22, 21, 25, 35, 24, 32, 27, 25, 0, 18, 24, 14, 34, 16, 25, 9, 15, 14, 2, 7, 9, 25, 18, 31, 5, 9, 5, 0, 2, 1, 24, 21, 6, 36, 35, 23, 11, 7, 14, 31, 10, 24, 10, 23, 9, 19, 19, 13, 29, 33, 6, 5, 17, 27, 21, 10, 7, 2, 18, 7, 11, 25, 28, 8, 4, 29, 17, 27, 3, 31, 20, 2, 12, 7, 36, 5, 0, 24, 6, 24, 36, 30, 35, 19, 8, 13, 6, 29, 9, 35, 0, 12, 18, 32, 6, 19, 14, 23, 22, 25, 36, 30, 8, 11, 13, 18, 4, 30, 0, 16, 31, 23, 27, 21, 0, 17, 27, 34, 0, 36, 29, 21, 4, 30, 8, 20, 5, 27, 9, 32, 18, 20, 30, 24, 0, 13, 30, 21, 4, 18, 36, 9, 34, 10, 29, 3, 16, 11, 33, 16, 13, 13, 8, 10, 21, 1, 16, 25, 30, 25, 12, 4, 32, 4, 24, 34, 5, 12, 17, 34, 33, 11, 31, 13, 22, 0, 27, 3, 31, 24, 4, 17, 7, 7, 31, 25, 26, 30, 16, 17, 10, 13, 9, 25, 33, 15, 3, 5, 35, 8, 10, 35, 15, 34, 33, 9, 0, 9, 19, 18, 28, 23, 8, 21, 14, 17, 17, 27, 5, 22, 16, 33, 22, 4, 21, 3, 18, 26, 12, 31, 29, 1, 31, 15, 5, 8, 13, 24, 3, 21, 18, 12, 27, 10, 12, 24, 32, 18, 22, 23, 36, 10, 35, 11, 0, 33, 5, 5, 10, 17, 26, 18, 36, 18, 4, 32, 26, 5, 19, 30, 0, 7, 19, 21, 29, 14, 29, 22, 35, 20, 21, 31, 10, 17, 16, 5, 0, 2, 9, 10, 14, 4, 35, 5, 5, 36, 34, 18, 23, 33, 7, 9, 1, 24, 12, 23, 21, 22, 24, 30, 36, 0, 27, 22, 1, 0, 17, 35, 11, 32, 30, 28, 27, 28, 35, 1, 2, 7, 22, 13, 11, 12, 6, 4, 0, 7, 26, 20, 23, 25, 21, 33, 31, 11, 18, 18, 11, 32, 15, 17, 4, 5, 23, 1, 33, 27, 19, 4, 18, 4, 11, 1, 23, 10, 5, 14, 9, 30, 12, 12, 33, 22, 13, 10, 3, 31, 31, 6, 6, 32, 33, 11, 14, 11, 3, 3, 27, 12, 17, 5, 1, 25, 18, 8, 18, 23, 30, 20, 17, 15, 12, 5, 7, 0, 24, 1, 9, 2, 0, 15, 29, 10, 15, 32, 22, 32, 18, 29, 23, 23, 17, 3, 32, 19, 26, 8, 23, 20, 17, 29, 32, 21, 36, 23, 27, 13, 27, 33, 34, 10, 6, 24, 30, 8, 14, 24, 15, 26, 22, 15, 16, 31, 33, 28, 32, 18, 1, 3, 21, 7, 23, 14, 23, 10, 11}

var List3 = []int{13, 7, 28, 28, 20, 18, 3, 18, 12, 32, 21, 25, 29, 21, 36, 23, 4, 11, 28, 34, 23, 33, 27, 13, 5, 21, 12, 8, 35, 20, 32, 15, 1, 28, 10, 15, 1, 27, 26, 24, 16, 14, 8, 10, 34, 25, 12, 30, 27, 19, 10, 33, 13, 7, 12, 15, 24, 26, 1, 21, 23, 21, 18, 34, 33, 1, 32, 1, 32, 20, 9, 30, 15, 35, 3, 2, 2, 36, 31, 1, 36, 25, 19, 25, 34, 16, 24, 21, 0, 19, 15, 12, 30, 27, 31, 11, 24, 36, 27, 14, 14, 17, 32, 27, 4, 30, 10, 0, 30, 33, 11, 12, 25, 13, 17, 26, 22, 0, 29, 31, 2, 22, 22, 27, 6, 18, 32, 8, 34, 32, 17, 11, 34, 17, 7, 8, 32, 2, 3, 17, 35, 7, 15, 26, 8, 13, 22, 23, 30, 28, 9, 16, 29, 18, 0, 21, 34, 30, 28, 5, 18, 23, 8, 3, 22, 12, 34, 10, 19, 11, 34, 13, 23, 30, 15, 4, 32, 12, 30, 6, 22, 13, 4, 36, 32, 8, 16, 23, 3, 3, 24, 16, 13, 33, 5, 7, 21, 20, 5, 32, 3, 30, 21, 12, 19, 2, 26, 22, 24, 23, 19, 31, 27, 7, 32, 3, 16, 28, 27, 4, 17, 26, 22, 20, 17, 16, 1, 29, 11, 23, 24, 7, 28, 28, 27, 0, 6, 6, 19, 27, 4, 22, 28, 20, 15, 10, 0, 14, 20, 28, 1, 34, 14, 12, 32, 11, 7, 34, 31, 9, 24, 23, 3, 34, 31, 13, 2, 4, 19, 9, 2, 32, 23, 32, 2, 12, 26, 10, 24, 36, 22, 33, 15, 21, 22, 11, 16, 31, 9, 34, 24, 15, 31, 20, 34, 23, 9, 32, 20, 31, 1, 3, 29, 13, 8, 34, 23, 32, 4, 26, 9, 17, 30, 31, 26, 19, 6, 14, 3, 1, 5, 32, 5, 12, 12, 18, 19, 30, 2, 30, 32, 21, 11, 28, 25, 6, 22, 13, 29, 18, 35, 19, 7, 34, 28, 12, 13, 26, 0, 25, 0, 36, 23, 20, 27, 25, 0, 4, 11, 11, 19, 2, 4, 2, 35, 2, 22, 29, 12, 27, 20, 3, 31, 29, 1, 11, 19, 12, 27, 22, 11, 26, 14, 14, 11, 35, 35, 10, 3, 7, 32, 2, 21, 34, 3, 10, 21, 31, 11, 16, 9, 25, 10, 27, 26, 1, 26, 16, 30, 24, 19, 30, 16, 34, 18, 16, 12, 10, 13, 17, 31, 36, 27, 16, 14, 21, 12, 22, 22, 11, 36, 8, 5, 28, 15, 26, 19, 33, 18, 22, 15, 29, 22, 35, 17, 9, 30, 11, 18, 17, 19, 34, 36, 28, 3, 5, 23, 32, 18, 19, 22, 0, 33, 35, 26, 28, 12, 36, 7, 27, 7, 4, 4, 12, 26, 28, 24, 15, 21, 27, 31, 5, 5, 28, 30, 21, 8, 19, 18, 24, 6, 2, 26, 4, 19, 9, 5, 19, 2, 7, 23, 36}

var List4 = []int{0, 19, 7, 6, 8, 22, 33, 29, 24, 20, 27, 10, 26, 3, 26, 27, 31, 22, 7, 19, 23, 13, 36, 4, 13, 0, 31, 28, 29, 32, 5, 33, 25, 12, 10, 13, 12, 10, 13, 7, 11, 26, 12, 22, 24, 2, 20, 1, 20, 35, 34, 8, 9, 1, 27, 2, 9, 30, 28, 4, 9, 32, 8, 31, 0, 9, 30, 17, 5, 33, 18, 19, 34, 19, 8, 35, 6, 3, 2, 9, 7, 24, 24, 20, 29, 0, 6, 31, 35, 22, 33, 18, 18, 4, 35, 18, 26, 23, 11, 3, 15, 5, 31, 29, 21, 2, 7, 18, 33, 18, 15, 8, 11, 14, 29, 2, 27, 2, 5, 3, 4, 24, 3, 16, 5, 34, 12, 5, 24, 12, 17, 12, 12, 32, 18, 19, 24, 10, 34, 8, 3, 30, 8, 28, 8, 36, 19, 14, 11, 24, 32, 28, 28, 35, 14, 22, 9, 24, 4, 9, 32, 7, 34, 9, 35, 36, 11, 27, 12, 11, 24, 12, 4, 9, 28, 5, 23, 3, 5, 5, 18, 32, 17, 35, 34, 22, 3, 28, 22, 15, 1, 36, 17, 30, 15, 0, 15, 27, 17, 2, 27, 22, 2, 22, 15, 36, 34, 33, 16, 14, 11, 24, 3, 26, 14, 32, 17, 23, 24, 10, 5, 32, 26, 2, 28, 30, 7, 4, 29, 21, 4, 7, 22, 25, 31, 31, 33, 5, 6, 22, 13, 12, 28, 5, 13, 18, 25, 9, 24, 31, 21, 35, 14, 35, 4, 30, 29, 36, 19, 10, 15, 18, 0, 5, 3, 35, 2, 11, 35, 29, 2, 13, 8, 31, 14, 15, 6, 4, 0, 25, 13, 12, 11, 32, 26, 15, 10, 21, 9, 31, 24, 24, 0, 33, 4, 11, 30, 9, 8, 6, 4, 1, 13, 16, 4, 25, 35, 20, 20, 31, 1, 33, 12, 18, 32, 16, 26, 22, 31, 19, 27, 35, 33, 16, 35, 26, 14, 1, 2, 18, 33, 8, 20, 32, 27, 19, 29, 3, 17, 24, 14, 34, 31, 20, 4, 7, 29, 20, 2, 23, 22, 27, 14, 33, 0, 2, 22, 3, 1, 18, 36, 2, 26, 21, 5, 35, 2, 5, 18, 8, 3, 12, 18, 8, 16, 28, 36, 31, 28, 21, 33, 32, 33, 12, 9, 1, 16, 33, 30, 1, 14, 32, 20, 14, 13, 7, 28, 28, 20, 18, 3, 18, 12, 32, 21, 25, 29, 21, 36, 23, 4, 11, 28, 34, 23, 33, 27, 13, 5, 21, 12, 8, 35, 20, 32, 15, 1, 28, 10, 15, 1, 27, 26, 24, 16, 14, 8, 10, 34, 25, 12, 30, 27, 19, 10, 33, 13, 7, 12, 15, 24, 26, 1, 21, 23, 21, 18, 34, 33, 1, 32, 1, 32, 20, 9, 30, 15, 35, 3, 2, 2, 36, 31, 1, 36, 25, 19, 25, 34, 16, 24, 21, 0, 19, 15, 12, 30, 27, 31, 11, 24, 36, 27, 14, 14, 17, 32, 27, 4, 30}

var List5 = []int{8, 9, 13, 12, 20, 24, 2, 1, 7, 3, 13, 29, 17, 7, 8, 22, 2, 21, 9, 1, 12, 35, 9, 24, 9, 12, 16, 9, 19, 22, 25, 21, 16, 18, 23, 7, 27, 27, 16, 9, 4, 34, 8, 13, 18, 33, 20, 7, 14, 33, 1, 29, 23, 30, 18, 20, 8, 36, 29, 26, 13, 28, 26, 9, 25, 34, 21, 35, 21, 31, 36, 9, 6, 21, 23, 17, 18, 36, 26, 30, 0, 28, 7, 15, 31, 22, 6, 18, 19, 8, 33, 15, 27, 16, 8, 30, 4, 19, 0, 10, 9, 9, 31, 10, 35, 16, 2, 13, 35, 16, 33, 17, 15, 26, 7, 5, 31, 0, 29, 11, 15, 29, 7, 7, 35, 21, 14, 19, 27, 31, 30, 1, 10, 31, 32, 20, 5, 17, 4, 19, 23, 32, 12, 7, 7, 27, 6, 24, 19, 19, 29, 0, 6, 25, 27, 30, 31, 35, 22, 19, 23, 10, 20, 36, 27, 29, 2, 24, 7, 11, 15, 24, 19, 8, 34, 14, 20, 12, 36, 29, 9, 17, 1, 8, 35, 0, 7, 21, 24, 20, 35, 14, 15, 23, 25, 14, 34, 27, 12, 14, 31, 32, 1, 17, 31, 3, 34, 26, 35, 1, 24, 1, 22, 18, 29, 36, 24, 9, 32, 3, 23, 36, 15, 27, 9, 14, 15, 10, 32, 27, 14, 7, 6, 20, 19, 11, 17, 2, 30, 35, 34, 31, 34, 18, 15, 31, 9, 9, 21, 11, 21, 9, 28, 23, 21, 27, 2, 15, 11, 16, 26, 20, 24, 9, 12, 0, 2, 5, 32, 31, 7, 33, 12, 36, 16, 36, 35, 15, 8, 33, 17, 13, 21, 21, 16, 14, 20, 27, 29, 10, 20, 5, 15, 23, 23, 19, 2, 8, 35, 20, 32, 33, 14, 24, 8, 5, 16, 32, 20, 23, 10, 24, 4, 13, 16, 16, 24, 28, 32, 2, 28, 14, 28, 27, 26, 9, 8, 21, 14, 7, 1, 12, 17, 16, 2, 10, 27, 21, 35, 1, 0, 10, 8, 6, 17, 29, 33, 29, 21, 36, 2, 24, 10, 34, 2, 9, 34, 5, 9, 10, 30, 26, 13, 8, 22, 28, 31, 7, 31, 2, 17, 25, 16, 33, 4, 2, 27, 10, 23, 33, 13, 30, 25, 31, 35, 17, 36, 26, 35, 16, 27, 19, 16, 34, 0, 29, 35, 22, 3, 12, 30, 13, 7, 26, 17, 31, 8, 26, 3, 11, 35, 18, 11, 2, 12, 34, 29, 26, 35, 15, 33, 7, 0, 19, 7, 6, 8, 22, 33, 29, 24, 20, 27, 10, 26, 3, 26, 27, 31, 22, 7, 19, 23, 13, 36, 4, 13, 0, 31, 28, 29, 32, 5, 33, 25, 12, 10, 13, 12, 10, 13, 7, 11, 26, 12, 22, 24, 2, 20, 1, 20, 35, 34, 8, 9, 1, 27, 2, 9, 30, 28, 4, 9, 32, 8, 31, 0, 9, 30, 17, 5, 33, 18, 19, 34, 19, 8, 35, 6, 3}

var List6 = []int{15, 1, 19, 6, 4, 31, 25, 26, 13, 20, 29, 26, 32, 19, 33, 24, 1, 22, 25, 36, 7, 8, 28, 9, 11, 36, 22, 30, 29, 15, 5, 33, 14, 17, 23, 11, 8, 22, 24, 17, 30, 16, 23, 1, 25, 30, 15, 7, 12, 31, 34, 18, 31, 36, 32, 34, 13, 5, 21, 7, 20, 24, 35, 21, 6, 20, 0, 32, 30, 32, 12, 15, 29, 21, 22, 18, 25, 16, 5, 19, 22, 23, 9, 31, 5, 1, 23, 33, 24, 18, 36, 24, 2, 16, 28, 33, 25, 17, 0, 4, 27, 1, 2, 33, 35, 17, 34, 6, 23, 30, 27, 35, 8, 35, 17, 29, 36, 20, 4, 31, 14, 31, 29, 2, 15, 24, 17, 12, 3, 6, 36, 23, 23, 21, 6, 8, 29, 7, 13, 8, 17, 13, 17, 13, 17, 2, 28, 16, 22, 36, 9, 21, 27, 24, 28, 14, 19, 31, 23, 8, 21, 30, 18, 8, 0, 7, 36, 28, 4, 8, 13, 14, 27, 18, 27, 23, 2, 20, 19, 14, 24, 8, 13, 8, 9, 2, 33, 20, 2, 32, 28, 33, 16, 26, 13, 1, 34, 0, 33, 5, 20, 33, 25, 14, 18, 11, 25, 16, 8, 9, 20, 30, 0, 29, 20, 22, 29, 12, 18, 30, 30, 17, 23, 28, 6, 0, 22, 0, 29, 27, 3, 22, 19, 11, 27, 31, 22, 16, 0, 16, 20, 27, 4, 1, 2, 26, 13, 24, 8, 18, 0, 8, 9, 13, 12, 20, 24, 2, 1, 7, 3, 13, 29, 17, 7, 8, 22, 2, 21, 9, 1, 12, 35, 9, 24, 9, 12, 16, 9, 19, 22, 25, 21, 16, 18, 23, 7, 27, 27, 16, 9, 4, 34, 8, 13, 18, 33, 20, 7, 14, 33, 1, 29, 23, 30, 18, 20, 8, 36, 29, 26, 13, 28, 26, 9, 25, 34, 21, 35, 21, 31, 36, 9, 6, 21, 23, 17, 18, 36, 26, 30, 0, 28, 7, 15, 31, 22, 6, 18, 19, 8, 33, 15, 27, 16, 8, 30, 4, 19, 0, 10, 9, 9, 31, 10, 35, 16, 2, 13, 35, 16, 33, 17, 15, 26, 7, 5, 31, 0, 29, 11, 15, 29, 7, 7, 35, 21, 14, 19, 27, 31, 30, 1, 10, 31, 32, 20, 5, 17, 4, 19, 23, 32, 12, 7, 7, 27, 6, 24, 19, 19, 29, 0, 6, 25, 27, 30, 31, 35, 22, 19, 23, 10, 20, 36, 27, 29, 2, 24, 7, 11, 15, 24, 19, 8, 34, 14, 20, 12, 36, 29, 9, 17, 1, 8, 35, 0, 7, 21, 24, 20, 35, 14, 15, 23, 25, 14, 34, 27, 12, 14, 31, 32, 1, 17, 31, 3, 34, 26, 35, 1, 24, 1, 22, 18, 29, 36, 24, 9, 32, 3, 23, 36, 15, 27, 9, 14, 15, 10, 32, 27, 14, 7, 6, 20, 19, 11, 17, 2, 30, 35, 34, 31, 34, 18, 15, 31, 9, 9, 21}

var Combined1 = []int{36, 12, 32, 23, 19, 2, 12, 36, 8, 15, 2, 3, 6, 32, 27, 35, 13, 0, 19, 12, 27, 25, 15, 22, 32, 4, 1, 2, 23, 16, 7, 33, 14, 24, 12, 13, 27, 29, 17, 11, 22, 14, 28, 18, 33, 31, 2, 25, 23, 4, 4, 22, 11, 19, 14, 30, 20, 9, 28, 7, 17, 10, 31, 25, 8, 25, 4, 10, 14, 30, 27, 18, 0, 29, 18, 22, 28, 30, 13, 13, 15, 12, 23, 31, 33, 17, 26, 13, 18, 13, 0, 7, 23, 35, 9, 18, 8, 22, 30, 23, 2, 5, 9, 26, 28, 16, 18, 1, 35, 36, 20, 21, 7, 15, 15, 10, 14, 26, 27, 5, 24, 27, 27, 11, 26, 9, 10, 8, 14, 35, 3, 9, 32, 17, 27, 29, 11, 13, 12, 1, 24, 13, 7, 22, 4, 16, 8, 0, 0, 0, 23, 6, 15, 8, 14, 25, 17, 34, 21, 21, 2, 1, 7, 26, 7, 3, 32, 20, 27, 11, 22, 3, 12, 29, 6, 36, 9, 32, 0, 10, 23, 21, 20, 16, 6, 11, 28, 4, 35, 12, 19, 23, 20, 8, 25, 22, 1, 33, 10, 2, 16, 13, 2, 11, 24, 7, 35, 30, 17, 16, 18, 22, 35, 33, 32, 4, 1, 26, 16, 33, 12, 33, 9, 4, 13, 4, 17, 10, 32, 26, 17, 31, 3, 4, 12, 32, 10, 35, 13, 13, 34, 23, 17, 12, 9, 32, 4, 28, 6, 0, 28, 17, 11, 17, 6, 7, 7, 7, 34, 18, 31, 15, 15, 22, 30, 36, 1, 16, 11, 9, 29, 31, 15, 5, 15, 23, 28, 4, 19, 4, 36, 25, 34, 28, 21, 11, 26, 24, 0, 7, 13, 17, 0, 2, 6, 30, 24, 2, 29, 33, 2, 0, 29, 24, 3, 19, 11, 12, 2, 19, 16, 4, 17, 4, 21, 17, 3, 12, 16, 6, 10, 0, 24, 29, 5, 15, 30, 29, 34, 7, 26, 0, 28, 3, 23, 26, 17, 9, 22, 7, 0, 35, 3, 27, 31, 17, 2, 3, 35, 22, 13, 22, 23, 7, 16, 27, 23, 32, 0, 25, 27, 34, 11, 23, 25, 1, 0, 14, 13, 18, 27, 34, 31, 1, 7, 20, 3, 30, 1, 23, 3, 35, 29, 14, 7, 17, 12, 1, 25, 35, 26, 35, 28, 12, 18, 3, 9, 27, 22, 12, 8, 30, 29, 31, 16, 2, 12, 29, 32, 35, 12, 7, 30, 34, 26, 7, 21, 36, 5, 22, 36, 24, 12, 14, 21, 20, 24, 11, 12, 26, 34, 31, 29, 30, 36, 25, 29, 15, 17, 6, 24, 3, 1, 4, 13, 30, 22, 26, 6, 4, 1, 33, 13, 7, 23, 22, 13, 4, 8, 15, 23, 18, 8, 33, 29, 12, 8, 5, 34, 3, 16, 3, 35, 25, 2, 3, 4, 14, 12, 8, 19, 14, 12, 10, 34, 16, 31, 22, 31, 9, 23, 7, 1, 34, 25, 19, 31, 21, 28, 12, 13, 22, 16, 28, 9, 4, 5, 31, 23, 17, 6, 27, 33, 36, 28, 9, 13, 10, 13, 14, 0, 30, 21, 33, 34, 23, 21, 8, 24, 9, 11, 10, 26, 5, 33, 31, 12, 15, 12, 16, 29, 7, 36, 26, 19, 20, 29, 26, 34, 17, 22, 24, 25, 29, 3, 4, 3, 29, 14, 33, 4, 23, 20, 30, 9, 13, 32, 11, 17, 34, 31, 27, 25, 1, 26, 3, 9, 2, 34, 11, 32, 34, 8, 10, 1, 26, 19, 4, 9, 23, 35, 35, 5, 6, 14, 12, 18, 36, 24, 8, 5, 21, 6, 23, 20, 19, 13, 35, 8, 7, 36, 23, 17, 27, 23, 20, 12, 5, 30, 31, 5, 23, 5, 0, 32, 22, 24, 19, 18, 28, 19, 10, 13, 6, 34, 3, 12, 0, 0, 4, 9, 29, 14, 20, 5, 10, 4, 9, 30, 21, 14, 25, 33, 6, 29, 2, 13, 13, 34, 12, 2, 30, 28, 13, 1, 27, 20, 7, 22, 0, 36, 8, 1, 13, 10, 28, 29, 8, 14, 18, 18, 9, 34, 11, 7, 2, 28, 4, 2, 11, 20, 13, 6, 31, 14, 8, 33, 35, 15, 20, 19, 30, 7, 12, 31, 18, 1, 24, 8, 31, 31, 32, 17, 24, 7, 22, 6, 18, 13, 21, 27, 26, 9, 31, 25, 25, 18, 35, 24, 23, 27, 15, 11, 26, 22, 3, 32, 15, 31, 12, 32, 11, 6, 32, 11, 32, 1, 32, 12, 24, 27, 21, 1, 33, 8, 18, 36, 33, 31, 24, 36, 32, 17, 18, 2, 28, 24, 22, 32, 32, 31, 29, 15, 0, 33, 0, 8, 3, 3, 22, 1, 32, 36, 12, 25, 36, 14, 30, 29, 34, 6, 3, 25, 22, 2, 31, 14, 18, 36, 9, 20, 23, 2, 9, 20, 27, 31, 25, 22, 25, 12, 33, 30, 19, 9, 32, 5, 10, 19, 17, 9, 27, 0, 31, 5, 20, 12, 8, 5, 11, 14, 4, 11, 21, 1, 4, 23, 29, 5, 25, 15, 11, 11, 9, 18, 30, 30, 22, 0, 27, 36, 16, 21, 12, 28, 36, 31, 36, 25, 3, 33, 8, 34, 21, 30, 11, 34, 34, 23, 24, 36, 13, 7, 22, 22, 12, 18, 30, 14, 6, 34, 5, 13, 10, 9, 34, 25, 28, 8, 22, 30, 14, 11, 0, 30, 7, 4, 27, 25, 11, 32, 29, 12, 26, 11, 7, 24, 1, 35, 32, 5, 32, 31, 36, 17, 14, 28, 36, 1, 21, 36, 34, 33, 12, 35, 5, 6, 14, 25, 28, 18, 11, 26, 36, 33, 0, 13, 19, 15, 18, 23, 3, 36, 15, 19, 7, 36, 23, 11, 28, 15, 14, 36, 5, 10, 33, 34, 10, 20, 20, 23, 21, 31, 9, 30, 32, 16, 6, 2, 20, 36, 25, 33, 4, 24, 24, 31, 12, 27, 33, 26, 4, 35, 32, 36, 23, 0, 13, 28, 3, 16, 25, 11, 16, 14, 25, 17, 23, 26, 27, 2, 9, 29, 29, 7, 7, 8, 26, 25, 2, 28, 29, 15, 6, 28, 35, 14, 16, 18, 12, 36, 26, 33, 28, 20, 20, 22, 35, 11, 1, 8, 8, 33, 28, 29, 6, 35, 29, 8, 23, 12, 35, 22, 18, 4, 27, 30, 21, 0, 15, 15, 2, 9, 2, 4, 17, 18, 18, 20, 24, 31, 18, 32, 3, 15, 30, 36, 34, 27, 5, 4, 5, 9, 7, 34, 24, 24, 29, 25, 9, 12, 14, 5, 22, 1, 7, 6, 29, 14, 15, 5, 30, 29, 10, 24, 16, 18, 1, 7, 6, 29, 14, 15, 5, 30, 29, 10, 24, 16, 18, 13, 0, 7, 7, 7, 13, 0, 25, 4, 9, 26, 13, 1, 25, 31, 24, 24, 8, 16, 24, 27, 32, 29, 11, 31, 35, 1, 13, 8, 18, 10, 4, 34, 35, 23, 17, 27, 16, 29, 9, 3, 29, 2, 17, 6, 15, 8, 30, 6, 16, 25, 12, 30, 15, 15, 1, 1, 22, 20, 7, 27, 15, 12, 36, 10, 21, 8, 27, 25, 26, 19, 24, 23, 35, 1, 9, 22, 19, 17, 31, 25, 32, 0, 3, 31, 34, 0, 14, 5, 20, 5, 23, 34, 30, 16, 15, 33, 30, 7, 24, 29, 5, 5, 16, 10, 21, 27, 28, 19, 0, 24, 17, 26, 2, 34, 26, 35, 15, 22, 16, 28, 8, 2, 31, 7, 1, 11, 32, 2, 19, 14, 0, 32, 13, 21, 28, 16, 17, 10, 1, 34, 18, 36, 5, 31, 19, 0, 16, 14, 19, 3, 12, 17, 9, 27, 12, 5, 18, 0, 1, 2, 33, 17, 30, 24, 16, 14, 20, 2, 36, 18, 22, 32, 32, 9, 15, 29, 24, 30, 20, 27, 22, 14, 26, 36, 15, 32, 17, 1, 4, 30, 21, 18, 22, 33, 15, 13, 13, 16, 23, 26, 12, 25, 10, 34, 13, 12, 14, 17, 0, 3, 36, 20, 18, 4, 17, 9, 23, 32, 13, 19, 27, 16, 28, 10, 31, 12, 24, 2, 27, 12, 15, 6, 17, 9, 22, 22, 30, 26, 20, 21, 18, 11, 17, 9, 12, 36, 30, 16, 12, 9, 17, 3, 14, 30, 12, 9, 10, 11, 35, 12, 13, 1, 8, 14, 24, 22, 12, 32, 5, 7, 32, 8, 27, 14, 33, 4, 19, 15, 35, 20, 14, 35, 1, 18, 15, 20, 24, 35, 3, 32, 11, 18, 31, 8, 28, 33, 25, 9, 0, 15, 10, 22, 25, 18, 20, 30, 2, 25, 20, 11, 22, 36, 9, 11, 24, 5, 33, 17, 17, 0, 1, 28, 25, 22, 26, 33, 8, 31, 0, 17, 2, 28, 24, 35, 29, 15, 6, 30, 3, 34, 18, 7, 1, 14, 21, 4, 14, 0, 35, 26, 24, 8, 22, 10, 5, 11, 22, 0, 12, 11, 1, 27, 17, 14, 6, 2, 19, 30, 24, 24, 32, 34, 8, 18, 10, 22, 23, 14, 22, 19, 7, 30, 3, 10, 5, 22, 24, 16, 29, 31, 4, 21, 31, 0, 29, 2, 1, 26, 36, 16, 5, 4, 28, 24, 12, 5, 9, 20, 5, 16, 5, 36, 14, 13, 20, 15, 14, 2, 14, 29, 27, 12, 30, 35, 9, 36, 36, 11, 21, 29, 27, 20, 21, 16, 22, 35, 6, 12, 8, 25, 1, 15, 11, 23, 1, 25, 9, 14, 28, 21, 9, 31, 29, 32, 34, 5, 31, 34, 24, 3, 18, 28, 35, 19, 7, 36, 3, 5, 25, 31, 36, 15, 20, 32, 27, 34, 7, 11, 31, 24, 4, 32, 35, 27, 36, 5}

var Combined2 = []int{20, 3, 22, 12, 33, 33, 14, 10, 15, 13, 7, 14, 10, 26, 26, 16, 1, 11, 14, 17, 22, 4, 30, 36, 10, 36, 18, 11, 25, 16, 1, 0, 34, 9, 26, 29, 18, 29, 13, 31, 8, 4, 14, 5, 32, 32, 32, 31, 16, 1, 26, 7, 8, 35, 35, 7, 9, 16, 9, 35, 21, 5, 36, 4, 3, 3, 30, 1, 30, 0, 19, 27, 9, 19, 5, 15, 18, 13, 25, 9, 23, 9, 1, 0, 4, 27, 16, 21, 7, 25, 8, 33, 5, 27, 26, 28, 23, 23, 17, 24, 10, 16, 32, 27, 36, 2, 23, 32, 10, 29, 24, 28, 32, 26, 26, 31, 21, 5, 21, 18, 8, 16, 0, 17, 3, 14, 30, 17, 0, 7, 17, 8, 3, 28, 0, 23, 21, 31, 33, 14, 0, 21, 4, 0, 28, 22, 5, 14, 23, 0, 30, 19, 11, 5, 8, 24, 3, 13, 4, 17, 6, 6, 25, 36, 26, 30, 10, 5, 32, 18, 14, 28, 31, 32, 10, 15, 13, 20, 16, 28, 11, 1, 9, 22, 35, 0, 14, 26, 7, 20, 21, 7, 9, 1, 27, 17, 26, 17, 0, 32, 4, 22, 31, 12, 33, 27, 31, 23, 11, 18, 36, 10, 28, 6, 11, 6, 6, 30, 29, 0, 13, 24, 23, 28, 24, 21, 14, 8, 26, 13, 30, 26, 29, 0, 20, 0, 14, 18, 17, 5, 11, 31, 9, 20, 24, 26, 19, 32, 19, 7, 11, 29, 2, 31, 1, 5, 20, 15, 31, 13, 0, 21, 6, 25, 30, 0, 10, 26, 18, 3, 3, 14, 17, 6, 29, 30, 11, 32, 10, 34, 5, 22, 36, 7, 19, 20, 0, 10, 3, 9, 16, 30, 6, 6, 11, 23, 26, 23, 36, 18, 2, 7, 21, 4, 1, 27, 0, 25, 17, 18, 30, 16, 4, 33, 22, 31, 16, 23, 4, 21, 10, 13, 32, 22, 35, 22, 11, 15, 31, 12, 18, 8, 34, 17, 32, 13, 16, 27, 18, 23, 18, 34, 6, 15, 5, 29, 19, 5, 26, 32, 17, 1, 8, 15, 32, 6, 30, 10, 5, 24, 35, 0, 25, 21, 27, 23, 1, 3, 23, 20, 30, 18, 15, 8, 15, 18, 24, 27, 22, 27, 10, 32, 5, 0, 30, 24, 12, 31, 2, 12, 12, 10, 15, 17, 28, 34, 32, 11, 10, 29, 33, 16, 4, 6, 11, 22, 3, 17, 9, 25, 13, 13, 13, 11, 8, 20, 0, 3, 35, 22, 36, 0, 24, 20, 1, 22, 33, 32, 8, 33, 14, 9, 2, 30, 29, 9, 17, 26, 6, 30, 14, 13, 24, 28, 20, 36, 4, 7, 13, 15, 36, 3, 22, 18, 14, 26, 34, 9, 7, 21, 8, 12, 3, 2, 20, 13, 11, 16, 20, 1, 4, 30, 29, 10, 12, 3, 16, 20, 10, 35, 23, 12, 31, 34, 20, 2, 1, 19, 36, 9, 23, 34, 15, 13, 15, 25, 34, 33, 11, 20, 24, 17, 28, 9, 2, 21, 32, 17, 4, 4, 0, 29, 22, 8, 25, 21, 1, 33, 15, 30, 10, 0, 29, 22, 0, 28, 7, 13, 8, 21, 12, 12, 2, 16, 26, 31, 18, 16, 0, 30, 12, 12, 5, 20, 4, 22, 3, 25, 11, 4, 15, 22, 24, 25, 31, 10, 15, 35, 7, 23, 13, 33, 17, 14, 15, 6, 10, 6, 7, 32, 12, 31, 3, 4, 5, 26, 12, 9, 17, 10, 31, 6, 36, 13, 21, 1, 9, 25, 30, 35, 22, 23, 21, 8, 28, 7, 34, 22, 36, 30, 7, 2, 16, 0, 12, 30, 23, 15, 23, 15, 34, 4, 16, 29, 0, 11, 27, 34, 35, 33, 8, 3, 18, 5, 10, 11, 26, 34, 25, 11, 29, 27, 24, 35, 14, 28, 30, 3, 7, 17, 35, 30, 30, 18, 28, 33, 3, 8, 36, 32, 7, 34, 21, 5, 21, 12, 6, 18, 28, 15, 1, 29, 29, 12, 35, 27, 18, 20, 8, 6, 4, 9, 6, 2, 19, 2, 13, 7, 18, 33, 35, 10, 19, 16, 27, 34, 7, 25, 21, 5, 15, 35, 32, 27, 21, 0, 29, 34, 18, 28, 32, 24, 22, 30, 20, 16, 29, 28, 22, 27, 34, 0, 11, 11, 16, 16, 24, 31, 31, 27, 17, 32, 21, 12, 18, 33, 8, 32, 28, 8, 34, 6, 14, 36, 12, 26, 33, 4, 11, 19, 10, 26, 6, 32, 24, 24, 31, 12, 17, 32, 34, 3, 34, 20, 23, 31, 35, 16, 36, 32, 20, 14, 12, 10, 0, 9, 26, 22, 30, 8, 27, 17, 30, 10, 27, 9, 33, 12, 26, 11, 5, 18, 4, 0, 28, 20, 21, 4, 16, 25, 31, 22, 15, 13, 20, 16, 29, 5, 13, 18, 13, 27, 20, 7, 6, 36, 5, 19, 16, 27, 20, 16, 28, 26, 2, 27, 20, 17, 29, 24, 3, 27, 23, 14, 6, 31, 19, 7, 1, 33, 8, 14, 33, 10, 19, 11, 7, 20, 7, 21, 34, 15, 29, 3, 1, 21, 16, 22, 19, 27, 5, 15, 21, 6, 14, 16, 26, 4, 28, 19, 10, 9, 3, 13, 20, 19, 30, 13, 31, 14, 4, 8, 5, 3, 27, 18, 12, 15, 9, 29, 2, 20, 1, 34, 28, 16, 21, 7, 23, 8, 14, 24, 19, 21, 10, 32, 26, 14, 15, 20, 23, 16, 30, 26, 24, 35, 22, 1, 2, 25, 20, 20, 33, 19, 23, 20, 29, 13, 7, 33, 6, 20, 10, 28, 17, 2, 36, 16, 4, 34, 32, 29, 27, 25, 13, 9, 30, 12, 27, 3, 14, 32, 29, 6, 10, 15, 10, 17, 0, 20, 19, 15, 12, 11, 4, 27, 14, 0, 9, 25, 1, 14, 20, 14, 9, 22, 34, 15, 34, 10, 1, 22, 20, 3, 24, 30, 8, 26, 13, 22, 31, 5, 11, 35, 19, 7, 7, 16, 36, 34, 27, 0, 8, 22, 10, 11, 32, 21, 19, 34, 12, 31, 22, 26, 30, 0, 30, 1, 2, 22, 30, 36, 14, 7, 28, 29, 7, 15, 0, 35, 21, 31, 24, 15, 20, 4, 19, 25, 34, 14, 31, 16, 36, 23, 20, 4, 11, 26, 22, 19, 25, 2, 9, 18, 16, 33, 25, 22, 31, 31, 21, 26, 5, 13, 11, 18, 0, 16, 16, 21, 34, 34, 5, 33, 24, 23, 21, 1, 24, 21, 11, 35, 9, 9, 29, 31, 6, 19, 7, 35, 28, 11, 30, 33, 28, 2, 14, 29, 25, 18, 9, 35, 6, 23, 24, 10, 20, 28, 26, 31, 3, 4, 3, 17, 26, 34, 10, 4, 16, 26, 29, 36, 13, 9, 24, 29, 4, 2, 19, 4, 35, 14, 18, 10, 32, 4, 13, 4, 27, 20, 28, 28, 5, 36}
