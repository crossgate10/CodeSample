using System;

namespace RealWorldCode
{
    class President : Approver
    {
        public override void ProcessRequest(Purchase p)
        {
            if (p.Amount < 100000.0)
            {
                Console.WriteLine($"{this.GetType().Name} approved request# {p.Number}");
            }
            else
            {
                Console.WriteLine($"Request# {p.Number} requires an executive meeting!");
            }
        }
    }
}