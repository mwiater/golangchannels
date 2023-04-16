# Golang-Channels

## Reference Article

See my article on Medium here: https://medium.com/@matt.wiater/golang-channels-goroutines-and-optimal-concurrency-demystifying-through-examples-a43ba6aee74f

## Application

### Convenience Commands

Type `make` for a list of commands (comments added for clarity):

```
Targets in this Makefile:

make golang-benchmark # Run benchmark test: ./dispatcher/dispatcher_test.go
make golang-build     # Build the application binary to: ./bin/golangchannels
make golang-godoc     # Start the documentation server
make golang-lint      # Lint the application
make golang-pprof     # Generate pprof profiles in ./pprof
make golang-run       # Compile and run the binary
make golang-test      # Run the tests

For details on these commands, see the bash scripts in the 'scripts/' directory.
```

## Run the application

You can use the .env file to override the default values. Blank values will become the defaults below (comments added for clarity):

```
DEBUG=                // Verbose console output when running the app, default: false
JOBNAME=              // Which job to execute, default: EmptySleepJob
STARTINGWORKERCOUNT=  // How many workers to start the app with, default: 1
MAXWORKERCOUNT=       // How many workers to end the app with, default: runtime.NumCPU()
TOTALJOBCOUNT=        // How many jobs to run, default: runtime.NumCPU() * 2
```

With the default values above, on an 8 CPU core system, the application:

* Runs a total of 8 passes:
  * The first pass with STARTINGWORKERCOUNT workers,  in this example: `1`
  * Add one worker for the next pass
  * Repeat until MAXWORKERCOUNT is reached, in this example: `8`
* Display a list of Summary Results: 

Type: `make golang-run`


**Example EmptySleepJob() Results:**

The table below represents running the app with `DEBUG=true` (although the summary table will print whether `DEBUG` is `true` or `false`):

* **1** to **8** workers running the `emptySleepJob()` function 16 times. Note: The output below **represents the final run with the maximum number of workers**. Runs with 1-7 workers were omitted for brevity. **The important part is the Summary Results table of all 8 runs**.

Expectations:
* ✅ As each job is just a `time.Sleep(1 * time.Second)` command, each job should take about **1 second to execute.**
* ✅ As expected, **1 worker completing 16 of these jobs takes about 16 seconds** and **8 workers runs approximately 8 times as fast.**

