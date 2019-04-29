using System;

namespace RealWorldCode
{
    internal class Lion : Carnivore
    {
        public override void Eat(Herbivore herbivore)
        {
            Console.WriteLine($"{this.GetType().Name} eats {herbivore.GetType().Name}");
        }
    }
}