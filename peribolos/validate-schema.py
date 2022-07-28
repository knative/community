import sys
from strictyaml import load, Map, Str, Int, Seq, Email, Bool, Optional,MapPattern, YAMLError

if len(sys.argv) != 2:
    raise RuntimeError(f"\n\tUsage:\n\t{sys.argv[0]} <path_to_peribolos_yml.")

with  open(sys.argv[1],"r") as f:
    yaml = f.read()


teamSchema = MapPattern(
    Str(), Map({
        "description": Str(),
        "privacy": Str(),
        Optional("members"): Seq(Str()),
        Optional("maintainers"): Seq(Str()),
        Optional("repos"): MapPattern(Str(), Str()),
        Optional("teams"): MapPattern(
            Str(),Map({
                "description": Str(),
                "privacy": Str(),
                Optional("members"): Seq(Str()),
                Optional("maintainers"): Seq(Str()),
            }),
        ),
    })
)

schema = Map({
    "orgs": MapPattern(
        Str(), Map({
            "description": Str(),
            "name": Str(),
            "email": Email(),
            Optional("billing_email"): Email(),
            "has_organization_projects": Bool(),
            "has_repository_projects": Bool(),
            "members_can_create_repositories": Bool(),
            "default_repository_permission": Str(),
            "admins": Seq(Str()),
            "members": Seq(Str()),
            "repos": MapPattern(
                Str(), Map({
                    "allow_merge_commit": Bool(),
                    "allow_rebase_merge": Bool(),
                    Optional("archived"): Bool(),
                    Optional("description"): Str(),
                    "has_projects": Bool(),
                    "has_wiki": Bool(),
                    Optional("homepage"): Str(),
                }),
            ),
            "teams": teamSchema,
        }),
    )
})

load(yaml, schema)
print(f"{sys.argv[1]} validated.")
