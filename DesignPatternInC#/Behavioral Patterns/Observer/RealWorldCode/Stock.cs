namespace RealWorldCode
{
    using System;
    using System.Collections.Generic;
    using System.Linq;
    using System.Text;
    using System.Threading.Tasks;

    abstract class Stock
    {
        private string symbol;
        private double price;
        private List<IInvestor> investors = new List<IInvestor>();

        public Stock(string symbol, double price)
        {
            this.symbol = symbol;
            this.price = price;
        }

        public void Attach(IInvestor investor)
        {
            this.investors.Add(investor);
        }

        public void Detach(IInvestor investor)
        {
            this.investors.Remove(investor);
        }

        public void Notify()
        {
            foreach (IInvestor investor in this.investors)
            {
                investor.Update(this);
            }
        }

        public double Price
        {
            get { return this.price; }
            set
            {
                if (this.price != value)
                {
                    this.price = value;
                    this.Notify();
                }
            }
        }

        public string Symbol { get; }
    }
}
