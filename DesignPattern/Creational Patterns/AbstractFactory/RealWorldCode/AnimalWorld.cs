namespace RealWorldCode
{
    internal class AnimalWorld
    {
        private Herbivore herbivore;
        private Carnivore carnivore;

        public AnimalWorld(ContinentFactory factory)
        {
            this.herbivore = factory.CreateHerbivore();
            this.carnivore = factory.CreateCarnivore();
        }

        internal void RunFoodChain()
        {
            this.carnivore.Eat(this.herbivore);
        }
    }
}