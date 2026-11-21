func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    st_color := image[sr][sc]

    if color == st_color {
        return image
    }

    dfs(image, color, sr, sc, st_color)
    return image
}

func dfs(image [][]int, color, sr, sc, st_color int) {
    image[sr][sc] = color

    if sr < len(image) - 1 && image[sr + 1][sc] == st_color {
        dfs(image, color, sr + 1, sc, st_color)
    }

    if sr > 0 && image[sr - 1][sc] == st_color {
        dfs(image, color, sr - 1, sc, st_color)
    }   

    if sc < len(image[0]) - 1 && image[sr][sc + 1] == st_color {
        dfs(image, color, sr, sc + 1, st_color)
    }

    if sc > 0 && image[sr][sc - 1] == st_color {
        dfs(image, color, sr, sc - 1, st_color)
    }

}

