# Automation of Google Groups maintenance for `knative.team`

- [Making changes](#making-changes)
- [Manual deploy](#manual-deploy)

## Making changes

- Edit your WGs's `groups.yaml`, e.g. [`wg-security/groups.yaml`][/groups/wg-security/groups.yaml]
- If adding or removing a group, edit [`restrictions.yaml`] to add or remove the group name
- Use `make test` to ensure the changes meet conventions
- Open a pull request
- When the pull request merges, the [post-k8sio-groups] job will deploy the changes


**The project name has a max length of 18 characters.**

## Manual deploy

- Must be run by someone who is a member of the productivity-infra-gcp-org@knative.team group
- Run `gcloud auth application-default login` to login
- Use `make run` to dry run the changes
- Use `make run -- --confirm` if the changes suggested in the previous step looks good

[post-k8sio-groups]: https://testgrid.k8s.io/r/knative-own-testgrid/utilities#knativeteam-groups-jobs

## How does this work?

- The groups are managed with the Google Admin SDK Groups API
- Google has a process called Domain Wide Delegation(DWD) that allows a Google Service Account to
  impersonate a google workspace user. https://developers.google.com/admin-sdk/directory/v1/guides/delegation
- Configuring DWD is one time process as long as the Google Service Account impersonating is not deleted.
