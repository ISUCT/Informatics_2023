namespace CodeWarsExercises.Codewars
{
    internal class SchoolPaperworkExercise : IExercise
    {
        public string Name => "School Paperwork";

        public static int GetPaperCount(int classmateCount, int pageCount)
        {
            if (classmateCount <= 0 || pageCount <= 0)
            {
                return 0;
            }

            return classmateCount * pageCount;
        }

        public void Run()
        {
            int classmateCount = ConsoleUtils.ReadInt("Enter classmate count: ");
            int paperCount = ConsoleUtils.ReadInt("Enter paper count: ");

            int blankPaperCount = GetPaperCount(classmateCount, paperCount);

            string blankSpelling = blankPaperCount != 1 ? "blanks" : "blank";
            Console.WriteLine($"You need {blankPaperCount} paper {blankSpelling}.");
        }
    }
}
