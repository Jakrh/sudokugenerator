# Sudoku Generator

`DIFFICULTY_LEVEL` ranges from 1 to 5, where 1 is easy and 5 is hard.

Run

```bash
go run sudokugenerator DIFFICULTY_LEVEL
```

e.g.

```bash
go run sudokugenerator 2
```

Random output

```
┌───────┬───────┬───────┐
│   1 5 │   7 2 │ 9 4   │
│ 7     │ 8 9 6 │       │
│   2 6 │ 1 4   │   3   │
├───────┼───────┼───────┤
│ 3 5 7 │     1 │   8 9 │
│ 2   9 │     3 │ 4 1   │
│   4 1 │ 9 5   │     3 │
├───────┼───────┼───────┤
│ 1 7 2 │ 5 8   │   6   │
│ 4 9 3 │ 6 1   │ 8 5 2 │
│ 5     │     4 │   9 7 │
└───────┴───────┴───────┘
```
