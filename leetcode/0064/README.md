dp,直接在原数组上面操作，先将第一行和第一列加上之前的数，然后遍历中间的二维数组，每次加上左边或者上边的较小数，最后返回grid[row-1][col-1]即可