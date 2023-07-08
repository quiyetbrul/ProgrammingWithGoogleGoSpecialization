# Concurrency

## Performance Limits
- Moore's Law used to help performance
  - number of transistors doubles every 18 months (outdated)
- more transistors used to lead to higher clock frequencies
- power/temperature constraints limit clock frequencies now

## Parallelism
- number of cores still increases over time
- multiple tasks may be performed at the same time on different cores
- difficulties with parallelism
  - when do tasks start/stop?
  - what if one task needs data from another task?
  - do tasks conflict in memory?
