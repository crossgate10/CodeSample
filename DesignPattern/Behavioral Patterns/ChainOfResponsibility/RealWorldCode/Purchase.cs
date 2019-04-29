namespace RealWorldCode
{
    class Purchase
    {
        private int number;
        private double amount;
        private string purpose;

        public Purchase(int number, double amount, string purpose)
        {
            this.number = number;
            this.amount = amount;
            this.purpose = purpose;
        }

        public double Amount
        {
            get { return this.amount; }
            set { this.amount = value; }
        }
        public int Number
        {
            get { return this.number; }
            set { this.number = value; }
        }
        public string Purpose
        {
            get { return this.purpose; }
            set { this.purpose = value; }
        }
    }
}