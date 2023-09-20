using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace CodeWarsExercises.Codewars {
    internal class EvenOrOdd : IExercise {
        public string Name => "Even or odd";

        public void Run() {
            // Note: 0 is even.
            Console.WriteLine("Please, enter number:");
            Console.WriteLine(ConsoleUtils.ReadInt() % 2 == 0 ? "Even" : "Odd");

        }
    }
}
