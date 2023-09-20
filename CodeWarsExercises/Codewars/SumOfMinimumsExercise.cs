namespace CodeWarsExercises.Codewars {
    internal class SumOfMinimumsExercise : IExercise {

        private interface IStrategy {
            string Name { get; }
            int Run();
        }

        private class NestedStrategy : IStrategy {
            public string Name => "nested arrays";

            public int Run() {
                int[][] matrix = ReadNested();
                return SumOfMinimums(matrix);
            }
        }

        private class TwoDimensionalStrategy : IStrategy {
            public string Name => "2d arrays";

            public int Run() {
                int[,] matrix = Read2D();
                return SumOfMinimums(matrix);
            }

        }

        public enum Mode {
            UseNestedArray,
            Use2DArrays
        }
        public string Name => $"Sum of minimums (using {m_strategy.Name})";

        IStrategy m_strategy;

        public SumOfMinimumsExercise(Mode mode) {
            switch (mode) {
                case Mode.UseNestedArray:
                    m_strategy = new NestedStrategy();
                    break;

                case Mode.Use2DArrays:
                    m_strategy = new TwoDimensionalStrategy();
                    break;
                default: throw new ArgumentException("Unknown mode.");
            }
        }

        public static int SumOfMinimums(int[][] nesetedArray) {
            int sum = 0;

            foreach (var row in nesetedArray) {

                if (row.Length <= 0)
                    continue;

                int min = int.MaxValue;

                foreach (var item in row) {
                    if (item < min)
                        min = item;
                }

                sum += min;
            }

            return sum;
        }

        public static int SumOfMinimums(int[,] twoDimensionalArray) {
            if (twoDimensionalArray.GetLength(1) <= 0)
                return 0;

            int sum = 0;

            for (int row = 0; row < twoDimensionalArray.GetLength(0); ++row) {
                int min = int.MaxValue;

                for (int column = 0; column < twoDimensionalArray.GetLength(1); ++column) {
                    int item = twoDimensionalArray[row, column];
                    if (item < min)
                        min = item;
                }

                sum += min;
            }

            return sum;
        }

        private static int[][] ReadNested() {
            int rows = ConsoleUtils.ReadInt("Enter row count: ");
            int[][] matrix = new int[rows][];

            for (int row = 0; row < rows; ++row)
                matrix[row] = ConsoleUtils.ReadOneDimensionalArray(RowGreeating(row));

            return matrix;
        }

        private static string RowGreeating(int row) => $"Enter items of row #{row + 1} as integer array where each item is space seperated: ";

        private static int[,] Read2D() {
            int rows = ConsoleUtils.ReadInt("Enter row count: ");
            int columns = ConsoleUtils.ReadInt("Enter column count: ");

            int[,] matrix = new int[rows, columns];

            for (int row = 0; row < rows; ++row) {
                int[] rowValues = ConsoleUtils.ReadOneDimensionalArray(RowGreeating(row));
                if (rowValues.Length != columns)
                    throw new ArgumentException($"Columns ({columns}) and items ({rowValues.Length}) missmath.");

                for (int column = 0; column < columns; ++column)
                    matrix[row, column] = rowValues[column];
            }

            return matrix;
        }

        public void Run() {
            int mininmum = m_strategy.Run();
            Console.WriteLine($"Sum of row's minimums is {mininmum}");
        }
    }
}
