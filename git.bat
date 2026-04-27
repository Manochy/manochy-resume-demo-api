@echo off

echo =========================
echo GIT AUTO PUSH
echo =========================

cd /d "%~dp0"

echo Checking status...
git status

echo.
echo Adding all files...
git add .

echo.
set /p msg=Commit message (default: update): 

if "%msg%"=="" set msg=update

echo.
echo Committing...
git commit -m "%msg%"

echo.
echo Pushing to origin main...
git push origin main

echo.
echo Done!
pause