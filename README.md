# Golang-Channels

**The Idea:** This project was born out of me trying to understand Golang channels and goroutine concurrency. Early hail Mary experiments failed show the benefits of of this pattern, so this is a bare bones benchmark and comparison written by eliminating as many variables as possible.

**The Initial Problem:** While my fundamental pattern was solid, the jobs that the workers were performing had too much variability and suffered from a lot of overhead from context switching, coordinating shared memory, and other processor-related overhead like synchronization. It was in writing this project that I was able to illustrate the fundamental nature of this pattern, as well as figure out why my initial attempts failed to show the benefits. **Just spawning more workers and goroutines will not necessarily benefit every type of processing job.** The overall takeaway for me what that sometimes the processor overhead created by running too many goroutines outweighs the performance increase.

**The Solution:** Separating out my code into distinct packages and narrowing the scope of the jobs performed (as well as making them consistent) was the key to seeing exactly how performant this pattern could be.

**Overview:** Fundamentally, this project does a few things:

1. Uses the pattern: `Dispatcher (spawns) -> Workers (process) -> Jobs`
2. Uses `runtime.NumCPU()` to see how many processors are available, in turn, generating **one worker per CPU** (or core). In my environment, I have a total of **8** cores available in total.
    1. By default the application will start with **1** worker to process the jobs and output the timing metrics.
    2. Then it will add **one more worker** and process the **same exact batch of jobs**.
    3. It will continue this process until it has reached the maximum number of CPUs as defined by `runtime.NumCPU()`. For my environment containing **8** CPUs, it will run **8** times, starting with **1** worker, adding another **single worker** with each iteration, and finally running the test with all **8** workers.
    4. Each worker and job gets its own unique UUID in order to see which worker is processing which job in the results.
3. The job being processed is just a 1 second `time.Sleep` function, so that there is very little variation in the time it takes a worker to execute it's number of jobs. By default, each worker runs twice as many jobs as the total number of workers, e.g.: `runtime.NumCPU() * 2`. This is a fairly arbitrary choice, but I felt that if this were a real world example, the workers would have more than a single queued job.

**EmptySleepJob():**

The important bit below (the "work") is: `time.Sleep(time.Duration(config.EmptySleepJobSleepTimeMs) * time.Millisecond)` The rest of the code is just Marshaling the timing data into a struct in order to collect the granular and overall timing of the job processing.

```
func (job Job) EmptySleepJob() (string, float64) {
	jobStartTime := time.Now()

	time.Sleep(time.Duration(config.EmptySleepJobSleepTimeMs) * time.Millisecond)

	jobEndTime := time.Now()
	jobElapsed := jobEndTime.Sub(jobStartTime)

	jobResult := structs.SleepJobResult{}
	jobResult.SleepTime = time.Duration(config.EmptySleepJobSleepTimeMs).String()
	jobResult.Elapsed = jobElapsed.String()
	jobResult.Status = strconv.FormatBool(true)

	jobResultString, err := json.Marshal(jobResult)
	if err != nil {
		fmt.Println(err)
	}

	return string(jobResultString), jobElapsed.Seconds()
}
```

**Example EmptySleepJob() Results:**

The table below represents:

* **1** to **8** workers running the `emptySleepJob()` function 16 times. Note: The output below **represents the final run with the maximum number of workers**. Runs with 1-7 workers were omitted for terseness. **The important part is the summary table of all 8 runs**.
* Starting with **1** worker `(baseline)`, after each completion the worker count is increased by **1**, until `runtime.NumCPU()` number of workers (**8** CPUs in the example) is reached.
* As each job is just a `time.Sleep(1 * time.Second)` command, each job should take about **1 second to execute.**
* As expected, **1 worker completing 16 of these jobs takes about 16 seconds** and **8 workers runs approximately 8 times as fast.**

