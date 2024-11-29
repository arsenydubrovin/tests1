[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ekaterina0sokolova_software_testing1&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=adfsghjkmgfwdasfgthyfjmhgs_tests1)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=adfsghjkmgfwdasfgthyfjmhgs_tests1&metric=bugs)](https://sonarcloud.io/summary/new_code?id=adfsghjkmgfwdasfgthyfjmhgs_tests1)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=adfsghjkmgfwdasfgthyfjmhgs_tests1&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=adfsghjkmgfwdasfgthyfjmhgs_tests1)

### План тестирования

| Название | Тип | Входные данные | Ожидается |
|---|---|---|---|
| different roots | Позитивный | `a = 1`, `b = 2`, `c = 0` | `root1 = 0`, `root2 = -2`, `hasRoots = true`, `err = nil` |
| equal roots | Позитивный | `a = 2`, `b = -4`, `c = 2` | `root1 = 1`, `root2 = 1`, `hasRoots = true`, `err = nil` |
| zero 'a' | Позитивный | `a = 0`, `b = 2`, `c = 1` | `root1 = -0.5`, `root2 = -0.5`, `hasRoots = true`, `err = nil` |
| zero 'a' and 'b' | Негативный | `a = 0`, `b = 0`, `c = 1` | `hasRoots = false`, `err = ErrZeroCoefs` |
| no roots | Негативный | `a = 1`, `b = 0`, `c = 1` | `hasRoots = false`, `err = ErrNoRoots` |