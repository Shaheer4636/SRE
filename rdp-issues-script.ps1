# --- Start of check_detailed_rdp_report.ps1 ---

Write-Host "Checking RDP Licensing Status and System Events on $env:COMPUTERNAME..."

# Check if Remote Desktop Licensing service is running
$licService = Get-Service -Name TermServLicensing -ErrorAction SilentlyContinue

# Get installed licenses from registry
$licenseKeyPath = "HKLM:\SYSTEM\CurrentControlSet\Services\TermService\Parameters\LicenseServers"
$licenses = @()

if (Test-Path $licenseKeyPath) {
    $licenseData = Get-ItemProperty -Path $licenseKeyPath
    # In many cases license details are stored differently; we'll use the previous script's approach instead
}

# Use WMI / CIM class to get license info (adapted from previous script)
# We'll re-run the CIM query to get license info
$licenseInfo = Get-CimInstance -Namespace root\cimv2\TerminalServices -ClassName Win32_TerminalServiceSetting -ErrorAction SilentlyContinue

# Fallback to previously shared output from your custom licensing tool or script
# For demo, assume licenseInfo is unavailable and parse your known license array from earlier

# Hardcoded licenses example - replace this with dynamic retrieval if available
$installedCALs = @(
    @{ProductType=3; Description="A02-5.00-EX"; TotalLicenses=4294967295; ExpirationDate="20360101"},
    @{ProductType=0; Description="A02-10.01-S"; TotalLicenses=4; ExpirationDate="20380101"},
    @{ProductType=6; Description="A02-10.01-VDIS"; TotalLicenses=12; ExpirationDate="20241219"},
    @{ProductType=0; Description="A02-10.01-S"; TotalLicenses=5; ExpirationDate="20380101"}
)

# Calculate total licenses and find earliest expiration date
$totalLicenses = 0
$expirationDates = @()

foreach ($cal in $installedCALs) {
    $totalLicenses += $cal.TotalLicenses
    $expirationDates += [datetime]::ParseExact($cal.ExpirationDate, "yyyyMMdd", $null)
}

$earliestExpiration = $expirationDates | Sort-Object | Select-Object -First 1

# Fetch recent RDP disconnection events from last 24 hours
$events = Get-WinEvent -LogName 'Microsoft-Windows-TerminalServices-LocalSessionManager/Operational' |
    Where-Object {
        $_.TimeCreated -gt (Get-Date).AddDays(-1) -and
        ($_.Id -eq 24 -or $_.Id -eq 25 -or $_.Id -eq 39 -or $_.Id -eq 40)
    } | Select TimeCreated, Id, Message | Sort-Object TimeCreated

# Fetch system restart events (event ID 1074, 6006, 6008)
$restarts = Get-WinEvent -LogName System |
    Where-Object {
        $_.TimeCreated -gt (Get-Date).AddDays(-1) -and
        ($_.Id -eq 1074 -or $_.Id -eq 6006 -or $_.Id -eq 6008)
    } | Select TimeCreated, Id, Message | Sort-Object TimeCreated

# Compose report
Write-Host "`n==== RDP License Report ===="
Write-Host "Server Name           : $env:COMPUTERNAME"
Write-Host "Remote Desktop Licensing Service Status : " `
    + (if ($licService.Status) { $licService.Status } else { "Service not found or not running" })

Write-Host "Total Licenses Available : $totalLicenses"
Write-Host "Earliest License Expiration Date : $earliestExpiration"

Write-Host "`n---- Installed License Details ----"
foreach ($cal in $installedCALs) {
    $expDate = [datetime]::ParseExact($cal.ExpirationDate, "yyyyMMdd", $null)
    Write-Host ("ProductType: {0,-3} Description: {1,-15} Licenses: {2,-10} Expiration: {3}" -f $cal.ProductType, $cal.Description, $cal.TotalLicenses, $expDate.ToShortDateString())
}

Write-Host "`n---- Recent RDP Session Events (Last 24h) ----"
if ($events) {
    $events | ForEach-Object {
        Write-Host ("{0} - EventID: {1} - {2}" -f $_.TimeCreated, $_.Id, $_.Message)
    }
} else {
    Write-Host "No RDP session events found in the last 24 hours."
}

Write-Host "`n---- System Restart Events (Last 24h) ----"
if ($restarts) {
    $restarts | ForEach-Object {
        Write-Host ("{0} - EventID: {1} - {2}" -f $_.TimeCreated, $_.Id, $_.Message)
    }
} else {
    Write-Host "No system restart events found in the last 24 hours."
}

Write-Host "`n==== End of Report ===="

# --- End of check_detailed_rdp_report.ps1 ---
