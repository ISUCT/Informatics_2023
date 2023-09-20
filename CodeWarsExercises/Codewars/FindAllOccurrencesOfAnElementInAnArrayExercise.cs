using System.Globalization;


namespace CodeWarsExercises.Codewars {
    internal class FindAllOccurrencesOfAnElementInAnArrayExercise : IExercise {
        public string Name => "Find all occurrences of an element in an array";

        public static int[] FindAllOccurrences(int[] array, int item) {
            List<int> list = new List<int>();

            for (int index = 0; index < array.Length; index++) {
                if (array[index] == item) {
                    list.Add(index);
                }
            }

            return list.ToArray();
        }

        public void Run() {
            Console.Write("Enter array of integers where each item is splited by space: ");


            int[] ints = Console.ReadLine()?.
                Split(' ', StringSplitOptions.RemoveEmptyEntries | StringSplitOptions.TrimEntries).
                Select(x =>
                    int.TryParse(x, NumberStyles.Number, CultureInfo.InvariantCulture, out int y) ?
                    y :
                    throw new ArgumentException($"'{x}' can't be parsed as int.")
                    ).
                    ToArray() ??
                    throw new InvalidOperationException("Execution was interrupted.");

            int item = ConsoleUtils.ReadInt("Enter integer to search: ");

            int[] indices = FindAllOccurrences(ints, item);

            Console.WriteLine($"Indices of item: {indices.FormatAsString()}");

        }
    }
}
