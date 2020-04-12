namespace Observer
{
    class ConcreteSubject : Subject
    {
        private string subjectState;

        public string SubjectState
        {
            get { return this.subjectState; }
            set { this.subjectState = value; }
        }
    }
}