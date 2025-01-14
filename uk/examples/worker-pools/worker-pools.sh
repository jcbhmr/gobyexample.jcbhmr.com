# Наша програма покаже 5 завдань - що виконуються
# різними працівниками. Програма працює близько
# 2 секунд, не зважаючи на роботу розраховану
# вцілому на 5 секунд - і все це завдяки 3'ом
# працівникам, що працюються одночасно.
$ time go run worker-pools.go
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
