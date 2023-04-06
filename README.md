# Golang-Channels

## Run

`go run .`

## Build binary

`go build -o ./bin/ .`

## Run binary

`./bin/golangchannels`

## Pprof

Run app.

In another terminal, use `pprof`, e.g.:

`go tool pprof http://192.168.0.99:6060/debug/pprof/profile?seconds=20`
`go tool pprof http://192.168.0.99:6060/debug/pprof/heap`
`go tool pprof http://192.168.0.99:6060/debug/pprof/goroutine`

## Results: emptySleepJob()

The table below represents:

* 1 to 8 workers running the `emptySleepJob()` function 16 times.
* Starting with 1 worker `(baseline)`, after each completion the worker count is increased by 1, until `runtime.NumCPU()` number of workers (8 CPUs in the example) is reached.
* As each job is just a `time.Sleep(1 * time.Second)` command, each job should take about **1 second to execute.**
* As expected, **1 worker completing 16 of these jobs takes about 16 seconds and 8 workers runs approximately 8 times as fast.**

```
Workers: 8:                          Job Name: emptySleepJob
--------------------------------------------------
  Workers In Use:                    8
  Workers Available:                 8
  Workers Idle:                      0
  Number of Jobs:                    16
--------------------------------------------------

Allocating Worker #1:                3d0dc960-9b9f-4588-9f4d-a5dbd071a19b
Allocating Worker #2:                6e524eb4-98f8-40df-b2e3-73a11615f4af
Allocating Worker #3:                12b5d356-c821-49b8-9c24-4eb2d44b948f
Allocating Worker #4:                372297a4-698b-42a5-ab85-e314b26b6c1c
Allocating Worker #5:                cd10b31d-be55-4390-87bf-20771318ea15
  Allocating Job #1:                 3489c34a-11a5-4b67-a5a2-d9381f694b5a
  Allocating Job #2:                 a68a348b-f4a6-47c3-bc2b-a3f8f132caa7
  Allocating Job #3:                 b77ca730-0343-4bf4-b4bd-2208c939bc6c
  Allocating Job #4:                 fd394b48-4bb4-481a-abe3-f2ae220519c5
  Allocating Job #5:                 ae9affc5-d518-4df3-9620-b8a5c75b922a
  JOB 1/16 STARTED:                  3489c34a-11a5-4b67-a5a2-d9381f694b5a with Worker: 6e524eb4-98f8-40df-b2e3-73a11615f4af
  JOB 4/16 STARTED:                  fd394b48-4bb4-481a-abe3-f2ae220519c5 with Worker: 12b5d356-c821-49b8-9c24-4eb2d44b948f
  Allocating Job #6:                 58bed1af-3a8f-4f80-b097-e6fd86b99990
  JOB 2/16 STARTED:                  a68a348b-f4a6-47c3-bc2b-a3f8f132caa7 with Worker: cd10b31d-be55-4390-87bf-20771318ea15
Allocating Worker #6:                1484cac6-7404-477a-9c34-f84a38b5d92f
Allocating Worker #7:                a5250688-d975-4874-ac73-b8a56c655274
  JOB 5/16 STARTED:                  ae9affc5-d518-4df3-9620-b8a5c75b922a with Worker: 372297a4-698b-42a5-ab85-e314b26b6c1c
  JOB 3/16 STARTED:                  b77ca730-0343-4bf4-b4bd-2208c939bc6c with Worker: 3d0dc960-9b9f-4588-9f4d-a5dbd071a19b
Allocating Worker #8:                4614814b-64b9-4829-befb-cb29fabbab1d
  Allocating Job #7:                 fede5656-5d42-4322-9933-4e1bad722d70
  Allocating Job #8:                 178d5d5a-e9d4-4eaa-9759-50c91efc8455
  Allocating Job #9:                 5ff567a9-66b0-4b2e-acc5-79e188e911a0
  Allocating Job #10:                47f426a7-8050-4650-b588-b8ed3d884807
  Allocating Job #11:                4d58289e-85cd-4c1a-a95e-4fedb8e78fef
  Allocating Job #12:                d9db676b-a972-410d-a9bb-5d5fd4cdbbcf
  Allocating Job #13:                51329299-adb0-4a96-8184-c010a4b47b10
  Allocating Job #14:                2b7b2d2e-190a-43d9-b684-a416f28a8944
  Allocating Job #15:                c40075e9-83c4-48c4-a105-b8f572062354
  Allocating Job #16:                2d002fc9-b2b9-43dd-9958-a2f6e9db4486
  JOB 8/16 STARTED:                  178d5d5a-e9d4-4eaa-9759-50c91efc8455 with Worker: 4614814b-64b9-4829-befb-cb29fabbab1d
  JOB 6/16 STARTED:                  58bed1af-3a8f-4f80-b097-e6fd86b99990 with Worker: 1484cac6-7404-477a-9c34-f84a38b5d92f
  JOB 7/16 STARTED:                  fede5656-5d42-4322-9933-4e1bad722d70 with Worker: a5250688-d975-4874-ac73-b8a56c655274
  JOB 9/16 STARTED:                  5ff567a9-66b0-4b2e-acc5-79e188e911a0 with Worker: a5250688-d975-4874-ac73-b8a56c655274
  JOB 10/16 STARTED:                 47f426a7-8050-4650-b588-b8ed3d884807 with Worker: 372297a4-698b-42a5-ab85-e314b26b6c1c
  JOB 11/16 STARTED:                 4d58289e-85cd-4c1a-a95e-4fedb8e78fef with Worker: 6e524eb4-98f8-40df-b2e3-73a11615f4af
  JOB 12/16 STARTED:                 d9db676b-a972-410d-a9bb-5d5fd4cdbbcf with Worker: 12b5d356-c821-49b8-9c24-4eb2d44b948f
    -> JOB 7/16 COMPLETED:           fede5656-5d42-4322-9933-4e1bad722d70 with Worker: a5250688-d975-4874-ac73-b8a56c655274 (Ran emptySleepJob in 1.003224526 Seconds)
  JOB 13/16 STARTED:                 51329299-adb0-4a96-8184-c010a4b47b10 with Worker: 4614814b-64b9-4829-befb-cb29fabbab1d
  JOB 15/16 STARTED:                 c40075e9-83c4-48c4-a105-b8f572062354 with Worker: 3d0dc960-9b9f-4588-9f4d-a5dbd071a19b
    -> JOB 5/16 COMPLETED:           ae9affc5-d518-4df3-9620-b8a5c75b922a with Worker: 372297a4-698b-42a5-ab85-e314b26b6c1c (Ran emptySleepJob in 1.003378164 Seconds)
  JOB 16/16 STARTED:                 2d002fc9-b2b9-43dd-9958-a2f6e9db4486 with Worker: cd10b31d-be55-4390-87bf-20771318ea15
    -> JOB 1/16 COMPLETED:           3489c34a-11a5-4b67-a5a2-d9381f694b5a with Worker: 6e524eb4-98f8-40df-b2e3-73a11615f4af (Ran emptySleepJob in 1.003463188 Seconds)
    -> JOB 4/16 COMPLETED:           fd394b48-4bb4-481a-abe3-f2ae220519c5 with Worker: 12b5d356-c821-49b8-9c24-4eb2d44b948f (Ran emptySleepJob in 1.003456132 Seconds)
    -> JOB 8/16 COMPLETED:           178d5d5a-e9d4-4eaa-9759-50c91efc8455 with Worker: 4614814b-64b9-4829-befb-cb29fabbab1d (Ran emptySleepJob in 1.003353988 Seconds)
    -> JOB 6/16 COMPLETED:           58bed1af-3a8f-4f80-b097-e6fd86b99990 with Worker: 1484cac6-7404-477a-9c34-f84a38b5d92f (Ran emptySleepJob in 1.003334243 Seconds)
    -> JOB 3/16 COMPLETED:           b77ca730-0343-4bf4-b4bd-2208c939bc6c with Worker: 3d0dc960-9b9f-4588-9f4d-a5dbd071a19b (Ran emptySleepJob in 1.003453667 Seconds)
    -> JOB 2/16 COMPLETED:           a68a348b-f4a6-47c3-bc2b-a3f8f132caa7 with Worker: cd10b31d-be55-4390-87bf-20771318ea15 (Ran emptySleepJob in 1.003505038 Seconds)
  JOB 14/16 STARTED:                 2b7b2d2e-190a-43d9-b684-a416f28a8944 with Worker: 1484cac6-7404-477a-9c34-f84a38b5d92f
    -> JOB 14/16 COMPLETED:          2b7b2d2e-190a-43d9-b684-a416f28a8944 with Worker: 1484cac6-7404-477a-9c34-f84a38b5d92f (Ran emptySleepJob in 1.001892706 Seconds)
    -> JOB 9/16 COMPLETED:           5ff567a9-66b0-4b2e-acc5-79e188e911a0 with Worker: a5250688-d975-4874-ac73-b8a56c655274 (Ran emptySleepJob in 1.002065712 Seconds)
    -> JOB 11/16 COMPLETED:          4d58289e-85cd-4c1a-a95e-4fedb8e78fef with Worker: 6e524eb4-98f8-40df-b2e3-73a11615f4af (Ran emptySleepJob in 1.002081539 Seconds)
    -> JOB 10/16 COMPLETED:          47f426a7-8050-4650-b588-b8ed3d884807 with Worker: 372297a4-698b-42a5-ab85-e314b26b6c1c (Ran emptySleepJob in 1.002088789 Seconds)
    -> JOB 12/16 COMPLETED:          d9db676b-a972-410d-a9bb-5d5fd4cdbbcf with Worker: 12b5d356-c821-49b8-9c24-4eb2d44b948f (Ran emptySleepJob in 1.002080951 Seconds)
    -> JOB 16/16 COMPLETED:          2d002fc9-b2b9-43dd-9958-a2f6e9db4486 with Worker: cd10b31d-be55-4390-87bf-20771318ea15 (Ran emptySleepJob in 1.002007645 Seconds)
    -> JOB 15/16 COMPLETED:          c40075e9-83c4-48c4-a105-b8f572062354 with Worker: 3d0dc960-9b9f-4588-9f4d-a5dbd071a19b (Ran emptySleepJob in 1.00203514 Seconds)
    -> JOB 13/16 COMPLETED:          51329299-adb0-4a96-8184-c010a4b47b10 with Worker: 4614814b-64b9-4829-befb-cb29fabbab1d (Ran emptySleepJob in 1.002061802 Seconds)

--------------------------------------------------
Total time taken:2.005891           Seconds


+-------------------+----------------+----------------+----------------+
| NUMBER OF WORKERS | NUMBER OF JOBS | EXECUTION TIME | SPEED INCREASE |
+-------------------+----------------+----------------+----------------+
|                 1 |             16 |      16.064956 | (baseline)     |
|                 2 |             16 |       8.054211 | +1.99x         |
|                 3 |             16 |       6.008089 | +2.67x         |
|                 4 |             16 |       4.012115 | +4x            |
|                 5 |             16 |       4.007524 | +4.01x         |
|                 6 |             16 |       3.021437 | +5.32x         |
|                 7 |             16 |       3.021649 | +5.32x         |
|                 8 |             16 |       2.005891 | +8.01x         |
+-------------------+----------------+----------------+----------------+
```