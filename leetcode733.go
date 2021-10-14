/*
 * @Descripttion:有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在 0 到 65535 之间。

给你一个坐标 (sr, sc) 表示图像渲染开始的像素值（行 ，列）和一个新的颜色值 newColor，让你重新上色这幅图像。

为了完成上色工作，从初始坐标开始，记录初始坐标的上下左右四个方向上像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对应四个方向上像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为新的颜色值。

最后返回经过上色渲染后的图像。

 * @version:
 * @Author: cm.d
 * @Date: 2021-10-14 14:43:09
 * @LastEditors: cm.d
 * @LastEditTime: 2021-10-14 17:21:56
*/
package myleetcode

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	image[sr][sc] = newColor
	flag := make([][]int, len(image))
	for i, j := range image {
		flag[i] = make([]int, len(j))
	}
	flag[sr][sc] = 1
	floodFillWithFlag(image, sr, sc+1, color, newColor, flag)
	floodFillWithFlag(image, sr+1, sc, color, newColor, flag)
	floodFillWithFlag(image, sr-1, sc, color, newColor, flag)
	floodFillWithFlag(image, sr, sc-1, color, newColor, flag)
	return image
}

func floodFillWithFlag(image [][]int, sr int, sc int, color, newColor int, flag [][]int) [][]int {

	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[sr]) {
		return image
	}
	if image[sr][sc] != color {
		return image
	}
	if flag[sr][sc] == 1 {
		return image
	}
	image[sr][sc] = newColor
	flag[sr][sc] = 1
	floodFillWithFlag(image, sr, sc+1, color, newColor, flag)
	floodFillWithFlag(image, sr+1, sc, color, newColor, flag)
	floodFillWithFlag(image, sr-1, sc, color, newColor, flag)
	floodFillWithFlag(image, sr, sc-1, color, newColor, flag)

	return image
}

func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	return floodFill(image, sr, sc, newColor)
}
