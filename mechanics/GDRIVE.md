# GSuite configuration for Knative.team


## Historical Context

The Knative Community GSuite tools used the google.com domain up until January 4th, 2021. With the GSuite org being under the google.com domain, this affected the following:



*   GDrive - Folks outside of Google can’t control the ownership of the community drive. Documents can't be shared with "everyone on the internet", only "with Google.com". This makes it hard to provide read view outside the Knative community.
*   Knative Community Calendar - Folks outside of Google have read/write access but don’t have admin privileges
*   Hangouts - Hangouts for WG meetings required someone from Google to let folks into the meeting
*   Mailing lists are all @googlegroups.com, rather than associated with Knative (branding/authentication challenge).

Due to the challenges of managing community artifacts, we now have a knative.team GSuite org.


#### GSuite gotchas

Google.com accounts have a corporate policy that prevents them from creating new drive documents outside the domain. Googlers should use a consumer (@gmail.com) or @knative.team account to create documents on the team drive (including uploading recordings). Googlers _can_ comment or edit existing documents without problems.


## Knative.team GSuite Setup


### Access/Privileges


#### Steering Committee

All steering committee members have “Super Admin” [privileges](https://support.google.com/a/answer/33325?hl=en) on the Knative.team GSuite account. This allows steering members to create new groups, manage Google Meet settings etc. In order to access the GSuite admin console, navigate to admin.google.com.


#### Technical Oversight Committee

The technical oversight committee members each have a knative.team GSuite account. This is to allow TOC members to assist with some of the GSuite automation and GDrive migration. In addition to this, TOC members will need a knative.team account to allow people without a knative.team account into public meetings hosted on Google Hangouts.


#### Working Group Leads

Each WG lead has a knative.team GSuite account. This is to allow WG leads to let people without a knative.team account into the public WG meetings hosted on Google Hangouts.

When each WG lead steps down, a GSuite admin will have to remove their account as there currently is no automation set up.


### Mailing lists

There are several mailing lists set up to manage permission groups (some could also be used for communication, but that is not the current usage). There is currently no automation to manage these lists.



*   wg-leads – all working group leads, gives update permission on the Knative official events calendar
*   steering – steering committee members
*   trademark – trademark committee members
*   (TODO) toc – tech oversight committee members


### Payment and Package

The Knative.team GSuite is a [Business Standard](https://workspace.google.com/intl/en/pricing.html) GSuite Plan. This is required so that WG meetings have the ability to record meetings. If we need changes made to the GSuite package-type on this account, [April Nassi](mailto:anassi@google.com) can help with this.

The GSuite org is currently paid by Google. [Mary Radomile](mailto:maryr@google.com) is the point of contact if we need help re: billing. 
