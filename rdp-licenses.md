# **Incident Report: Unexpected RDP Disconnections and Licensing Concerns**

**Date of Incident:** May 22, 2025  
**Prepared By:** CRP\Administrator

---

## 1. Incident Overview

On May 22, 2025, during mid-day working hours, several users reported unexpected disconnections from Remote Desktop (RDP) sessions on the Windows Server `EC2AMAZ-UEM8JU1`. These disconnections occurred simultaneously across multiple user accounts, raising immediate concerns about licensing constraints or system-wide policy enforcements that might be causing sessions to terminate unexpectedly.

---

## 2. Problem Description

Upon receiving the reports, the infrastructure team initially suspected that the disconnections may have been caused by the expiration of the Remote Desktop Services (RDS) grace period or exhaustion of available RDP licenses (CALs). These assumptions were driven by the sudden and widespread nature of the disconnections, which resembled behavior typical of expired grace periods or unlicensed RDP sessions.

---

## 3. Investigation Process

A full diagnostic script was executed to verify the status of installed RDP licenses, their expiration dates, and the grace period. The registry was queried for the presence of the `GracePeriod` key under `HKLM\SYSTEM\CurrentControlSet\Control\Terminal Server\RCM`, and it was confirmed that the key still exists — indicating the system remains within its default 120-day grace window.

Additionally, all installed CALs were enumerated and checked for expiration. The licenses found were all valid, with expiration dates ranging between December 2024 and January 2038. No evidence was found of license exhaustion or misconfiguration. Concurrent login tests using 11 different user accounts were also conducted successfully without restriction, confirming that the license server is functioning correctly or that grace mode is still permitting unrestricted access.

Event logs were reviewed for the 24-hour window during which the incident occurred. It was discovered that a manual system reboot was initiated on the server at approximately 2:43 PM by the user `CRP\Administrator`. This action immediately terminated all active sessions. Log entries also showed a number of disconnection events with reason codes such as `12` (user-initiated or idle timeout), `5` (reconnect), and non-standard hexadecimal values (e.g., `3489660929`), which, while confusing, were not indicative of license errors.

---

## 4. Root Cause

The root cause of the incident was a **manual server restart** that occurred during working hours. This restart caused all active RDP sessions to be forcefully disconnected. The event was mistakenly interpreted as a licensing or grace period failure due to the lack of visibility into the license state at that moment and the atypical event logs observed immediately after the reboot. However, further investigation conclusively showed that both license availability and the grace period were intact and functioning as expected.

---

## 5. Resolution and Corrective Actions

To prevent similar confusion in the future, the following corrective actions were taken:

- Local Group Policy and registry configurations enforcing RDP session timeouts or idle disconnections were reviewed and explicitly disabled using PowerShell.
- The licensing status was documented using script-based audits for future reference.
- A test of concurrent RDP sessions was performed and passed successfully, ensuring no license-based denial was in effect.
- The system's Event Log was archived for further monitoring and to maintain a record of the incident.

---

## 6. Conclusion

The RDP disconnections observed on May 22 were not the result of license expiration or grace period expiry. Instead, they were caused by a deliberate system reboot that was not coordinated with end users. All subsequent technical checks confirm that the server is operating within a valid licensing state, the grace period is still active, and no license server communication issues have occurred. System policies have been adjusted to avoid unintended session terminations, and users can continue using RDP without restriction.
