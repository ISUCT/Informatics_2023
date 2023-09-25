namespace CodeWarsExercises.Codewars
{
    internal class CountingSheepExercise : IExercise
    {
        public string Name => "Counting sheep...";

        public static int CountSheeps(bool[] places)
        {
            if (places is null)
            {
                throw new ArgumentNullException(nameof(places));
            }

            int counter = 0;

            foreach (var sheep in places)
            {
                if (sheep)
                {
                    ++counter;
                }
            }

            return counter;
        }

        public void Run()
        {
            {
                string message =
                    "Describe sheep places splited by space. \n" +
                    "enter 'true' per place if sheep is in it's place or 'false' otherwise: ";

                Console.Write(message);
            }

            bool[] sheeps = Console.ReadLine()?.
                Split(
                ' ',
                StringSplitOptions.RemoveEmptyEntries | StringSplitOptions.TrimEntries).
                Select(x =>
                {
                    return x.ToLower() switch
                    {
                        "true" => true,
                        "false" => false,
                        _ => throw new InvalidOperationException($"Can't parse {x} as true or false."),
                    };
                }).ToArray() ?? throw new InvalidOperationException("Input was interrupted.");

            Console.WriteLine($"Sheep count: {CountSheeps(sheeps)}");
        }
    }
}
