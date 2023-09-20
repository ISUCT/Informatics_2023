using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace CodeWarsExercises {
    internal static class FormatUtils {

        public static string FormatAsString<T>(this T[] values) {
            if (values.Length == 0)
                return "[]";
            
            StringBuilder sb = new StringBuilder("[");
            foreach(T v in values) {
                sb.Append(v?.ToString() ?? "*null*");
                sb.Append(", ");
            }

            sb.Length -= 2;
            sb.Append("]");

            return sb.ToString();
        }
    }
}
