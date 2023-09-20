using CodeWarsExercises;
using CodeWarsExercises.Codewars;
using System.Globalization;

using static System.Console;

IExercise[] exercises = new IExercise[] {
    new EvenOrOdd()
};

while (true) {
    WriteLine("Choose an exercise you wish to test:");

    for (int index = 0; index < exercises.Length; ++index) {
        WriteLine($"{index:00}. {exercises[index].Name}");
    }

    WriteLine();

    while (true) {
        Write("Please enter a number or 'e' symbol to exit: ");

        string? userInput = ReadLine()?.Trim();

        if (userInput is null || userInput == "e")
            return; // user cut their input, so stop execution.

        bool inputIsCorrect = int.TryParse(
            userInput,
            NumberStyles.Integer,
            CultureInfo.InvariantCulture,
            out int index
            );

        if (inputIsCorrect) {
            if (index < 0 || index >= exercises.Length) {
                WriteLine($"The index must lies between 0 and {exercises.Length}.");
                WriteLine($"[0, {exercises.Length})");
                continue;
            }

            Clear();
            exercises[index].Run();

            WriteLine("Hit Enter to continue...");
            ReadLine();
            Clear();
        }
    }
}