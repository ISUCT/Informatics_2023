namespace CodeWarsExercises.Codewars
{
    internal class TenGreenBottlesExercise : IExercise
    {
        private readonly string[] letters = new string[]
        {
            "No",
            "One",
            "Two",
            "Three",
            "Four",
            "Five",
            "Six",
            "Seven",
            "Eeight",
            "Nine",
            "Ten",
        };

        public string Name => "Ten Green Bottles";

        public void Run()
        {
            int bottleCount = ConsoleUtils.ReadInt("Enter numeber from 1 to 10: ");

            if (bottleCount < 1 || bottleCount > 10)
            {
                throw new ArgumentException("number must be lie between 1 and 10.");
            }

            while (bottleCount > 0)
            {
                string bottleSpelling;

                bottleSpelling = GetBottleSpelling(bottleCount);
                string lineOne = $"{GetTextRepresentation(bottleCount, false)} green {bottleSpelling} hanging on the wall,";

                Console.WriteLine(lineOne);
                Console.WriteLine(lineOne);

                Console.WriteLine($"And if one green bottle should accidentally fall,");

                --bottleCount; // Oh no, bottle fall!

                bottleSpelling = GetBottleSpelling(bottleCount);
                Console.WriteLine($"There'll be {GetTextRepresentation(bottleCount, true)} green {bottleSpelling} hanging on the wall.");
                Console.WriteLine();
            }
        }

        private static string GetBottleSpelling(int bottleCount) => bottleCount != 1 ? "bottles" : "bottle";

        private string GetTextRepresentation(int n, bool lower)
        {
            if (lower)
            {
                return letters[n].ToLower();
            }

            return letters[n];
        }
    }
}
