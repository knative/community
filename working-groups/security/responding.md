# Vulnerability Disclosure Response Policy

The project follows the following process when responding to vulnerabilities
raised via the Security Disclosure Process.

## Product Security Taskforce Membership

The subset of maintainers who are subscribed to the security@knative.team email
address form the Product Security Taskforce.
This group will be responsible for fixing vulnerabilities and managing the
disclosure process.

The Product Security Taskforce is distinct from the embargo list: it is not
necessary (or helpful) to join the PST in order to receive early notice of
vulnerabilities. The PST is intentionally small to avoid inadvertent
disclosures of vulnerabilities (while being large enough to spread the work of
responding to vulnerabilities) and should be composed of those who will
actively work on the fix and disclosure process.

Anyone who is an Approver, Lead or TOC Member in a Knative Project, or who is
an employee of a member of the Embargo list working in a security-focused or
knative-focused role may apply to be a member of this group.
The application is approved or rejected by the TOC.

## Identifying a Fix Lead

The PST Fix Lead Role is responsible for initially responding to
vulnerabilities sent to security@knative.team, and for coordinating and
managing the response with the rest of the PST.

We use a rota system to assign the default PST Lead role.
All members of the PST are expected to sign up to the rota.
Regardless of the rota, another member of the PST may volunteer to be the PST
lead for a particular vulnerability.
This may be appropriate, for example, when the vulnerability is in a particular
area of the code that PST member is familiar with, or when the rota-ed PST lead
is already handing another vulnerability.
**To avoid any confusion, if the PST Lead role is transferred, [positive hand-off](https://www.faasafety.gov/gslac/ALC/course_content.aspx?cID=36&sID=196&preview=true)
should be performed ("you are the PST Lead for X", "Ok, I am the PST Lead for X", "Great, you are the PST lead for X!")**.

## Organising the Fix Team

The Fix Lead will quickly identify relevant engineers from the affected
projects/areas/working groups and the PST and add those engineers to the
security disclosure / email thread. This group becomes the Fix Team for the
vulnerability. To prevent accidental disclosure and to maximize efficiency of
response, the Fix Team should be the minimal size required for efficient
response.

Note: Given the current size of the project it is likely that in many cases the
Fix Team will be the entire PST, possibly with an additional engineer or
engineers for the affected area if the PST does not have relevant expertise
itself.

If the Fix Team decides the disclosure is in fact a vulnerability which should
be fixed, the Fix Lead will create a Github Security Advisory and provide the
Fix Team access to a private branch to develop the fix.

## Developing a Fix and Managing Disclosures

### Within 1-7 Days of Disclosure

Vulnerabilities are evaluated based on the Knative Threat Model.

1. The Fix Team will create a CVSS using the [CVSS calculator](https://www.first.org/cvss/calculator/3.0).
   The Fix Lead makes the final call on the calculated CVSS; it is better to
   move quickly than making the CVSS perfect.
1. **If the Fix Team determines this is not a vulnerability based on the Knative
   Threat Model, the discloser is notified and the process ends.**
   (Note that in all cases it is important to inform the reporter that
   the disclosure has been seen and is being addressed as quickly as possible).
1. Otherwise, the Fix Lead creates a private branch for development of the Fix,
   and a Github Security Advisory to coordinate the disclosure.
1. The Fix Lead responds to the reporter to let them know the work has begun to
   address the vulnerability. **If appropriate, and if the reporter is willing,
   the reporter is added to the Security Advisory and can help to develop the
   fix. The reporter may also wish to help with writing the CVE entry. They
   should in all cases be asked if they would like credit for identifying the
   vulnerability in the CVE and disclosure**.
1. The Fix Team urgently develops a fix or mitigation so that a realistic
   timeline can be communicated to users.
1. The Fix Lead completes the Github Security Advisory and requests a CVE from Github.
1. If the CVSS score is below 4.0, the Fix Team may - after appropriate
   discussion and based on developer bandwidth, public holidays etc - decide to
   slow down the release process.
1. The fact that a vulnerability has been disclosed and a fix will be released
   is communicated by the Fix Lead (working with the Fix Team) to the
   knative-security-announce mailing list with an exact date and time planned
   for the disclosure. If the optional embargoed disclosure step will be
   followed (see next section), this date must be three weeks from when the
   disclosure is released to the embargo list.
1. This email should be actionable: it should be clear when users
   should block time to apply patches, what mitigations are available (if
   appropriate), etc.

### Optional, Within 1-14 Days of Disclosure (Private, Embargoed Disclosure)

1. If, in the judgement of the Fix Lead, a private disclosure is warranted, the
   Embargo List is emailed with details of the fix (including any proposed
   patch/mitigation and actionable instructions) and the date it will be
   publicly disclosed and released.
1. If this is the case, the Public Disclosure should not be
   released until at least three weeks after this notification is sent, to give
   vendors time to prepare to release the patch. In exceptional circumstances
   this may be extended to four weeks to account for significant public
   holidays etc.
1. Note: This step is optional and should be reserved for significant
   vulnerabilities that warrant this treatment. In general these are
   vulnerabilities that could lead to remotely exploitable or privilege
   escalation issues.
1. In the unlikely event that a member of the PST list breaks embargo, or if
   the exploit otherwise becomes public early, the PST Lead will decide whether
   to accelerate the schedule for public disclosure and release of the fix.
   When in doubt push forward and go public ASAP.

### Within 1-28 Days of Disclosure (Fix Release Day)

1. The Fix Lead cherry-picks the fix on to the main branch and all relevant
   release branches and creates PRs.
1. Maintainers should approve this PR as quickly as possible (it should already
   have been reviewed on the private branch).
1. If necessary, Release jobs should be manually triggered to ensure
   a release is available as quickly as possible.
1. The Issue should be disclosed to the knative-security-annnounce mailing
   list, the knative-users mailing list, the knative-dev mailing list
   and the embargo list including relevant PRs, mitigations and the release
   numbers containing the fix.
1. The Fix Lead removes any permissions from private branches granted for the
   purposes of the Fix.

### Retrospective

1. A [blameless retrospective](https://landing.google.com/sre/book/chapters/postmortem-culture.html)
   should be conducted 1-3 days after the release.
1. The Fix Lead will send a retrospective of the process to knative-dev
   including details on everyone involved, the timeline of the process, links
   to relevant PRs that introduced the issue, if relevant, and any critiques of
   the response and release process.
1. Maintainers and Fix Team are also encouraged to send their own feedback on
   the process to knative-dev. Honest critique is the only way we are going to
   get good at this as a community.

## References

This document borrows heavily and gratefully from the [Envoy Vulnerability
Response Process](https://github.com/envoyproxy/envoy/blob/main/SECURITY.md)
and the [Google OSS Vulnerability Guide](https://github.com/google/oss-vulnerability-guide/blob/main/guide.md).
