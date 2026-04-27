@echo off
cd /d "%~dp0"

echo =========================
echo GIT AUTO PUSH SAFE MODE
echo =========================

git status

git add .

set msg=update
if not "%1"=="" set msg=%1

echo Commit: %msg%

git commit -m "%msg%"

git push origin main

echo DONE
pause