```
...
--------------------------------------
|  Spawning workers for test 8 of 8  |
--------------------------------------

Workers: 8                           Job Name: EmptySleepJob
------------------------------------------------------------
  Workers In Use:                    8
  Workers Available:                 8
  Workers Idle:                      0
  Number of Jobs:                    16
------------------------------------------------------------

Allocating Worker #1:                ba9a7d7b-49da-4be9-9b85-50ac6b1a2fad
Allocating Worker #2:                94229bc1-3036-4a15-a5da-287fc3c16764
Allocating Worker #3:                995decd7-e2c9-4cf5-9571-71fe30761777
Allocating Worker #4:                8546abec-2440-4277-95d3-44a6adcc6fb2
Allocating Worker #5:                5a8e888f-e7ec-413b-a95e-66c9f4b07ee0
  Allocating Job #1:                 866a101b-ca22-41e4-bf5e-9c982a3c712a
  Allocating Job #2:                 4509977e-5fb9-4167-801d-080e492ff3d1
  Allocating Job #3:                 d3160123-2eeb-4df6-9717-c3dfedda97e8
  Allocating Job #4:                 73c99f65-f1f7-4bce-9897-748e71bea4f0
  Allocating Job #5:                 1b8f827a-2a9b-49ac-9eab-6948b4059f29
  Allocating Job #6:                 14b35713-a785-484d-b3aa-a98b3e288da1
  Allocating Job #7:                 02b5d815-abaa-478a-9d19-8c49e9b246ec
  Allocating Job #8:                 4c564808-6701-43c5-9830-9d65dad63f1a
  Allocating Job #9:                 8d3255c8-1051-4c97-b9d9-caa03772151f
  Allocating Job #10:                0e4bf9fd-3b3d-4d5e-858a-4de1f1b543de
  Allocating Job #11:                4d73dfa2-1f3d-488d-9a6c-d89fa2002d60
  Allocating Job #12:                11c0bbba-749b-4c0d-a8c3-b6a65aa9bed4
  JOB 5/16 STARTED:                  1b8f827a-2a9b-49ac-9eab-6948b4059f29 with Worker: 94229bc1-3036-4a15-a5da-287fc3c16764
Allocating Worker #6:                c968512f-2d98-4c4d-b473-3bff9da7c8f2
Allocating Worker #7:                97b602a4-638a-4ddd-a801-94bf74b77518
Allocating Worker #8:                1325f602-90c2-4a85-94fd-b841b030efcc
  JOB 1/16 STARTED:                  866a101b-ca22-41e4-bf5e-9c982a3c712a with Worker: 5a8e888f-e7ec-413b-a95e-66c9f4b07ee0
  Allocating Job #13:                348e4b45-7fbe-4d67-8206-890b12a2e071
  JOB 4/16 STARTED:                  73c99f65-f1f7-4bce-9897-748e71bea4f0 with Worker: ba9a7d7b-49da-4be9-9b85-50ac6b1a2fad
  JOB 3/16 STARTED:                  d3160123-2eeb-4df6-9717-c3dfedda97e8 with Worker: 8546abec-2440-4277-95d3-44a6adcc6fb2
  Allocating Job #14:                4ef775e3-4238-49cb-b306-1f1685ee0f62
  Allocating Job #15:                04f3e0f8-14ee-4c0f-9387-27fdf2f61930
  Allocating Job #16:                36b8fda4-8736-4a17-960d-4023e8952f58
  JOB 8/16 STARTED:                  4c564808-6701-43c5-9830-9d65dad63f1a with Worker: 97b602a4-638a-4ddd-a801-94bf74b77518
  JOB 6/16 STARTED:                  14b35713-a785-484d-b3aa-a98b3e288da1 with Worker: c968512f-2d98-4c4d-b473-3bff9da7c8f2
  JOB 2/16 STARTED:                  4509977e-5fb9-4167-801d-080e492ff3d1 with Worker: 995decd7-e2c9-4cf5-9571-71fe30761777
  JOB 7/16 STARTED:                  02b5d815-abaa-478a-9d19-8c49e9b246ec with Worker: 1325f602-90c2-4a85-94fd-b841b030efcc
  JOB 16/16 STARTED:                 36b8fda4-8736-4a17-960d-4023e8952f58 with Worker: c968512f-2d98-4c4d-b473-3bff9da7c8f2
  JOB 9/16 STARTED:                  8d3255c8-1051-4c97-b9d9-caa03772151f with Worker: 1325f602-90c2-4a85-94fd-b841b030efcc
    -> JOB 7/16 COMPLETED:           02b5d815-abaa-478a-9d19-8c49e9b246ec with Worker: 1325f602-90c2-4a85-94fd-b841b030efcc (Ran EmptySleepJob in 1.00196453 Seconds)
    -> JOB 1/16 COMPLETED:           866a101b-ca22-41e4-bf5e-9c982a3c712a with Worker: 5a8e888f-e7ec-413b-a95e-66c9f4b07ee0 (Ran EmptySleepJob in 1.002029206 Seconds)
    -> JOB 5/16 COMPLETED:           1b8f827a-2a9b-49ac-9eab-6948b4059f29 with Worker: 94229bc1-3036-4a15-a5da-287fc3c16764 (Ran EmptySleepJob in 1.002131037 Seconds)
    -> JOB 2/16 COMPLETED:           4509977e-5fb9-4167-801d-080e492ff3d1 with Worker: 995decd7-e2c9-4cf5-9571-71fe30761777 (Ran EmptySleepJob in 1.002018658 Seconds)
    -> JOB 4/16 COMPLETED:           73c99f65-f1f7-4bce-9897-748e71bea4f0 with Worker: ba9a7d7b-49da-4be9-9b85-50ac6b1a2fad (Ran EmptySleepJob in 1.002054231 Seconds)
    -> JOB 3/16 COMPLETED:           d3160123-2eeb-4df6-9717-c3dfedda97e8 with Worker: 8546abec-2440-4277-95d3-44a6adcc6fb2 (Ran EmptySleepJob in 1.002048102 Seconds)
    -> JOB 8/16 COMPLETED:           4c564808-6701-43c5-9830-9d65dad63f1a with Worker: 97b602a4-638a-4ddd-a801-94bf74b77518 (Ran EmptySleepJob in 1.002046132 Seconds)
    -> JOB 6/16 COMPLETED:           14b35713-a785-484d-b3aa-a98b3e288da1 with Worker: c968512f-2d98-4c4d-b473-3bff9da7c8f2 (Ran EmptySleepJob in 1.002046722 Seconds)
  JOB 12/16 STARTED:                 11c0bbba-749b-4c0d-a8c3-b6a65aa9bed4 with Worker: 995decd7-e2c9-4cf5-9571-71fe30761777
  JOB 13/16 STARTED:                 348e4b45-7fbe-4d67-8206-890b12a2e071 with Worker: ba9a7d7b-49da-4be9-9b85-50ac6b1a2fad
  JOB 10/16 STARTED:                 0e4bf9fd-3b3d-4d5e-858a-4de1f1b543de with Worker: 5a8e888f-e7ec-413b-a95e-66c9f4b07ee0
  JOB 15/16 STARTED:                 04f3e0f8-14ee-4c0f-9387-27fdf2f61930 with Worker: 97b602a4-638a-4ddd-a801-94bf74b77518
  JOB 11/16 STARTED:                 4d73dfa2-1f3d-488d-9a6c-d89fa2002d60 with Worker: 94229bc1-3036-4a15-a5da-287fc3c16764
  JOB 14/16 STARTED:                 4ef775e3-4238-49cb-b306-1f1685ee0f62 with Worker: 8546abec-2440-4277-95d3-44a6adcc6fb2
    -> JOB 14/16 COMPLETED:          4ef775e3-4238-49cb-b306-1f1685ee0f62 with Worker: 8546abec-2440-4277-95d3-44a6adcc6fb2 (Ran EmptySleepJob in 1.00486978 Seconds)
    -> JOB 13/16 COMPLETED:          348e4b45-7fbe-4d67-8206-890b12a2e071 with Worker: ba9a7d7b-49da-4be9-9b85-50ac6b1a2fad (Ran EmptySleepJob in 1.004910002 Seconds)
    -> JOB 16/16 COMPLETED:          36b8fda4-8736-4a17-960d-4023e8952f58 with Worker: c968512f-2d98-4c4d-b473-3bff9da7c8f2 (Ran EmptySleepJob in 1.00507389 Seconds)
    -> JOB 9/16 COMPLETED:           8d3255c8-1051-4c97-b9d9-caa03772151f with Worker: 1325f602-90c2-4a85-94fd-b841b030efcc (Ran EmptySleepJob in 1.005068561 Seconds)
    -> JOB 12/16 COMPLETED:          11c0bbba-749b-4c0d-a8c3-b6a65aa9bed4 with Worker: 995decd7-e2c9-4cf5-9571-71fe30761777 (Ran EmptySleepJob in 1.004927677 Seconds)
    -> JOB 15/16 COMPLETED:          04f3e0f8-14ee-4c0f-9387-27fdf2f61930 with Worker: 97b602a4-638a-4ddd-a801-94bf74b77518 (Ran EmptySleepJob in 1.004904105 Seconds)
    -> JOB 10/16 COMPLETED:          0e4bf9fd-3b3d-4d5e-858a-4de1f1b543de with Worker: 5a8e888f-e7ec-413b-a95e-66c9f4b07ee0 (Ran EmptySleepJob in 1.004913756 Seconds)
    -> JOB 11/16 COMPLETED:          4d73dfa2-1f3d-488d-9a6c-d89fa2002d60 with Worker: 94229bc1-3036-4a15-a5da-287fc3c16764 (Ran EmptySleepJob in 1.004899232 Seconds)

-----------------------------------------------------------------------
Total time taken:                    2.007584 Seconds




Summary Results: EmptySleepJob
+---------+------+--------------+-----------------+--------+
| WORKERS | JOBS | AVG JOB TIME | TOTAL PROC TIME |  +/-   |
+---------+------+--------------+-----------------+--------+
|       1 |   16 | 1.002449s    | 16.040965s      | (1x)*  |
|       2 |   16 | 1.011412s    | 8.092529s       | +1.98x |
|       3 |   16 | 1.005287s    | 6.039008s       | +2.66x |
|       4 |   16 | 1.004454s    | 4.018751s       | +3.99x |
|       5 |   16 | 1.014734s    | 4.049062s       | +3.96x |
|       6 |   16 | 1.028390s    | 3.078182s       | +5.21x |
|       7 |   16 | 1.020040s    | 3.049111s       | +5.26x |
|       8 |   16 | 1.003494s    | 2.007584s       | +7.99x |
+---------+------+--------------+-----------------+--------+

* Baseline: All subsequent +/- tests are compared to this.

```

