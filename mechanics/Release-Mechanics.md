# Release Mechanics

Knative core components (repo) releases from master are coordinated for the
following reasons:

1. Coordinated release makes it easier to release documentation and other
   cross-cutting repositories.
1. Coordinated release makes it easier for operators to install and describe
   what release they are running.
1. Coordinated release makes it easier to test compatibility across repos,
   though cross-repo dependencies should generally be supported at least +/- 1
   revision (i.e. revision N+1 of repo X should work with revision N of repo Y
   and vice-versa).

**Patch (cherrypick) releases are not coordinated across repos**; it is expected
that each repo will perform patch releases for critical bugs and security issues
for the last 2 releases.

## Release Schedule

Knative releases will be cut (automatically) every 6 weeks. Any change to this
schedule should be coordinated by the TOC.

<!-- TODO: do we have a calendar for this? What about releases during holidays like Christmas? -->

## Release Naming

Regular releases from master will be named by the date of the release in ISO
date format (e.g. `2019-08-06`). Patch releases will be named as `2019-08-06.1`,
`2019-08-06.2`, etc. The head release will be named without a ".0" suffix; the
first patch (cherrypick) release should be named ".1", with each subsequent
patch incrementing the number.

Each repo should define a mapping from the canonical date-based format to a
traditional [Semantic Versioning](https://semver.org) minor version number. Note
that these mappings may not be the same between repos; for example,
`knative/serving` might map `2019-11-05` to "1.1", but `knative/observability`
might map the same release to "0.10".

Date-based version numbers will be used for all of the following:

- release branch name
- docs page dropdown
- tag in the container images (the date-based tag will point to the most recent
  patch release of that date-based branch)

Date + patch version numbers will be used for the following:

- tag in github (as there may be multiple tagged releases on the same branch in
  the case of patches)
- label on k8s artifacts (so that users can determine whether their resources
  are up to date with the latest patch release)
- tag in the container images (the exact date+patch tag should be kept both for
  reference by released yaml and to be able to work backwards from a docker
  image to the release information)

Semantic versions should (at least) be referenced in the main documentation
page, e.g. https://knative.dev/docs/serving

## Historical Background

Historically (up to 0.8), Knative releases were named using 0.X (major version =
0, minor version incrementing on each release.) Since around January 2019, the
Knative release schedule has been every 6 weeks; around March 2019, this was
automated.
