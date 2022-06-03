# Knative Automation Bots

The peribolos configuration in this directory is used to configure users and groups in GitHub.
Among those accounts configured are several bot accounts that are used for automated tasks
in the Knative testing and release infrastructure. Those user accounts are listed here:

- knative-automation
- knative-metrics-robot
- knative-prow-robot
- knative-prow-releaser-robot
- knative-prow-updater-robot
- knative-test-reporter-robot

## Knative Automation [@knative-automation](https://github.com/knative-automation)

Most of the activity by this bot involves pull requests that go through a normal approval process. This bot should be covered by EasyCLA.

- Synchronizes GitHub actions and automated workflows.
- Updates snapshot dependency versions between projects.
- Updates common dependencies across all projects.
- Synchronizes GitHub configuration across all projects.

## Knative Metrics Robot [@knative-metrics-robot](https://github.com/knative-metrics-robot)

No recent activity. This bot is being considered for removal from the organization. No coverage by EasyCLA is needed.

## Knative Prow Robot [@knative-prow-robot](https://github.com/knative-prow-robot)

This bot writes to release branches, but does so using the original author's commits via cherry picks. Coverage by EasyCLA is probably not required, but would be good just in case.

- Synchronizes cherry picked commits across release branches.
- Owns the Knative Prow app https://github.com/apps/knative-prow

## Knative Prow Releaser Robot [@knative-prow-releaser-robot](https://github.com/knative-prow-releaser-robot)

This bot's activity is restricted to creating releases for the projects via Prow and the GitHub API. Releases themselves do not modify the project source code in any way. This bot does not require coverage by EasyCLA.

- Creates releases via Prow and the GitHub API across all projects

## Knative Prow Updater Robot [@knative-prow-updater-robot](https://github.com/knative-prow-updater-robot)

This bot automates the update of Prow itself. In doing so, it creates pull requests in the `knative/test-infra` repository (where Prow is configured) that are automatically merged. An example: https://github.com/apps/knative-prow. This bot should be covered by EasyCLA.

- Automates updates of Prow via pull requests.

## Knative Test Reporter Robot [@knative-test-reporter-robot](https://github.com/knative-test-reporter-robot)

This bot opens issues related to flaky test via a `flaky-test-reporter` action. Because issues do not modify the repository, this bot does not need coverage with EasyCLA.

- Reports test failures potentially across all projects by opening GitHub Issues. 

