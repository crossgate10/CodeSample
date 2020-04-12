using Factory_Strategy.Protocol;
using Factory_Strategy.Service;
using System;

namespace Factory_Strategy
{
    class Program
    {
        static void Main(string[] args)
        {
            IDep1 dep1 = new MockDep();
            IDep2 dep2 = new MockDep();
            IDep3 dep3 = new MockDep();
            IDep4 dep4 = new MockDep();
            IDep5 dep5 = new MockDep();
            IDep6 dep6 = new MockDep();

            var strategy = new CarStrategy(new ICarFactory[] {
                new Car1Factory(dep1, dep2, dep3),
                new Car2Factory(dep4, dep5, dep6)
            });

            // And then once it is injected, you would simply do this.
            // Note that you could use a magic string or some other 
            // data type as the parameter if you prefer.
            var car1 = strategy.CreateCar(typeof(Car1));
            var car2 = strategy.CreateCar(typeof(Car2));

            Console.ReadKey();
        }
    }
}
