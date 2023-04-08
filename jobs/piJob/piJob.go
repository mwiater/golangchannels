// Package piJob is used to calulate Pi
package piJob

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"time"

	"github.com/mattwiater/golangchannels/config"
	"github.com/mattwiater/golangchannels/structs"
)

type Job structs.Job

func (job Job) PiJob() (string, float64) {
	jobStartTime := time.Now()

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 1000
	randNumber := rand.Intn(max-min+1) + min

	temp, _ := strconv.ParseUint("2000", 10, 32)
	precision := uint(temp+uint64(randNumber)) * 3

	PI := Pi(precision)

	jobEndTime := time.Now()
	jobElapsed := jobEndTime.Sub(jobStartTime)

	jobResult := structs.SleepJobResult{}
	jobResult.SleepTime = time.Duration(config.EmptySleepJobSleepTimeMs).String()
	jobResult.Elapsed = jobElapsed.String()
	jobResult.Status = strconv.FormatBool(true)
	jobResult.Data = PI

	jobResultString, err := json.Marshal(jobResult)
	if err != nil {
		fmt.Println(err)
	}

	return string(jobResultString), jobElapsed.Seconds()
}

func Pi(precision uint) *big.Float {
	k := 0
	pi := new(big.Float).SetPrec(precision).SetFloat64(0)
	k1k2k3 := new(big.Float).SetPrec(precision).SetFloat64(0)
	k4k5k6 := new(big.Float).SetPrec(precision).SetFloat64(0)
	temp := new(big.Float).SetPrec(precision).SetFloat64(0)
	minusOne := new(big.Float).SetPrec(precision).SetFloat64(-1)
	total := new(big.Float).SetPrec(precision).SetFloat64(0)

	two2Six := math.Pow(2, 6)
	two2SixBig := new(big.Float).SetPrec(precision).SetFloat64(two2Six)
	for {
		if k > int(precision) {
			break
		}
		t1 := float64(float64(1) / float64(10*k+9))
		k1 := new(big.Float).SetPrec(precision).SetFloat64(t1)
		t2 := float64(float64(64) / float64(10*k+3))
		k2 := new(big.Float).SetPrec(precision).SetFloat64(t2)
		t3 := float64(float64(32) / float64(4*k+1))
		k3 := new(big.Float).SetPrec(precision).SetFloat64(t3)
		k1k2k3.Sub(k1, k2)
		k1k2k3.Sub(k1k2k3, k3)

		t4 := float64(float64(4) / float64(10*k+5))
		k4 := new(big.Float).SetPrec(precision).SetFloat64(t4)
		t5 := float64(float64(4) / float64(10*k+7))
		k5 := new(big.Float).SetPrec(precision).SetFloat64(t5)
		t6 := float64(float64(1) / float64(4*k+3))
		k6 := new(big.Float).SetPrec(precision).SetFloat64(t6)
		k4k5k6.Add(k4, k5)
		k4k5k6.Add(k4k5k6, k6)
		k4k5k6 = k4k5k6.Mul(k4k5k6, minusOne)
		temp.Add(k1k2k3, k4k5k6)

		k7temp := new(big.Int).Exp(big.NewInt(-1), big.NewInt(int64(k)), nil)
		k8temp := new(big.Int).Exp(big.NewInt(1024), big.NewInt(int64(k)), nil)

		k7 := new(big.Float).SetPrec(precision).SetFloat64(0)
		k7.SetInt(k7temp)
		k8 := new(big.Float).SetPrec(precision).SetFloat64(0)
		k8.SetInt(k8temp)

		t9 := float64(256) / float64(10*k+1)
		k9 := new(big.Float).SetPrec(precision).SetFloat64(t9)
		k9.Add(k9, temp)
		total.Mul(k9, k7)
		total.Quo(total, k8)
		pi.Add(pi, total)

		k = k + 1
	}
	pi.Quo(pi, two2SixBig)
	return pi
}
