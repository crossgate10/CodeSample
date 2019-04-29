using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Reflection;
using System.Text;
using System.Threading.Tasks;

namespace ViewModels
{
    public class MainWindowViewModel : ViewModelBase
    {
        /// <summary>
        ///     程式名稱
        /// </summary>
        private string appVersion;
        public string AppVersion
        {
            get
            {
                return appVersion;
            }
            set
            {
                if (appVersion != value)
                {
                    appVersion = value;
                    RaisePropertyChanged("AppVersion");
                }
            }
        }

        public MainWindowViewModel()
        {
            Assembly assembly = AppDomain.CurrentDomain.GetAssemblies().Where(a => a.ManifestModule.Name == "AutoUpdateSample.exe").Single();
            FileVersionInfo fileInfo = FileVersionInfo.GetVersionInfo(assembly.Location);
            AppVersion = fileInfo.FileVersion;
        }
    }
}