```
--------------------------------------
|  Spawning workers for test 8 of 8  |
--------------------------------------

Workers: 8                           Job Name: emptySleepJob
------------------------------------------------------------
  Workers In Use:                    8
  Workers Available:                 8
  Workers Idle:                      0
  Number of Jobs:                    16
------------------------------------------------------------

Allocating Worker #1:                391755d9-bc82-4c9c-8914-4137d2f528af
Allocating Worker #2:                39f98b30-d928-4268-a620-c083f8f09855
Allocating Worker #3:                1cd12e59-af15-4ff4-99b5-80323473139d
Allocating Worker #4:                bf4fa589-80ba-4b56-9c00-4417163fd201
Allocating Worker #5:                8625e635-1e1a-4e63-a33d-854be7a0de96
Allocating Worker #6:                dd1b66d7-c5cc-4c30-a31d-17d6c0e9e605
Allocating Worker #7:                22032e71-9928-41b7-894f-5b2d4d4297f1
Allocating Worker #8:                f46f2080-62dc-45fb-8582-d4d0fec75252
  Allocating Job #1:                 fd7adabd-3110-4936-9373-5476d811204e
  Allocating Job #2:                 004eae10-a38a-4464-8142-1fd5fa104a43
  Allocating Job #3:                 c3ec6f42-a509-44ef-bc0b-e8e85a985b90
  Allocating Job #4:                 c9cc844a-d031-49e5-b8c3-c21e22c3f77c
  Allocating Job #5:                 2bd2783d-bb7c-4d6a-9f28-46a304df11a1
  Allocating Job #6:                 fc6624a6-a86c-40ed-a199-a02674f6a8cc
  Allocating Job #7:                 867f7c61-f4c3-488f-b3ac-9e9fe2ea3aa8
  Allocating Job #8:                 0f264516-1bfd-44d6-ba0b-b98910578258
  Allocating Job #9:                 bcb65aa7-6b72-49b1-86de-01d6c34a203a
  Allocating Job #10:                a09aef1a-12a5-4ccf-8b1e-0c82604ec07f
  Allocating Job #11:                783cee5a-2959-4a5b-a7c1-c31017442a14
  Allocating Job #12:                47e23ff7-4001-402f-a94a-50da80b501cb
  Allocating Job #13:                3daca53a-a6d1-41d3-abd9-76fb152da200
  Allocating Job #14:                097ae54f-3611-45f9-90eb-9788390e98d5
  Allocating Job #15:                59c73a0f-54e4-488c-b196-260d1d29fea3
  Allocating Job #16:                1dc2c78d-d051-49e6-8985-45b9710d073f
  JOB 1/16 STARTED:                  fd7adabd-3110-4936-9373-5476d811204e with Worker: f46f2080-62dc-45fb-8582-d4d0fec75252
  JOB 2/16 STARTED:                  004eae10-a38a-4464-8142-1fd5fa104a43 with Worker: 391755d9-bc82-4c9c-8914-4137d2f528af
  JOB 3/16 STARTED:                  c3ec6f42-a509-44ef-bc0b-e8e85a985b90 with Worker: 39f98b30-d928-4268-a620-c083f8f09855
  JOB 4/16 STARTED:                  c9cc844a-d031-49e5-b8c3-c21e22c3f77c with Worker: 1cd12e59-af15-4ff4-99b5-80323473139d
  JOB 5/16 STARTED:                  2bd2783d-bb7c-4d6a-9f28-46a304df11a1 with Worker: bf4fa589-80ba-4b56-9c00-4417163fd201
  JOB 6/16 STARTED:                  fc6624a6-a86c-40ed-a199-a02674f6a8cc with Worker: 8625e635-1e1a-4e63-a33d-854be7a0de96
  JOB 7/16 STARTED:                  867f7c61-f4c3-488f-b3ac-9e9fe2ea3aa8 with Worker: dd1b66d7-c5cc-4c30-a31d-17d6c0e9e605
  JOB 8/16 STARTED:                  0f264516-1bfd-44d6-ba0b-b98910578258 with Worker: 22032e71-9928-41b7-894f-5b2d4d4297f1
  JOB 9/16 STARTED:                  bcb65aa7-6b72-49b1-86de-01d6c34a203a with Worker: 22032e71-9928-41b7-894f-5b2d4d4297f1
    -> JOB 8/16 COMPLETED:           0f264516-1bfd-44d6-ba0b-b98910578258 with Worker: 22032e71-9928-41b7-894f-5b2d4d4297f1 (Ran emptySleepJob in 1.001227428 Seconds)
  JOB 10/16 STARTED:                 a09aef1a-12a5-4ccf-8b1e-0c82604ec07f with Worker: bf4fa589-80ba-4b56-9c00-4417163fd201
    -> JOB 5/16 COMPLETED:           2bd2783d-bb7c-4d6a-9f28-46a304df11a1 with Worker: bf4fa589-80ba-4b56-9c00-4417163fd201 (Ran emptySleepJob in 1.001321472 Seconds)
    -> JOB 4/16 COMPLETED:           c9cc844a-d031-49e5-b8c3-c21e22c3f77c with Worker: 1cd12e59-af15-4ff4-99b5-80323473139d (Ran emptySleepJob in 1.001343448 Seconds)
  JOB 11/16 STARTED:                 783cee5a-2959-4a5b-a7c1-c31017442a14 with Worker: 1cd12e59-af15-4ff4-99b5-80323473139d
  JOB 15/16 STARTED:                 59c73a0f-54e4-488c-b196-260d1d29fea3 with Worker: f46f2080-62dc-45fb-8582-d4d0fec75252
  JOB 14/16 STARTED:                 097ae54f-3611-45f9-90eb-9788390e98d5 with Worker: 391755d9-bc82-4c9c-8914-4137d2f528af
    -> JOB 7/16 COMPLETED:           867f7c61-f4c3-488f-b3ac-9e9fe2ea3aa8 with Worker: dd1b66d7-c5cc-4c30-a31d-17d6c0e9e605 (Ran emptySleepJob in 1.001342991 Seconds)
    -> JOB 6/16 COMPLETED:           fc6624a6-a86c-40ed-a199-a02674f6a8cc with Worker: 8625e635-1e1a-4e63-a33d-854be7a0de96 (Ran emptySleepJob in 1.00134668 Seconds)
  JOB 16/16 STARTED:                 1dc2c78d-d051-49e6-8985-45b9710d073f with Worker: 39f98b30-d928-4268-a620-c083f8f09855
  JOB 12/16 STARTED:                 47e23ff7-4001-402f-a94a-50da80b501cb with Worker: dd1b66d7-c5cc-4c30-a31d-17d6c0e9e605
  JOB 13/16 STARTED:                 3daca53a-a6d1-41d3-abd9-76fb152da200 with Worker: 8625e635-1e1a-4e63-a33d-854be7a0de96
    -> JOB 2/16 COMPLETED:           004eae10-a38a-4464-8142-1fd5fa104a43 with Worker: 391755d9-bc82-4c9c-8914-4137d2f528af (Ran emptySleepJob in 1.001376729 Seconds)
    -> JOB 1/16 COMPLETED:           fd7adabd-3110-4936-9373-5476d811204e with Worker: f46f2080-62dc-45fb-8582-d4d0fec75252 (Ran emptySleepJob in 1.001385292 Seconds)
    -> JOB 3/16 COMPLETED:           c3ec6f42-a509-44ef-bc0b-e8e85a985b90 with Worker: 39f98b30-d928-4268-a620-c083f8f09855 (Ran emptySleepJob in 1.001376315 Seconds)
    -> JOB 15/16 COMPLETED:          59c73a0f-54e4-488c-b196-260d1d29fea3 with Worker: f46f2080-62dc-45fb-8582-d4d0fec75252 (Ran emptySleepJob in 1.001151321 Seconds)
    -> JOB 13/16 COMPLETED:          3daca53a-a6d1-41d3-abd9-76fb152da200 with Worker: 8625e635-1e1a-4e63-a33d-854be7a0de96 (Ran emptySleepJob in 1.001117998 Seconds)
    -> JOB 9/16 COMPLETED:           bcb65aa7-6b72-49b1-86de-01d6c34a203a with Worker: 22032e71-9928-41b7-894f-5b2d4d4297f1 (Ran emptySleepJob in 1.001285919 Seconds)
    -> JOB 10/16 COMPLETED:          a09aef1a-12a5-4ccf-8b1e-0c82604ec07f with Worker: bf4fa589-80ba-4b56-9c00-4417163fd201 (Ran emptySleepJob in 1.001233061 Seconds)
    -> JOB 12/16 COMPLETED:          47e23ff7-4001-402f-a94a-50da80b501cb with Worker: dd1b66d7-c5cc-4c30-a31d-17d6c0e9e605 (Ran emptySleepJob in 1.001143418 Seconds)
    -> JOB 16/16 COMPLETED:          1dc2c78d-d051-49e6-8985-45b9710d073f with Worker: 39f98b30-d928-4268-a620-c083f8f09855 (Ran emptySleepJob in 1.001150912 Seconds)
    -> JOB 11/16 COMPLETED:          783cee5a-2959-4a5b-a7c1-c31017442a14 with Worker: 1cd12e59-af15-4ff4-99b5-80323473139d (Ran emptySleepJob in 1.001207857 Seconds)
    -> JOB 14/16 COMPLETED:          097ae54f-3611-45f9-90eb-9788390e98d5 with Worker: 391755d9-bc82-4c9c-8914-4137d2f528af (Ran emptySleepJob in 1.001135284 Seconds)

-----------------------------------------------------------------------
Total time taken:                    2.003099 Seconds

Summary:
+-------------------+----------------+----------------+----------------+
| NUMBER OF WORKERS | NUMBER OF JOBS | EXECUTION TIME | SPEED INCREASE |
+-------------------+----------------+----------------+----------------+
|                 1 |             16 |      16.048032 | (baseline)     |
|                 2 |             16 |       8.046808 | +1.99x         |
|                 3 |             16 |       6.026783 | +2.66x         |
|                 4 |             16 |       4.028621 | +3.98x         |
|                 5 |             16 |       4.021466 | +3.99x         |
|                 6 |             16 |       3.022764 | +5.31x         |
|                 7 |             16 |       3.007792 | +5.34x         |
|                 8 |             16 |       2.003099 | +8.01x         |
+-------------------+----------------+----------------+----------------+
```

