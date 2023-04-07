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

Below are the results for running 32 workers, or `runtime.NumCPU() * 4`:

```
+-------------------+----------------+----------------+----------------+
| NUMBER OF WORKERS | NUMBER OF JOBS | EXECUTION TIME | SPEED INCREASE |
+-------------------+----------------+----------------+----------------+
|                 1 |             16 |      16.030430 | (baseline)     |
|                 2 |             16 |       8.019319 | +2x            |
|                 3 |             16 |       6.034474 | +2.66x         |
|                 4 |             16 |       4.007691 | +4x            |
|                 5 |             16 |       4.028187 | +3.98x         |
|                 6 |             16 |       3.009318 | +5.33x         |
|                 7 |             16 |       3.023195 | +5.3x          |
|                 8 |             16 |       2.042858 | +7.85x         |
|                 9 |             16 |       2.032193 | +7.89x         |
|                10 |             16 |       2.017141 | +7.95x         |
|                11 |             16 |       2.019986 | +7.94x         |
|                12 |             16 |       2.005618 | +7.99x         |
|                13 |             16 |       2.008364 | +7.98x         |
|                14 |             16 |       2.003370 | +8x            |
|                15 |             16 |       2.007190 | +7.99x         |
|                16 |             16 |       1.014751 | +15.8x         |
|                17 |             16 |       1.002632 | +15.99x        |
|                18 |             16 |       1.003825 | +15.97x        |
|                19 |             16 |       1.001292 | +16.01x        |
|                20 |             16 |       1.000506 | +16.02x        |
|                21 |             16 |       1.000729 | +16.02x        |
|                22 |             16 |       1.019471 | +15.72x        |
|                23 |             16 |       1.020171 | +15.71x        |
|                24 |             16 |       1.010650 | +15.86x        |
|                25 |             16 |       1.002440 | +15.99x        |
|                26 |             16 |       1.002254 | +15.99x        |
|                27 |             16 |       1.001130 | +16.01x        |
|                28 |             16 |       1.001569 | +16.01x        |
|                29 |             16 |       1.006593 | +15.93x        |
|                30 |             16 |       1.003033 | +15.98x        |
|                31 |             16 |       1.001035 | +16.01x        |
|                32 |             16 |       1.001349 | +16.01x        |
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

`go tool pprof http://192.168.0.99:6060/debug/pprof/profile?seconds=20`
`go tool pprof http://192.168.0.99:6060/debug/pprof/heap`
`go tool pprof http://192.168.0.99:6060/debug/pprof/goroutine`
