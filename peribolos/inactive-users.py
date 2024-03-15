import sys
from strictyaml import load, Map, Str, Int, Seq, Email, Bool, Optional,MapPattern, YAMLError, as_document

if len(sys.argv) != 3:
    raise RuntimeError(f"\n\tUsage:\n\t{sys.argv[0]} <path_to_peribolos_yml> <list of users>")

with open(sys.argv[1],"r") as f:
    yaml = f.read()

with open(sys.argv[2],"r") as f:
    active_users = set(line.strip().lower() for line in f)

# add bots
active_users.add('knative-automation')
active_users.add('knative-metrics-robot')
active_users.add('knative-prow-releaser-robot')
active_users.add('knative-prow-robot')
active_users.add('knative-prow-updater-robot')
active_users.add('knative-test-reporter-robot')

config=load(yaml).data
in_active=set()

for name, org in config['orgs'].items():
    for member in org['members']:
        if member.lower() not in active_users:
            in_active.add(member)

    for name, team in org['teams'].items():
        if 'members' in team:
            for member in team['members']:
                if member.lower() not in active_users:
                    in_active.add(member)

        if 'maintainers' in team:
            for member in team['maintainers']:
                if member.lower() not in active_users:
                    in_active.add(member)

        if 'teams' in team:
            for name, team in team['teams'].items():
                if 'members' in team:
                    for member in team['members']:
                        if member.lower() not in active_users:
                            in_active.add(member)

                if 'maintainers' in team:
                    for member in team['maintainers']:
                        if member.lower() not in active_users:
                            in_active.add(member)

print('\n'.join(sorted(in_active)))

