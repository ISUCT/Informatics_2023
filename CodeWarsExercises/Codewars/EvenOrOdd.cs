namespace CodeWarsExercises.Codewars {
    internal class EvenOrOdd : IExercise {
        public string Name => "Even or odd";

        public void Run() {
            // Note: 0 is even.
            Console.WriteLine(ConsoleUtils.ReadInt("Please, enter number: ") % 2 == 0 ? "Even" : "Odd");
        }
    }
}
