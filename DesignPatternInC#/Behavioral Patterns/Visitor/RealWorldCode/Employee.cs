using System;

namespace RealWorldCode
{
    internal class Employee : Element
    {
        private string name;
        private double income;
        private int vacationDays;

        public Employee(string name, double income, int vacationDays)
        {
            this.name = name;
            this.income = income;
            this.vacationDays = vacationDays;
        }

        public double Income { get { return this.income; } set { this.income = value; } }
        public string Name { get { return this.name; } set { this.name = value; } }
        public int VacationDays { get { return this.vacationDays; } set { this.vacationDays = value; } }

        public override void Accept(IVisitor visitor)
        {
            visitor.Visit(this);
        }
    }
}