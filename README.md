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

Allocating Worker #1:                4add73e2-da2b-4637-a4fb-51543d5509b4
Allocating Worker #2:                6916721e-d3da-493e-bce8-71cd4eb3de10
Allocating Worker #3:                32332fcb-64f7-43c7-8c3f-26e18648384e
Allocating Worker #4:                d5be5f1c-2477-4d5d-8322-2ef428b0f535
Allocating Worker #5:                41f951d9-a878-4f5e-9265-9924ec2f815a
Allocating Worker #6:                81e7fd26-7e5b-45f7-811c-7d356f140593
Allocating Worker #7:                d21cdd06-f55e-4e3f-9c01-ed79fe543711
Allocating Worker #8:                da2e47c2-a7bf-428f-b029-2c4ea51422d4
  Allocating Job #1:                 8c183aa8-5b10-49aa-b1cd-9dfa0b1db194
  Allocating Job #2:                 697b8947-0273-46a3-bf42-0995b2daed30
  Allocating Job #3:                 f2df6e2a-3f85-4b27-b7cd-b4aa1f6883b8
  Allocating Job #4:                 919f3133-50d8-4152-bf59-92cc7eca50b4
  JOB 2/16 STARTED:                  697b8947-0273-46a3-bf42-0995b2daed30 with Worker: 6916721e-d3da-493e-bce8-71cd4eb3de10
  JOB 3/16 STARTED:                  f2df6e2a-3f85-4b27-b7cd-b4aa1f6883b8 with Worker: 32332fcb-64f7-43c7-8c3f-26e18648384e
  JOB 4/16 STARTED:                  919f3133-50d8-4152-bf59-92cc7eca50b4 with Worker: da2e47c2-a7bf-428f-b029-2c4ea51422d4
  JOB 1/16 STARTED:                  8c183aa8-5b10-49aa-b1cd-9dfa0b1db194 with Worker: 4add73e2-da2b-4637-a4fb-51543d5509b4
  Allocating Job #5:                 e388ad83-4ba2-454d-925d-db40988fd161
  Allocating Job #6:                 0eda2855-7f5b-4f61-82fd-161e809892ec
  Allocating Job #7:                 c8d8bd76-f463-4968-bdda-3113b06adca0
  Allocating Job #8:                 78fa697a-42e6-45b8-80b2-c4a7415b8cc4
  Allocating Job #9:                 b9220f71-dab5-4329-b1c6-5f49a117d616
  Allocating Job #10:                ee86c5d7-8ff2-4c76-8dbd-78ad66d5aa1a
  Allocating Job #11:                c2336100-de1f-4433-b82d-683f48861e64
  Allocating Job #12:                765de944-9ec6-4b50-af3a-b3decd81e967
  Allocating Job #13:                f581ddd6-aec2-423f-be13-9c6bc0fddeaa
  Allocating Job #14:                0684858e-9d93-42fd-a203-f4b73c40dfab
  Allocating Job #15:                e8740614-c7cf-41b1-b946-6077ba5b8a20
  Allocating Job #16:                a782c1e8-6569-47bf-99c0-15de8846ed30
  JOB 6/16 STARTED:                  0eda2855-7f5b-4f61-82fd-161e809892ec with Worker: d5be5f1c-2477-4d5d-8322-2ef428b0f535
  JOB 5/16 STARTED:                  e388ad83-4ba2-454d-925d-db40988fd161 with Worker: 41f951d9-a878-4f5e-9265-9924ec2f815a
  JOB 8/16 STARTED:                  78fa697a-42e6-45b8-80b2-c4a7415b8cc4 with Worker: 81e7fd26-7e5b-45f7-811c-7d356f140593
  JOB 7/16 STARTED:                  c8d8bd76-f463-4968-bdda-3113b06adca0 with Worker: d21cdd06-f55e-4e3f-9c01-ed79fe543711
  JOB 9/16 STARTED:                  b9220f71-dab5-4329-b1c6-5f49a117d616 with Worker: 4add73e2-da2b-4637-a4fb-51543d5509b4
  JOB 10/16 STARTED:                 ee86c5d7-8ff2-4c76-8dbd-78ad66d5aa1a with Worker: d21cdd06-f55e-4e3f-9c01-ed79fe543711
  JOB 13/16 STARTED:                 f581ddd6-aec2-423f-be13-9c6bc0fddeaa with Worker: da2e47c2-a7bf-428f-b029-2c4ea51422d4
    -> JOB 1/16 COMPLETED:           8c183aa8-5b10-49aa-b1cd-9dfa0b1db194 with Worker: 4add73e2-da2b-4637-a4fb-51543d5509b4 (Ran emptySleepJob in 1.000478511 Seconds)
  JOB 14/16 STARTED:                 0684858e-9d93-42fd-a203-f4b73c40dfab with Worker: 81e7fd26-7e5b-45f7-811c-7d356f140593
  JOB 11/16 STARTED:                 c2336100-de1f-4433-b82d-683f48861e64 with Worker: 6916721e-d3da-493e-bce8-71cd4eb3de10
    -> JOB 7/16 COMPLETED:           c8d8bd76-f463-4968-bdda-3113b06adca0 with Worker: d21cdd06-f55e-4e3f-9c01-ed79fe543711 (Ran emptySleepJob in 1.000410017 Seconds)
  JOB 16/16 STARTED:                 a782c1e8-6569-47bf-99c0-15de8846ed30 with Worker: d5be5f1c-2477-4d5d-8322-2ef428b0f535
  JOB 15/16 STARTED:                 e8740614-c7cf-41b1-b946-6077ba5b8a20 with Worker: 41f951d9-a878-4f5e-9265-9924ec2f815a
    -> JOB 2/16 COMPLETED:           697b8947-0273-46a3-bf42-0995b2daed30 with Worker: 6916721e-d3da-493e-bce8-71cd4eb3de10 (Ran emptySleepJob in 1.000579614 Seconds)
    -> JOB 3/16 COMPLETED:           f2df6e2a-3f85-4b27-b7cd-b4aa1f6883b8 with Worker: 32332fcb-64f7-43c7-8c3f-26e18648384e (Ran emptySleepJob in 1.000560273 Seconds)
    -> JOB 4/16 COMPLETED:           919f3133-50d8-4152-bf59-92cc7eca50b4 with Worker: da2e47c2-a7bf-428f-b029-2c4ea51422d4 (Ran emptySleepJob in 1.000556052 Seconds)
    -> JOB 8/16 COMPLETED:           78fa697a-42e6-45b8-80b2-c4a7415b8cc4 with Worker: 81e7fd26-7e5b-45f7-811c-7d356f140593 (Ran emptySleepJob in 1.000437111 Seconds)
    -> JOB 5/16 COMPLETED:           e388ad83-4ba2-454d-925d-db40988fd161 with Worker: 41f951d9-a878-4f5e-9265-9924ec2f815a (Ran emptySleepJob in 1.000455047 Seconds)
    -> JOB 6/16 COMPLETED:           0eda2855-7f5b-4f61-82fd-161e809892ec with Worker: d5be5f1c-2477-4d5d-8322-2ef428b0f535 (Ran emptySleepJob in 1.000459854 Seconds)
  JOB 12/16 STARTED:                 765de944-9ec6-4b50-af3a-b3decd81e967 with Worker: 32332fcb-64f7-43c7-8c3f-26e18648384e
    -> JOB 11/16 COMPLETED:          c2336100-de1f-4433-b82d-683f48861e64 with Worker: 6916721e-d3da-493e-bce8-71cd4eb3de10 (Ran emptySleepJob in 1.001107541 Seconds)
    -> JOB 13/16 COMPLETED:          f581ddd6-aec2-423f-be13-9c6bc0fddeaa with Worker: da2e47c2-a7bf-428f-b029-2c4ea51422d4 (Ran emptySleepJob in 1.001115304 Seconds)
    -> JOB 15/16 COMPLETED:          e8740614-c7cf-41b1-b946-6077ba5b8a20 with Worker: 41f951d9-a878-4f5e-9265-9924ec2f815a (Ran emptySleepJob in 1.001170553 Seconds)
    -> JOB 16/16 COMPLETED:          a782c1e8-6569-47bf-99c0-15de8846ed30 with Worker: d5be5f1c-2477-4d5d-8322-2ef428b0f535 (Ran emptySleepJob in 1.001130361 Seconds)
    -> JOB 12/16 COMPLETED:          765de944-9ec6-4b50-af3a-b3decd81e967 with Worker: 32332fcb-64f7-43c7-8c3f-26e18648384e (Ran emptySleepJob in 1.00112921 Seconds)
    -> JOB 9/16 COMPLETED:           b9220f71-dab5-4329-b1c6-5f49a117d616 with Worker: 4add73e2-da2b-4637-a4fb-51543d5509b4 (Ran emptySleepJob in 1.001230054 Seconds)
    -> JOB 14/16 COMPLETED:          0684858e-9d93-42fd-a203-f4b73c40dfab with Worker: 81e7fd26-7e5b-45f7-811c-7d356f140593 (Ran emptySleepJob in 1.001195592 Seconds)
    -> JOB 10/16 COMPLETED:          ee86c5d7-8ff2-4c76-8dbd-78ad66d5aa1a with Worker: d21cdd06-f55e-4e3f-9c01-ed79fe543711 (Ran emptySleepJob in 1.001210098 Seconds)

