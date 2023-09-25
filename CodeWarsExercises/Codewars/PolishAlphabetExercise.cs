namespace CodeWarsExercises.Codewars
{
    using System.Text;

    internal class PolishAlphabetExercise : IExercise
    {
        private static readonly Dictionary<char, char> Alphabet = new Dictionary<char, char>()
        {
            { 'ą', 'a' },
            { 'ć', 'c' },
            { 'ę', 'e' },
            { 'ł', 'l' },
            { 'ń', 'n' },
            { 'ó', 'o' },
            { 'ś', 's' },
            { 'ź', 'z' },
            { 'ż', 'z' },
        };

        public string Name => "Polish alphabet";

        public static string Englify(string input)
        {
            var output = new StringBuilder(input);

            for (int i = 0; i < input.Length; ++i)
            {
                if (Alphabet.TryGetValue(input[i], out char c))
                {
                    output[i] = c;
                }
            }

            return output.ToString();
        }

        public void Run()
        {
            Console.Write("Enter string to englify: ");
            Console.WriteLine(Englify(Console.ReadLine() !));
        }
    }
}
