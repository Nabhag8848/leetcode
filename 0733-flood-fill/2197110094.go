func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    st_color := image[sr][sc]

    if color == st_color {
        return image
    }

    visited := make([][]bool, len(image))

    for i := range visited {
        visited[i] = make([]bool, len(image[i]))
    }

    dfs(image, visited, color, sr, sc, st_color)
    return image
}

func dfs(image [][]int, visited [][]bool, color, sr, sc, st_color int) {
    visited[sr][sc] = true
    image[sr][sc] = color

    if sr < len(image) - 1 && !visited[sr + 1][sc] && image[sr + 1][sc] == st_color {
        dfs(image, visited, color, sr + 1, sc, st_color)
    }

    if sr > 0 && !visited[sr - 1][sc] && image[sr - 1][sc] == st_color {
        dfs(image, visited, color, sr - 1, sc, st_color)
    }   

    if sc < len(image[0]) - 1 && !visited[sr][sc + 1] && image[sr][sc + 1] == st_color {
        dfs(image, visited, color, sr, sc + 1, st_color)
    }

    if sc > 0 && !visited[sr][sc - 1] && image[sr][sc - 1] == st_color {
        dfs(image, visited, color, sr, sc - 1, st_color)
    }

}

