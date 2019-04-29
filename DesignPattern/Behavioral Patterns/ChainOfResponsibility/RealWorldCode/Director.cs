using System;

namespace RealWorldCode
{
    class Director : Approver
    {
        public override void ProcessRequest(Purchase p)
        {
            if (p.Amount < 10000.0)
            {
                Console.WriteLine($"{this.GetType().Name} approved request# {p.Number}");
            }
            else if (this.successor != null)
            {
                this.successor.ProcessRequest(p);
            }
        }
    }
}