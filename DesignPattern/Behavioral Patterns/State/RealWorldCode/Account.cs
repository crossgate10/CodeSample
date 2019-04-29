using System;

namespace RealWorldCode
{
    class Account
    {
        private State state;
        private string owner;

        public Account(string owner)
        {
            this.owner = owner;
            this.state = new SilverState(0.0, this);
        }

        public double Balance
        {
            get { return state.Balance; }
        }

        public State State
        {
            get { return state; }
            set { state = value; }
        }

        public void Deposit(double amount)
        {
            state.Deposit(amount);
            Console.WriteLine("Deposited {0:C} --- ", amount);
            Console.WriteLine(" Balance = {0:C}", this.Balance);
            Console.WriteLine(" Status = {0}", this.State.GetType().Name);
            Console.WriteLine("");
        }

        public void PayInterest()
        {
            this.state.PayInterest();
            Console.WriteLine("Interest Paid --- ");
            Console.WriteLine(" Balance = {0:C}", this.Balance);
            Console.WriteLine(" Status = {0}\n", this.State.GetType().Name);
        }

        public void Withdraw(double amount)
        {
            this.state.Withdraw(amount);
            Console.WriteLine("Deposited {0:C} --- ", amount);
            Console.WriteLine(" Balance = {0:C}", this.Balance);
            Console.WriteLine(" Status = {0}\n", this.State.GetType().Name);
        }
    }
}