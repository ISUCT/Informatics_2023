using System.Globalization;

namespace CodeWarsExercises {
    internal static class ConsoleUtils {

        public static int ReadInt(string? message = null) {
            // TODO: Make this function exception-safe.
            if (message is object)
                Console.Write(message);

            return int.Parse(Console.ReadLine()!, CultureInfo.InvariantCulture);
        }
    }
}
