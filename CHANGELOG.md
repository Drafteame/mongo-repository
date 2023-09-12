## v0.4.4 (2023-09-12)


- Merge pull request #17 from Drafteame/feat/remove-default-sort-by-id
- fix: remove unnecesary default sort
- fix: remove unnecesary default sort

## v0.4.3 (2023-06-29)


- Merge pull request #11 from Drafteame/fix/remove_search_limit_in_constructor
- fix: remove search limit from search options constructor
- fix: remove search limit from search options constructor

## v0.4.2 (2023-06-26)


- Merge pull request #10 from Drafteame/reafactor/add-so-as-generic-type
- refactor: decouple search signature and generify search options type
- refactor: decouple search signature and generify search options type

## v0.4.1 (2023-06-26)


- Merge pull request #9 from Drafteame/reafactor/add-so-as-generic-type
- refactor: add search orders to search options as generic type
- refactor: add search orders to search options as generic type

## v0.4.0 (2023-06-26)


- feat: add search interfaces to be able to replace defined structs (#7)
- * feat: add search interfaces to be able to replace defined structs

* chore: add interface assertion to builtin structs

## v0.3.0 (2023-06-26)


- build(deps): bump go.mongodb.org/mongo-driver from 1.11.6 to 1.12.0 (#8)
- Bumps [go.mongodb.org/mongo-driver](https://github.com/mongodb/mongo-go-driver) from 1.11.6 to 1.12.0.
- [Release notes](https://github.com/mongodb/mongo-go-driver/releases)
- [Commits](https://github.com/mongodb/mongo-go-driver/compare/v1.11.6...v1.12.0)

---
updated-dependencies:
- dependency-name: go.mongodb.org/mongo-driver
  dependency-type: direct:production
  update-type: version-update:semver-minor
...

Signed-off-by: dependabot[bot] <support@github.com>
Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>

## v0.2.0 (2023-06-01)


- feat: add mongo driver options constructor (#5)
- * feat: add mongo driver options constructor

* fix: error on empty client

## v0.1.1 (2023-05-31)


- fix: no timestamps conf (#4)
- * feat: add fill featured generic repo

* feat: add HardDelete method

* feat: add HardDeleteMany function

* refactor: generify SortOrders to be able to build precedence

* feat: add order constants

* refactor: examples

* refactor: raname internal log methods

* fix: logger tests

* refactor: use primitive.ObjectID instead of pointer

* fix: add testing and implementation for no timestamps

* fix: add updated add field on soft deletes

* deps: upgrade

* fix: assert dates

* ci: remove unneded lines

## v0.1.0 (2023-05-30)


- feat: full featured (#1)
- * feat: add fill featured generic repo

* ci: add ci configuration

* feat: add HardDelete method

* feat: add HardDeleteMany function

* refactor: generify SortOrders to be able to build precedence

* feat: add order constants

* refactor: examples

* docs: change readme

* deps: upgrade precommit

* fix: simplify validations delete many

* refactor: raname internal log methods

* fix: logger tests

* refactor: generify SearchOptions to a static struct instead of implementation

* refactor: use Now helper instead clock.Now property

* feat: add projection to search options
- Create README.md
- Initial commit
