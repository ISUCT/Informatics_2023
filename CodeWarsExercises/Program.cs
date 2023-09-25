namespace CodeWarsExercises
{
    using System.Globalization;
    using CodeWarsExercises.Codewars;

    using static System.Console;

    public static class Program
    {
        // This way we don't need to care about exercise's internal state undefined.
        private static readonly Func<IExercise>[] ExerciseFactories = new Func<IExercise>[]
        {
            () => new EvenOrOddExercise(),
            () => new CountingSheepExercise(),
            () => new CountTheMonkeysExercise(optimized: false),
            () => new CountTheMonkeysExercise(optimized: true),
            () => new TenGreenBottlesExercise(),
            () => new SchoolPaperworkExercise(),
            () => new IsHeGonnaSurviveExercise(),
            () => new NumberToStringExercise(),
            () => new PolishAlphabetExercise(),
            () => new FindAllOccurrencesOfAnElementInAnArrayExercise(),
            () => new SumOfMinimumsExercise(SumOfMinimumsExercise.Mode.UseNestedArray),
            () => new SumOfMinimumsExercise(SumOfMinimumsExercise.Mode.Use2DArrays),
        };

        public static void Main()
        {
            IExercise[] exerciseInstances = new IExercise[ExerciseFactories.Length];

            while (true)
            {
                WriteLine("Choose an exercise you wish to test:");

                for (int index = 0; index < ExerciseFactories.Length; ++index)
                {
                    IExercise instance = ExerciseFactories[index].Invoke();
                    exerciseInstances[index] = instance;

                    WriteLine($"{index:00}. {instance.Name}");
                }

                WriteLine();

                while (true)
                {
                    Write("Please enter a number or 'e' symbol to exit: ");

                    string? userInput = ReadLine()?.Trim();

                    if (userInput is null || userInput == "e")
                    {
                        // user cut their input, so stop execution.
                        return;
                    }

                    bool inputIsCorrect = int.TryParse(
                        userInput,
                        NumberStyles.Integer,
                        CultureInfo.InvariantCulture,
                        out int index);

                    if (inputIsCorrect)
                    {
                        if (index < 0 || index >= ExerciseFactories.Length)
                        {
                            WriteLine($"The index must lies between 0 and {ExerciseFactories.Length}.");
                            WriteLine($"[0, {ExerciseFactories.Length})");
                            continue;
                        }

                        Clear();

                        try
                        {
                            IExercise selectedExercise = exerciseInstances[index];

                            WriteLine($"===[{selectedExercise.Name}]===");
                            WriteLine();

                            selectedExercise.Run();
                        }
                        catch (Exception e)
                        {
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
        }
    }
}