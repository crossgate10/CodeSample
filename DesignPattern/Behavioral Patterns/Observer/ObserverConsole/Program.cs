using System;

namespace Observer
{
    class Program
    {
        static void Main(string[] args)
        {
            NewspaperOffice office = new NewspaperOffice();

            Customer customer1 = new Customer("David");
            office.SubscribeNewspaper(customer1);
            Customer customer2 = new Customer("Peter");
            office.SubscribeNewspaper(customer2);

            office.SendNewspaper("Morning Breaking News!!!");

            office.UnsubscribeNewspaper(customer1);

            office.SendNewspaper("Evening Breaking News!!!");

            Console.ReadLine();
        }
    }
}
