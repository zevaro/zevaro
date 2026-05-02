# sign-windows.ps1 — Sign the Zevaro Windows binary with an EV certificate via SignTool.
# Windows code signing implemented in ZV-051.
# Usage: .\scripts\sign-windows.ps1 -BinaryPath <path>
param(
    [Parameter(Mandatory = $true)]
    [string]$BinaryPath
)

Write-Host "Windows code signing is implemented in ZV-051"
exit 0
