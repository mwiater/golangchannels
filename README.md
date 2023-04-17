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

Below are the results for running 32 workers, or `runtime.NumCPU() * 8`:

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
```

The defaults are listed below when no values are set. **To override any/all, just set the env vars below.**

**DEBUG:** More verbose console output, default is `false`
**STARTINGWORKERCOUNT:** By default this will be set to `1`
**MAXWORKERCOUNT:** By default this will be set to `runtime.NumCPU()`
**TOTALJOBCOUNT:** By default this will be set to `runtime.NumCPU()*2`


## Running the application

`go run .`

Or: `make golang-run` Note: This builds and executes the binary, not quite equivalent to: `go run .`

## Building the application binary

`make golang-build` Outputs binary to `./bin/` directory.

## Running the application binary

`./bin/golangchannels`

Or: `make golang-run` Note: This builds and executes the binary, not quite equivalent to: `go run .`

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


```
Clearing test cache...
...Complete.

Running tests...
EMPTY .
EMPTY dispatcher
PASS common.TestConsoleColumnWidth (0.00s)
PASS common.TestSplitStringLines (0.00s)
PASS common
PASS config.TestAppConfig (0.00s)
PASS config
PASS jobs/emptySleepJob.TestEmptySleepJob (1.00s)
PASS jobs/emptySleepJob
PASS jobs/piJob.TestPiJob (15.10s)
PASS jobs/piJob
PASS jobs/ioJob.TestIoJob (16.03s)
PASS jobs/ioJob
EMPTY structs
EMPTY workers

DONE 6 tests in 22.292s
...Complete.
```

#### Code coverage:

`go test ./... -coverprofile ./coverage/coverprofile.out`
`go tool cover -html=./coverage/coverprofile.out -o ./coverage/cover.html`
