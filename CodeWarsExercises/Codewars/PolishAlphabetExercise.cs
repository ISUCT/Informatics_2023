using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace CodeWarsExercises.Codewars {
    internal class PolishAlphabetExercise : IExercise {
        public string Name => "Polish alphabet";

        private static readonly Dictionary<char, char> Alphabet = new() {
            { 'ą', 'a' },
            { 'ć', 'c' },
            { 'ę', 'e' },
            { 'ł', 'l' },
            { 'ń', 'n' },
            { 'ó', 'o' },
            { 'ś', 's' },
            { 'ź', 'z' },
            { 'ż', 'z' }
        };

        public static string Englify(string input) {
            var output = new StringBuilder(input);

            for (int i = 0; i < input.Length; ++i)
                if (Alphabet.TryGetValue(input[i], out char c))
                    output[i] = c;

            return output.ToString();
        }

        public void Run() {
            Console.Write("Enter string to englify: ");
            Console.WriteLine(Englify(Console.ReadLine()!));
        }
    }
}
