using Factory_Strategy.Protocol;
using System;

namespace Factory_Strategy.Service
{
    public class Car2Factory :ICarFactory
    {
        private readonly IDep4 dep4;
        private readonly IDep5 dep5;
        private readonly IDep6 dep6;

        public Car2Factory(IDep4 dep4, IDep5 dep5, IDep6 dep6)
        {
            if (dep4 == null)
            {
                throw new ArgumentNullException("dep4");
            }
            if (dep5 == null)
            {
                throw new ArgumentNullException("dep5");
            }
            if (dep6 == null)
            {
                throw new ArgumentNullException("dep6");
            }

            this.dep4 = dep4;
            this.dep5 = dep5;
            this.dep6 = dep6;
        }

        public bool AppliesTo(Type type)
        {
            return typeof(Car2).Equals(type);
        }

        public ICar CreateCar()
        {
            return new Car2(this.dep4, this.dep5, this.dep6);
        }
    }
}
