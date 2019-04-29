@echo off

for /f "tokens=5 delims= " %%c in ('netstat -aon^|findstr "YourPort"') do (
    REM echo %%c
	setlocal
	set isFind = tasklist|findstr %%c
	if  "%isFind%" neq 0 taskkill /F /PID %%c
	endlocal
)

pause