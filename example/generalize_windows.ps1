#!powershell

# Wait for Windows Guest Agent to be fully installed
while ((Get-Service RdAgent).Status -ne 'Running') { Start-Sleep -s 5 }
while ((Get-Service WindowsAzureTelemetryService).Status -ne 'Running') { Start-Sleep -s 5 }
while ((Get-Service WindowsAzureGuestAgent).Status -ne 'Running') { Start-Sleep -s 5 }

# Stop all Windows Guest Agent
Get-Service WindowsAzure* | Where {$_.status -eq 'Running'} | Stop-Service -Force

# Cleanup
Remove-Item -Recurse -Force 'C:\AzureData'
Remove-Item -Recurse -Force 'C:\WindowsAzure\Config'
Remove-Item -Recurse -Force 'HKLM:\SOFTWARE\Microsoft\Windows Azure'
Remove-Item -Recurse -Force 'HKLM:\SOFTWARE\Microsoft\Azure\DSC'

# Run and wait for Sysgrep completion
if (Test-Path $Env:SystemRoot\windows\system32\Sysprep\unattend.xml) {
  rm $Env:SystemRoot\windows\system32\Sysprep\unattend.xml -Force
}

& $env:SystemRoot\System32\Sysprep\Sysprep.exe /oobe /generalize /quiet /quit

while ($true) {
  $imageState = Get-ItemProperty 'HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Setup\State' | Select ImageState
  if ($imageState.ImageState -ne 'IMAGE_STATE_GENERALIZE_RESEAL_TO_OOBE') {
    Write-Output $imageState.ImageState
    Start-Sleep -s 10
  } else {
    break
  }
}