Benchmark comparison: `make golang-benchmark`

```
Clearing test cache...
...Complete.

Running benchmarks...

  BENCHMARK SETUP:
  #=> Job: EmptySleepJob
  #=> Starting Worker Count:       1
  #=> Max Worker Count:            8
  #=> Job Count:                   16
    #=> Benchmark Elapsed:         46.104529262s

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkRun/Worker_Count:_1-8                 1        46104598869 ns/op         235712 B/op       4203 allocs/op

  BENCHMARK SETUP:
  #=> Job: EmptySleepJob
  #=> Starting Worker Count:       2
  #=> Max Worker Count:            8
  #=> Job Count:                   16
    #=> Benchmark Elapsed:         30.039865569s

BenchmarkRun/Worker_Count:_2-8                 1        30039887054 ns/op         200440 B/op       3576 allocs/op

  BENCHMARK SETUP:
  #=> Job: EmptySleepJob
  #=> Starting Worker Count:       3
  #=> Max Worker Count:            8
  #=> Job Count:                   16
    #=> Benchmark Elapsed:         22.033414346s

BenchmarkRun/Worker_Count:_3-8                 1        22033433374 ns/op         217856 B/op       3075 allocs/op

  BENCHMARK SETUP:
  #=> Job: EmptySleepJob
  #=> Starting Worker Count:       4
  #=> Max Worker Count:            8
  #=> Job Count:                   16
    #=> Benchmark Elapsed:         16.016102428s

BenchmarkRun/Worker_Count:_4-8                 1        16016143164 ns/op         118288 B/op       2587 allocs/op

  BENCHMARK SETUP:
  #=> Job: EmptySleepJob
  #=> Starting Worker Count:       5
  #=> Max Worker Count:            8
  #=> Job Count:                   16
    #=> Benchmark Elapsed:         12.015069931s

BenchmarkRun/Worker_Count:_5-8                 1        12015090970 ns/op          94864 B/op       2089 allocs/op

  BENCHMARK SETUP:
  #=> Job: EmptySleepJob
  #=> Starting Worker Count:       6
  #=> Max Worker Count:            8
  #=> Job Count:                   16
    #=> Benchmark Elapsed:          8.008436433s

BenchmarkRun/Worker_Count:_6-8                 1        8008455986 ns/op           75736 B/op       1573 allocs/op

  BENCHMARK SETUP:
  #=> Job: EmptySleepJob
  #=> Starting Worker Count:       7
  #=> Max Worker Count:            8
  #=> Job Count:                   16
    #=> Benchmark Elapsed:         5.002237243s

BenchmarkRun/Worker_Count:_7-8                 1        5002258990 ns/op           50944 B/op       1083 allocs/op

  BENCHMARK SETUP:
  #=> Job: EmptySleepJob
  #=> Starting Worker Count:       8
  #=> Max Worker Count:            8
  #=> Job Count:                   16
    #=> Benchmark Elapsed:         2.007394911s

BenchmarkRun/Worker_Count:_8-8                 1        2007433030 ns/op           25672 B/op        556 allocs/op
PASS
ok      command-line-arguments  141.239s
...Complete.

```