------------------------------------------------------
Total time taken:                    2.002621 Seconds

+-------------------+----------------+----------------+----------------+
| NUMBER OF WORKERS | NUMBER OF JOBS | EXECUTION TIME | SPEED INCREASE |
+-------------------+----------------+----------------+----------------+
|                 1 |             16 |      16.125234 | (baseline)     |
|                 2 |             16 |       8.045123 | +2x            |
|                 3 |             16 |       6.035669 | +2.67x         |
|                 4 |             16 |       4.003083 | +4.03x         |
|                 5 |             16 |       4.008031 | +4.02x         |
|                 6 |             16 |       3.015694 | +5.35x         |
|                 7 |             16 |       3.014996 | +5.35x         |
|                 8 |             16 |       2.002621 | +8.05x         |
+-------------------+----------------+----------------+----------------+
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
STARTINGWORKERCOUNT=
MAXWORKERCOUNT=
TOTALJOBCOUNT=
PPROF=
PPROFIP=
PPROFPORT=
```

The defaults are listed below when no values are set. **To override any/all, just set the env vars below.**

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

### Pprof

Run app.

In another terminal, use `pprof`, e.g.:

`go tool pprof http://192.168.0.99:6060/debug/pprof/profile?seconds=5`
`go tool pprof http://192.168.0.99:6060/debug/pprof/heap`
`go tool pprof http://192.168.0.99:6060/debug/pprof/goroutine`
`go tool pprof http://192.168.0.99:6060/debug/pprof/block`
`go tool pprof http://192.168.0.99:6060/debug/pprof/threadcreate`
`go tool pprof http://192.168.0.99:6060/debug/pprof/mutex`


go tool pprof -png http://192.168.0.99:6060/debug/pprof/heap > heap_1.png
go tool pprof -png http://192.168.0.99:6060/debug/pprof/heap > heap_8.png
go tool pprof -png http://192.168.0.99:6060/debug/pprof/heap > heap_16.png


go tool pprof -png http://192.168.129.52:6060/debug/pprof/heap > heap_16.png


curl -sK -v http://192.168.129.52:6060/debug/pprof/heap > heap.out
go tool pprof heap.out


curl -sK -v http://192.168.129.52:6060/debug/pprof/goroutine > goroutines-1.out
clear && go tool pprof goroutines-1.out

curl -sK -v http://192.168.129.52:6060/debug/pprof/goroutine > goroutines-4.out
clear && go tool pprof goroutines-4.out

curl -sK -v http://192.168.129.52:6060/debug/pprof/goroutine > goroutines-16.out
clear && go tool pprof goroutines-16.out