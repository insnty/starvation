# Concurrent Starvation
Concurrent process starvation demonstrated with two concurrent processes, where one process (starving_proc) starves the other (normal_proc) on the shared resource 'mutex'.

Example output:
```
2020/05/30 02:35:31 normal_proc: finished 425404 operations
2020/05/30 02:35:31 starving_proc: finished 688710 operations
2020/05/30 02:35:31 all processes finished
2020/05/30 02:35:31 starving_proc finished 263306 more operations than normal_proc