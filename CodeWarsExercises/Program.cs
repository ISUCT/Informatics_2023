using CodeWarsExercises;
using CodeWarsExercises.Codewars;
using System.Globalization;

using static System.Console;

// This way we don't need to care about exercise's internal state undefined.
Func<IExercise>[] exerciseFactories = new Func<IExercise>[] {
    () => new EvenOrOddExercise(),
    () => new CountingSheepExercise(),
    () => new CountTheMonkeysExercise(optimized: false),
    () => new CountTheMonkeysExercise(optimized: true),
    () => new TenGreenBottlesExercise(),
    () => new SchoolPaperworkExercise(),
    () => new IsHeGonnaSurviveExercise(),
    () => new NumberToStringExercise(),
    () => new PolishAlphabetExercise(),
    () => new FindAllOccurrencesOfAnElementInAnArrayExercise()
};

IExercise[] exerciseInstances = new IExercise[exerciseFactories.Length];

while (true) {
    WriteLine("Choose an exercise you wish to test:");

    for (int index = 0; index < exerciseFactories.Length; ++index) {
        IExercise instance = exerciseFactories[index].Invoke();
        exerciseInstances[index] = instance;

        WriteLine($"{index:00}. {instance.Name}");
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
            if (index < 0 || index >= exerciseFactories.Length) {
                WriteLine($"The index must lies between 0 and {exerciseFactories.Length}.");
                WriteLine($"[0, {exerciseFactories.Length})");
                continue;
            }

            Clear();

            try {
                IExercise selectedExercise = exerciseInstances[index];
                
                WriteLine($"===[{selectedExercise.Name}]===");
                WriteLine();

                selectedExercise.Run();
            }
            catch(Exception e) {
                WriteLine("Exception was occured during execution:");
                WriteLine(e.ToString());
            }

            WriteLine("Hit Enter to continue...");
            ReadLine();
            Clear();
            break;
        }
    }
}