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

- Synchronizes GitHub actions and automated workflows.
- Updates snapshot dependency versions between projects.
- Updates common dependencies across all projects.
- Synchronizes GitHub configuration across all projects.

## Knative Metrics Robot [@knative-metrics-robot](https://github.com/knative-metrics-robot)

- No recent activity

## Knative Prow Robot [@knative-prow-robot](https://github.com/knative-prow-robot)

- Used to synchronize cherry picked commits across release branches

## Knative Prow Releaser Robot [@knative-prow-releaser-robot](https://github.com/knative-prow-releaser-robot)

- No recent activity

## Knative Prow Updater Robot [@knative-prow-updater-robot](https://github.com/knative-prow-updater-robot)

- Automates updates of Prow itself.

## Knative Test Reporter Robot [@knative-test-reporter-robot](https://github.com/knative-test-reporter-robot)

- Reports test failures potentially across all projects by opening GitHub Issues. 