Benchmark comparison: `make-golang-benchmark`

```
Running benchmarks...
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkRun/Worker_Count:_1-8                 1        46110535159 ns/op         223928 B/op       3882 allocs/op
BenchmarkRun/Worker_Count:_2-8                 1        30029084843 ns/op         197120 B/op       3336 allocs/op
BenchmarkRun/Worker_Count:_3-8                 1        22048430705 ns/op         217528 B/op       2856 allocs/op
BenchmarkRun/Worker_Count:_4-8                 1        16022740159 ns/op         111064 B/op       2379 allocs/op
BenchmarkRun/Worker_Count:_5-8                 1        12008797437 ns/op          87400 B/op       1908 allocs/op
BenchmarkRun/Worker_Count:_6-8                 1        8010978770 ns/op           67784 B/op       1434 allocs/op
BenchmarkRun/Worker_Count:_7-8                 1        5005642671 ns/op           45944 B/op        966 allocs/op
BenchmarkRun/Worker_Count:_8-8                 1        2008275612 ns/op           24672 B/op        505 allocs/op
PASS
ok      command-line-arguments  141.255s
...Complete.
```

**Overloading EmptySleepJob() Results:**

In the above example, when limiting the maximum number or workers to `runtime.NumCPU()`, the application processing speed increases as expected. But what if we pass this limit?

Below are the results for running 64 workers, or `runtime.NumCPU() * 8`:

