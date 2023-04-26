# Golang-Channels

## Reference Article

**See my accompanying article on Medium here:** https://medium.com/@matt.wiater/golang-channels-goroutines-and-optimal-concurrency-demystifying-through-examples-a43ba6aee74f

## To Do

 - [ ] Benchmark docs and examples
 - [ ] Better tests
 - [ ] More documentation via godocs after tests are updated

## Installation

```
git clone git@github.com:mwiater/golangchannels.git
cd golangchannels
go get
```

Optional:

* For linting with `make golang-lint`: **[golangci-lint](https://golangci-lint.run/usage/install/) must be installed**
* For testing with `make golang-test`: **[gotestsum](https://github.com/gotestyourself/gotestsum#install) must be installed.**

## Application

### Convenience Commands

**NOTE:** _This repository was created to run on a linux x86 machine. There are example `make` commands and scripts to cross-compile to other architectures, but these convenience methods have only been tested from a linux x86 source (though the corss-compiled binaries have been successfully tested on the appropriate source architectures). If you are on another source architecture, native `go` commands should work fine, but the build scripts included here may not. See the `./scripts` directory to see what the bash scripts are doing--you should be able to infer the native `go` commands for your system, if you have issues executing any of the convenience scripts._

Type `make` for a list of commands (comments added for clarity):

```
Targets in this Makefile:

make golang-benchmark       # Run benchmark test: ./dispatcher/dispatcher_test.go
make golang-build-arm64     # Build the application binary to: ./bin/golangchannels-arm64
make golang-build-linux64   # Build the application binary to: ./bin/golangchannels
make golang-build-windows64 # Build the application binary to: ./bin/bin/golangchannels.exe
make golang-godoc           # Start the documentation server
make golang-lint            # Lint the application
make golang-pprof           # Generate pprof profiles in ./pprof
make golang-run             # Compile and run the binary
make golang-test            # Run the tests

For details on these commands, see the bash scripts in the 'scripts/' directory.
```

## Run the application

Type: `make golang-run`

## `.env`: Override default settings

You can use the `.env` file to override the default values. Blank values will become the defaults below (comments added for clarity):

```
DEBUG=true              // Verbose console output, default: false
JOBNAME=                // Job: EmptySleepJob (default), PiJob, or IoJob
STARTINGWORKERCOUNT=    // Workers to start the test with, default: 1
MAXWORKERCOUNT=         // Workers to ramp up to, default: runtime.NumCPU()
TOTALJOBCOUNT=          // Jobs to run, default: runtime.NumCPU() * 2
```

With the default values above, on an 8 CPU core system, the application:

* Runs a total of 8 passes:
  * The first pass with STARTINGWORKERCOUNT workers,  in this example: **1**
  * Add one worker for the next pass
  * Repeat until MAXWORKERCOUNT is reached, in this example: **8**
  * Total jobs to process: **16**
* Display a list of Summary Results: 

Type: `make golang-run`

**Example 01: `EmptySleepJob` Results:**

With the `.env` file shown above, the results are:

```
(tests 1-7 ommitted from this example)
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

Allocating Worker #1:                5602333e-76eb
Allocating Worker #2:                54e6a5f7-2738
Allocating Worker #3:                6cbb7138-3d46
Allocating Worker #4:                06ff3344-6433
Allocating Worker #5:                90813e84-d391
Allocating Worker #6:                8d2a9ce9-6f64
Allocating Worker #7:                6e12a298-b1cc
Allocating Worker #8:                0a1c3778-092a
  Allocating Job #1:                 e827c11d-49c7
  Allocating Job #2:                 cb0ac4b7-7c95
  Allocating Job #3:                 5f4528c1-f523
  Allocating Job #4:                 b7b22f30-b9cb
  Allocating Job #5:                 a9706aec-7fc7
  Allocating Job #6:                 9b2c68c1-647f
  Allocating Job #7:                 f89c0bb5-9915
  Allocating Job #8:                 03f4c06b-484c
  Allocating Job #9:                 f19123a8-f5e4
  Allocating Job #10:                0fe68097-d6ad
  Allocating Job #11:                46b954cc-9485
  Allocating Job #12:                3b93deb7-b296
  Allocating Job #13:                75f7cd25-bced
  Allocating Job #14:                905956e4-acd9
  Allocating Job #15:                3e852446-2e27
  JOB 4/16 STARTED:                  b7b22f30-b9cb with Worker: 06ff3344-6433
  JOB 5/16 STARTED:                  a9706aec-7fc7 with Worker: 90813e84-d391
  JOB 3/16 STARTED:                  5f4528c1-f523 with Worker: 6cbb7138-3d46
  JOB 8/16 STARTED:                  03f4c06b-484c with Worker: 6e12a298-b1cc
  JOB 7/16 STARTED:                  f89c0bb5-9915 with Worker: 0a1c3778-092a
  JOB 1/16 STARTED:                  e827c11d-49c7 with Worker: 5602333e-76eb
  JOB 6/16 STARTED:                  9b2c68c1-647f with Worker: 8d2a9ce9-6f64
  JOB 2/16 STARTED:                  cb0ac4b7-7c95 with Worker: 54e6a5f7-2738
  Allocating Job #16:                37fe0fe1-84dd
  JOB 9/16 STARTED:                  f19123a8-f5e4 with Worker: 5602333e-76eb
  JOB 11/16 STARTED:                 46b954cc-9485 with Worker: 6e12a298-b1cc
  JOB 10/16 STARTED:                 0fe68097-d6ad with Worker: 8d2a9ce9-6f64
  JOB 12/16 STARTED:                 3b93deb7-b296 with Worker: 54e6a5f7-2738
    -> JOB 1/16 COMPLETED:           ✓ Ran EmptySleepJob in: 1.000 Seconds / Used 0.006MB of Memory
    -> JOB 6/16 COMPLETED:           ✓ Ran EmptySleepJob in: 1.000 Seconds / Used 0.006MB of Memory
    -> JOB 8/16 COMPLETED:           ✓ Ran EmptySleepJob in: 1.000 Seconds / Used 0.006MB of Memory
    -> JOB 2/16 COMPLETED:           ✓ Ran EmptySleepJob in: 1.000 Seconds / Used 0.006MB of Memory
  JOB 13/16 STARTED:                 75f7cd25-bced with Worker: 0a1c3778-092a
    -> JOB 7/16 COMPLETED:           ✓ Ran EmptySleepJob in: 1.001 Seconds / Used 0.006MB of Memory
  JOB 14/16 STARTED:                 905956e4-acd9 with Worker: 6cbb7138-3d46
    -> JOB 3/16 COMPLETED:           ✓ Ran EmptySleepJob in: 1.001 Seconds / Used 0.006MB of Memory
  JOB 15/16 STARTED:                 3e852446-2e27 with Worker: 90813e84-d391
    -> JOB 4/16 COMPLETED:           ✓ Ran EmptySleepJob in: 1.001 Seconds / Used 0.006MB of Memory
    -> JOB 5/16 COMPLETED:           ✓ Ran EmptySleepJob in: 1.001 Seconds / Used 0.006MB of Memory
  JOB 16/16 STARTED:                 37fe0fe1-84dd with Worker: 06ff3344-6433
    -> JOB 12/16 COMPLETED:          ✓ Ran EmptySleepJob in: 1.002 Seconds / Used 0.006MB of Memory
    -> JOB 11/16 COMPLETED:          ✓ Ran EmptySleepJob in: 1.002 Seconds / Used 0.006MB of Memory
    -> JOB 13/16 COMPLETED:          ✓ Ran EmptySleepJob in: 1.002 Seconds / Used 0.006MB of Memory
    -> JOB 9/16 COMPLETED:           ✓ Ran EmptySleepJob in: 1.002 Seconds / Used 0.006MB of Memory
    -> JOB 16/16 COMPLETED:          ✓ Ran EmptySleepJob in: 1.002 Seconds / Used 0.006MB of Memory
    -> JOB 10/16 COMPLETED:          ✓ Ran EmptySleepJob in: 1.002 Seconds / Used 0.006MB of Memory
    -> JOB 15/16 COMPLETED:          ✓ Ran EmptySleepJob in: 1.002 Seconds / Used 0.006MB of Memory
    -> JOB 14/16 COMPLETED:          ✓ Ran EmptySleepJob in: 1.002 Seconds / Used 0.006MB of Memory

-----------------------------------------------------------------------
Total time taken:                    2.004093 Seconds


Summary Results: EmptySleepJob
+---------+------+--------------+-------------------+-------------+--------+
| WORKERS | JOBS | AVG JOB TIME | TOTAL WORKER TIME | AVG MEM USE |  +/-   |
+---------+------+--------------+-------------------+-------------+--------+
|       1 |   16 | 1.00s        | 16.03s            | 0.003Mb     | (1x)*  |
|       2 |   16 | 1.00s        | 8.02s             | 0.004Mb     | +2x    |
|       3 |   16 | 1.00s        | 6.01s             | 0.004Mb     | +2.67x |
|       4 |   16 | 1.00s        | 4.01s             | 0.004Mb     | +4x    |
|       5 |   16 | 1.00s        | 4.01s             | 0.005Mb     | +4x    |
|       6 |   16 | 1.00s        | 3.01s             | 0.005Mb     | +5.33x |
|       7 |   16 | 1.00s        | 3.01s             | 0.005Mb     | +5.33x |
|       8 |   16 | 1.00s        | 2.00s             | 0.006Mb     | +8x    |
+---------+------+--------------+-------------------+-------------+--------+

* Baseline: All subsequent +/- tests are compared to this.

```

Expectations:
* ✅ As each job is just a `time.Sleep(1 * time.Second)` command, each job should take about **1 second to execute.**
* ✅ As expected, **1 worker completing 16 of these jobs takes about 16 seconds** and **8 workers runs approximately 8 times as fast: 2 seconds.**


**Example 02: `PiJob` Results:**

To run a similar test, but with a different job, just change the `JOBNAME` (options: `EmptySleepJob`, `PiJob`, or `IoJob`)

```
DEBUG=true
JOBNAME=PiJob
STARTINGWORKERCOUNT=
MAXWORKERCOUNT=
TOTALJOBCOUNT=
```

The `PiJob` represents a CPU-bound test, so the overall concurrency performance will be much lower than the `EmptySleepJob`:


```
(tests 1-7 ommitted from this example)
...
--------------------------------------
|  Spawning workers for test 8 of 8  |
--------------------------------------

Workers: 8                           Job Name: PiJob
----------------------------------------------------
  Workers In Use:                    8
  Workers Available:                 8
  Workers Idle:                      0
  Number of Jobs:                    16
----------------------------------------------------

Allocating Worker #1:                7fe7b934-58be
Allocating Worker #2:                5110d081-61cf
Allocating Worker #3:                eea33792-edeb
Allocating Worker #4:                0350c02e-be14
Allocating Worker #5:                b97cfa56-57be
Allocating Worker #6:                9beb8f92-f783
  Allocating Job #1:                 ee278e6e-2ea9
  Allocating Job #2:                 116065ce-b313
  Allocating Job #3:                 ec344b16-83ce
  Allocating Job #4:                 9e42773f-40c9
  JOB 2/16 STARTED:                  116065ce-b313 with Worker: 7fe7b934-58be
  JOB 4/16 STARTED:                  9e42773f-40c9 with Worker: eea33792-edeb
  JOB 3/16 STARTED:                  ec344b16-83ce with Worker: 5110d081-61cf
Allocating Worker #7:                2841b0aa-4751
Allocating Worker #8:                41e06f3b-fe55
  Allocating Job #5:                 f9e6ccb1-43b7
  JOB 5/16 STARTED:                  f9e6ccb1-43b7 with Worker: 0350c02e-be14
  JOB 1/16 STARTED:                  ee278e6e-2ea9 with Worker: 9beb8f92-f783
  Allocating Job #6:                 437387c3-3be9
  Allocating Job #7:                 a3179e53-0a4b
  Allocating Job #8:                 d43e1761-1509
  Allocating Job #9:                 eb33c4a2-f160
  Allocating Job #10:                64f260c6-dbe0
  Allocating Job #11:                575bd2c6-91f3
  Allocating Job #12:                af5ef86a-25d3
  Allocating Job #13:                102f8717-9da3
  Allocating Job #14:                abb6fe18-f7db
  Allocating Job #15:                14d0af91-1b65
  Allocating Job #16:                7d84a170-a8be
  JOB 6/16 STARTED:                  437387c3-3be9 with Worker: b97cfa56-57be
  JOB 8/16 STARTED:                  d43e1761-1509 with Worker: 2841b0aa-4751
  JOB 7/16 STARTED:                  a3179e53-0a4b with Worker: 41e06f3b-fe55
  JOB 9/16 STARTED:                  eb33c4a2-f160 with Worker: eea33792-edeb
    -> JOB 4/16 COMPLETED:           ✓ Ran PiJob in: 2.438 Seconds / Used 0.085MB of Memory
  JOB 10/16 STARTED:                 64f260c6-dbe0 with Worker: 41e06f3b-fe55
    -> JOB 7/16 COMPLETED:           ✓ Ran PiJob in: 2.547 Seconds / Used 0.015MB of Memory
  JOB 11/16 STARTED:                 575bd2c6-91f3 with Worker: b97cfa56-57be
    -> JOB 6/16 COMPLETED:           ✓ Ran PiJob in: 2.561 Seconds / Used 0.015MB of Memory
  JOB 12/16 STARTED:                 af5ef86a-25d3 with Worker: 2841b0aa-4751
    -> JOB 8/16 COMPLETED:           ✓ Ran PiJob in: 2.612 Seconds / Used 0.020MB of Memory
  JOB 13/16 STARTED:                 102f8717-9da3 with Worker: 0350c02e-be14
    -> JOB 5/16 COMPLETED:           ✓ Ran PiJob in: 2.633 Seconds / Used 0.015MB of Memory
  JOB 14/16 STARTED:                 abb6fe18-f7db with Worker: 7fe7b934-58be
    -> JOB 2/16 COMPLETED:           ✓ Ran PiJob in: 2.645 Seconds / Used 0.015MB of Memory
  JOB 15/16 STARTED:                 14d0af91-1b65 with Worker: 9beb8f92-f783
    -> JOB 1/16 COMPLETED:           ✓ Ran PiJob in: 2.707 Seconds / Used 0.014MB of Memory
  JOB 16/16 STARTED:                 7d84a170-a8be with Worker: 5110d081-61cf
    -> JOB 3/16 COMPLETED:           ✓ Ran PiJob in: 2.727 Seconds / Used 0.015MB of Memory
    -> JOB 11/16 COMPLETED:          ✓ Ran PiJob in: 2.367 Seconds / Used 0.014MB of Memory
    -> JOB 12/16 COMPLETED:          ✓ Ran PiJob in: 2.331 Seconds / Used 0.014MB of Memory
    -> JOB 9/16 COMPLETED:           ✓ Ran PiJob in: 2.532 Seconds / Used 0.014MB of Memory
    -> JOB 10/16 COMPLETED:          ✓ Ran PiJob in: 2.462 Seconds / Used 0.014MB of Memory
    -> JOB 13/16 COMPLETED:          ✓ Ran PiJob in: 2.389 Seconds / Used 0.014MB of Memory
    -> JOB 15/16 COMPLETED:          ✓ Ran PiJob in: 2.361 Seconds / Used 0.015MB of Memory
    -> JOB 14/16 COMPLETED:          ✓ Ran PiJob in: 2.424 Seconds / Used 0.015MB of Memory
    -> JOB 16/16 COMPLETED:          ✓ Ran PiJob in: 2.358 Seconds / Used 0.014MB of Memory

-----------------------------------------------------------------------
Total time taken:                    5.092568 Seconds




Summary Results: PiJob
+---------+------+--------------+-------------------+-------------+--------+
| WORKERS | JOBS | AVG JOB TIME | TOTAL WORKER TIME | AVG MEM USE |  +/-   |
+---------+------+--------------+-------------------+-------------+--------+
|       1 |   16 | 1.01s        | 16.19s            | 0.009Mb     | (1x)*  |
|       2 |   16 | 1.08s        | 8.67s             | 0.010Mb     | +1.87x |
|       3 |   16 | 1.23s        | 7.21s             | 0.011Mb     | +2.25x |
|       4 |   16 | 1.54s        | 6.20s             | 0.014Mb     | +2.61x |
|       5 |   16 | 1.64s        | 6.08s             | 0.016Mb     | +2.66x |
|       6 |   16 | 1.83s        | 5.48s             | 0.015Mb     | +2.96x |
|       7 |   16 | 2.12s        | 5.67s             | 0.020Mb     | +2.85x |
|       8 |   16 | 2.51s        | 5.09s             | 0.019Mb     | +3.18x |
+---------+------+--------------+-------------------+-------------+--------+

* Baseline: All subsequent +/- tests are compared to this.

```

## Other convenience commands:

List available commands: `make`

```
Targets in this Makefile:

make golang-benchmark
make golang-build-arm64
make golang-build-linux64
make golang-build-windows64
make golang-godoc
make golang-lint
make golang-pprof
make golang-run
make golang-test

For details on these commands, see the bash scripts in the 'scripts/' directory.
```

### make golang-benchmark

`make golang-benchmark` To benchmark the application via: `./dispatcher/dispatcher_test.go`

### make golang-build

`make golang-build` Build the application binary to: `./bin/golangchannels`

make golang_build_arm64
make golang_build_linux64
make golang_build_windows64

### make golang-godoc

`make golang-godoc` To start the docs server: `Starting godoc server on port: 6060...` The root directory for this app is: `/pkg/github.com/mattwiater/golangchannels/`

### make golang-lint

`make golang-lint` **[golangci-lint](https://golangci-lint.run/usage/install/) must be installed** to lint the application.

### make golang-pprof

`make golang-pprof` **[pprof](https://github.com/google/pprof#building-pprof) must be installed** to generate `cpuprofile`, `memprofile`, `blockprofile`, and, `mutexprofile` profiles. This command will generate 1-8 profiles of each type, which can then be analyzed via the command line (e.g.: `go tool pprof ./pprof/cpuprofile-08-workers.out`) or visually via a browser (e.g.: `go tool pprof -http='{your-ip-address}:8081' ./pprof/cpuprofile-08-workers.out`).

Example Output (Just showing the first of 8 benchmarks):

```
Running pprofs for 8 CPUs...
Running pprof: 1 Workers

  BENCHMARK SETUP:
  #=> Job:   PiJob
  #=> Starting Worker Count:       1
  #=> Max Worker Count:            1
  #=> Job Count:                   64
    #=> Benchmark Elapsed:        1m4.354124664s

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkRun/Worker_Count:_1-8                 1        64354143863 ns/op       35979108096 B/op        23537220 allocs/op
PASS
ok      command-line-arguments  64.552s
    Generated Profile: ./pprof/cpuprofile-01-workers.out
    Generated Profile: ./pprof/memprofile-01-workers.out
    Generated Profile: ./pprof/blockprofile-01-workers.out
    Generated Profile: ./pprof/mutexprofile-01-workers.out

...
```

### make golang-run

`make golang-run` To run the application. **NOTE: This command compiles and executes the binary in: `./bin/golangchannels` If you just want to run the application, use: `go run .`**

### make golang-test

`make golang-test` **[gotestsum](https://github.com/gotestyourself/gotestsum#install) must be installed** to test the application.


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

## Code coverage:

`go test ./... -coverprofile ./coverage/coverprofile.out`

`go tool cover -html=./coverage/coverprofile.out -o ./coverage/cover.html`

The commands above will generate a test coverage report: `./coverage/cover.html`
