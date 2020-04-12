namespace AbstractFactory
{
    public class Client
    {
        private AbstractProductA abstractProductA;
        private AbstractProductB abstractProductB;

        public Client(AbstractFactory factory)
        {
            this.abstractProductA = factory.CreateProductA();
            this.abstractProductB = factory.CreateProductB();
        }

        public void Run()
        {
            this.abstractProductB.Interact(this.abstractProductA);
        }
    }
}
