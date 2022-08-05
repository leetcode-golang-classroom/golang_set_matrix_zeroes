# golang_set_matrix_zeroes

Given an `m x n` integer matrix `matrix`, if an element is `0`, set its entire row and column to `0`'s.

You must do it [in place](https://en.wikipedia.org/wiki/In-place_algorithm).

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/08/17/mat1.jpg](https://assets.leetcode.com/uploads/2020/08/17/mat1.jpg)

```
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/08/17/mat2.jpg](https://assets.leetcode.com/uploads/2020/08/17/mat2.jpg)

```
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

```

**Constraints:**

- `m == matrix.length`
- `n == matrix[0].length`
- `1 <= m, n <= 200`
- `231 <= matrix[i][j] <= 231 - 1`

**Follow up:**

- A straightforward solution using `O(mn)` space is probably a bad idea.
- A simple improvement uses `O(m + n)` space, but still not the best solution.
- Could you devise a constant space solution?

## 解析

給定一個整數 matrix 

對每個 matrix[i][j] = 0 對應的 i row 與 j column 的值全部更改為 0

要求寫一個演算法去做以上的操作


![](https://i.imgur.com/PeXmEPO.png)

為了避免修改後的值 去影響判斷式

所以最直觀的作法是先複製原 matrix 來作為參照值

如下圖

![](https://i.imgur.com/cwWeFj3.png)

然後逐步依據所在 row, col 中是否有需要改成零做更改

然而可以發現 其實每次只要有一個 matrix[row][col] 影響的就是整個 row 與 col

所以其實只需要紀錄 哪個 row 以及哪一個 col 需要做修改

可以簡化如下

![](https://i.imgur.com/u2ZNdq1.png)

這樣來做就可以讓空間複雜度從 O(m*n) 降低至 O(m+n)

而時間複雜度是 O(m*n)

然而 可以發現只要從上而下 由左至右來做比較

可以利用第0 列第0行來做標記

標記完再根據標記過的行列來做修改

需要注意的是 因為 matrix[0][0] 因為會同時在第 0 列與第0行交會點

所以需要額外多一個儲存空間來標記 

以下多使用一個 firstColumnZero 用來標記第 0 column 是否需要改成 0

matrix[0][0] 代表第 0 row 是否要改成 0

如下

![](https://i.imgur.com/qRR0REr.png)

透過這個方法 先標記

然後在根據標記去修改對應的值

就可以把 空間複雜度降低到 O(1)

而時間複雜度因為需要跑完整個 matrix 所以還是 O(m*n)

## 程式碼
```go
package sol

func setZeroes(matrix [][]int) {
	firstColumnZero := false
	ROW, COL := len(matrix), len(matrix[0])
	// mark matrix[0][col] = 0 , matrix[row][0] = 0 if matrix[row][col] = 0
	for row := 0; row < ROW; row++ {
		for col := 0; col < COL; col++ {
			if matrix[row][col] == 0 {
				// mark row
				matrix[row][0] = 0
				// mark col
				if col == 0 {
					firstColumnZero = true
				} else {
					matrix[0][col] = 0
				}
			}
		}
	}
	// 1..ROW-1, 1..COL-1 , if matrix[row][0] = 0 || matrix[0][col] = 0, matrix[row][col] = 0
	for row := 1; row < ROW; row++ {
		for col := 1; col < COL; col++ {
			if matrix[row][0] == 0 || matrix[0][col] == 0 {
				if matrix[row][col] != 0 {
					matrix[row][col] = 0
				}
			}
		}
	}
	// check first row
	if matrix[0][0] == 0 {
		for col := 0; col < COL; col++ {
			if matrix[0][col] != 0 {
				matrix[0][col] = 0
			}
		}
	}
	if firstColumnZero {
		for row := 0; row < ROW; row++ {
			if matrix[row][0] != 0 {
				matrix[row][0] = 0
			}
		}
	}
}
```
## 困難點

1. 需要利用比較順序思考出利用原本矩陣來做標記的作法

## Solve Point

- [x]  初始化 firstColumnZero := false 用來紀錄第1列是否需要改成 0
- [x]  初始化 ROW = len(matrix), COL = len(matrix[0])
- [x]  從 row := 0 , col := 0 比較到 row = ROW - 1, col = COL - 1
- [x]  先把 matrix[row][col] = 0 的值 如果 col ≠ 0 , 則matrix[0][col] = 0 還有 matrix[row][0] = 0, 如果 col = 0 則標記在 firstColumnZero = true
- [x]  先把 row := 1..ROW-1, col := 1..COL-1, if  matrix[0][col] = 0 || matrix[row][0] = 0 更新 matrix[row][col] = 0
- [x]  if matrix[0][0] = 0 更新第0 row 所有值 = 0
- [x]  if firstColumnZero = 0 更新第0 column 所有值 = 0