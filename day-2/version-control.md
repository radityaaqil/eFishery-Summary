# VERSION CONTROL

Git is free and open source software for distributed version control: tracking changes in any set of files, usually used for coordinating work among programmers collaboratively developing source code during software development.

## Some most useful and used commands

| Commands                           | Function                                                             |
| ---------------------------------- | -------------------------------------------------------------------- |
| `git init <options> `              | Initialize Git in local file                                         |
| `git add <file> / <directory>`     | Moves changes from the working directory to the staging area.        |
| `git remote <options> `            | Add remote repository                                                |
| `git status `                      | Check local repo status                                              |
| `git commit <options> `            | Save changes before pushing into remote repo                         |
| `git push <options> `              | Upload local file into remote repo                                   |
| `git fetch <remote> [options] `    | Get latest update from remote repo                                   |
| `git pull <remote> <branch-name> ` | Get latest update and merge with local branch                        |
| `git branch <branch> `             | List local branches                                                  |
| `git merge <branch-name> `         | Merge two branches                                                   |
| `git stash <command> `             | Save and kinda hide our update in local                              |

## KARMA Convention

In committing our working files, there are rules to be followed to help organizing and tracking changes through the files.

Allowed type values: 
- feat for a new feature for the user, not a new feature for build script. Such commit will trigger a release bumping a MINOR version.
- fix for a bug fix for the user, not a fix to a build script. Such commit will trigger a release bumping a PATCH version.
- perf for performance improvements. Such commit will trigger a release bumping a PATCH version.
- docs for changes to the documentation.
- style for formatting changes, missing semicolons, etc.
- refactor for refactoring production code, e.g. renaming a variable.
- test for adding missing tests, refactoring tests; no production code change.
- build for updating build configuration, development tools or other changes irrelevant to the user.

Example scope values: 
- init
- runner
- watcher
- config
- web-server
- proxy
- etc.
The <scope> can be empty (e.g. if the change is a global or difficult to assign to a single component), in which case the parentheses are omitted. In smaller projects such as Karma plugins, the <scope> is empty.


### SEMANTIC VERSIONING
Example : version x.y.z
x major, y minor, z patch


### GIT MANAGEMENT
TRUNK BASED DEVELOPMENT --> single branch (main or master)
Why should we use trunk based development?

- Using this development method encourages all changes to come back to the mainline quickly. It can eliminate unnecessary divergence. And with a smaller team, it can keep overall software development costs down. This is because admins and engineers are not trying to maintain several “main” branches.

- Because the master branch is the only long-lived branch, all other branches have a defined and limited lifespan. In theory, this can help minimize merge conflicts because no branches are hanging out in development for too long.

- It enables developers to move fast. They do not need to deal with multiple steps before merging their code into the mainline. Because the fixes are applied right to the master branch, your code should always be production ready and waiting for the next release.
