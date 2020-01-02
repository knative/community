# Knative Feature Tracks

This document outlines the Knative process for adding non-trivial features.  The 
intent of this process is to articulate the best practices many successful 
features have followed across a variety of working groups.  This process may be 
elided at the discretion of working group leads, but its application is strongly 
encouraged for any feature that may have subtleties or contentious elements to 
front-load their consideration.

It is notable that just by adhering to this process there is no guarantee that 
your feature will land, for several reasons.  Reviewing the artifacts enumerated 
here takes time and the bandwidth of leads and approvers is limited.  Every 
feature that lands increases our support burden and spreads us more thinly, and 
so leads must be careful about prioritizing scope creep.

## Step 1: The Problem Statement

Every feature exists to solve a problem.  The process starts by articulating 
that problem in the form of a **Github issue** 
([template](https://docs.google.com/document/d/1wcQj6SBvIegfoeBT6Q7Wa_snNdsGy5j-bwwSOcTcejY/edit#heading=h.n8a530nnrb)).
 Be sure that a particular issue is tagged with the appropriate areas to bring 
it to the attention of the pertinent working group leads for triage.  In Prow, 
this can be done via:

```
# Assigns the label "area/networking"
/area networking
```

Contributors are encouraged to engage early/often at this stage, and ask 
clarifying questions.  Ultimately the leads are responsible for making a 
determination of how a given problem weighs against the broader working group's 
priorities.

Features may span working groups, but often have a "primary" working group 
focus.  If there is substantial work involved in multiple work groups, then the 
leads for all involved working groups should be involved throughout the process.

## Step 2: The Proposal

All non-trivial<sup>[2.1](#2.1)</sup> features deserve a design for 
articulating the solution to their problem statement. This should be done in a 
**Google doc**
([template](https://docs.google.com/document/d/1s6IIU98bi5FlRNmmBaLAn1rgoleK_ovcL746L7NHq0c/edit) -- make a copy!) 
under the appropriate folder in the [Team 
Drive](https://drive.google.com/corp/drive/u/0/folders/0APnJ_hRs30R2Uk9PVA)<sup>[2.2](#2.2)</sup>.
If the solution is clear and non-contentious, then the doc may be very short!
For major<sup>[2.3](#2.3)</sup> designs, once the proposal is accepted it should be converted to
markdown and committed to the appropriate GitHub repo.

However, the bigger the feature the less likely this is to be true.  Designs 
help us call out nuance early, and the best designs are often informed by 
prototypes that help surface these nuances.  **Do not expect to check in the 
prototype**; enter into these prototypes with the expectation that ALL of it 
will be thrown out.

The design should summarize the problem statement but the core focus should be 
on the end state (the **"what"**); think about how we might document the feature 
in the end state and give us a preview of what that might look like.  For some 
features there is complexity in getting from where we are to that end state.  
Such features will also need a detailed migration plan (the **"how"**<sup>[2.4](#2.4)</sup>).
Designs should also list reviewers that they expect to sign-off (the **"who"**),
this should be the WG leads (or their delegate).

> <a name="2.1"><sup>2.1</sup></a> - For the moment, this is at the discretion of the Working Group Lead.

> <a name="2.2"><sup>2.2</sup></a> - For read/write access join the [knative-dev](https://groups.google.com/forum/#!forum/knative-dev) Google group.

> <a name="2.3"><sup>2.3</sup></a> - This is at the discretion of WG leads, since doing this for every change will quickly flood the repo and make things _less_ discoverable.

> <a name="2.4"><sup>2.4</sup></a> - For a detailed outline of how to stage component changes, please [read this](https://github.com/knative/serving/issues/2639).

## Step 3: The Review

Each proposal, however non-contentious, should be called out during the working 
group meeting / slack for broader consideration and commentary.  The contributor 
driving this process should work with a lead for the appropriate area to get it 
onto the WG's agenda.  For particularly active WGs, this may take time, and 
lower priority features may get bumped by higher priority features at the lead's 
discretion.  

The review need not follow any prescribed process, but the lead should make sure 
things stay constructive and on-topic.  Leads should be up-front about how a 
decision will be made (e.g. eventing often uses the [Coinbase decision 
making](https://medium.com/@barmstrong/how-we-make-decisions-at-coinbase-cd6c630322e9) 
guide); by default a majority of the WG leads shall decide.  The outcome of a 
review may not be an immediate decision, and the contributor may get sent back 
to gather more information and iterate.

When a proposal is accepted, the leads should designate one or more 
reviewers<sup>[3.1](#3.1)</sup> within the WG as "sponsors" for the 
feature to help shepherd it through the process.  It is recommended that at 
least one sponsor be an approver 
([e.g.](https://github.com/knative/serving/blob/2018fcd98c18922cb1ce8b0207aa9aa6bef5eed1/OWNERS_ALIASES#L19)), 
but if non-approvers 
([e.g.](https://github.com/knative/serving/blob/2018fcd98c18922cb1ce8b0207aa9aa6bef5eed1/OWNERS_ALIASES#L25)) 
are listed, they should be considered the primary reviewer(s) so that they can 
hone their review skills and work towards approver.

> <a name="^3.1"><sup>3.1</sup></a> - Leads should try to be sensitive to the relative timezone of contributors with their sponsors to reduce cycle times on reviews.

## Step 4: The Breakdown

At a minimum, large features should be broken down into reasonably sized Pull 
Requests.  However, large features (e.g. multi-milestone) may need to be broken 
down into a number of smaller issues (no bigger than a milestone).  For these 
large features, the design document should eventually<sup>[4.1](#4.1)</sup> include, and the
broad strokes of how the work will be broken down. For these very large features,
the sponsor would then help set up a Github project containing issues for all of
the pertinent items.  Until all of this completes, the feature is not done.

> <a name="4.1"><sup>4.1</sup></a> - WG Leads can decide how much of the work breakdown to front-load, but what is important is that all parties (feature owner, lead, and sponsors) understand the full-scope of the work to be done.


## Step 5: The Planning Process

As leads are planning each milestone, they should draw issues from the Github 
Projects for in-flight features by working with feature owners and sponsors.

## Step 6: The Prestige

When all of the associated issues close, your feature is GA!
