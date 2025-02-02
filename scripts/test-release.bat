@echo off

:: Clean dist directory
rd /s /q dist 2>nul

:: Run goreleaser in test mode
goreleaser build --clean --snapshot

:: Test the binary
cd dist\ariaj_windows_amd64_v1
ariaj.exe --help | findstr "Usage: ariaj [prompt]"

if %ERRORLEVEL% EQU 0 (
    echo Local test passed!
) else (
    echo Local test failed!
    exit /b 1
)
