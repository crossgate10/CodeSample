using System;
using Observer.Protocol;
using System.Collections.Generic;

namespace Observer
{
    public class NewspaperOffice : ISubject
    {
        List<IObserver> observers;

        public NewspaperOffice()
        {
            this.observers = new List<IObserver>();
        }

        public void NotifyObservers(string content)
        {
            foreach (IObserver observer in this.observers)
            {
                observer.Update(content);
            }
        }

        public void RegisterObserver(IObserver observer)
        {
            this.observers.Add(observer);
        }

        public void RemoveObserver(IObserver observer)
        {
            if (this.observers.IndexOf(observer) >= 0)
            {
                this.observers.Remove(observer);
            }
        }

        public void SubscribeNewspaper(IObserver customer)
        {
            this.RegisterObserver(customer);
        }

        public void UnsubscribeNewspaper(IObserver customer)
        {
            this.RemoveObserver(customer);
        }

        public void SendNewspaper(string content)
        {
            Console.WriteLine("Send News..");
            this.NotifyObservers(content);
        }
    }
}
