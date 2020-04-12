using System;

namespace RealWorldCode
{
    class Investor : IInvestor
    {
        private string name;
        private Stock stock;

        public Investor(string name)
        {
            this.name = name;
        }

        public void Update(Stock stock)
        {
            Console.WriteLine("Notified {0} of {1}'s change to {2:C}", this.name, stock.Symbol, stock.Price);
        }

        public Stock Stock
        {
            get { return this.stock; }
            set { this.stock = value; }
        }
    }
}