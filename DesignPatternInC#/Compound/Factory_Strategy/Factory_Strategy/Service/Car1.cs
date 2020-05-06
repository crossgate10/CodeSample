using Factory_Strategy.Protocol;

namespace Factory_Strategy.Service
{
    public class Car1 : ICar
    {
        private IDep1 dep1;
        private IDep2 dep2;
        private IDep3 dep3;

        public Car1(IDep1 dep1, IDep2 dep2, IDep3 dep3)
        {
            this.dep1 = dep1;
            this.dep2 = dep2;
            this.dep3 = dep3;
        }
    }
}