```
+-------------------+----------------+----------------+----------------+
| NUMBER OF WORKERS | NUMBER OF JOBS | EXECUTION TIME | SPEED INCREASE |
+-------------------+----------------+----------------+----------------+
|                 1 |             16 |      16.032758 | (baseline)     |
|                 2 |             16 |       8.031666 | +2x            |
|                 3 |             16 |       6.023406 | +2.66x         |
|                 4 |             16 |       4.013871 | +3.99x         |
|                 5 |             16 |       4.023615 | +3.98x         |
|                 6 |             16 |       3.017553 | +5.31x         |
|                 7 |             16 |       3.014057 | +5.32x         |
|                 8 |             16 |       2.004924 | +8x            |
|                 9 |             16 |       2.004691 | +8x            |
|                10 |             16 |       2.012549 | +7.97x         |
|                11 |             16 |       2.019926 | +7.94x         |
|                12 |             16 |       2.002311 | +8.01x         |
|                13 |             16 |       2.015261 | +7.96x         |
|                14 |             16 |       2.003148 | +8x            |
|                15 |             16 |       2.022867 | +7.93x         |
|                16 |             16 |       1.004523 | +15.96x        |
|                17 |             16 |       1.001547 | +16.01x        |
|                18 |             16 |       1.013735 | +15.82x        |
|                19 |             16 |       1.021544 | +15.69x        |
|                20 |             16 |       1.026882 | +15.61x        |
|                21 |             16 |       1.011589 | +15.85x        |
|                22 |             16 |       1.014409 | +15.81x        |
|                23 |             16 |       1.001913 | +16x           |
|                24 |             16 |       1.001759 | +16x           |
|                25 |             16 |       1.000737 | +16.02x        |
|                26 |             16 |       1.000726 | +16.02x        |
|                27 |             16 |       1.005954 | +15.94x        |
|                28 |             16 |       1.005635 | +15.94x        |
|                29 |             16 |       1.011591 | +15.85x        |
|                30 |             16 |       1.002436 | +15.99x        |
|                31 |             16 |       1.001021 | +16.02x        |
|                32 |             16 |       1.029714 | +15.57x        |
|                33 |             16 |       1.001840 | +16x           |
|                34 |             16 |       1.001056 | +16.02x        |
|                35 |             16 |       1.009254 | +15.89x        |
|                36 |             16 |       1.002040 | +16x           |
|                37 |             16 |       1.001370 | +16.01x        |
|                38 |             16 |       1.001989 | +16x           |
|                39 |             16 |       1.002544 | +15.99x        |
|                40 |             16 |       1.007401 | +15.91x        |
|                41 |             16 |       1.004792 | +15.96x        |
|                42 |             16 |       1.005586 | +15.94x        |
|                43 |             16 |       1.000805 | +16.02x        |
|                44 |             16 |       1.001293 | +16.01x        |
|                45 |             16 |       1.002134 | +16x           |
|                46 |             16 |       1.018746 | +15.74x        |
|                47 |             16 |       1.039963 | +15.42x        |
|                48 |             16 |       1.005729 | +15.94x        |
|                49 |             16 |       1.016810 | +15.77x        |
|                50 |             16 |       1.009616 | +15.88x        |
|                51 |             16 |       1.008329 | +15.9x         |
|                52 |             16 |       1.000620 | +16.02x        |
|                53 |             16 |       1.000831 | +16.02x        |
|                54 |             16 |       1.001612 | +16.01x        |
|                55 |             16 |       1.001994 | +16x           |
|                56 |             16 |       1.002233 | +16x           |
|                57 |             16 |       1.002621 | +15.99x        |
|                58 |             16 |       1.004138 | +15.97x        |
|                59 |             16 |       1.002829 | +15.99x        |
|                60 |             16 |       1.001385 | +16.01x        |
|                61 |             16 |       1.001539 | +16.01x        |
|                62 |             16 |       1.001430 | +16.01x        |
|                63 |             16 |       1.002649 | +15.99x        |
|                64 |             16 |       1.002106 | +16x           |
+-------------------+----------------+----------------+----------------+


```


## .env

The supplied .env file is blank by default, and should run properly without making changes:

```
DEBUG=
STARTINGWORKERCOUNT=
MAXWORKERCOUNT=
TOTALJOBCOUNT=
PPROF=false
PPROFIP=
PPROFPORT=
```

The defaults are listed below when no values are set. **To override any/all, just set the env vars below.**

**DEBUG:** More verbose console output, default is `false`
**STARTINGWORKERCOUNT:** By default this will be set to `1`
**MAXWORKERCOUNT:** By default this will be set to `runtime.NumCPU()`
**TOTALJOBCOUNT:** By default this will be set to `runtime.NumCPU()*2`
**PPROF:** By default this will be set to `false` and the pprof server will not start.
**PPROFIP:** If PPROF is set to `true`, the app will try to figure out your local IP address. If it does not do this correctly, simply set your local IP Address here.
**PPROFPORT:** By default this will be set to `6060`


## Running the application

`go run .`

Or: `make golang-run` Note: This builds and executes the binary, not equivalent to: `go run .`

## Building the application binary

`make golang-build` Outputs binary to `./bin/` directory.

## Running the application binary

`./bin/golangchannels`

Or: `make golang-run` Note: This builds and executes the binary, not equivalent to: `go run .`

## Other convenience commands:

List available commands: `make`

```
Targets in this Makefile:

make golang-build
make golang-godoc
make golang-lint
make golang-run
make golang-test

For details on these commands, see the bash scripts in the 'scripts/' directory.
```

### make golang-lint

