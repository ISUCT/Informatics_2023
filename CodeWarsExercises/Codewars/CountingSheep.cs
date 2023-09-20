namespace CodeWarsExercises.Codewars {
    internal class CountingSheep : IExercise {
        public string Name => "Counting sheep...";

        public static int CountSheeps(bool?[] places) {
            if (places is null)
                throw new ArgumentNullException(nameof(places));

            int counter = 0;

            foreach (var sheep in places) {
                if (!sheep.HasValue)
                    throw new ArgumentException("There's musn't be nulls for sheeps.");

                if (sheep.Value)
                    ++counter;
            }

            return counter;
        }

        public void Run() {

            Console.Write("Describe sheep places splited by space. \n" +
                "enter 'true' per place if sheep is in it's place or 'false' otherwise: ");

            bool?[]? sheeps = Console.ReadLine()?.
                Split(
                ' ',
                StringSplitOptions.RemoveEmptyEntries | StringSplitOptions.TrimEntries
                ).Select<string, bool?>(x => {
                    switch (x.ToLower()) {
                        case "true": return true;
                        case "false": return false;
                        default: return null;
                    }
                })?.ToArray();

            Console.WriteLine($"Sheep count: {CountSheeps(sheeps!)}");
        }
    }
}
