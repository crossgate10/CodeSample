using Factory_Strategy.Protocol;
using System;

namespace Factory_Strategy.Service
{
    public class Car1Factory :ICarFactory
    {
        private readonly IDep1 dep1;
        private readonly IDep2 dep2;
        private readonly IDep3 dep3;

        public Car1Factory(IDep1 dep1, IDep2 dep2, IDep3 dep3)
        {
            if (dep1 == null)
            {
                throw new ArgumentNullException("dep1");
            }
            if (dep2 == null)
            {
                throw new ArgumentNullException("dep2");
            }
            if (dep3 == null)
            {
                throw new ArgumentNullException("dep3");
            }

            this.dep1 = dep1;
            this.dep2 = dep2;
            this.dep3 = dep3;
        }

        public bool AppliesTo(Type type)
        {
            return typeof(Car1).Equals(type);
        }

        public ICar CreateCar()
        {
            return new Car1(this.dep1, this.dep2, this.dep3);
        }
    }
}
