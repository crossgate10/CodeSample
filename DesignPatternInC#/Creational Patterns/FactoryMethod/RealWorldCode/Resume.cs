namespace RealWorldCode
{
    internal class Resume : Document
    {
        public override void CreatePages()
        {
            this.Pages.Add(new SkillPage());
            this.Pages.Add(new EducationPage());
            this.Pages.Add(new ExperiencePage());
        }
    }
}