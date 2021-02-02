---
title: "Security Policy"
linkTitle: "Security policy"
weight: 10
type: "docs"
aliases:
  - /contributing/security-policy/
---

# Introduction

Security is core to our values, and we value the input of developers, operators,
security researchers, and others acting in good faith to help us maintain a high
standard for the security and privacy for our users, and their users. This
includes encouraging responsible vulnerability research and disclosure. This
policy sets out our definition of good faith in the context of finding and
reporting vulnerabilities, as well as what you can expect from us in return.

*If you have identified a vulnerability in Knative, please do not post an issue 
in the public issue tracker. Instead, please follow the responsible disclosure 
policy described in this document.*

# Expectations

When working with us according to this policy, you can expect us to:

- Extend [Safe Harbor](#safe-harbor) for your vulnerability research that is
  related to this policy;
- Work with you to understand and validate your report, including a timely
  initial response to the submission;
- Work to remediate discovered vulnerabilities in
  [supported versions](#scope-and-supported-versions) in a timely manner;
- Publicly disclose accepted vulnerabilities in a timely manner; and
- Recognize your contribution to improving our security if you are the first to
  report a unique vulnerability, and your report triggers a code or
  configuration change.

# Ground Rules

To encourage vulnerability research and to avoid any confusion between
good-faith hacking and malicious attack, we ask that you:

- Play by the rules. This includes following this policy, as well as any other
  relevant agreements, particularly the Knative
  [Code of Conduct](https://knative.dev/community/contributing/code-of-conduct/).
  If there is any inconsistency between this policy and any other relevant
  terms, the terms of this policy will prevail;
- Report any vulnerability you’ve discovered promptly;
- Avoid violating the privacy of others, disrupting out-of-scope systems,
  destroying data, and/or harming user experience;
- Use only the [Official Channels](#reporting-a-vulnerability) to discuss
  vulnerability information with us;
- Keep the details of any discovered vulnerabilities confidential until they are
  fixed, according to the Disclosure Policy;
- Perform testing only on in-scope systems, and respect systems and activities
  which are out-of-scope;
- If a vulnerability provides unintended access to data: Limit the amount of
  data you access to the minimum required for effectively demonstrating a Proof
  of Concept; and cease testing and submit a report immediately if you encounter
  any user data during testing, such as Personally Identifiable Information
  (PII), Personal Healthcare Information (PHI), credit card data, or proprietary
  information;
- You should only interact with test accounts and Knative instances you own or
  with explicit permission from the owner or account holder; and
- Do not engage in extortion.

# Safe Harbor

When conducting vulnerability research according to this policy, we consider
this research to be:

- Authorized in accordance with the Computer Fraud and Abuse Act (CFAA) (and/or
  similar state laws), and we will not initiate or support legal action against
  you for accidental, good faith violations of this policy;
- Exempt from the Digital Millennium Copyright Act (DMCA), and we will not bring
  a claim against you for circumvention of technology controls;
- Exempt from restrictions in our Terms & Conditions that would interfere with
  conducting security research, and we waive those restrictions on a limited
  basis for work done under this policy; and
- Lawful, helpful to the overall security of the Internet, and conducted in good
  faith.

You are expected, as always, to comply with all applicable laws.

If at any time you have concerns or are uncertain whether your security research
is consistent with this policy, please submit a report through one of our
[Official Channels](#reporting-a-vulnerability) before going any further.

# Scope and Supported Versions

We (the community) support the 4 most recent minor versions of Knative. For more
details on what this means, see our
[release principles](https://knative.dev/community/contributing/mechanics/release-versioning-principles/).

While we encourage you to report vulnerabilities in any Knative repo, at HEAD or
the tip of any release branch (release-X.Y), we commit to patching
vulnerabilities only at HEAD and in currently supported release branches.

Repos with releases have the releases listed in the Github UI, e.g.:

- [serving](https://github.com/knative/serving/releases)
- [eventing](https://github.com/knative/eventing/releases)

Repos in Knative Sandbox that have no releases listed may have a higher 
bar for accepting vulnerability reports. No support is provided for archived repos. 

# Reporting a Vulnerability

Please email the
[Knative Security list](mailto:knative-security@googlegroups.com) with the following information:

* the repo and version/branch that contains the vulnerability
* a description of the type of vulnerability (e.g., buffer overflow)

_Please avoid including any details which would allow reproduction of the issue at this stage._
Details will be requested subsequently, over encrypted communications.

Someone from the list will get back to you within 3 business days. Accepted reports 
will be fixed within 90 calendar days via one or more releases for versions that are
both affected and supported at that point.

We ask that you keep reports confidential for the full 90 days, or until the
security list members release you from the confidentiality, whichever comes
first, as we coordinate disclosures with Knative partners who may need some time
after the vulnerability is addressed in Knative in order to update their
systems.

## What's in a report?

Initial reports via email should include the limited information described
[above](#reporting-a-vulnerability). Followup details should include a 
description of the vulnerability, a
proof-of-concept exploit demonstrating the vulnerability, and a statement of
what this vulnerability puts at risk. Please send reports in English.

The security list members may make a best-effort attempt to address reports that
do not meet these standards.

## Coordinated Vulnerability Disclosure (CVD)

Knative Security engages with partners, including vendors that host
Knative-based products, to responsibly disclose newly discovered
vulnerabilities.

If you believe your organization should participate in Knative's CVD process,
please contact Knative Security to discuss this.

For more information on CVD, please review the information provided in the
following links:

[Guidelines and Practices for Multi-Party Vulnerability Coordination and Disclosure](https://www.first.org/global/sigs/vulnerability-coordination/multiparty/FIRST-Multiparty-Vulnerability-Coordination-latest.pdf)

[The CERT Guide to Vulnerability Disclosure](https://resources.sei.cmu.edu/asset_files/SpecialReport/2017_003_001_503340.pdf)

[PSIRT Services Framework](https://www.first.org/standards/frameworks/psirts/FIRST_PSIRT_Services_Framework_v1.0.pdf)

# Vulnerability Rewards

Knative itself has no vulnerability report reward program.

Some vulnerability reports may be eligible for rewards through programs from
other organizations, including those from vendors of
[Knative-based offerings](https://knative.dev/docs/knative-offerings/).
Reporters should comply with program rules to receive these rewards; the Knative
community will not submit reports for rewards on the reporter's behalf.

Vendors with relevant reward programs:

- [Google](https://www.google.com/about/appsecurity/reward-program/index.html)

# Sources

Adapted from
[disclose.io USA core terms](https://github.com/disclose/terms/blob/master/regional/USA-core-terms.md)
