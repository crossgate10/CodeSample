namespace Observer
{
    using System;

    class ConcreteObserver : Observer
    {
        private ConcreteSubject subject;
        private string name;
        private string observerState;

        public ConcreteObserver(ConcreteSubject subject, string name)
        {
            this.subject = subject;
            this.name = name;
        }

        public void Notify()
        {
            throw new NotImplementedException();
        }

        public override void Update()
        {
            this.observerState = this.subject.SubjectState;
            Console.WriteLine($"Observer {this.name}'s new state is {this.observerState}");
        }

        ConcreteSubject Subject
        {
            get { return this.subject; }
            set { this.subject = value; }
        }
    }
}