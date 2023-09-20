using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace CodeWarsExercises {
    internal interface IExercise {
        
        public string Name { get; }
        
        void Run();
    }
}
