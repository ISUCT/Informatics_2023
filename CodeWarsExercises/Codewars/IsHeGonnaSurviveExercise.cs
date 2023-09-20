namespace CodeWarsExercises.Codewars {
    internal class IsHeGonnaSurviveExercise : IExercise {
        public string Name => "Is He Gonna Survive?";

        public const int DefaultDragonHealth = 2;

        public static bool IsHeGonnaSurvive(int bulletCount, int dragonCount, int dragonHealth) {
            if (bulletCount < 0)
                throw new ArgumentException("bulletCount must be >= 0");

            if (dragonCount < 0)
                throw new ArgumentException("dragonCount must be >= 0");

            if (dragonHealth <= 0)
                throw new ArgumentException("dragonHealth must be >= 1");

            return bulletCount >= dragonCount * dragonHealth;
        }

        public void Run() {
            int bulletCount = ConsoleUtils.ReadInt("Enter which bullet count you prefer to take, Hero: ");
            int dragonCount = ConsoleUtils.ReadInt("Enter count of dragons in the castle, Hero: ");

            Console.WriteLine(
                IsHeGonnaSurvive(bulletCount, dragonCount, DefaultDragonHealth) ?
                "You'll survive, my brave hero." :
                "I'm sorry to hear you're dead, my poor hero..."
                );
        }
    }
}
