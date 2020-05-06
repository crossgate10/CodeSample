using System;
using Observer.Protocol;

namespace Observer
{
    public class Customer : IObserver
    {
        public string Name { private get; set; }

        public Customer(string name)
        {
            this.Name = name;
        }

        public void Update(string message)
        {
            Console.WriteLine($"{this.Name} receive a new message: {message}");
        }
    }
}