**Overloading EmptySleepJob() Results:**

In the above example, when limiting the maximum number or workers to `runtime.NumCPU()`, the application processing speed increases as expected. But what if we pass this limit?

Below are the results for running 64 workers, or `runtime.NumCPU() * 8`:

```
Summary Results: PiJob
+---------+------+--------------+-----------------+--------+
| WORKERS | JOBS | AVG JOB TIME | TOTAL PROC TIME |  +/-   |
+---------+------+--------------+-----------------+--------+
|       1 |   64 | 0.988064s    | 63.359632s      | (1x)*  |
|       2 |   64 | 1.080282s    | 34.639627s      | +1.83x |
|       3 |   64 | 1.259059s    | 27.533909s      | +2.3x  |
|       4 |   64 | 1.438016s    | 23.145630s      | +2.74x |
|       5 |   64 | 1.537516s    | 20.044856s      | +3.16x |
|       6 |   64 | 1.880944s    | 20.562739s      | +3.08x |
|       7 |   64 | 2.121186s    | 20.138135s      | +3.15x |
|       8 |   64 | 2.540276s    | 20.421168s      | +3.1x  |
|       9 |   64 | 2.659540s    | 19.800304s      | +3.2x  |
|      10 |   64 | 2.890543s    | 19.415050s      | +3.26x |
|      11 |   64 | 3.307533s    | 19.838304s      | +3.19x |
|      12 |   64 | 3.587793s    | 20.066332s      | +3.16x |
|      13 |   64 | 4.049644s    | 20.471609s      | +3.1x  |
|      14 |   64 | 4.312730s    | 20.940306s      | +3.03x |
|      15 |   64 | 4.474387s    | 20.171849s      | +3.14x |
|      16 |   64 | 5.073989s    | 20.682957s      | +3.06x |
|      17 |   64 | 5.282467s    | 20.934039s      | +3.03x |
|      18 |   64 | 5.468803s    | 20.829361s      | +3.04x |
|      19 |   64 | 5.673584s    | 20.477446s      | +3.09x |
|      20 |   64 | 5.968736s    | 20.362165s      | +3.11x |
|      21 |   64 | 6.371592s    | 20.383568s      | +3.11x |
|      22 |   64 | 6.541627s    | 19.899953s      | +3.18x |
|      23 |   64 | 6.518885s    | 19.513377s      | +3.25x |
|      24 |   64 | 6.797447s    | 19.832938s      | +3.19x |
|      25 |   64 | 6.876581s    | 19.447455s      | +3.26x |
|      26 |   64 | 6.877876s    | 18.923182s      | +3.35x |
|      27 |   64 | 7.118303s    | 18.781495s      | +3.37x |
|      28 |   64 | 7.227538s    | 18.356835s      | +3.45x |
|      29 |   64 | 7.767756s    | 18.708181s      | +3.39x |
|      30 |   64 | 7.875914s    | 18.318063s      | +3.46x |
|      31 |   64 | 8.218964s    | 18.289108s      | +3.46x |
|      32 |   64 | 8.425439s    | 17.590077s      | +3.6x  |
+---------+------+--------------+-----------------+--------+

* Baseline: All subsequent +/- tests are compared to this.
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

make golang-benchmark
make golang-build
make golang-godoc
make golang-lint
make golang-pprof
make golang-run
make golang-test


For details on these commands, see the bash scripts in the 'scripts/' directory.
```

