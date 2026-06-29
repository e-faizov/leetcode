package easy

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	oldColor := image[sr][sc]
	if color == oldColor {
		return image
	}
	image[sr][sc] = color

	if sr > 0 && image[sr-1][sc] == oldColor {
		image = floodFill(image, sr-1, sc, color)
	}
	if sr+1 < len(image) && image[sr+1][sc] == oldColor {
		image = floodFill(image, sr+1, sc, color)
	}
	if sc > 0 && image[sr][sc-1] == oldColor {
		image = floodFill(image, sr, sc-1, color)
	}
	if sc+1 < len(image[0]) && image[sr][sc+1] == oldColor {
		image = floodFill(image, sr, sc+1, color)
	}

	return image
}
