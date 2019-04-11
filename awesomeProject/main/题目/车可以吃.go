package 题目

// 计算一次移动中可以吃到的数量，上下左右都遍历一遍
func numRookCaptures(board [][]byte) int {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'R' {
				return cap(board, i, j, 0, 1) + cap(board, i, j, 0, -1) + cap(board, i, j, 1, 0) + cap(board, i, j, -1, 0)
			}
		}
	}
}
func cap(board [][]byte, x, y, dx, dy int) int {
	for x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && board[x][y] != 'B' {
		if board[x][y] == 'p' {
			return 1
		}
		x += dx
		y += dy
	}
	return 0
}

//class Solution {
//    public int numRookCaptures(char[][] board) {
//        for(int i=0;i<board.length;i++){
//            for(int j=0;j<board[i].length;j++){
//                //找到R的位置
//                if(board[i][j]=='R'){
//                    //以R 为原点建立坐标系
//                    //依次向上找,向下找,向右找,向左找
//                    return cap(board,i,j,0,1)+cap(board,i,j,0,-1)+cap(board,i,j,1,0)+cap(board,i,j,-1,0);
//                }
//            }
//        }
//        return 0;
//    }
//    public int cap(char[][] a,int x,int y,int dx,int dy){
//        /*参数说明
//         *a为原数组矩阵
//         *x,y为R的坐标
//         *dx,dy为增长步长
//        */
//        while(x>=0 && x<a.length && y>=0 && y<a[x].length && a[x][y]!='B'){
//            if(a[x][y]=='p'){
//                return 1;
//            }
//            x+=dx;
//            y+=dy;
//        }
//        return 0;
//    }
//}
