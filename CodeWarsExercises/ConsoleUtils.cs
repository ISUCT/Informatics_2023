using System.Globalization;

namespace CodeWarsExercises {
    internal static class ConsoleUtils {

        public static int ReadInt() {
            // TODO: Make this function exception-safe.
            return int.Parse(Console.ReadLine()!, CultureInfo.InvariantCulture);
        }
    }
}