### make golang-lint

`make golang-lint` To lint the application.  **[golangci-lint](https://golangci-lint.run/usage/install/) must be installed**

### make golang-test

`make golang-test` To test the application using `gotestsum` **[gotestsum](https://github.com/gotestyourself/gotestsum#install) must be installed.**

#### Code coverage:

`go test ./... -coverprofile ./coverage/coverprofile.out`
`go tool cover -html=./coverage/coverprofile.out -o ./coverage/cover.html`

## Benchmark Tests

Results of running the application (`make golang-run`): 8 Iterations, starting with 1 Worker, ending with 8 Workers, all processing 64 of the same job (PiJob)

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



### Results of the benchmarks (`make golang-benchmark`):

The following is a lot of commands and benchmark output. The idea here is that the output roughly matches the timing of running the application. As there is a small variability in how long it takes the jobs to execute--usually much less than a second--the numbers won't be exact. The following is simply confirmation that we're getting the results that we expect.

The following tests were run on:

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -jobName=PiJob -startingWorkerCount=1 -maxWorkerCount=1 -jobCount=64` #=>

```
BENCHMARK SETUP:
  #=> Job:   PiJob
  #=> Starting Worker Count:       1
  #=> Max Worker Count:            1
  #=> Job Count:                   64

BenchmarkRun/Worker_Count:_1-8                 1        64292688948 ns/op       35976337824 B/op        23511673 allocs/op
PASS
ok      command-line-arguments  64.467s
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -jobName=PiJob -startingWorkerCount=2 -maxWorkerCount=2 -jobCount=64` #=>

```
BENCHMARK SETUP:
  #=> Job:   PiJob
  #=> Starting Worker Count:       2
  #=> Max Worker Count:            2
  #=> Job Count:                   64

BenchmarkRun/Worker_Count:_2-8                 1        38487971462 ns/op       36086760640 B/op        23605364 allocs/op
PASS
ok      command-line-arguments  38.683s
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -jobName=PiJob -startingWorkerCount=3 -maxWorkerCount=3 -jobCount=64` #=>

```
BENCHMARK SETUP:
  #=> Job:   PiJob
  #=> Starting Worker Count:       3
  #=> Max Worker Count:            3
  #=> Job Count:                   64

BenchmarkRun/Worker_Count:_3-8                 1        31649478509 ns/op       36098221112 B/op        23631627 allocs/op
PASS
ok      command-line-arguments  31.873s
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -jobName=PiJob -startingWorkerCount=4 -maxWorkerCount=4 -jobCount=64` #=>

```
BENCHMARK SETUP:
  #=> Job:   PiJob
  #=> Starting Worker Count:       4
  #=> Max Worker Count:            4
  #=> Job Count:                   64

BenchmarkRun/Worker_Count:_4-8                 1        25412844723 ns/op       36074032056 B/op        23617726 allocs/op
PASS
ok      command-line-arguments  25.625s
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -jobName=PiJob -startingWorkerCount=5 -maxWorkerCount=5 -jobCount=64` #=>

```
BENCHMARK SETUP:
  #=> Job:   PiJob
  #=> Starting Worker Count:       5
  #=> Max Worker Count:            5
  #=> Job Count:                   64

BenchmarkRun/Worker_Count:_5-8                 1        20767157101 ns/op       36104025472 B/op        23640474 allocs/op
PASS
ok      command-line-arguments  20.914s
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -jobName=PiJob -startingWorkerCount=6 -maxWorkerCount=6 -jobCount=64` #=>

```
BENCHMARK SETUP:
  #=> Job:   PiJob
  #=> Starting Worker Count:       6
  #=> Max Worker Count:            6
  #=> Job Count:                   64

BenchmarkRun/Worker_Count:_6-8                 1        18937268938 ns/op       36135256720 B/op        23655645 allocs/op
PASS
ok      command-line-arguments  19.092s
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -jobName=PiJob -startingWorkerCount=7 -maxWorkerCount=7 -jobCount=64` #=>

```
BENCHMARK SETUP:
  #=> Job:   PiJob
  #=> Starting Worker Count:       7
  #=> Max Worker Count:            7
  #=> Job Count:                   64

BenchmarkRun/Worker_Count:_7-8                 1        18182033241 ns/op       36135543352 B/op        23654627 allocs/op
PASS
ok      command-line-arguments  18.386s
```

`go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -jobName=PiJob -startingWorkerCount=8 -maxWorkerCount=8 -jobCount=64` #=>

```
BENCHMARK SETUP:
  #=> Job:   PiJob
  #=> Starting Worker Count:       8
  #=> Max Worker Count:            8
  #=> Job Count:                   64

BenchmarkRun/Worker_Count:_8-8                 1        17796937403 ns/op       36178650928 B/op        23679568 allocs/op
PASS
ok      command-line-arguments  18.024s
```

---

# Comparing CPU profiles via pprof:

## Web

go tool pprof -http='192.168.0.99:8081' ./pprof/_memprofile-08-workers.out


## Commandline

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

