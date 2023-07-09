package arrays;

public class MatrixZeros {

    public void setZeroes(int[][] matrix) {

        // 1. find the rows and columns that have 0s
        // 2. set those rows and columns to 0s
        // 3. return the matrix

        int rows = matrix.length;
        int cols = matrix[0].length;

        boolean[] row = new boolean[rows];
        boolean[] col = new boolean[cols];

        for (int i = 0; i < rows; i++) {
            for (int j = 0; j < cols; j++) {
                if (matrix[i][j] == 0) {
                    row[i] = true;
                    col[j] = true;
                }
            }
        }
        for (int i = 0; i < rows; ++i) {
            if (row[i]) {
                for (int j = 0; j < cols; ++j) {
                    matrix[i][j] = 0;
                }
            }
        }
        for (int i = 0; i < cols; ++i) {
            if (col[i]) {
                for (int j = 0; j < rows; ++j) {
                    matrix[j][i] = 0;
                }
            }
        }
    }

    public static void main(String[] args) {
        MatrixZeros mz = new MatrixZeros();
        int[][] matrix = { { 1, 1, 1 }, { 1, 0, 1 }, { 1, 1, 1 } };
        mz.setZeroes(matrix);
        for (int i = 0; i < matrix.length; i++) {
            for (int j = 0; j < matrix[0].length; j++)
                System.out.print(matrix[i][j] + " ");
            System.out.println();
        }
    }
}