`make golang-lint` To lint the application.  **[golangci-lint](https://golangci-lint.run/usage/install/) must be installed**

### make golang-test

`make golang-test` To test the application using `gotestsum` **[gotestsum](https://github.com/gotestyourself/gotestsum#install) must be installed.**

## Benchmark Tests

Results of running the application: 8 Iterations, starting with 1 Worker, ending with 8 Workers, all processing 64 of the same job (PiJob)

```
Summary:
+-------------------+----------------+----------------------------+-----------------------------+----------------+
| NUMBER OF WORKERS | NUMBER OF JOBS | AVERAGE JOB EXECUTION TIME | TOTAL WORKER EXECUTION TIME | SPEED INCREASE |
+-------------------+----------------+----------------------------+-----------------------------+----------------+
|                 1 |             64 |                   1.005249 |                   64.455789 | (baseline)     |
|                 2 |             64 |                   2.125810 |                   35.977528 | +1.79x         |
|                 3 |             64 |                   3.469966 |                   29.420817 | +2.19x         |
|                 4 |             64 |                   5.061737 |                   25.585365 | +2.52x         |
|                 5 |             64 |                   6.913418 |                   24.061442 | +2.68x         |
|                 6 |             64 |                   8.898033 |                   21.736235 | +2.97x         |
|                 7 |             64 |                  11.084222 |                   20.750771 | +3.11x         |
|                 8 |             64 |                  13.642186 |                   20.672078 | +3.12x         |
+-------------------+----------------+----------------------------+-----------------------------+----------------+
```

### Results of the benchmarks:

The following is a lot of commands and benchmark output. The idea here is that the output roughly matches the timing of runing the application. As there is a small variability in how long it takes the jobs to execute--usually much less than a second--the numbers won't be exact. The following is simply confirmation that we're getting the results that we expect.

The following tests were run on:

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -startingWorkerCount=1 -maxWorkerCount=1 -jobCount=64` #=>

```
startingWorkerCount: 1 maxWorkerCount: 1 jobCount: 64
BenchmarkRun/Worker_Count:_1-8                 1        64292688948 ns/op       35976337824 B/op        23511673 allocs/op
PASS
ok      command-line-arguments  64.467s
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -startingWorkerCount=2 -maxWorkerCount=2 -jobCount=64` #=>

```
startingWorkerCount: 2 maxWorkerCount: 2 jobCount: 64
BenchmarkRun/Worker_Count:_2-8                 1        38487971462 ns/op       36086760640 B/op        23605364 allocs/op
PASS
ok      command-line-arguments  38.683s
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -startingWorkerCount=3 -maxWorkerCount=3 -jobCount=64` #=>

```
startingWorkerCount: 3 maxWorkerCount: 3 jobCount: 64
BenchmarkRun/Worker_Count:_3-8                 1        31649478509 ns/op       36098221112 B/op        23631627 allocs/op
PASS
ok      command-line-arguments  31.873s
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -startingWorkerCount=4 -maxWorkerCount=4 -jobCount=64` #=>

```
startingWorkerCount: 4 maxWorkerCount: 4 jobCount: 64
BenchmarkRun/Worker_Count:_4-8                 1        25412844723 ns/op       36074032056 B/op        23617726 allocs/op
PASS
ok      command-line-arguments  25.625s
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -startingWorkerCount=5 -maxWorkerCount=5 -jobCount=64` #=>

```
startingWorkerCount: 5 maxWorkerCount: 5 jobCount: 64
BenchmarkRun/Worker_Count:_5-8                 1        20767157101 ns/op       36104025472 B/op        23640474 allocs/op
PASS
ok      command-line-arguments  20.914s
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -startingWorkerCount=6 -maxWorkerCount=6 -jobCount=64` #=>

```
startingWorkerCount: 6 maxWorkerCount: 6 jobCount: 64
BenchmarkRun/Worker_Count:_6-8                 1        18937268938 ns/op       36135256720 B/op        23655645 allocs/op
PASS
ok      command-line-arguments  19.092s
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -startingWorkerCount=7 -maxWorkerCount=7 -jobCount=64` #=>

```
startingWorkerCount: 7 maxWorkerCount: 7 jobCount: 64
BenchmarkRun/Worker_Count:_7-8                 1        18182033241 ns/op       36135543352 B/op        23654627 allocs/op
PASS
ok      command-line-arguments  18.386s
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -startingWorkerCount=8 -maxWorkerCount=8 -jobCount=64` #=>

```
startingWorkerCount: 8 maxWorkerCount: 8 jobCount: 64
BenchmarkRun/Worker_Count:_8-8                 1        17796937403 ns/op       36178650928 B/op        23679568 allocs/op
PASS
ok      command-line-arguments  18.024s
```

---

# Comparing CPU profiles via pperf:

clear && go tool pprof pprof/cpuprofile-01.out

```
File: dispatcher.test
Type: cpu
Time: Apr 12, 2023 at 12:53pm (PDT)
Duration: 64.44s, Total samples = 68.93s (106.97%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top1000
Showing nodes accounting for 62.46s, 90.61% of 68.93s total
Dropped 261 nodes (cum <= 0.34s)
      flat  flat%   sum%        cum   cum%
    41.56s 60.29% 60.29%     41.56s 60.29%  math/big.addMulVVW
     3.69s  5.35% 65.65%     46.83s 67.94%  math/big.basicSqr
     2.81s  4.08% 69.72%      2.81s  4.08%  math/big.mulAddVWW
     1.89s  2.74% 72.46%      1.89s  2.74%  math/big.subVV
     1.78s  2.58% 75.05%      1.78s  2.58%  runtime.memclrNoHeapPointers
     1.59s  2.31% 77.35%      1.59s  2.31%  math/big.shlVU
     1.11s  1.61% 78.96%      2.10s  3.05%  runtime.scanobject
     0.93s  1.35% 80.31%      1.20s  1.74%  runtime.scanblock
     0.88s  1.28% 81.59%      0.88s  1.28%  math/big.addVV
     0.78s  1.13% 82.72%      0.78s  1.13%  runtime.futex
     0.72s  1.04% 83.77%      0.73s  1.06%  math/big.nat.norm (inline)
     0.62s   0.9% 84.67%      0.62s   0.9%  runtime.madvise
     0.53s  0.77% 85.43%      0.53s  0.77%  runtime.tgkill
     0.50s  0.73% 86.16%      0.50s  0.73%  runtime.procyield
     0.49s  0.71% 86.87%      4.81s  6.98%  math/big.nat.divBasic
     0.40s  0.58% 87.45%      0.40s  0.58%  runtime.memmove
     0.37s  0.54% 87.99%      0.50s  0.73%  runtime.findObject
     0.35s  0.51% 88.50%      6.09s  8.84%  runtime.mallocgc
     0.15s  0.22% 88.71%      5.56s  8.07%  math/big.nat.divRecursiveStep
     0.14s   0.2% 88.92%      0.53s  0.77%  runtime.(*mheap).allocSpan
     0.10s  0.15% 89.06%     52.27s 75.83%  math/big.nat.expNN
     0.10s  0.15% 89.21%      1.80s  2.61%  math/big.nat.mul
     0.08s  0.12% 89.32%      0.93s  1.35%  runtime.bgscavenge
     0.07s   0.1% 89.42%     62.90s 91.25%  github.com/mattwiater/golangchannels/jobs/piJob.Pi
     0.07s   0.1% 89.53%     37.74s 54.75%  math/big.karatsubaSqr
     0.06s 0.087% 89.61%      0.48s   0.7%  runtime.gentraceback
     0.06s 0.087% 89.70%      6.13s  8.89%  runtime.makeslice
     0.05s 0.073% 89.77%      1.54s  2.23%  math/big.(*Float).Add
     0.05s 0.073% 89.84%      0.40s  0.58%  runtime.(*sweepLocked).sweep
     0.05s 0.073% 89.92%      2.83s  4.11%  runtime.gcDrain
     0.04s 0.058% 89.98%      0.41s  0.59%  math/big.karatsubaAdd
     0.04s 0.058% 90.03%      0.46s  0.67%  runtime.sweepone
     0.03s 0.044% 90.08%      5.65s  8.20%  math/big.nat.make (inline)
     0.03s 0.044% 90.12%      1.40s  2.03%  math/big.nat.mulAddWW
     0.03s 0.044% 90.16%      0.86s  1.25%  runtime.schedule
     0.02s 0.029% 90.19%      7.49s 10.87%  math/big.(*Float).uquo
     0.02s 0.029% 90.22%      0.81s  1.18%  math/big.(*Float).usub
     0.02s 0.029% 90.25%     52.29s 75.86%  math/big.(*Int).Exp
     0.02s 0.029% 90.28%      0.62s   0.9%  math/big.fnorm
     0.02s 0.029% 90.31%         7s 10.16%  math/big.nat.divLarge
     0.02s 0.029% 90.34%      5.87s  8.52%  math/big.nat.divRecursive
     0.02s 0.029% 90.37%     50.70s 73.55%  math/big.nat.sqr
     0.02s 0.029% 90.40%      0.98s  1.42%  runtime.(*mcache).nextFree
     0.02s 0.029% 90.43%      0.95s  1.38%  runtime.(*mcache).refill
     0.02s 0.029% 90.45%      0.68s  0.99%  runtime.(*pageAlloc).scavenge
     0.01s 0.015% 90.47%      1.05s  1.52%  math/big.(*Float).SetInt
     0.01s 0.015% 90.48%      0.74s  1.07%  math/big.(*Float).uadd
     0.01s 0.015% 90.50%      0.64s  0.93%  runtime.(*mcentral).cacheSpan
     0.01s 0.015% 90.51%      0.59s  0.86%  runtime.(*mheap).alloc
     0.01s 0.015% 90.53%      0.67s  0.97%  runtime.bgsweep
     0.01s 0.015% 90.54%      0.39s  0.57%  runtime.findrunnable
     0.01s 0.015% 90.56%      2.85s  4.13%  runtime.gcBgMarkWorker.func2
     0.01s 0.015% 90.57%      0.57s  0.83%  runtime.goschedImpl
     0.01s 0.015% 90.58%      0.46s  0.67%  runtime.preemptone
     0.01s 0.015% 90.60%      7.20s 10.45%  runtime.systemstack
     0.01s 0.015% 90.61%      0.61s  0.88%  runtime.wakep
         0     0% 90.61%     63.06s 91.48%  github.com/mattwiater/golangchannels/jobs/piJob.Job.PiJob
         0     0% 90.61%     63.06s 91.48%  github.com/mattwiater/golangchannels/workers.PerformJob (inline)
         0     0% 90.61%     63.06s 91.48%  github.com/mattwiater/golangchannels/workers.Worker
         0     0% 90.61%      7.49s 10.87%  math/big.(*Float).Quo
         0     0% 90.61%      7.01s 10.17%  math/big.nat.div
         0     0% 90.61%      0.69s  1.00%  math/big.nat.set (inline)
         0     0% 90.61%      0.75s  1.09%  math/big.nat.shl
         0     0% 90.61%      0.38s  0.55%  runtime.(*gcWork).balance
         0     0% 90.61%      0.56s  0.81%  runtime.(*mcentral).grow
         0     0% 90.61%      0.58s  0.84%  runtime.(*mheap).alloc.func1
         0     0% 90.61%      0.62s   0.9%  runtime.(*pageAlloc).scavenge.func1
         0     0% 90.61%      0.62s   0.9%  runtime.(*pageAlloc).scavengeOne
         0     0% 90.61%      0.62s   0.9%  runtime.(*pageAlloc).scavengeOneFast
         0     0% 90.61%      0.62s   0.9%  runtime.(*pageAlloc).scavengeRangeLocked
         0     0% 90.61%      0.67s  0.97%  runtime.futexwakeup
         0     0% 90.61%      1.80s  2.61%  runtime.gcAssistAlloc
         0     0% 90.61%      1.79s  2.60%  runtime.gcAssistAlloc.func1
         0     0% 90.61%      1.79s  2.60%  runtime.gcAssistAlloc1
         0     0% 90.61%      3.13s  4.54%  runtime.gcBgMarkWorker
         0     0% 90.61%      1.70s  2.47%  runtime.gcDrainN
         0     0% 90.61%      0.68s  0.99%  runtime.gcStart
         0     0% 90.61%      0.48s   0.7%  runtime.gcStart.func2
         0     0% 90.61%      0.38s  0.55%  runtime.gopreempt_m
         0     0% 90.61%      1.87s  2.71%  runtime.markroot
         0     0% 90.61%      0.71s  1.03%  runtime.markroot.func1
         0     0% 90.61%      1.16s  1.68%  runtime.markrootBlock
         0     0% 90.61%      0.57s  0.83%  runtime.mcall
         0     0% 90.61%      0.49s  0.71%  runtime.morestack
         0     0% 90.61%      0.49s  0.71%  runtime.newstack
         0     0% 90.61%      0.66s  0.96%  runtime.notewakeup
         0     0% 90.61%      0.37s  0.54%  runtime.park_m
         0     0% 90.61%      0.55s   0.8%  runtime.preemptM (inline)
         0     0% 90.61%      0.55s   0.8%  runtime.signalM
         0     0% 90.61%      0.51s  0.74%  runtime.startTheWorldWithSema
         0     0% 90.61%      0.60s  0.87%  runtime.startm
         0     0% 90.61%      0.43s  0.62%  runtime.suspendG
         0     0% 90.61%      0.62s   0.9%  runtime.sysUnused
```



clear && go tool pprof pprof/cpuprofile-08.out

```
File: dispatcher.test
Type: cpu
Time: Apr 12, 2023 at 12:56pm (PDT)
Duration: 18s, Total samples = 96.72s (537.40%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top1000
Showing nodes accounting for 87.25s, 90.21% of 96.72s total
Dropped 302 nodes (cum <= 0.48s)
      flat  flat%   sum%        cum   cum%
    49.58s 51.26% 51.26%     49.58s 51.26%  math/big.addMulVVW
     5.31s  5.49% 56.75%      5.31s  5.49%  runtime.memclrNoHeapPointers
     4.22s  4.36% 61.11%     56.89s 58.82%  math/big.basicSqr
     4.17s  4.31% 65.43%      4.17s  4.31%  runtime.madvise
     3.33s  3.44% 68.87%      3.33s  3.44%  math/big.mulAddVWW
     2.86s  2.96% 71.83%      2.86s  2.96%  runtime.procyield
     2.25s  2.33% 74.15%      2.25s  2.33%  math/big.subVV
     1.95s  2.02% 76.17%      1.95s  2.02%  math/big.shlVU
     1.95s  2.02% 78.18%      1.95s  2.02%  runtime.tgkill
     1.31s  1.35% 79.54%      1.31s  1.35%  math/big.addVV
     1.15s  1.19% 80.73%      2.32s  2.40%  runtime.scanobject
     0.96s  0.99% 81.72%      1.06s  1.10%  math/big.nat.norm (inline)
     0.95s  0.98% 82.70%      0.95s  0.98%  runtime.futex
     0.89s  0.92% 83.62%      0.91s  0.94%  runtime.asyncPreempt
     0.61s  0.63% 84.25%      5.79s  5.99%  math/big.nat.divBasic
     0.60s  0.62% 84.87%      0.60s  0.62%  runtime.memmove
     0.56s  0.58% 85.45%     11.82s 12.22%  runtime.mallocgc
     0.55s  0.57% 86.02%      0.74s  0.77%  runtime.findObject
     0.36s  0.37% 86.39%      0.59s  0.61%  runtime.(*spanSet).push
     0.34s  0.35% 86.75%      1.13s  1.17%  runtime.(*mheap).allocSpan
     0.29s   0.3% 87.05%      2.77s  2.86%  runtime.lock2
     0.26s  0.27% 87.31%      1.01s  1.04%  math/big.getNat
     0.23s  0.24% 87.55%      0.82s  0.85%  sync.(*Pool).Put
     0.18s  0.19% 87.74%     65.84s 68.07%  math/big.nat.expNN
     0.14s  0.14% 87.88%     82.60s 85.40%  github.com/mattwiater/golangchannels/jobs/piJob.Pi
     0.14s  0.14% 88.03%      7.33s  7.58%  math/big.nat.divRecursiveStep
     0.14s  0.14% 88.17%      0.79s  0.82%  runtime.gentraceback
     0.13s  0.13% 88.31%      2.67s  2.76%  math/big.nat.mul
     0.13s  0.13% 88.44%      2.04s  2.11%  runtime.sweepone
     0.12s  0.12% 88.56%      3.17s  3.28%  runtime.(*mcentral).cacheSpan
     0.12s  0.12% 88.69%      1.83s  1.89%  runtime.(*sweepLocked).sweep
     0.11s  0.11% 88.80%      3.77s  3.90%  runtime.(*mcache).refill
     0.11s  0.11% 88.92%      0.54s  0.56%  sync.(*Pool).Get
     0.09s 0.093% 89.01%      1.58s  1.63%  math/big.(*Float).uadd
     0.09s 0.093% 89.10%     45.48s 47.02%  math/big.karatsubaSqr
     0.09s 0.093% 89.20%     63.66s 65.82%  math/big.nat.sqr
     0.08s 0.083% 89.28%      1.68s  1.74%  runtime.schedule
     0.05s 0.052% 89.33%      4.18s  4.32%  runtime.gcDrain
     0.05s 0.052% 89.38%      2.85s  2.95%  runtime.newstack
     0.04s 0.041% 89.42%      3.04s  3.14%  math/big.(*Float).Add
     0.04s 0.041% 89.46%     65.90s 68.13%  math/big.(*Int).Exp
     0.04s 0.041% 89.51%     10.77s 11.14%  math/big.nat.make (inline)
     0.04s 0.041% 89.55%      1.99s  2.06%  math/big.nat.mulAddWW
     0.04s 0.041% 89.59%     11.75s 12.15%  runtime.makeslice
     0.04s 0.041% 89.63%      2.71s  2.80%  runtime.morestack
     0.03s 0.031% 89.66%      2.22s  2.30%  math/big.(*Float).SetInt
     0.03s 0.031% 89.69%      0.62s  0.64%  math/big.karatsubaAdd
     0.03s 0.031% 89.72%      9.45s  9.77%  math/big.nat.div
     0.03s 0.031% 89.75%      2.14s  2.21%  math/big.nat.shl
     0.03s 0.031% 89.78%     13.89s 14.36%  runtime.systemstack
     0.02s 0.021% 89.81%     10.29s 10.64%  math/big.(*Float).Quo
     0.02s 0.021% 89.83%     10.27s 10.62%  math/big.(*Float).uquo
     0.02s 0.021% 89.85%      1.73s  1.79%  math/big.(*Float).usub
     0.02s 0.021% 89.87%      0.73s  0.75%  math/big.fnorm
     0.02s 0.021% 89.89%      0.85s  0.88%  math/big.putNat (inline)
     0.02s 0.021% 89.91%      4.25s  4.39%  runtime.bgscavenge
     0.02s 0.021% 89.93%      1.26s  1.30%  runtime.forEachP
     0.02s 0.021% 89.95%      5.79s  5.99%  runtime.gcBgMarkWorker
     0.02s 0.021% 89.97%      0.69s  0.71%  runtime.gcDrainN
     0.02s 0.021% 89.99%      2.80s  2.89%  runtime.lock (inline)
     0.01s  0.01% 90.00%      0.54s  0.56%  math/big.nat.clear (inline)
     0.01s  0.01% 90.01%      9.41s  9.73%  math/big.nat.divLarge
     0.01s  0.01% 90.02%      7.98s  8.25%  math/big.nat.divRecursive
     0.01s  0.01% 90.03%      3.79s  3.92%  runtime.(*mcache).nextFree
     0.01s  0.01% 90.04%      1.21s  1.25%  runtime.(*mcentral).grow
     0.01s  0.01% 90.05%      0.60s  0.62%  runtime.(*mcentral).uncacheSpan
     0.01s  0.01% 90.06%      1.22s  1.26%  runtime.(*mheap).alloc
     0.01s  0.01% 90.07%      1.20s  1.24%  runtime.(*mheap).alloc.func1
     0.01s  0.01% 90.08%      1.17s  1.21%  runtime.(*mheap).freeSpan.func1
     0.01s  0.01% 90.10%      4.16s  4.30%  runtime.(*pageAlloc).scavengeOneFast
     0.01s  0.01% 90.11%      0.78s  0.81%  runtime.gcAssistAlloc
     0.01s  0.01% 90.12%      0.71s  0.73%  runtime.gcAssistAlloc1
     0.01s  0.01% 90.13%      0.55s  0.57%  runtime.gcstopm
     0.01s  0.01% 90.14%      2.64s  2.73%  runtime.gopreempt_m
     0.01s  0.01% 90.15%      2.78s  2.87%  runtime.lockWithRank
     0.01s  0.01% 90.16%      0.53s  0.55%  runtime.park_m
     0.01s  0.01% 90.17%      2.01s  2.08%  runtime.preemptM (inline)
     0.01s  0.01% 90.18%      1.61s  1.66%  runtime.preemptone
     0.01s  0.01% 90.19%         2s  2.07%  runtime.signalM
     0.01s  0.01% 90.20%      0.56s  0.58%  runtime.stopTheWorldWithSema
     0.01s  0.01% 90.21%      1.02s  1.05%  runtime.suspendG
         0     0% 90.21%     82.73s 85.54%  github.com/mattwiater/golangchannels/jobs/piJob.Job.PiJob
         0     0% 90.21%     82.73s 85.54%  github.com/mattwiater/golangchannels/workers.PerformJob (inline)
         0     0% 90.21%     82.73s 85.54%  github.com/mattwiater/golangchannels/workers.Worker
         0     0% 90.21%      1.64s  1.70%  math/big.nat.set (inline)
         0     0% 90.21%      1.17s  1.21%  runtime.(*mheap).freeSpan
         0     0% 90.21%      4.23s  4.37%  runtime.(*pageAlloc).scavenge
         0     0% 90.21%      4.19s  4.33%  runtime.(*pageAlloc).scavenge.func1
         0     0% 90.21%      4.19s  4.33%  runtime.(*pageAlloc).scavengeOne
         0     0% 90.21%      4.18s  4.32%  runtime.(*pageAlloc).scavengeRangeLocked
         0     0% 90.21%      1.64s  1.70%  runtime.deductSweepCredit
         0     0% 90.21%      0.74s  0.77%  runtime.futexwakeup
         0     0% 90.21%      0.71s  0.73%  runtime.gcAssistAlloc.func1
         0     0% 90.21%      4.18s  4.32%  runtime.gcBgMarkWorker.func2
         0     0% 90.21%      1.65s  1.71%  runtime.gcMarkDone
         0     0% 90.21%      0.88s  0.91%  runtime.gcMarkDone.func1
         0     0% 90.21%      0.49s  0.51%  runtime.gcMarkTermination
         0     0% 90.21%      0.54s  0.56%  runtime.gcStart
         0     0% 90.21%      2.74s  2.83%  runtime.goschedImpl
         0     0% 90.21%      2.03s  2.10%  runtime.markroot
         0     0% 90.21%      1.68s  1.74%  runtime.markroot.func1
         0     0% 90.21%      0.68s   0.7%  runtime.mcall
         0     0% 90.21%      0.56s  0.58%  runtime.notewakeup
         0     0% 90.21%      1.37s  1.42%  runtime.preemptall
         0     0% 90.21%      0.66s  0.68%  runtime.scanstack
         0     0% 90.21%      4.17s  4.31%  runtime.sysUnused
```