namespace CodeWarsExercises.Codewars {
    internal class CountTheMonkeysExercise : IExercise {

        public string Name => "Count the monkeys" + (m_optimized ? " (optimized)" : "");

        private bool m_optimized;

        public CountTheMonkeysExercise(bool optimized) {
            m_optimized = optimized;
        }

        public static int[] CountTheMonkeys(int n) {
            if (n <= 0) throw new ArgumentException("n > 0");

            int[] monkeys = new int[n];

            for (int x = 0; x < n; ++x)
                monkeys[x] = x + 1;

            return monkeys;
        }

        /// <summary>
        /// Produces monkey's output using CountTheMonkeys function under the hood.
        /// </summary>
        /// <param name="n">Count of monkeys.</param>
        private void ProduceOutput(int n) {
            int[] monkey = CountTheMonkeys(n);

            Console.Write(monkey.Length);
            Console.Write(" --> [");

            for (int x = 0; x < monkey.Length; ++x) {
                Console.Write(monkey[x]);
                bool itemIsLast = x >= monkey.Length - 1;

                if (!itemIsLast)
                    Console.Write(", ");
            }

            Console.WriteLine("]");
        }

        /// <summary>
        /// Produces monkey's output without any array allocation under the hood.
        /// </summary>
        /// <param name="n">Count of monkeys.</param>
        private void ProduceOutputOptimized(int n) {
            Console.Write(n);
            Console.Write(" --> [");

            for (int x = 1; x <= n; ++x) {
                Console.Write(x);
                bool itemIsLast = x >= n;

                if (!itemIsLast)
                    Console.Write(", ");
            }

            Console.WriteLine("]");
        }

        public void Run() {
            int n = ConsoleUtils.ReadInt("Enter the monkey count: ");

            if (m_optimized)
                ProduceOutputOptimized(n);
            else
                ProduceOutput(n);
        }
    }
}
