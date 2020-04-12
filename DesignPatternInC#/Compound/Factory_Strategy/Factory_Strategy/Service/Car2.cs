using Factory_Strategy.Protocol;

namespace Factory_Strategy.Service
{
    public class Car2 : ICar
    {
        private IDep4 dep4;
        private IDep5 dep5;
        private IDep6 dep6;

        public Car2(IDep4 dep4, IDep5 dep5, IDep6 dep6)
        {
            this.dep4 = dep4;
            this.dep5 = dep5;
            this.dep6 = dep6;
        }
    }
}
