using System.Globalization;

namespace CodeWarsExercises {
    internal static class ConsoleUtils {

        public static int ReadInt(string? message = null) {
            // TODO: Make this function exception-safe.
            if (message is object)
                Console.Write(message);

            return int.Parse(Console.ReadLine()!, CultureInfo.InvariantCulture);
        }

        public static int[] ReadOneDimensionalArray(string? v = null) {
            if (v is object)
                Console.Write(v);

            return Console.ReadLine()?.
                            Split(' ', StringSplitOptions.RemoveEmptyEntries | StringSplitOptions.TrimEntries).
                            Select(x =>
                                int.TryParse(x, NumberStyles.Number, CultureInfo.InvariantCulture, out int y) ?
                                y :
                                throw new ArgumentException($"'{x}' can't be parsed as int.")
                                ).
                                ToArray() ??
                                throw new InvalidOperationException("Execution was interrupted.");
        }
    }
}
